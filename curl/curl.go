package curl

import (
	"fmt"
	"net/http"
	"time"
)

func Curl() {
	var url string
	fmt.Println("Bem vindo ao curl do Tutuixa :)")
	fmt.Println("Coloque uma URL para rodar o curl nela:")
	fmt.Scanln(&url) // URL
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Get(url)
	fmt.Println(err)
	if resp != nil {
		fmt.Println("RESP 200")
		return
	}
}
