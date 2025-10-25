package internal

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/marcuwynu23/terraform-provider-nmap/internal/datasources"
)

// NewProvider initializes the provider entrypoint.
func NewProvider() provider.Provider {
	return &nmapProvider{}
}

// nmapProvider implements the Terraform provider interface.
type nmapProvider struct{}

// Metadata sets the provider type name.
func (p *nmapProvider) Metadata(_ context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "nmap"
}

// Schema defines provider-level configuration fields.
func (p *nmapProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Provider for running nmap scans via Terraform.",
		Attributes: map[string]schema.Attribute{
			"nmap_path": schema.StringAttribute{
				Optional:    true,
				Description: "Path to the nmap binary (defaults to system path).",
			},
		},
	}
}

// Configure sets up provider configuration. Optional here.
func (p *nmapProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// You could load custom config here (e.g., custom nmap path)
}

// Resources defines the provider's supported resources (none yet).
func (p *nmapProvider) Resources(_ context.Context) []func() resource.Resource {
	return nil
}

// DataSources defines the provider's supported data sources.
func (p *nmapProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		datasources.NewNmapScanDataSource,
	}
}
