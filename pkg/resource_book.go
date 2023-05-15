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
		UpdateContext: resourceBookUpdate,
		DeleteContext: resourceBookDelete,
		Schema: map[string]*schema.Schema{
			"book": &schema.Schema{
				Type:     schema.TypeList,
				MaxItems: 1,
				Required: true,
				Elem: schema.Resource{
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
				},
			},
		},
	}
}

func resourceBookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

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

func resourceBookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func resourceBookUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}
func resourceBookDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	return nil
}

func setItems(d *schema.ResourceData, book Book) {
	d.Set("id", book.ID)
	d.Set("author", book.Author)
	d.Set("title", book.Author)
	d.Set("price", book.Price)
}
