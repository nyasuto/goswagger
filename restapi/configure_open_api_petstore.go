// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"goswagger/restapi/operations/node"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/go-openapi/runtime/security"
	"goswagger/restapi/operations"
)

//go:generate swagger generate server --target .. --name  --spec ../test.yaml

func configureFlags(api *operations.OpenAPIPetstoreAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.OpenAPIPetstoreAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	// api.XMLConsumer = runtime.XMLConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// api.XMLProducer = runtime.XMLProducer()

	// Applies when the "api_key" header is set
	api.APIKeyAuth = func(token string) (interface{}, error) {
		return "yasu", nil
		// return nil, errors.NotImplemented("api key auth (api_key) api_key from header param [api_key] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	api.APIAuthorizer = security.Authorized()

	//api.PetGetPetByIDHandler = pet.Search
	/*
	api.PetGetPetByIDHandler = pet.GetPetByIDHandlerFunc(func(params pet.GetPetByIDParams, principal interface{}) middleware.Responder {
		return pet.Search(params)
	})*/

	api.NodeGetNodesHandler = node.GetNodesHandlerFunc(func(params node.GetNodesParams, principal interface{}) middleware.Responder {
		return node.Search()
	})

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
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
