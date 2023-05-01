package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBook() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBookCreate,
	}
}

func resourceBookCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	
}
