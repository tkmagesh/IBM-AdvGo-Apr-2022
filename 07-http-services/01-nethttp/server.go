package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to Go Web Server"))
	})
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			w.Write([]byte("All the products will be returned"))
		case "POST":
			w.Write([]byte("The submitted product will be added"))
		case "DELETE":
			w.Write([]byte("The given product will be removed"))
		}
	})
	http.ListenAndServe(":8080", nil)
}

/*
	/products
		GET = /
			get the data from db
			Serialize the data into JSON
			send the response
		GET = /P-101
		POST =
			get the data from the request
			deserialize the data
			add the data to the db
			send the response
		PUT
		DELETE
		PATCH =

*/
