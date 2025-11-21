package main

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsulterConsult/tofulint-plugin-sdk/tflint"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &tflint.BuiltinRuleSet{
			Name:    "foo",
			Version: "0.1.0",
			Rules:   []tflint.Rule{},
		},
	})
}
