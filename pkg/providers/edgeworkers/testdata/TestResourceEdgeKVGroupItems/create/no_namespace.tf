provider "akamai" {
  edgerc = "../../test/edgerc"
}

resource "akamai_edgekv_group_items" "test" {
  network    = "staging"
  group_name = "1234"
  items = {
    key1 = "value1"
    key2 = "value2"
  }
}

