// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	"codementordev/hello-johnhckuo/internal/handlers"
	"codementordev/hello-johnhckuo/models"

	"codementordev/hello-johnhckuo/restapi/operations"

	"codementordev/hello-johnhckuo/restapi/operations/timeslot"
)

//go:generate swagger generate server --target ../../hello-johnhckuo --name CodementordevHelloJohnhckuo --spec ../swagger.yml --principal interface{}

func configureFlags(api *operations.CodementordevHelloJohnhckuoAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.CodementordevHelloJohnhckuoAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.TimeslotAddTimeslotHandler = timeslot.AddTimeslotHandlerFunc(func(params timeslot.AddTimeslotParams) middleware.Responder {
		return middleware.NotImplemented("operation timeslot.AddTimeslot has not yet been implemented")
	})

	api.TimeslotAddUserHandler = timeslot.AddUserHandlerFunc(func(params timeslot.AddUserParams) middleware.Responder {
		var id int64
		var err error
		if id, err = handlers.CreateUser(params.Username); err != nil {
			return timeslot.NewAddUserDefault(500).WithPayload(&models.Error{Code: 500, Message: swag.String(err.Error())})
		}
		return timeslot.NewAddUserCreated().WithPayload(&models.User{ID: id, Name: &params.Username})
	})

	api.TimeslotDestroyTimeslotHandler = timeslot.DestroyTimeslotHandlerFunc(func(params timeslot.DestroyTimeslotParams) middleware.Responder {
		return middleware.NotImplemented("operation timeslot.DestroyTimeslot has not yet been implemented")
	})

	api.TimeslotGetTimeslotHandler = timeslot.GetTimeslotHandlerFunc(func(params timeslot.GetTimeslotParams) middleware.Responder {
		return middleware.NotImplemented("operation timeslot.GetTimeslot has not yet been implemented")
	})

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
