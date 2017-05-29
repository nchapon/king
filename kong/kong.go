package kong

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ApiInfo struct {
	Name string `json:"name"`
}

type ApiInfoer interface {
	GetApiInfo(string) (string, error)
}

type KongApiInfoer struct {
}

// CallRequest ...
func (kg KongApiInfoer) GetApiInfo(name string) (string, error) {

	apiUrl := fmt.Sprintf("http://localhost:8001/apis/%s", name)

	response, err := http.Get(apiUrl)
	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return "", err
	}

	fmt.Printf("%s", body)

	apiInfo := new(ApiInfo)
	json.Unmarshal(body, &apiInfo)

	return apiInfo.Name, nil
}

// GetApiInfo ...
func GetApiInfo(ai ApiInfoer, name string) (string, error) {

	apiName, err := ai.GetApiInfo(name)

	if err != nil {
		return "", fmt.Errorf("Error query api : %s", err)
	}
	return apiName, nil

}
