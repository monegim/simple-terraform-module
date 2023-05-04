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
				DefaultFunc: schema.EnvDefaultFunc("SIMPLE_GIN_HOST", "localhost:8080"),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"simple_gin_book": resourceBook(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"simple_gin_book": dataSourceBook(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	// var host *string
	return nil, nil
}
