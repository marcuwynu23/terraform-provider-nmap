terraform {
  required_providers {
    nmap = {
      source  = "local/nmap"
      version = "0.1.0"
    }
  }
}

provider "nmap" {}

data "nmap_scan" "example" {
  target = "scanme.nmap.org"
}

output "scan_result" {
  value = data.nmap_scan.example.result
}
