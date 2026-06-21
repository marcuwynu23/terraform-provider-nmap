---
page_title: "Data Source: nmap_scan"
description: |-
  Runs an Nmap scan on a target and returns discovered hosts and open ports.
---

# Data Source: nmap_scan

Runs an Nmap scan on a given target (hostname or IP) and returns the list of discovered
hosts with their open ports.

## Example Usage

```hcl
data "nmap_scan" "example" {
  target = "scanme.nmap.org"
}

output "hosts" {
  value = data.nmap_scan.example.hosts
}
```

## Schema

### Required

- `target` (String) - Target hostname, IP address, or CIDR network to scan.

### Read-Only

- `id` (String) - Unique identifier for the scan result.
- `hosts` (List of Object) - List of discovered hosts, each containing:
  - `ip` (String) - IP address of the discovered host.
  - `ports` (List of String) - List of open ports in `<port>/<protocol>` format (e.g., `22/tcp`).
