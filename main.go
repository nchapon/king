package main

import (
	"fmt"
	"king/kong"
)

// main ...
func main() {
	fmt.Printf("List all apis \n")

	status, output := kong.CallRequest("http://localhost:8001/apis")

	fmt.Printf("Code status %v \n", status)
	fmt.Printf("Body %s \n", output)

}
