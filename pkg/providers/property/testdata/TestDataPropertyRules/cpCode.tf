provider "akamai" {
  edgerc = "~/.edgerc"
}

data "akamai_property_rules" "rules" {
  rules {
    behavior {
      name = "cpCode"
      option {
        key   = "id"
        value = "cpc_12345"
      }
    }
  }
}
