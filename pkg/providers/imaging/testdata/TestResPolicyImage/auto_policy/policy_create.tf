provider "akamai" {
  edgerc = "../../test/edgerc"
}

resource "akamai_imaging_policy_image" "policy" {
  policy_id    = ".auto"
  contract_id  = "test_contract"
  policyset_id = "test_policy_set"
  json = jsonencode({
    "breakpoints" : {
      "widths" : [
        320,
        640,
        1024,
        2048,
        5000
      ]
    },
    "output" : {
      "perceptualQuality" : "mediumHigh"
    },
    "transformations" : [
      {
        "colors" : 2,
        "transformation" : "MaxColors"
      }
    ]
  })
}