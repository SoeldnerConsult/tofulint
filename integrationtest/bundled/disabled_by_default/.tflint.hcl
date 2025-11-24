config {
  disabled_by_default = true
}


plugin "opentofu" {
  enabled = true
  version = "0.0.8"
  source = "github.com/SoeldnerConsult/tofulint-ruleset-opentofu"
}

rule "terraform_unused_declarations" {
  enabled = true
}