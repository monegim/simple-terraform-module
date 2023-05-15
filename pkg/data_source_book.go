package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceBookRead,
		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
	}
}

func dataSourceBookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	bookID := d.Get("id").(int)

	book, err := c.GetBook(bookID)
	if err != nil {
		return diag.FromErr(err)
	}
	setItems(d, *book)
	return diags
}
