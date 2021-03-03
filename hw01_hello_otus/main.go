package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	got := stringutil.Reverse("Hello, OTUS!")
	fmt.Print(got)
}
