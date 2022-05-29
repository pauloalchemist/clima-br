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
	cli.CheckError(err)

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	cli.CheckError(err)

	conditions, err := cli.ParseResponse(data)
	cli.CheckError(err)

	fmt.Println("Cidade:", conditions.City)
	fmt.Println("Temperatura (°C):", conditions.Temp)
}
