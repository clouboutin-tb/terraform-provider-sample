package provider

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSampleResource(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + testAccExampleResourceConfig("test example"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify the lonely attribute
					resource.TestCheckResourceAttr("sample_sample.test", "string", "test example"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("sample_sample.test", "last_updated"),
					resource.TestCheckResourceAttrSet("sample_sample.test", "result"),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig + testAccExampleResourceConfig("update example"),
				Check: resource.ComposeAggregateTestCheckFunc(
					// Verify the lonely attribute
					resource.TestCheckResourceAttr("sample_sample.test", "string", "update example"),
					// Verify dynamic values have any value set in the state.
					resource.TestCheckResourceAttrSet("sample_sample.test", "last_updated"),
					resource.TestCheckResourceAttrSet("sample_sample.test", "result"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccExampleResourceConfig(configurableAttribute string) string {
	return fmt.Sprintf(`
resource "sample_sample" "test" {
  string = %[1]q
}
`, configurableAttribute)
}
