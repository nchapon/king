package main

import (
	"fmt"
	"king/kong"
)

// main ...
func main() {
	fmt.Printf("Request api \n")

	kg := kong.KongApiInfoer{}
	result, err := kong.GetApiInfo(kg, "catalog")

	fmt.Println("Response \n", result)
	if err != nil {
		fmt.Println(err)
	}

}
