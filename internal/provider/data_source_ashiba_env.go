package provider

import (
	"context"
	"os"
	"regexp"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceAshibaEnv() *schema.Resource {
	return &schema.Resource{
		// This description is used by the documentation generator and the language server.
		Description: "Sample data source in the Terraform provider scaffolding.",

		ReadContext: dataSourceScaffoldingRead,

		Schema: map[string]*schema.Schema{
			"environment": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"workspace": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"project": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeString,
				Computed:    true,
			},
			"is_workspace": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
			"is_pr": {
				// This description is used by the documentation generator and the language server.
				Description: "Sample attribute.",
				Type:        schema.TypeBool,
				Computed:    true,
			},
		},
	}
}

func dataSourceScaffoldingRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// use the meta value to retrieve your client from the provider configure method
	client := meta.(*apiClient)

	var diags diag.Diagnostics

	project := os.Getenv("CBR_PROJECT_KEY")
	workspace := os.Getenv("CBR_APP_WORKSPACE")
	environment := os.Getenv("CBR_APP_ENV")

	if err := d.Set("project", project); err != nil {
		// extend on description of the error
		return diag.FromErr(err)
	}

	if err := d.Set("workspace", workspace); err != nil {
		// extend on description of the error
		return diag.FromErr(err)
	}

	if err := d.Set("environment", environment); err != nil {
		// extend on description of the error
		return diag.FromErr(err)
	}

	d.Set("is_workspace", false)
	if workspace != "default" {
		d.Set("is_workspace", true)
	}

	matched, err := regexp.Match(client.prPattern, []byte(workspace))

	if err != nil {
		return diag.FromErr(err)
	}

	d.Set("is_pr", false)
	if matched {
		d.Set("is_pr", true)
	}

	// client.prPattern

	// ignore for time being
	// idFromAPI := "my-id"
	// d.SetId(idFromAPI)

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
