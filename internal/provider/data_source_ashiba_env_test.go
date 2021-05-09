package provider

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// resource.TestMatchResourceAttr(
// 	"data.scaffolding_data_source.foo", "sample_attribute", regexp.MustCompile("^ba")),

func TestAccDataSourceAshiba(t *testing.T) {
	os.Setenv("CBR_PROJECT_KEY", "my_project_key")
	os.Setenv("CBR_APP_WORKSPACE", "my_workspace")
	os.Setenv("CBR_APP_ENV", "play")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAshiba,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "project", "my_project_key"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "workspace", "my_workspace"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "environment", "play"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "is_workspace", "true"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "is_pr", "false"),
				),
			},
		},
	})
}

func TestAccDataSourceAshibaPr(t *testing.T) {
	os.Setenv("CBR_PROJECT_KEY", "my_project_key")
	os.Setenv("CBR_APP_WORKSPACE", "pr123")
	os.Setenv("CBR_APP_ENV", "play")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAshiba,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "project", "my_project_key"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "workspace", "pr123"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "environment", "play"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "is_workspace", "true"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "is_pr", "true"),
				),
			},
		},
	})
}

func TestAccDataSourceAshibaDefaultWorkspace(t *testing.T) {
	os.Setenv("CBR_PROJECT_KEY", "my_project_key")
	os.Setenv("CBR_APP_WORKSPACE", "default")
	os.Setenv("CBR_APP_ENV", "play")

	resource.UnitTest(t, resource.TestCase{
		PreCheck:          func() { testAccPreCheck(t) },
		ProviderFactories: providerFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceAshiba,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "project", "my_project_key"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "workspace", "default"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "environment", "play"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "is_workspace", "false"),
					resource.TestCheckResourceAttr(
						"data.ashiba_env.foo", "is_pr", "false"),
				),
			},
		},
	})
}

const testAccDataSourceAshiba = `
data "ashiba_env" "foo" {
}
`
