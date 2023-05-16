terraform {
  required_providers {
    simplegin = {
      version = "0.0.1"
      source  = "hashicorp.com/edu/simplegin"
    }
  }
}
provider "simplegin" {
  
}
resource "simplegin_book" "new" {
  book {
  id     = 1
  title  = "Atomic Habits"
  author = "Not Sure"
  price  = 30000 
  }
}
