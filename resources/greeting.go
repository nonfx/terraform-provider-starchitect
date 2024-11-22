package resources

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	resschema "github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// GreetingResource defines the resource implementation.
type GreetingResource struct{}

// GreetingResourceModel describes the resource data model.
type GreetingResourceModel struct {
	Name     types.String `tfsdk:"name"`
	Greeting types.String `tfsdk:"greeting"`
	Id       types.String `tfsdk:"id"`
}

func NewGreetingResource() resource.Resource {
	return &GreetingResource{}
}

func (r *GreetingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_greeting"
}

func (r *GreetingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = resschema.Schema{
		Description: "Manages a greeting message.",
		Attributes: map[string]resschema.Attribute{
			"name": resschema.StringAttribute{
				Description: "Name of the person to greet",
				Required:    true,
			},
			"greeting": resschema.StringAttribute{
				Description: "Generated greeting message",
				Computed:    true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"id": resschema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *GreetingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan GreetingResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate greeting
	name := plan.Name.ValueString()
	greeting := fmt.Sprintf("Hello from %s!", name)

	// Set computed values
	plan.Greeting = types.StringValue(greeting)
	plan.Id = types.StringValue(name)

	tflog.Info(ctx, "Created greeting resource", map[string]any{
		"name":     name,
		"greeting": greeting,
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *GreetingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state GreetingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Re-generate greeting (in a real provider, you might fetch this from an API)
	name := state.Name.ValueString()
	greeting := fmt.Sprintf("Hello from %s!", name)
	state.Greeting = types.StringValue(greeting)

	tflog.Info(ctx, "Read greeting resource", map[string]any{
		"name":     name,
		"greeting": greeting,
	})

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}

func (r *GreetingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan GreetingResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Generate new greeting
	name := plan.Name.ValueString()
	greeting := fmt.Sprintf("Hello from %s!", name)
	plan.Greeting = types.StringValue(greeting)

	tflog.Info(ctx, "Updated greeting resource", map[string]any{
		"name":     name,
		"greeting": greeting,
	})

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
}

func (r *GreetingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state GreetingResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Info(ctx, "Deleted greeting resource", map[string]any{
		"name": state.Name.ValueString(),
	})
}
