package Post

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Post_1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	data := url.Values{}
	data.Set("name", "John")
	data.Set("age", "30")

	client := &http.Client{}
	req, err := http.NewRequest("POST", "https://jsonplaceholder.typicode.com/posts", strings.NewReader(data.Encode()))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func Post_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	jsonStr := []byte(`{"title":"foo","body":"bar","userId":1}`)
	req, err := http.NewRequest("POST", "localhost:8080/posts", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	_, err = w.Write(body)
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
