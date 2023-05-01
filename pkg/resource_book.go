package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBookCreate,
		Schema:        map[string]*schema.Schema{
			"book": &schema.Schema{
				Type: schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							Required: true,
						},
						"title": &schema.Schema{
							Type: schema.TypeMap,
						},
						"author": &schema.Schema{
							Type: schema.TypeString,
						},
						"price": &schema.Schema{
							Type: schema.TypeInt,
						},
					},
				},
			},
		},
	}
}

func resourceBookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
