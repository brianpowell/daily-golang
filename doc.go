package daily

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	query "github.com/google/go-querystring/query"
)

// Helper Functions
func (c *Client) makeRequest(method string, endpoint string, body interface{}, out interface{}) error {

	var queryString string
	// Handle the Query String
	if body != nil {
		v, _ := query.Values(body)
		queryString = "?" + v.Encode()
	}

	// Build the new request
	req, err := http.NewRequest(method, API+endpoint+queryString, nil)
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	// Headers
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)

	// Make actual request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return err
	}

	// Parse out the response to JSON
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err
	}
	fmt.Printf("client: response body: %s\n", resBody)

	// Marshal out the response
	return json.Unmarshal(resBody, &out)
}

// Helper Functions
func (c *Client) makeBodyRequest(method string, endpoint string, body interface{}, out interface{}) error {

	// Handle the Body information
	jsonData, _ := json.Marshal(body)

	// Build the new request
	req, err := http.NewRequest(method, API+endpoint, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("client: could not create request: %s\n", err)
		return err
	}

	// Headers
	fmt.Println("Token", c.Config.Token)
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")

	// Make actual request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Printf("client: error making http request: %s\n", err)
		return err
	}

	// Parse out the response to JSON
	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		return err
	}
	fmt.Printf("client: response body: %s\n", resBody)

	// Marshal out the response
	return json.Unmarshal(resBody, &out)
}
