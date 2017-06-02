package model

import (
	"testing"
)

func TestNewApi(t *testing.T) {

	api := NewApiFromJSON(`{
			  "name":"catalog"
			   }`)

	expectedName := "catalog"

	name := api.Name
	if name != expectedName {
		t.Fatalf("Expected %s but got %s", expectedName, name)
	}

}
