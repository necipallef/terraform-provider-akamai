provider "akamai" {
  edgerc = "../../test/edgerc"
}

resource "akamai_clientlist_activation" "activation_ASN_LIST_1" {
  list_id                 = "12_AB"
  version                 = 2
  network                 = "STAGING"
  comments                = "Activation Comments"
  notification_recipients = ["user@example.com"]
  siebel_ticket_id        = "ABC-12345-UPDATED"
}
