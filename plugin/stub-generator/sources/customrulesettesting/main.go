package main

import (
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/SoeldnerConsult/tofulint/plugin/stub-generator/sources/customrulesettesting/custom"
	"github.com/SoeldnerConsult/tofulint/plugin/stub-generator/sources/customrulesettesting/rules"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &custom.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:    "customrulesettesting",
				Version: "0.1.0",
				Rules: []tflint.Rule{
					rules.NewAwsInstanceExampleTypeRule(),
				},
			},
		},
	})
}
