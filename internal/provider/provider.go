package provider

import (
	"context"
	"os"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &sampleProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &sampleProvider{
			version: version,
		}
	}
}

// sampleProvider is the provider implementation.
type sampleProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// sampleProviderModel maps provider schema data to a Go type.
type sampleProviderModel struct {
	Prefix    types.String `tfsdk:"prefix"`
	Delimiter types.String `tfsdk:"delimiter"`
}

// Metadata returns the provider type name.
func (p *sampleProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "sample"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *sampleProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"prefix": schema.StringAttribute{
				Description: "String added as a prefix to the resource's string",
				Optional:    true,
			},
			"delimiter": schema.StringAttribute{
				Description: "Delimiter joining the prefix and the string defined in resource",
				Optional:    true,
			},
		},
	}
}

func (p *sampleProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config sampleProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Prefix.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("prefix"),
			"Unknown Prefix",
			"The provider has unknown configuration value for the prefix. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the SAMPLE_PREFIX environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	prefix := os.Getenv("SAMPLE_PREFIX")

	if !config.Prefix.IsNull() {
		prefix = config.Prefix.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if prefix == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("prefix"),
			"Missing sample Prefix",
			"The provider has unknown configuration value for the prefix. "+
				"Set the prefix value in the configuration or use the SAMPLE_PREFIX environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	delimiter := os.Getenv("SAMPLE_DELIMITER")

	if !config.Delimiter.IsNull() {
		delimiter = config.Delimiter.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if prefix == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("delimiter"),
			"Missing sample Delimiter",
			"The provider has unknown configuration value for the delimiter. "+
				"Set the delimiter value in the configuration or use the SAMPLE_DELIMITER environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	//     resp.DataSourceData = prefix
	resp.ResourceData = prefix + delimiter
}

// DataSources defines the data sources implemented in the provider.
func (p *sampleProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *sampleProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewSampleResource,
	}
}
