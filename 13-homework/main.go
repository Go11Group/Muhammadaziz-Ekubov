package main

import (
	"Go11/13-homework/Delete"
	"Go11/13-homework/Get"
	"Go11/13-homework/Post"
	"Go11/13-homework/Put"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/basic-get", Get.Get_1)
	http.HandleFunc("/basic-get-param", Get.Get_2)
	http.HandleFunc("/basic-get-header", Get.Get_3)

	http.HandleFunc("/basic-delete", Delete.Delete_1)
	http.HandleFunc("/delete-with-param/", Delete.Delete_2)
	http.HandleFunc("/delete-with-header", Delete.Delete_3)

	http.HandleFunc("/basic-post", Post.Post_1)
	http.HandleFunc("/post-form-data", Post.Post_2)

	http.HandleFunc("/basic-put", Put.Put_1)
	http.HandleFunc("/put-with-param/", Put.Put_2)
	http.HandleFunc("/put-with-header", Put.Put_3)

	fmt.Println("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
