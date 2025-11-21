package cmd

import (
	"fmt"

	"github.com/SoeldnerConsult/tofulint-plugin-sdk/plugin"
	"github.com/SoeldnerConsulterConsult/tofulint-plugin-sdk/tflint"
	"github.com/SoeldnerConsulterConsult/tofulint-ruleset-opentofu/project"
	"github.com/SoeldnerConsulterConsult/tofulint-ruleset-opentofu/rules"
	"github.com/SoeldnerConsulterConsult/tofulint-ruleset-opentofu/terraform"
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
