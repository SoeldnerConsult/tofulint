package main

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:       "incompatiblehost",
			Version:    "0.1.0",
			Constraint: ">= 1.0",
		},
	})
}
