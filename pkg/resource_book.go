package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBookCreate,
		ReadContext:   resourceBookRead,
		Schema: map[string]*schema.Schema{
			"book": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"title": &schema.Schema{
							Type: schema.TypeString,
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
	c := m.(Client)

	var diags diag.Diagnostics
	id := d.Get("id").(int)
	title := d.Get("title").(string)
	author := d.Get("author").(string)
	price := d.Get("price").(int)

	b := Book{
		ID:     id,
		Title:  title,
		Author: author,
		Price:  price,
	}
	err := c.CreateBook(b)
	if err != nil {
		return diag.FromErr(err)
	}
	resourceBookRead(ctx, d, m)
	return diags
}

