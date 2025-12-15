# TofuLint
[![build](https://github.com/SoeldnerConsult/tofulint/actions/workflows/build.yml/badge.svg)](https://github.com/SoeldnerConsult/tofulint/actions/workflows/build.yml)
[![GitHub release](https://img.shields.io/github/release/SoeldnerConsult/tofulint.svg)](https://github.com/SoeldnerConsult/tofulint/releases/latest)
[![Opentofu Compatibility](https://img.shields.io/badge/opentofu-%3E%3D%201.0-blue)](docs/user-guide/compatibility.md)
[![License: MPL 2.0](https://img.shields.io/badge/License-MPL%202.0-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/SoeldnerConsult/tofulint)](https://goreportcard.com/report/github.com/SoeldnerConsult/tofulint)
[![Homebrew](https://img.shields.io/badge/dynamic/json.svg?url=https://formulae.brew.sh/api/formula/tflint.json&query=$.versions.stable&label=homebrew)](https://formulae.brew.sh/formula/tflint)

A **pluggable** [OpenTofu](https://opentofu.org/) linter inspired by TFLint.

## Disclaimer
`TofuLint` is an **experimental** fork of `TFLint` that replaces Terraform internals with **OpenTofu**.  
It is **highly experimental** and **not production-ready**. Use at your own risk.

## Features

TofuLint is a modular framework where each feature is provided via plugins. Key features include:

- Detect potential errors (e.g., invalid instance types) for major cloud providers: AWS, Azure, GCP.  
- Warn about deprecated syntax and unused declarations.  
- Enforce best practices and naming conventions.  

### TofuLint specific features
In comparison to [TFLint](https://github.com/terraform-linters/tflint/) TofuLint introduces some enhancements. 
Most of the enhancements evolve around opentofu specific filetypes.
- support for `.tf` as well as `.tofu` files
- support for `.tflint.hcl` as well as `tofulint.hcl` module config files.
- support for the equivalent `.json` files
- the whole terraform fork in the core of tflint was replaced by an fork of opentofu.

While this fork now is up to date with the basic tflint functionality, some more opentofu specific features can be introduced:
- **features evolving around state-encryption are not supported yet**: Linting for the `encryption` block in `terraform` configurations, including validation of `key_provider` consistency, method syntax (e.g. AES-GCM), and policy enforcement for prod environments.
- **linting for dynamic provider-defined functions**: Static analysis of provider-defined function calls, checking argument types/signatures against provider schemas and detecting usage of unknown values where only constants are allowed.
- **linting for removed blocks**: Validation of `removed` blocks for consistency with resource/module definitions, warnings on non-matching or obsolete removals, and policy rules to prevent accidental deprovisioning in shared modules.
- **linting for loopable import blocks**: Checks on `for_each`/`count` expressions in `import` blocks for valid mappings against data sources, resource ID format compatibility, and avoidance of duplicate imports with manual state entries.


## Pipeline Integration

For two distinct examples of integrating `tofulint` into a CI/CD workflow - one that **fails the pipeline** on errors and one that **always succeeds** but **reports issues** to the GitHub Security Status page - please check out the dedicated **[demo repository](https://github.com/SoeldnerConsult/tofulint-test-repo)**.

## Installation
Currently, only one installation method is available:

### Bash (Linux)
```bash
curl -s https://raw.githubusercontent.com/SoeldnerConsult/tofulint/master/install_linux.sh | bash
````

### Docker
A Docker-based installation will be available in a future release.

## Getting Started
TofuLint comes bundled with a [Terraform language ruleset](https://github.com/SoeldnerConsult/tofulint-ruleset-opentofu), enabling recommended rules by default.

### Enabling the Opentofu Plugin
Declare the plugin block in your `.tflint.hcl` or `.tofulint.hcl`:

```hcl
plugin "opentofu" {
  enabled = true
  version = "0.0.9"
  source = "github.com/SoeldnerConsult/tofulint-ruleset-opentofu"
}
```
> Even though tofulint currently comes with the opentofu plugin pre-packaged, it is still necessary to enable th plugin manually with the given plugin source. This is due to a bug in tofulint source code.

More details: [TFLint Terraform Ruleset Configuration](https://github.com/SoeldnerConsult/tofulint-ruleset-opentofu/blob/main/docs/configuration.md)


## Known Issues

### Broken Pre-bundled Installation

**Status:** Investigation Ongoing

**Description**  
Currently, the automatic installation of the `tofulint-ruleset-opentofu` plugin fails during the standard `tofulint` setup. While this ruleset is intended to be pre-bundled, the application currently fails to initialize it upon startup.

**Symptoms**  
When running `tofulint` immediately after installation, the application fails to initialize the plugin and suggests running the initialization command manually:

```bash
$ tofulint
'Failed to initialize plugins; Plugin "opentofu" not found. Did you run "tofulint --init"?'
```

**Technical Details**  
Attempting to resolve the issue by running `tofulint --init` reveals a malformed URL error during the fetch process. The installer appears to be constructing a GitHub API request without the necessary repository or host information:

```bash
$ tofulint --init
Installing "opentofu" plugin...
Failed to install a plugin; Failed to fetch GitHub releases: Get "https:///api/v3/repos///releases/tags/v0.0.7": http: no Host in request URL
```

> **Note:** The specific code responsible for generating the malformed URL (`http: no Host in request URL`) has not yet been identified.

**Fix**  
Right now the only available fix is to manually install the opentofu plugin. Therefore define an `.tflint.hcl` or `.tofulint.hcl` file and manually define the necessary opentofu plugin:
``` hcl
plugin "opentofu" {
  enabled = true
  version = "0.0.9"
  source = "github.com/SoeldnerConsult/tofulint-ruleset-opentofu"
}
```


### Cloud Provider Plugins

If you use a cloud provider, install the corresponding plugin:

* [AWS](https://github.com/SoeldnerConsult/tofulint-ruleset-aws)
* [GCP](https://github.com/SoeldnerConsult/tofulint-ruleset-google)

Other plugins can be added via `.tflint.hcl` and installed with:

```bash
tofulint --init
```

### Write your own plugins
To write your own plugins the syntax and workflow is the same as in `tflint`. This also enables to modify existing `tflint` plugins to work with tofulint. Therefor fork the wanted plugin, enable all github-actions via th github web interface, and replace all `terraform-linters/tflint-plugin-sdk` with `SoeldnerConsult/tofulint-plugin-sdk`. Also replace all `terraform-linters/tofulint-plugin-<pluginName>` occurances with the new repository where the plugin is located now, as well as the version in `project/main.go`. 
Commit and push the changes, create a new tag or empty release with the Syntax `vX.X.X` (best wopuld be to use the same Version as above) and the rest will be handled by the corresponsing actions.

Use the plugin by including the following in your `.tofulint.hcl` / `.tflint.hcl` file:
```hcl
plugin "<pluginName>" {
  enabled = true
  version = "version (e.g. 0.0.7)"
  source = "github.com/<owner>/tofulint-ruleset-<pluginName>"
}
```
Install with 
```bash
tofulint --init
```

### Example Plugin Configuration

```hcl
plugin "foo" {
  enabled = true
  version = "0.1.0"
  source  = "github.com/org/tflint-ruleset-foo"

  signing_key = <<-KEY
  -----BEGIN PGP PUBLIC KEY BLOCK-----
  ...
  KEY
}
```

For custom rules, create your own plugin or use Rego policies:

* [Writing Plugins](docs/developer-guide/plugins.md)


## Usage

By default, TofuLint inspects files in the current directory. Example options:

```bash
$ tofulint --help
Usage:
  tofulint --chdir=DIR/--recursive [OPTIONS]

Application Options:
  -v, --version                         Print TofuLint version
      --init                            Install plugins
      --langserver                      Start language server
  -f, --format=[default|json|checkstyle|junit|compact|sarif] Output format
  -c, --config=FILE                     Config file name (default: .tflint.hcl)
      --ignore-module=SOURCE            Ignore module sources
      --enable-rule=RULE_NAME           Enable rules from the command line
      --disable-rule=RULE_NAME          Disable rules from the command line
      --only=RULE_NAME                  Enable only this rule
      --enable-plugin=PLUGIN_NAME       Enable plugins from the command line
      --var-file=FILE                    Terraform variable file
      --var='foo=bar'                    Set a Terraform variable
      --call-module-type=[all|local|none] Types of module to call (default: local)
      --chdir=DIR                        Change working directory
      --recursive                        Run recursively in subdirectories
      --filter=FILE                       Filter issues by file names/globs
      --force                             Return zero exit code even if issues found
      --minimum-failure-severity=[error|warning|notice] Minimum severity for non-zero exit
      --color                             Enable colorized output
      --no-color                          Disable colorized output
      --fix                               Automatically fix issues
      --no-parallel-runners               Disable parallelism

Help Options:
  -h, --help                             Show this help message
```

## 


See [User Guide](docs/user-guide) for more details.

## Debugging

Enable detailed logs using the `TFLINT_LOG` environment variable:

```bash
$ TFLINT_LOG=debug tofulint
```
## Developing

See [Developer Guide](docs/developer-guide) for instructions on contributing and building plugins.

## Security

For reporting security issues, refer to our [security policy](SECURITY.md).

