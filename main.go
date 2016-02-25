package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	InitConfig()
	InitRouter()
	log.Println(fmt.Sprintf("Serving at localhost:%d...", config.Port))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
