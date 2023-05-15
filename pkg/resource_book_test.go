package pkg

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestBookBasic(t *testing.T) {
	id := 1
	title := "Atomic Habits"
	author := "Not Sure"
	price := 30000

	resource.Test(t, resource.TestCase{
		PreCheck:     func() {},
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckBookDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckBookConfigBasic(id, title, author, price),
				Check:  resource.ComposeTestCheckFunc(
					testAccCheckBookExists("simple_gin_book.atomic"),
				),
			},
		},
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

func testAccCheckBookDestroy(s *terraform.State) error {
	return nil
}

func testAccCheckBookExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		res, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if res.Primary.ID == "" {
			return fmt.Errorf("No bookID set")
		}
		return nil
	}
}
