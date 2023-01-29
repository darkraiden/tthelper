# tthelper

A helper library for Gruntwork's Terratest, specifically for Terraform unit-tests on Azure.

## Installation

go get github.com/darkraiden/tthelper

## Usage

Here is an example of how to use the library in a test:

```go
package main

import (
    "testing"

    "github.com/darkraiden/tthelper"
    "github.com/gruntwork-io/terratest/modules/terraform"
)

func TestTerraform(t *testing.T) {
    tt := tthelper.New(t)
    terraformOptions := tt.TerraformOptions("./fixture", "./variables.tfvars")

    defer terraform.Destroy(t, terraformOptions)
    terraform.InitAndApply(t, terraformOptions)
}
```

## Environment Variables

The library requires two environment variables to be set:

- `TEST_AZURE_SUBSCRIPTION_ID`
- `TEST_AZURE_TENANT_ID`

## API Reference

### func New

```go
func New(t *testing.T) *Terratest
```

New creates a new instance of Terratest by checking if the required environment variables,
`TEST_AZURE_SUBSCRIPTION_ID` and `TEST_AZURE_TENANT_ID`, are set. If either of these are missing,
the function will call t.Fatal with an error message. Otherwise, it will return a new instance of Terratest with the values of the environment variables
set as the SubscriptionID and TenantID fields.

### func (tt *Terratest) TerraformOptions

```go
func (tt *Terratest) TerraformOptions(fixtureFolder string, varFiles ...string) *terraform.Options
```

TerraformOptions returns Terraform options with the fields for `TerraformDir`, `VarFiles`, and `Vars` set.
`TerraformDir` is set to the given `fixtureFolder`, `VarFiles` is set to the given `varFiles` and `Vars` is set
to a map containing the `subscription_id` and `tenant_id` fields from the Terratest instance.

## Author

[Davide Di Mauro](https://github.com/darkraiden)

## License

MIT
