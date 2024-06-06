package Delete

import (
	"fmt"
	"net/http"
	"strings"
)

func Delete_1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	req, err := http.NewRequest("DELETE", "localhost:8080/delete", nil)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer resp.Body.Close()

	_, err = w.Write([]byte("Status: " + resp.Status))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func Delete_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	parts := strings.Split(r.URL.Path, "/")
	id := parts[len(parts)-1]

	req, err := http.NewRequest("DELETE", fmt.Sprintf("localhost:8080/delete/%s", id), nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer resp.Body.Close()

	_, err = w.Write([]byte("Status: " + resp.Status))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func Delete_3(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	req, err := http.NewRequest("DELETE", "https://jsonplaceholder.typicode.com/posts/1", nil)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	req.Header.Set("Authorization", "Bearer YOUR_TOKEN")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	defer resp.Body.Close()

	_, err = w.Write([]byte("Status: " + resp.Status))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}
