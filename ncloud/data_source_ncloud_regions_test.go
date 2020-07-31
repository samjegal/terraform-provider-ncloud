package ncloud

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccDataSourceNcloudRegionsClassic(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudRegionsConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID("data.ncloud_regions.regions"),
				),
			},
		},
	})
}

func TestAccDataSourceNcloudRegionsVPC(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceNcloudRegionsVPCConfig,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDataSourceID("data.ncloud_regions.regions"),
				),
			},
		},
	})
}

var testAccDataSourceNcloudRegionsConfig = `
data "ncloud_regions" "regions" {}
`

var testAccDataSourceNcloudRegionsVPCConfig = `
data "ncloud_regions" "regions" {
	platform_type = "vpc"
}
`
