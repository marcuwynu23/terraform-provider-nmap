package datasources

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	nm "github.com/Ullaakut/nmap/v3"
)

type nmapScanDataSource struct{}

func NewNmapScanDataSource() datasource.DataSource {
	return &nmapScanDataSource{}
}

type nmapScanModel struct {
	ID     types.String `tfsdk:"id"`
	Target types.String `tfsdk:"target"`
	Hosts  []hostModel  `tfsdk:"hosts"`
}

type hostModel struct {
	IP    types.String   `tfsdk:"ip"`
	Ports []types.String `tfsdk:"ports"`
}

func (d *nmapScanDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_scan"
}

func (d *nmapScanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Runs an nmap scan on a given target and returns discovered hosts and open ports.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:    true,
				Description: "Unique identifier for the scan result.",
			},
			"target": schema.StringAttribute{
				Required:    true,
				Description: "Target host or network to scan.",
			},
			"hosts": schema.ListNestedAttribute{
				Computed:    true,
				Description: "List of discovered hosts and open ports.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"ip": schema.StringAttribute{
							Computed:    true,
							Description: "IP address of the discovered host.",
						},
						"ports": schema.ListAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "List of open ports in <port>/<protocol> format.",
						},
					},
				},
			},
		},
	}
}

func (d *nmapScanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state nmapScanModel

	diags := req.Config.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	target := state.Target.ValueString()

	// Context with timeout
	ctxTimeout, cancel := context.WithTimeout(ctx, 60*time.Second)
	defer cancel()

	// Use the new API: first param is context
	scanner, err := nm.NewScanner(
		ctxTimeout,
		nm.WithTargets(target),
		nm.WithSkipHostDiscovery(),
	)
	if err != nil {
		resp.Diagnostics.AddError("Failed to create nmap scanner", err.Error())
		return
	}

result, warnings, err := scanner.Run()
if err != nil {
	resp.Diagnostics.AddError("Nmap scan failed", err.Error())
	return
}

// Handle warnings properly
if warnings != nil && len(*warnings) > 0 {
	resp.Diagnostics.AddWarning("Nmap warnings", fmt.Sprintf("%v", *warnings))
}


	var hosts []hostModel

	for _, host := range result.Hosts {
		if len(host.Addresses) == 0 {
			continue
		}
		ip := host.Addresses[0].Addr

		var ports []types.String
		for _, port := range host.Ports {
			if port.State.State == "open" {
				ports = append(ports, types.StringValue(fmt.Sprintf("%d/%s", port.ID, port.Protocol)))
			}
		}

		hosts = append(hosts, hostModel{
			IP:    types.StringValue(ip),
			Ports: ports,
		})
	}

	state.ID = types.StringValue(fmt.Sprintf("%d", time.Now().Unix()))
	state.Hosts = hosts

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
}
