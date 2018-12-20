// Copyright 2018 Bloomberg Finance L.P.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// HealthzHandlerFunc turns a function with the right signature into a healthz handler
type HealthzHandlerFunc func(HealthzParams) middleware.Responder

// Handle executing the request and returning a response
func (fn HealthzHandlerFunc) Handle(params HealthzParams) middleware.Responder {
	return fn(params)
}

// HealthzHandler interface for that can handle valid healthz params
type HealthzHandler interface {
	Handle(HealthzParams) middleware.Responder
}

// NewHealthz creates a new http.Handler for the healthz operation
func NewHealthz(ctx *middleware.Context, handler HealthzHandler) *Healthz {
	return &Healthz{Context: ctx, Handler: handler}
}

/*Healthz swagger:route GET /healthz healthz

The healthcheck endpoint provides detailed information about the health of a web service. If each of the components required by the service are healthy, then the service is considered healthy and will return a 200 OK response. If any of the components needed by the service are unhealthy, then a 503 Service Unavailable response will be provided.

*/
type Healthz struct {
	Context *middleware.Context
	Handler HealthzHandler
}

func (o *Healthz) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewHealthzParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
