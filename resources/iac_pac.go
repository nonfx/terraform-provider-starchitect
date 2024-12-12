package resources

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"terraform-provider-starchitect/resources/utils"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	resschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// IACPACResource defines the resource implementation.
type IACPACResource struct{}

// IACPACResourceModel describes the resource data model.
type IACPACResourceModel struct {
	IACPath    types.String `tfsdk:"iac_path"`
	PACPath    types.String `tfsdk:"pac_path"`
	PACVersion types.String `tfsdk:"pac_version"`
	LogPath    types.String `tfsdk:"log_path"`
	ScanResult types.String `tfsdk:"scan_result"`
	Score      types.String `tfsdk:"score"`
	Threshold  types.String `tfsdk:"threshold"`
}

type RegulaRuleResult struct {
	Controls        []string          `json:"controls"`
	Families        []string          `json:"families"`
	Filepath        string            `json:"filepath"`
	InputType       string            `json:"input_type"`
	Provider        string            `json:"provider"`
	ResourceID      string            `json:"resource_id"`
	ResourceType    string            `json:"resource_type"`
	ResourceTags    map[string]string `json:"resource_tags"`
	RuleDescription string            `json:"rule_description"`
	RuleID          string            `json:"rule_id"`
	RuleMessage     string            `json:"rule_message"`
	RuleName        string            `json:"rule_name"`
	RuleRawResult   bool              `json:"rule_raw_result"`
	RuleResult      string            `json:"rule_result"`
	RuleSeverity    string            `json:"rule_severity"`
	RuleSummary     string            `json:"rule_summary"`
}

type RegulaOutput struct {
	RuleResults []RegulaRuleResult `json:"rule_results"`
}

func NewIACPACResource() resource.Resource {
	return &IACPACResource{}
}

func NewIACPACResourceWithModifyPlan() resource.ResourceWithModifyPlan {
	return &IACPACResource{}
}

func (r *IACPACResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	var plan IACPACResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	iacPath := plan.IACPath.ValueString()
	pacPath := plan.PACPath.ValueString()
	pacVersion := plan.PACVersion.ValueString()
	threshold := plan.Threshold.ValueString()
	logPath := plan.LogPath.ValueString()

	scanResult, score := GetScanResult(iacPath, pacPath, pacVersion, logPath)
	plan.ScanResult = types.StringValue(scanResult)
	plan.Score = types.StringValue(score)

	// Check threshold if specified
	if threshold != "" {
		thresholdValue, err := strconv.ParseFloat(threshold, 64)
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid threshold value",
				fmt.Sprintf("Could not parse threshold value: %v", err),
			)
			return
		}

		// Extract score value
		scoreStr := strings.TrimSpace(score)
		parts := strings.Split(scoreStr, "Score: ")
		if len(parts) != 2 {
			resp.Diagnostics.AddError(
				"Invalid score format",
				fmt.Sprintf("Could not parse score from: %s", score),
			)
			return
		}

		scoreStr = strings.TrimSuffix(parts[1], " percent")
		scoreValue, err := strconv.ParseFloat(scoreStr, 64)
		if err != nil {
			resp.Diagnostics.AddError(
				"Invalid score value",
				fmt.Sprintf("Could not parse score value: %v", err),
			)
			return
		}

		if scoreValue < thresholdValue {
			resp.Diagnostics.AddError(
				"Security Score Below Threshold",
				fmt.Sprintf("Security score (%.2f%%) is below the required threshold (%.2f%%)", scoreValue, thresholdValue),
			)
			return
		}
	}

	diags = resp.Plan.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IACPACResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iac_pac"
}

func (r *IACPACResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resschema.Schema{
		Description: "accepts IAC and PAC path to run policies",
		Attributes: map[string]resschema.Attribute{
			"iac_path": resschema.StringAttribute{
				Description: "IAC path",
				Required:    true,
			},
			"pac_path": resschema.StringAttribute{
				Description: "PAC path",
				Optional:    true,
			},
			"pac_version": resschema.StringAttribute{
				Description: "default PAC version",
				Optional:    true,
			},
			"log_path": resschema.StringAttribute{
				Description: "Path to store log files",
				Optional:    true,
			},
			"threshold": resschema.StringAttribute{
				Description: "Minimum required security score (0-100)",
				Optional:    true,
			},
			"scan_result": resschema.StringAttribute{
				Description: "Generated scan result",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"score": resschema.StringAttribute{
				Description: "Generated score. evaluated from scan result",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *IACPACResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IACPACResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	iacPath := plan.IACPath.ValueString()
	pacPath := plan.PACPath.ValueString()
	pacVersion := plan.PACVersion.ValueString()
	logPath := plan.LogPath.ValueString()

	scanResult, score := GetScanResult(iacPath, pacPath, pacVersion, logPath)
	plan.ScanResult = types.StringValue(scanResult)
	plan.Score = types.StringValue(score)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IACPACResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IACPACResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	iacPath := state.IACPath.ValueString()
	pacPath := state.PACPath.ValueString()
	pacVersion := state.PACVersion.ValueString()
	logPath := state.LogPath.ValueString()

	scanResult, score := GetScanResult(iacPath, pacPath, pacVersion, logPath)
	state.ScanResult = types.StringValue(scanResult)
	state.Score = types.StringValue(score)

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *IACPACResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IACPACResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	iacPath := plan.IACPath.ValueString()
	pacPath := plan.PACPath.ValueString()
	pacVersion := plan.PACVersion.ValueString()
	logPath := plan.LogPath.ValueString()

	scanResult, score := GetScanResult(iacPath, pacPath, pacVersion, logPath)
	plan.ScanResult = types.StringValue(scanResult)
	plan.Score = types.StringValue(score)

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *IACPACResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IACPACResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func formatRegulaOutput(regulaOutput RegulaOutput) string {
	var formatted strings.Builder

	// Add timestamp
	formatted.WriteString(fmt.Sprintf("Scan Time: %s\n", time.Now().Format(time.RFC3339)))
	formatted.WriteString("====================\n\n")

	// Calculate summary
	var passCount, failCount int
	for _, rule := range regulaOutput.RuleResults {
		if rule.RuleResult == "PASS" {
			passCount++
		} else if rule.RuleResult == "FAIL" {
			failCount++
		}
	}

	// Add summary
	formatted.WriteString("Summary:\n")
	formatted.WriteString(fmt.Sprintf("PASSED: %d\n", passCount))
	formatted.WriteString(fmt.Sprintf("FAILED: %d\n", failCount))
	formatted.WriteString("\nDetailed Results:\n")
	formatted.WriteString("----------------\n")

	for _, rule := range regulaOutput.RuleResults {
		formatted.WriteString(fmt.Sprintf("\nRule ID: %s\n", rule.RuleID))
		formatted.WriteString(fmt.Sprintf("Name: %s\n", rule.RuleName))
		formatted.WriteString(fmt.Sprintf("Result: %s\n", rule.RuleResult))
		formatted.WriteString(fmt.Sprintf("Severity: %s\n", rule.RuleSeverity))
		formatted.WriteString(fmt.Sprintf("Summary: %s\n", rule.RuleSummary))
		formatted.WriteString(fmt.Sprintf("Description: %s\n", rule.RuleDescription))

		if rule.ResourceType != "" {
			formatted.WriteString(fmt.Sprintf("Resource Type: %s\n", rule.ResourceType))
		}
		if rule.ResourceID != "" {
			formatted.WriteString(fmt.Sprintf("Resource ID: %s\n", rule.ResourceID))
		}
		if rule.RuleMessage != "" {
			formatted.WriteString(fmt.Sprintf("Message: %s\n", rule.RuleMessage))
		}

		if len(rule.Controls) > 0 {
			formatted.WriteString("Controls:\n")
			for _, control := range rule.Controls {
				formatted.WriteString(fmt.Sprintf("  - %s\n", control))
			}
		}

		if len(rule.Families) > 0 {
			formatted.WriteString("Families:\n")
			for _, family := range rule.Families {
				formatted.WriteString(fmt.Sprintf("  - %s\n", family))
			}
		}

		formatted.WriteString("---\n")
	}

	return formatted.String()
}

func writeToLogFiles(rawOutput string, formattedOutput string, logPath string) error {
	timestamp := time.Now().Format("20060102_150405")

	// Create log directory if it doesn't exist
	if logPath != "" {
		if err := os.MkdirAll(logPath, 0755); err != nil {
			return fmt.Errorf("failed to create log directory: %v", err)
		}
	}

	// Write raw output to JSON file
	rawFileName := fmt.Sprintf("%s_starchitect_raw.json", timestamp)
	if logPath != "" {
		rawFileName = filepath.Join(logPath, rawFileName)
	}
	if err := os.WriteFile(rawFileName, []byte(rawOutput), 0644); err != nil {
		return fmt.Errorf("failed to write raw output: %v", err)
	}

	// Write formatted summary to log file
	summaryFileName := fmt.Sprintf("%s_starchitect_summary.log", timestamp)
	if logPath != "" {
		summaryFileName = filepath.Join(logPath, summaryFileName)
	}
	if err := os.WriteFile(summaryFileName, []byte(formattedOutput), 0644); err != nil {
		return fmt.Errorf("failed to write summary: %v", err)
	}

	return nil
}

func calculateScore(regulaOutput RegulaOutput) string {
	var passCount, failCount int
	for _, rule := range regulaOutput.RuleResults {
		if rule.RuleResult == "PASS" {
			passCount++
		} else if rule.RuleResult == "FAIL" {
			failCount++
		}
	}

	total := passCount + failCount
	if total == 0 {
		return "no PASS or FAIL results found"
	}

	score := (float64(passCount) / float64(total)) * 100
	return fmt.Sprintf("PASSED: %d FAILED: %d Score: %.2f percent", passCount, failCount, score)
}

func GetScanResult(iacPath, pacPath, pacVersion, logPath string) (string, string) {

	var err error
	if pacPath == "" {
		pacPath, err = utils.GetDefaultPAC(iacPath, pacVersion)
		if err != nil {
			return err.Error(), ""
		}
	}

	var stderr bytes.Buffer
	tempDir, err := os.MkdirTemp("", "regula-scan")
	if err != nil {
		return fmt.Sprintf("Error creating temporary directory: %v\n", err), ""
	}
	defer os.RemoveAll(tempDir)

	// Create the output file path in the temporary directory
	outputFile := filepath.Join(tempDir, "results.json")

	cmd := exec.Command(
		"regula",
		"run",
		"-i", pacPath,
		iacPath,
		"-n",
		"-f", "json",
	)

	// Redirect the output to the temporary file
	output, err := os.Create(outputFile)
	if err != nil {
		return fmt.Sprintf("Error creating output file: %v\n", err), ""
	}
	defer output.Close()

	cmd.Stdout = output
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		if bytes.Contains(stderr.Bytes(), []byte("rego_type_error")) {
			return fmt.Sprintf("Error: rego_type_error encountered. %v", string(stderr.String())), ""
		}
		err = nil
	}

	// Read the raw output
	content, err := os.ReadFile(outputFile)
	if err != nil {
		return fmt.Sprintf("Error reading output file: %s %v\n", outputFile, err), ""
	}

	rawOutput := string(content)

	// Parse the JSON content
	var regulaOutput RegulaOutput
	if err := json.Unmarshal(content, &regulaOutput); err != nil {
		return fmt.Sprintf("Error parsing JSON output: %v\n", err), ""
	}

	// Calculate score
	score := calculateScore(regulaOutput)

	// Format the summary output
	formattedOutput := formatRegulaOutput(regulaOutput)

	// Write both outputs to separate files
	if err := writeToLogFiles(rawOutput, formattedOutput, logPath); err != nil {
		log.Printf("Warning: Failed to write to log files: %v", err)
	}

	return formattedOutput, score
}
