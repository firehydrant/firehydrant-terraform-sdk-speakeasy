---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "firehydrant-terraform-sdk Provider"
subcategory: ""
description: |-
  
---

# firehydrant-terraform-sdk Provider



## Example Usage

```terraform
terraform {
  required_providers {
    firehydrant-terraform-sdk = {
      source  = "speakeasy/firehydrant-terraform-sdk"
      version = "0.0.5"
    }
  }
}

provider "firehydrant-terraform-sdk" {
  # Configuration options
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `api_key` (String, Sensitive)
- `server_url` (String) Server URL (defaults to https://api.firehydrant.io/)
