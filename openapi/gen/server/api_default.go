/*
 * openapi-tryout
 *
 * Try OpenAPI 3
 *
 * API version: 0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// A DefaultApiController binds http requests to an api service and writes the service results to the http response
type DefaultApiController struct {
	service DefaultApiServicer
}

// NewDefaultApiController creates a default api controller
func NewDefaultApiController(s DefaultApiServicer) Router {
	return &DefaultApiController{ service: s }
}

// Routes returns all of the api route for the DefaultApiController
func (c *DefaultApiController) Routes() Routes {
	return Routes{ 
		{
			"AccountGet",
			strings.ToUpper("Get"),
			"/api/v1/account",
			c.AccountGet,
		},
		{
			"EchoGet",
			strings.ToUpper("Get"),
			"/api/v1/echo",
			c.EchoGet,
		},
	}
}

// AccountGet - 
func (c *DefaultApiController) AccountGet(w http.ResponseWriter, r *http.Request) { 
	query := r.URL.Query()
	offset, err := parseInt32Parameter(query.Get("offset"))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	limit, err := parseInt32Parameter(query.Get("limit"))
	if err != nil {
		w.WriteHeader(500)
		return
	}
	
	userName := query.Get("userName")
	result, err := c.service.AccountGet(r.Context(), offset, limit, userName)
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}

// EchoGet - 
func (c *DefaultApiController) EchoGet(w http.ResponseWriter, r *http.Request) { 
	result, err := c.service.EchoGet(r.Context())
	//If an error occured, encode the error with the status code
	if err != nil {
		EncodeJSONResponse(err.Error(), &result.Code, w)
		return
	}
	//If no error, encode the body and the result code
	EncodeJSONResponse(result.Body, &result.Code, w)
	
}
