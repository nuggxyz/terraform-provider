package typer

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"

	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/nuggxyz/terraform-provider/internal/client"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &typerDataSource{}
	_ datasource.DataSourceWithConfigure = &typerDataSource{}
)

// NewTyperDataSource is a helper function to simplify the provider implementation.
func NewTyperDataSource() datasource.DataSource {
	return &typerDataSource{}
}

// typerDataSource is the data source implementation.
type typerDataSource struct {
	client *client.Client
}

// typerDataSourceModel maps the data source schema data.
type typerDataSourceModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// typerIngredientsModel maps coffee ingredients data.
type typerIngredientsModel struct {
	ID types.Int64 `tfsdk:"id"`
}

// Metadata returns the data source type name.
func (d *typerDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_typer"
}

// Schema defines the schema for the data source.
func (d *typerDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Fetches the list of typer.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Description: "Numeric identifier of the coffee.",
				Computed:    true,
			},
			"name": schema.StringAttribute{
				Description: "Product name of the coffee.",
				Computed:    true,
			},
		},
	}
}

// Configure adds the provider configured client to the data source.
func (d *typerDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*client.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *client.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

// Read refreshes the Terraform state with the latest data.
func (d *typerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {

	coffeeState := typerDataSourceModel{
		ID:   types.Int64Value(1),
		Name: types.StringValue("placeholder"),
	}

	// Set state
	diags := resp.State.Set(ctx, &coffeeState)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}
