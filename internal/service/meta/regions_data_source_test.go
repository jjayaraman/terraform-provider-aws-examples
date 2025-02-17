package meta_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	tfmeta "github.com/hashicorp/terraform-provider-aws/internal/service/meta"
)

func TestAccMetaRegionsDataSource_basic(t *testing.T) {
	dataSourceName := "data.aws_regions.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, tfmeta.PseudoServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRegionsDataSourceConfig_empty(),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "names.#", "0"),
				),
			},
		},
	})
}

func TestAccMetaRegionsDataSource_filter(t *testing.T) {
	dataSourceName := "data.aws_regions.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, tfmeta.PseudoServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRegionsDataSourceConfig_optInStatusFilter("opt-in-not-required"),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "names.#", "0"),
				),
			},
		},
	})
}

func TestAccMetaRegionsDataSource_allRegions(t *testing.T) {
	dataSourceName := "data.aws_regions.test"

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.PreCheck(t) },
		ErrorCheck:               acctest.ErrorCheck(t, tfmeta.PseudoServiceID),
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccRegionsDataSourceConfig_allRegions(),
				Check: resource.ComposeTestCheckFunc(
					acctest.CheckResourceAttrGreaterThanValue(dataSourceName, "names.#", "0"),
				),
			},
		},
	})
}

func testAccRegionsDataSourceConfig_empty() string {
	return `
data "aws_regions" "test" {}
`
}

func testAccRegionsDataSourceConfig_allRegions() string {
	return `
data "aws_regions" "test" {
  all_regions = "true"
}
`
}

func testAccRegionsDataSourceConfig_optInStatusFilter(filter string) string {
	return fmt.Sprintf(`
data "aws_regions" "test" {
  filter {
    name   = "opt-in-status"
    values = [%[1]q]
  }
}
`, filter)
}
