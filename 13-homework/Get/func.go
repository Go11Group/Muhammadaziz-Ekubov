package Get

import (
	"fmt"
	"net/http"
)

func Get_1(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	_, err := w.Write([]byte("hi student"))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func Get_2(w http.ResponseWriter, r *http.Request) {
	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	words := ""
	fmt.Scan(&words)

	_, err := w.Write([]byte(words))
	if err != nil {
		fmt.Println("Error writing response:", err)
	}
}

func Get_3(w http.ResponseWriter, r *http.Request) {

	fmt.Println("URL:", r.URL)
	fmt.Println("Host:", r.Host)
	fmt.Println("Method:", r.Method)

	word := ""

	for {
		fmt.Scan(&word)
		if word == "exit" {
			return
		}
		_, err := w.Write([]byte(word))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	}
}
