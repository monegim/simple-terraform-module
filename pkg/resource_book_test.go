package pkg

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestBookBasic(t *testing.T) {
	// id := 1
	// title := "Atomic Habits"
	// author := "Not Sure"
	// price := 30000

	resource.Test(t, resource.TestCase{
		PreCheck: func() {},
	})
}

func testAccCheckBookConfigBasic(bookId int, title, author string, price int) string {
	return fmt.Sprintf(`
	resource "simple_gin_book" "atomic" {
		id = %d
		title = %s
		author = %s
		price = %d
	}
	`, bookId, title, author, price)
}
