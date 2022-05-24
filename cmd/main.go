package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

const baseUrl = "https://apiprevmet3.inmet.gov.br/previsao/2910800"

func main() {
	resp, err := http.Get(baseUrl)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	io.Copy(os.Stdout, resp.Body)
}
