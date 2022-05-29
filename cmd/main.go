package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/pauloalchemist/climabr/internal/cli"
)

const baseURL = "http://apiadvisor.climatempo.com.br"

func main() {
	key := os.Getenv("API_KEY")
	if key == "" {
		log.Fatal("Set the environment variable API_KEY.")
	}

	//Feira de Santana is 7554
	URL := fmt.Sprintf("%s/api/v1/weather/locale/7554/current?token=%s", baseURL, key)

	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	conditions, err := cli.ParseResponse(data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Cidade:", conditions.City)
	fmt.Println("Temperatura (°C):", conditions.Temp)
}
