---
page_title: "Provider: Nmap"
description: |-
  The Nmap provider allows Terraform to run network scans using Nmap and use the results as data sources.
---

# Nmap Provider

The Nmap provider integrates [Nmap](https://nmap.org/) network scanning into Terraform,
allowing you to perform network discovery and port scanning as part of your infrastructure
automation workflows.

Use the navigation to the left to read about available data sources.

## Example Usage

```hcl
terraform {
  required_providers {
    nmap = {
      source  = "registry.terraform.io/marcuwynu23/nmap"
      version = "~> 0.1.0"
    }
  }
}

provider "nmap" {}

data "nmap_scan" "example" {
  target = "scanme.nmap.org"
}

output "scan_result" {
  value = data.nmap_scan.example.hosts
}
```

## Schema

### Optional

- `nmap_path` (String) - Path to the nmap binary. Defaults to the system PATH.
