package restapi

import (
	"net/http"

	errors "github.com/go-swagger/go-swagger/errors"
	httpkit "github.com/go-swagger/go-swagger/httpkit"
	middleware "github.com/go-swagger/go-swagger/httpkit/middleware"

	"restapi/operations"
)

// This file is safe to edit. Once it exists it will not be overwritten

func configureFlags(api *operations.SoftLayerBaremetalProvisioningAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.SoftLayerBaremetalProvisioningAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	api.JSONConsumer = httpkit.JSONConsumer()

	api.JSONProducer = httpkit.JSONProducer()

	api.GetBmsDeploymentNameHandler = operations.GetBmsDeploymentNameHandlerFunc(func(params operations.GetBmsDeploymentNameParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetBmsDeploymentName has not yet been implemented")
	})
	api.GetInfoHandler = operations.GetInfoHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation .GetInfo has not yet been implemented")
	})
	api.GetLoginUsernamePasswordHandler = operations.GetLoginUsernamePasswordHandlerFunc(func(params operations.GetLoginUsernamePasswordParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetLoginUsernamePassword has not yet been implemented")
	})
	api.GetSlPackagePackageIDOptionsHandler = operations.GetSlPackagePackageIDOptionsHandlerFunc(func(params operations.GetSlPackagePackageIDOptionsParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetSlPackagePackageIDOptions has not yet been implemented")
	})
	api.GetSlPackagesHandler = operations.GetSlPackagesHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation .GetSlPackages has not yet been implemented")
	})
	api.GetStemcellsHandler = operations.GetStemcellsHandlerFunc(func() middleware.Responder {
		return middleware.NotImplemented("operation .GetStemcells has not yet been implemented")
	})
	api.GetTasksHandler = operations.GetTasksHandlerFunc(func(params operations.GetTasksParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetTasks has not yet been implemented")
	})
	api.GetTasksTaskIDTxtLevelHandler = operations.GetTasksTaskIDTxtLevelHandlerFunc(func(params operations.GetTasksTaskIDTxtLevelParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetTasksTaskIDTxtLevel has not yet been implemented")
	})
	api.PutBaremetalServerIDStateHandler = operations.PutBaremetalServerIDStateHandlerFunc(func(params operations.PutBaremetalServerIDStateParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutBaremetalServerIDState has not yet been implemented")
	})
	api.PutBaremetalsHandler = operations.PutBaremetalsHandlerFunc(func(params operations.PutBaremetalsParams) middleware.Responder {
		return middleware.NotImplemented("operation .PutBaremetals has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
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
