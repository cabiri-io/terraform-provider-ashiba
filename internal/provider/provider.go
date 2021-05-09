package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func init() {
	// Set descriptions to support markdown syntax, this will be used in document generation
	// and the language server.
	schema.DescriptionKind = schema.StringMarkdown

	// Customize the content of descriptions when output. For example you can add defaults on
	// to the exported descriptions if present.
	// schema.SchemaDescriptionBuilder = func(s *schema.Schema) string {
	// 	desc := s.Description
	// 	if s.Default != nil {
	// 		desc += fmt.Sprintf(" Defaults to `%v`.", s.Default)
	// 	}
	// 	return strings.TrimSpace(desc)
	// }
}

func New(version string) func() *schema.Provider {
	return func() *schema.Provider {
		p := &schema.Provider{
			Schema: map[string]*schema.Schema{
				"pr_pattern": {
					Type:     schema.TypeString,
					Optional: true,
					Default:  "pr[0-9]+",
				},
				"alb_default_priority": {
					Type:     schema.TypeInt,
					Optional: true,
					Default:  49000,
				},
			},
			DataSourcesMap: map[string]*schema.Resource{
				"ashiba_env": dataSourceAshibaEnv(),
			},
			ResourcesMap: map[string]*schema.Resource{
				"ashiba_resource": resourceScaffolding(),
			},
		}

		p.ConfigureContextFunc = configure(version, p)

		return p
	}
}

// What provider has to provide:
// - workspace verification
// - ability to use env variable as datasource
// - ability to determine if it is workspace or pr
// - service name creation

type apiClient struct {
	// Add whatever fields, client or connection info, etc. here
	// you would need to setup to communicate with the upstream
	// API.
	prPattern          string
	albDefaultPriority int
}

func configure(version string, p *schema.Provider) func(context.Context, *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return func(c context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
		// Setup a User-Agent for your API client (replace the provider name for yours):
		// userAgent := p.UserAgent("terraform-provider-scaffolding", version)
		// TODO: myClient.UserAgent = userAgent
		prPattern := d.Get("pr_pattern").(string)
		albDefaultPriority := d.Get("alb_default_priority").(int)

		return &apiClient{
			prPattern:          prPattern,
			albDefaultPriority: albDefaultPriority,
		}, nil
	}
}
