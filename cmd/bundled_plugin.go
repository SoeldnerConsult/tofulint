package cmd

import (
	"fmt"

	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsult/tofulint-plugin-sdk/tflint"
	"github.com/SoeldnerConsult/tofulint-ruleset-opentofu/project"
	"github.com/SoeldnerConsult/tofulint-ruleset-opentofu/rules"
	"github.com/SoeldnerConsult/tofulint-ruleset-opentofu/terraform"
)

func (cli *CLI) actAsBundledPlugin() int {
	plugin.Serve(&plugin.ServeOpts{
		RuleSet: &terraform.RuleSet{
			BuiltinRuleSet: tflint.BuiltinRuleSet{
				Name:       "opentofu",
				Version:    fmt.Sprintf("%s-bundled", project.Version),
				Constraint: ">= 0.0.1",
			},
			PresetRules: rules.PresetRules,
		},
	})
	return ExitCodeOK
}
