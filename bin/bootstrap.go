package bin

import (
	"database/sql"
	"github.com/go-chi/chi"
	"github.com/maestre3d/go-mq-example/common"
	delivery "github.com/maestre3d/go-mq-example/delivery/http"
	"log"
)

// Bootstrap Application's Bootstrapper
type Bootstrap struct {
	db *sql.DB
}

// InitApplication Start application
func (b *Bootstrap) InitApplication() {
	log.Print("APPLICATION: Starting ...")
	// Start application's environment
	common.InitEnvironment()
	log.Print("ENVIRONMENT: Started environment")
	// Start required services
	log.Print(common.ServicesStarted)
}

// StartHTTPServer Run an HTTP Server
func (b *Bootstrap) StartHTTPServer() *chi.Mux {
	router := delivery.NewRouter()
	log.Printf("HTTP: HTTP Server started on localhost%s", common.HTTPPort)
	return router.Mux
}

// StartAMQPBroker Run an AMQP Broker
func (b *Bootstrap) StartAMQPBroker() {

}
