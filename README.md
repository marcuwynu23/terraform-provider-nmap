# Terraform Provider for Nmap

A custom Terraform provider that integrates [Nmap](https://nmap.org/) — the powerful open-source network scanner — into Terraform.
This provider lets you perform network discovery and port scanning as part of your infrastructure automation workflows.

---

## 🌐 Overview

`terraform-provider-nmap` allows Terraform configurations to run Nmap scans and expose the results as Terraform data sources.
You can use it to automatically audit networks, detect open ports, or validate host accessibility during provisioning.

---

## ⚙️ Features

- Perform basic Nmap scans directly from Terraform
- Return discovered hosts and their open ports
- Integrate network scanning into your IaC (Infrastructure as Code) workflows
- Support for Windows and Linux environments

---

## 📦 Installation

### Terraform Registry (Recommended)

This provider is published on the [Terraform Registry](https://registry.terraform.io/providers/marcuwynu23/nmap).

Add the following to your Terraform configuration:

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
```

Then run:

```bash
terraform init
```

### Local Development

To build and use the provider locally without the registry:

```bash
go build -o terraform-provider-nmap.exe
```

Then configure Terraform with a dev override in `%APPDATA%\terraform.d\cli_config.tfrc` (Windows) or `~/.terraform.d/cli_config.tfrc` (Linux/macOS):

```hcl
provider_installation {
  dev_overrides {
    "registry.terraform.io/marcuwynu23/nmap" = "D:/path/to/terraform-provider-nmap"
  }
  direct {}
}
```

---

## 🧪 Usage Example

In your Terraform configuration (`main.tf`):

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

Run:

```bash
terraform init
terraform apply
```

Example output:

```
Outputs:

scan_result = [
  {
    ip    = "45.33.32.156"
    ports = ["22/tcp", "80/tcp"]
  }
]
```

---

## 🧱 Provider Structure

```
terraform-provider-nmap/
├── .goreleaser.yaml       # GoReleaser configuration for cross-platform builds
├── .github/
│   └── workflows/
│       └── release.yml    # GitHub Actions release workflow
├── docs/
│   ├── index.md           # Terraform Registry documentation
│   └── data-sources/
│       └── nmap_scan.md   # Data source documentation
├── internal/
│   ├── provider.go        # Defines provider configuration and data sources
│   └── datasources/
│       └── nmap_scan.go   # Implements the actual Nmap scan logic
├── main.go                # Provider entrypoint
├── test/                  # Example Terraform configuration for local testing
└── README.md
```

---

## 🧰 Development Notes

### Dependencies

- Go 1.21+
- Terraform Plugin Framework v1+
- [Ullaakut/nmap](https://github.com/Ullaakut/nmap) Go library

Install dependencies:

```bash
go mod tidy
```

### Build & Test Locally

```bash
go build -o terraform-provider-nmap.exe
cd test
terraform init
terraform apply
```

---

## 🚀 Roadmap

- [ ] Support for additional Nmap options (e.g. `-sV`, `-O`, timing templates)
- [ ] JSON output parsing
- [ ] Cross-platform binary builds
- [x] Publish to Terraform Registry

---

## 🧑‍💻 Author

**Marc Wayne Menorca (@marcuwynu23)**
Lead Developer & Open Source Contributor
[GitHub](https://github.com/marcuwynu23)

---

## ⚖️ License

This project is licensed under the **MIT License** — feel free to use, modify, and distribute.
