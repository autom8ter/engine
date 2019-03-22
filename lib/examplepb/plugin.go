//+build

package main

import (
	"github.com/autom8ter/engine/lib/examplepb/server"
)

/*
rpc Echo(SimpleMessage) returns (SimpleMessage) {
		option (google.api.http) = {
			post: "/v1/example/echo/{id}"
			additional_bindings {
				get: "/v1/example/echo/{id}/{num}"
			}
			additional_bindings {
				get: "/v1/example/echo/{id}/{num}/{lang}"
			}
			additional_bindings {
				get: "/v1/example/echo1/{id}/{line_num}/{status.note}"
			}
			additional_bindings {
				get: "/v1/example/echo2/{no.note}"
			}
		};
	}
	// EchoBody method receives a simple message and returns it.
	rpc EchoBody(SimpleMessage) returns (SimpleMessage) {
		option (google.api.http) = {
			post: "/v1/example/echo_body"
			body: "*"
		};
	}
	// EchoDelete method receives a simple message and returns it.
	rpc EchoDelete(SimpleMessage) returns (SimpleMessage) {
		option (google.api.http) = {
			delete: "/v1/example/echo_delete"
		};
	}
 */
var Plugin server.Example
