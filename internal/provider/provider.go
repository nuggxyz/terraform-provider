package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/nuggxyz/terraform-provider/internal/client"
	"github.com/nuggxyz/terraform-provider/internal/datasources/typer"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &ProviderClient{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &ProviderClient{
			version: version,
		}
	}
}

// ProviderClient is the provider implementation.
type ProviderClient struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// ProviderClientModel maps provider schema data to a Go type.
type ProviderClientModel struct {
}

// Metadata returns the provider type name.
func (p *ProviderClient) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nugg.xyz"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *ProviderClient) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A provider built by nugg.xyz",
		Attributes:  map[string]schema.Attribute{},
	}
}

func (p *ProviderClient) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring nugg.xyz client")

	// Retrieve provider data from configuration
	var config ProviderClientModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "nugg.xyz_password")

	tflog.Debug(ctx, "Creating nugg.xyz client")

	// Create a new nugg.xyz client using the configuration values
	clnt, err := client.NewClient()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create nugg.xyz API Client",
			"An unexpected error occurred when creating the nugg.xyz API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"nugg.xyz Client Error: "+err.Error(),
		)
		return
	}

	// Make the nugg.xyz client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = clnt
	resp.ResourceData = clnt

	tflog.Info(ctx, "Configured nugg.xyz client", map[string]any{"success": true})
}

// DataSources defines the data sources implemented in the provider.
func (p *ProviderClient) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		typer.NewTyperDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *ProviderClient) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}
