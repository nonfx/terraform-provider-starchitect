package starchitect

import (
	"context"

	"terraform-provider-starchitect/resources"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	prschema "github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &starchitectProvider{
			version: version,
		}
	}
}

type starchitectProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *starchitectProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "starchitect"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *starchitectProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = prschema.Schema{
		Attributes: map[string]prschema.Attribute{
			"host": prschema.StringAttribute{
				Optional: true,
			},
			"username": prschema.StringAttribute{
				Optional: true,
			},
			"password": prschema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *starchitectProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *starchitectProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return nil
}

// Resources defines the resources implemented in the provider.
func (p *starchitectProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		resources.NewGreetingResource,
		resources.NewIACPACResource,
	}
}
