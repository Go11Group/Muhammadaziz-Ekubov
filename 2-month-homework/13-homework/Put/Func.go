package Put

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func Put_1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	jsonStr := []byte(`{"id":1,"title":"foo","body":"bar","userId":1}`)
	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonStr))
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

func Put_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	jsonStr := []byte(`{"title":"foo updated","body":"bar updated","userId":1}`)
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%s", id), bytes.NewBuffer(jsonStr))
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

func Put_3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	jsonStr := []byte(`{"id":1,"title":"foo","body":"bar","userId":1}`)
	req, err := http.NewRequest("PUT", "https://jsonplaceholder.typicode.com/posts/1", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

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
