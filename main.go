package main

import (
	"flag"
	"fmt"
)

type Checker interface {
	Check()
}

func main() {
	url := flag.String("u", "", "Enter the target url")
	port := flag.String("p", "", "Enter specific port")
	server := flag.Bool("s", false, "Start rest api server")

	flag.Parse()

	if url != nil && *url != "" {
		domain := Domain{
			Url: *url,
		}
		if port != nil && *port != "" {
			domain.Port = *port
		} else {
			domain.Port = "80"
		}

		domain.Check()
	}

	if server != nil && *server != false {
		fmt.Println("Start the server")
	}
}
