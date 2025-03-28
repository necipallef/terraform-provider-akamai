provider "akamai" {
  edgerc = "../../test/edgerc"
}

data "akamai_cloudlets_audience_segmentation_match_rule" "test" {
  match_rules {
    matches {
      match_type     = "clientip"
      match_operator = "invalid_operator"
      object_match_value {
        type  = "simple"
        value = ["a"]
      }
    }
    forward_settings {
    }
  }
}