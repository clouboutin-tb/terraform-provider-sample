package provider

import (
	"context"
	"fmt"
	//     "strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource = &sampleResource{}
)

// sampleResource is the resource implementation.
type sampleResource struct {
	prefix_with_delimiter string
}

// sampleResourceModel maps the resource schema data.
type sampleResourceModel struct {
	String      types.String `tfsdk:"string"`
	LastUpdated types.String `tfsdk:"last_updated"`
	Result      types.String `tfsdk:"result"`
}

// NewSampleResource is a helper function to simplify the provider implementation.
func NewSampleResource() resource.Resource {
	return &sampleResource{}
}

// Metadata returns the resource type name.
func (r *sampleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_sample"
}

// Schema defines the schema for the resource.
func (r *sampleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Basic resource that does nothing other than interact with the Terraform state",
		Attributes: map[string]schema.Attribute{
			"string": schema.StringAttribute{
				Description: "A simple string",
				Required:    true,
			},
			"last_updated": schema.StringAttribute{
				Description: "Resource update date in RFC850 format ",
				Computed:    true,
			},
			"result": schema.StringAttribute{
				Description: "Concatenation of provider.prefix, provider.delimiter and the resource string",
				Computed:    true,
			},
		},
	}
}

// Create a new resource.
func (r *sampleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan sampleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	plan.Result = types.StringValue(r.prefix_with_delimiter + plan.String.ValueString())
	//     plan.Result = r.prefix + plan.String.String()

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read resource information.
func (r *sampleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state sampleResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *sampleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values from plan
	var plan sampleResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))
	plan.Result = types.StringValue(r.prefix_with_delimiter + plan.String.ValueString())

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *sampleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}

// Configure adds the provider configured client to the resource.
func (r *sampleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {

	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	prefix_with_delimiter, ok := req.ProviderData.(string)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected string, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.prefix_with_delimiter = prefix_with_delimiter
}
