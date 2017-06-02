package model

import (
	"encoding/json"
)

type Api struct {
	Name        string `json:"name"`
	StripUri    bool   `json:"strip_uri"`
	UpstreamUrl string `json:"upstream_url"`
}

// NewApiFromJSON ...
func NewApiFromJSON(str string) Api {
	res := Api{}
	json.Unmarshal([]byte(str), &res)
	return res
}
