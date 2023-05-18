package pkg

import (
	"context"
	"strconv"

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
			"bookid": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"title": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"author": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"price": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			// "book": &schema.Schema{
			// 	Type:     schema.TypeList,
			// 	MaxItems: 1,
			// 	Elem: &schema.Resource{
			// 		Schema: map[string]*schema.Schema{
			// 			"bookID": &schema.Schema{
			// 				Type:     schema.TypeInt,
			// 				Required: true,
			// 			},
			// 			"title": &schema.Schema{
			// 				Type: schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"author": &schema.Schema{
			// 				Type: schema.TypeString,
			// 				Required: true,
			// 			},
			// 			"price": &schema.Schema{
			// 				Type: schema.TypeInt,
			// 				Required: true,
			// 			},
			// 		},
			// 	},
			// },
		},
	}
}

func resourceBookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics
	bookID := d.Get("bookid").(int)
	title := d.Get("title").(string)
	author := d.Get("author").(string)
	price := d.Get("price").(int)

	b := Book{
		BookID: bookID,
		Title:  title,
		Author: author,
		Price:  price,
	}
	err := c.CreateBook(b)
	if err != nil {
		return diag.Errorf("something has happened: %s", err)
	}
	d.SetId(strconv.Itoa(bookID))
	setItems(d, b)
	resourceBookRead(ctx, d, m)
	return diags
}

func resourceBookRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*Client)

	var diags diag.Diagnostics

	bookID, err := strconv.Atoi(d.Id())
	if err != nil {
		diag.FromErr(err)
	}
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
	d.Set("bookid", book.BookID)
	d.Set("author", book.Author)
	d.Set("title", book.Title)
	d.Set("price", book.Price)
}
