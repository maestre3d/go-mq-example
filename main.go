package main

import (
	"log"
	"net/http"

	"github.com/maestre3d/go-mq-example/bin"
	"github.com/maestre3d/go-mq-example/common"
)

func main() {

	// Start application
	bootstrap := new(bin.Bootstrap)

	// Init HTTP
	err := http.ListenAndServe(common.HTTPPort, bootstrap.StartHTTPServer())
	if err != nil {
		log.Panic(err)
	}

	// Init AMQP
	// PD: Use go routine
}
