provider "akamai" {
  edgerc        = "../../test/edgerc"
  cache_enabled = false
}

resource "akamai_botman_challenge_injection_rules" "test" {
  config_id                 = 43253
  challenge_injection_rules = <<-EOF
{
  "testKey": "updated_testValue3"
}
EOF
}