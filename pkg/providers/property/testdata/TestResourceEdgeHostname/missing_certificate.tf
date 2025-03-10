provider "akamai" {
  edgerc = "../../test/edgerc"
}

resource "akamai_edge_hostname" "edgehostname" {
  contract_id   = "ctr_2"
  group_id      = "grp_2"
  product_id    = "prd_2"
  edge_hostname = "test.edgekey.net"
  ip_behavior   = "IPV6_PERFORMANCE"
}
