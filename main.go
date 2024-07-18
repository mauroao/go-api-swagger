package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/mauroao/go-api-swagger/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Swagger Example API
//	@description	This is a sample server.

func main() {
	addr := "127.0.0.1:8080"

	http.HandleFunc("GET /api/hello", helloHandler)
	http.HandleFunc("GET /api/swagger/*", httpSwagger.WrapHandler)

	fmt.Printf("Listening on %s\r", addr)

	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

// @Summary		Hello
// @Description	Hello example
// @Produce		plain
// @Success		200			{string}	string	"hello"
// @Router		/api/hello 	[get]
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello!!!")
}
