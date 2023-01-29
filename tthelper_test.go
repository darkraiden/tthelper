package tthelper

import (
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Parallel()

	setTestEnvVars()
	tt := New(t)

	assert.NotEmpty(t, tt.SubscriptionID, "something went wrong; tt.SubscriptionID isn't supposed to be empty")
	assert.NotEmpty(t, tt.TenantID, "something went wrong; tt.TenantID isn't supposed to be empty")
}

func TestTerraformConfig(t *testing.T) {
	tests := []struct {
		fixtureFolder struct {
			actual   string
			expected string
		}
		varFiles struct {
			actual        []string
			expectedCount int
		}
	}{
		{
			fixtureFolder: struct{ actual, expected string }{
				actual:   "somewhereTerraformBelongs",
				expected: "somewhereTerraformBelongs",
			},
			varFiles: struct {
				actual        []string
				expectedCount int
			}{
				actual: []string{
					"file1",
					"file2",
				},
				expectedCount: 2,
			},
		},
		{
			fixtureFolder: struct{ actual, expected string }{
				actual:   "",
				expected: "./fixture",
			},
			varFiles: struct {
				actual        []string
				expectedCount int
			}{
				expectedCount: 0,
			},
		},
	}

	for _, test := range tests {
		test := test
		t.Run("TestTerraformConfig", func(t *testing.T) {
			t.Parallel()

			setTestEnvVars()
			tt := New(t)

			terraformOptions := tt.TerraformOptions(test.fixtureFolder.actual, test.varFiles.actual...)

			assert.Equal(t, test.fixtureFolder.expected, terraformOptions.TerraformDir, fmt.Sprintf("fixture folder values don't match. Expected %s, got %s", test.fixtureFolder.expected, terraformOptions.TerraformDir))
			assert.Equal(t, test.varFiles.expectedCount, len(terraformOptions.VarFiles), fmt.Sprintf("unexpected number of VarFiles. Expected %d, got %d", test.varFiles.expectedCount, len(terraformOptions.VarFiles)))
		})
	}
}

func setTestEnvVars() {
	os.Setenv("TEST_AZURE_SUBSCRIPTION_ID", uuid.New().String())
	os.Setenv("TEST_AZURE_TENANT_ID", uuid.New().String())
}
