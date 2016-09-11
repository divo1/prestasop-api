package main

import (
	"fmt"
	"github.com/divo1/psapi"
)

func main() {
	s := psapi.NewService("https://cattaneo.pl/api", "5QINNHUJE64QS4T6AB2G7GPQ3E9KA8AK")

	object := s.GetProducts()
	categories := s.GetCategories()

	fmt.Printf("\n\n%#v", object)
	fmt.Printf("\n\n%#v", categories)
}