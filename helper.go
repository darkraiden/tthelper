package tthelper

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
)

// Terratest represents an instance of a Terraform testing structure, with fields for SubscriptionID and TenantID.
// These fields are used to configure Terraform tests by passing values for the Azure subscription and tenant IDs.
type Terratest struct {
	SubscriptionID string
	TenantID       string
}

// New creates a new instance of Terratest by checking if the required environment variables,
// TEST_AZURE_SUBSCRIPTION_ID and TEST_AZURE_TENANT_ID, are set. If either of these are missing,
// the function will call t.Fatal with an error message.
// Otherwise, it will return a new instance of Terratest with the values of the environment variables
// set as the SubscriptionID and TenantID fields.
func New(t *testing.T) *Terratest {
	t.Helper()

	if getSubscriptionID() == "" || getTenantID() == "" {
		t.Fatal("both `TEST_AZURE_SUBSCRIPTION_ID` and `TEST_AZURE_TENANT_ID` environment variables must be set")
	}

	return &Terratest{
		SubscriptionID: getSubscriptionID(),
		TenantID:       getTenantID(),
	}
}

// TerraformOptions returns a pointer to a Terraform Options struct configured with the specified fixture folder, variables, and variable files.
// If the input `fixtureFolder` is an empty string, it defaults to "./fixture". If the input `vars` map is nil, it initializes an empty map.
// The values for the keys "subscription_id" and "tenant_id" are added to the `vars` map based on the values stored in the Terratest struct.
func (tt *Terratest) TerraformOptions(fixtureFolder string, vars map[string]interface{}, varFiles ...string) *terraform.Options {
	if fixtureFolder == "" {
		fixtureFolder = "./fixture"
	}
	if vars == nil {
		vars = make(map[string]interface{})
	}
	vars["subscription_id"] = tt.SubscriptionID
	vars["tenant_id"] = tt.TenantID

	return &terraform.Options{
		TerraformDir: fixtureFolder,
		VarFiles:     varFiles,
		Vars:         vars,
	}
}

func getSubscriptionID() string {
	return os.Getenv("TEST_AZURE_SUBSCRIPTION_ID")
}

func getTenantID() string {
	return os.Getenv("TEST_AZURE_TENANT_ID")
}
