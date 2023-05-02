package pkg

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceBook(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	return nil, nil
}
