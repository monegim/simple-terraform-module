package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("simplegin_HOST", "localhost:8080"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"simplegin_book": resourceBook(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"simplegin_book": dataSourceBook(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var host *string

	hval, ok := d.GetOk("host")
	if ok {
		tempHost := hval.(string)
		host = &tempHost
	}
	var diags diag.Diagnostics
	c, err := NewClient(host)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create client",
		})
		return nil, diags
	}
	return c, diags
}
