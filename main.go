package main

import (
    "context"
    "log"

    "github.com/hashicorp/terraform-plugin-framework/providerserver"
    "github.com/marcuwynu23/terraform-provider-nmap/internal"
)

func main() {
    opts := providerserver.ServeOpts{
        // Address: "registry.terraform.io/marcuwynu23/nmap", // for local dev use any namespace
        Address: "local/nmap", // for local dev use any namespace
    }

    if err := providerserver.Serve(context.Background(), internal.NewProvider, opts); err != nil {
        log.Fatalf("error serving provider: %s", err)
    }
}
