package kong

import (
	"io/ioutil"
	"log"
	"net/http"
)

// CallRequest ...
func CallRequest(request string) (int, []byte) {

	res, err := http.Get(request)
	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return res.StatusCode, body
}
