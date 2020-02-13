package gocurl

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string, headers map[string]string)(code int, response string) {

	//create a new request
	r, err := http.NewRequest("GET", url, nil)



	//if the new request didn't b0rk
	if err == nil {
		//add any supplied headers
		for k, v := range headers {
			fmt.Printf("Key: %s Value: %s\n", k, v)
			r.Header.Add(k,v)
		}
		//get a response [jazz hands]

		res, err := http.DefaultClient.Do(r)


		if err == nil {
			//defer the closing of the request until safe innit
			defer res.Body.Close()
			// get the response body
			code = res.StatusCode
			response, err := ioutil.ReadAll(res.Body)

			if err == nil {

				fmt.Println(string(response))

			} else {
				fmt.Printf("[ERROR] reading the response body : %v\n", err.Error())
			}

		} else {
			fmt.Printf("[ERROR] Requesting url %v : %v\n", url, err.Error())
		}

	} else {

		fmt.Printf("[ERROR] creating new http request : %v\n", err.Error())
	}

	return
}
