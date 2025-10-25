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

## 📦 Installation (Local Development)

Since this is a local provider, you must configure Terraform to load it from a local path.

### 1. Build the Provider

```bash
go build -o terraform-provider-nmap.exe
```

This will create the provider binary.

### 2. Place the Binary

Terraform expects the following directory structure:

```
plugins/
└── local/
    └── nmap/
        └── 0.1.0/
            └── windows_amd64/
                └── terraform-provider-nmap.exe
```

Or on Linux/macOS:

```
plugins/
└── local/
    └── nmap/
        └── 0.1.0/
            └── linux_amd64/
                └── terraform-provider-nmap
```

### 3. Create the Terraform CLI Config

Create a file:

- **Windows:** `%APPDATA%\terraform.d\cli_config.tfrc`
- **Linux/macOS:** `~/.terraform.d/cli_config.tfrc`

Add this content:

```hcl
provider_installation {
  dev_overrides {
    "local/nmap" = "D:/Projects/experiments/terraform-provider/terraform-provider-nmap/plugins/local/nmap/0.1.0/windows_amd64"
  }
  direct {}
}
```

Make sure the path matches your local setup.

---

## 🧪 Usage Example

In your Terraform configuration (`main.tf`):

```hcl
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
- [ ] Publish to Terraform Registry

---

## 🧑‍💻 Author

**Marc Wayne Menorca (@marcuwynu23)**
Lead Developer & Open Source Contributor
[GitHub](https://github.com/marcuwynu23)

---

## ⚖️ License

This project is licensed under the **MIT License** — feel free to use, modify, and distribute.
