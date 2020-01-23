package common

const (
	// ServiceStarting Starting a service
	ServiceStarting string = "SERVICE: Starting service %s ... "
	// ServiceStarted Service has started
	ServiceStarted string = "SERVICE: %s service started"
	// ServiceFailure Failed to start service
	ServiceFailure string = "SERVICE: Failed to start %s service"
	//ServicesStarted All services started
	ServicesStarted string = "SERVICE: All required services started"
)

// Response Generic response model
type Response struct {
	Message string `json:"message"`
}
