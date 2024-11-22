package resources

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

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
	ScanResult types.String `tfsdk:"scan_result"`
}

func NewIACPACResource() resource.Resource {
	return &IACPACResource{}
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
				Required:    true,
			},
			"scan_result": resschema.StringAttribute{
				Description: "Generated scan result",
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

	scanResult := GetScanResult(iacPath, pacPath)
	plan.ScanResult = types.StringValue(scanResult)

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

	scanResult := GetScanResult(iacPath, pacPath)
	state.ScanResult = types.StringValue(scanResult)

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

	scanResult := GetScanResult(iacPath, pacPath)
	plan.ScanResult = types.StringValue(scanResult)

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

func GetScanResult(iacPath, pacPath string) string {
	var stderr bytes.Buffer
	tempDir, err := os.MkdirTemp("", "regula-scan")
	if err != nil {
		return fmt.Sprintf("Error creating temporary directory: %v\n", err)
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
		return fmt.Sprintf("Error creating output file: %v\n", err)
	}
	defer output.Close()

	cmd.Stdout = output
	cmd.Stderr = &stderr

	err = cmd.Run()
	if err != nil {
		if bytes.Contains(stderr.Bytes(), []byte("rego_type_error")) {
			return fmt.Sprintf("Error: rego_type_error encountered. %v", string(stderr.String()))
		}
		err = nil
	}

	content, err := os.ReadFile(outputFile)
	if err != nil {
		return fmt.Sprintf("Error reading output file: %s %v\n", outputFile, err)
	}

	return string(content)
}
