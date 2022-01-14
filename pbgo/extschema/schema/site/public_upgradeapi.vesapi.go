//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package site

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gogo/protobuf/proto"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"

	"gopkg.volterra.us/stdlib/client"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

var (
	_ = fmt.Sprintf("dummy for fmt import use")
)

// Create UpgradeAPI GRPC Client satisfying server.CustomClient
type UpgradeAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient UpgradeAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *UpgradeAPIGrpcClient) doRPCUpgradeOS(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &UpgradeOSRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.UpgradeOSRequest", yamlReq)
	}
	rsp, err := c.grpcClient.UpgradeOS(ctx, req, opts...)
	return rsp, err
}

func (c *UpgradeAPIGrpcClient) doRPCUpgradeSW(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &UpgradeSWRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.UpgradeSWRequest", yamlReq)
	}
	rsp, err := c.grpcClient.UpgradeSW(ctx, req, opts...)
	return rsp, err
}

func (c *UpgradeAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}
	if cco.YAMLReq == "" {
		return nil, fmt.Errorf("Error, empty request body")
	}
	ctx = client.AddHdrsToCtx(cco.Headers, ctx)

	rsp, err := rpcFn(ctx, cco.YAMLReq, cco.GrpcCallOpts...)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using GRPC")
	}
	if cco.OutCallResponse != nil {
		cco.OutCallResponse.ProtoMsg = rsp
	}
	return rsp, nil
}

func NewUpgradeAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &UpgradeAPIGrpcClient{
		conn:       cc,
		grpcClient: NewUpgradeAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["UpgradeOS"] = ccl.doRPCUpgradeOS

	rpcFns["UpgradeSW"] = ccl.doRPCUpgradeSW

	ccl.rpcFns = rpcFns

	return ccl
}

// Create UpgradeAPI REST Client satisfying server.CustomClient
type UpgradeAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *UpgradeAPIRestClient) doRPCUpgradeOS(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &UpgradeOSRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.UpgradeOSRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		var op string
		if hm == "post" {
			op = http.MethodPost
		} else {
			op = http.MethodPut
		}
		newReq, err := http.NewRequest(op, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrapf(err, "Creating new HTTP %s request for custom API", op)
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("version", fmt.Sprintf("%v", req.Version))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &UpgradeOSResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.site.UpgradeOSResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *UpgradeAPIRestClient) doRPCUpgradeSW(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &UpgradeSWRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.site.UpgradeSWRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		var op string
		if hm == "post" {
			op = http.MethodPost
		} else {
			op = http.MethodPut
		}
		newReq, err := http.NewRequest(op, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrapf(err, "Creating new HTTP %s request for custom API", op)
		}
		hReq = newReq
	case "get":
		newReq, err := http.NewRequest(http.MethodGet, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP GET request for custom API")
		}
		hReq = newReq
		q := hReq.URL.Query()
		_ = q
		q.Add("name", fmt.Sprintf("%v", req.Name))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("version", fmt.Sprintf("%v", req.Version))

		hReq.URL.RawQuery += q.Encode()
	case "delete":
		newReq, err := http.NewRequest(http.MethodDelete, url, nil)
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP DELETE request for custom API")
		}
		hReq = newReq
	default:
		return nil, fmt.Errorf("Error, invalid/empty HTTPMethod(%s) specified, should be POST|DELETE|GET", callOpts.HTTPMethod)
	}
	hReq = hReq.WithContext(ctx)
	hReq.Header.Set("Content-Type", "application/json")
	client.AddHdrsToReq(callOpts.Headers, hReq)

	rsp, err := c.client.Do(hReq)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient")
	}
	defer rsp.Body.Close()

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &UpgradeSWResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.site.UpgradeSWResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *UpgradeAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
	rpcFn, exists := c.rpcFns[rpc]
	if !exists {
		return nil, fmt.Errorf("Error, no such rpc %s", rpc)
	}
	cco := server.NewCustomCallOpts()
	for _, opt := range opts {
		opt(cco)
	}

	rsp, err := rpcFn(ctx, cco)
	if err != nil {
		return nil, errors.Wrap(err, "Doing custom RPC using Rest")
	}
	return rsp, nil
}

func NewUpgradeAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &UpgradeAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["UpgradeOS"] = ccl.doRPCUpgradeOS

	rpcFns["UpgradeSW"] = ccl.doRPCUpgradeSW

	ccl.rpcFns = rpcFns

	return ccl
}

// Create UpgradeAPIInprocClient

// INPROC Client (satisfying UpgradeAPIClient interface)
type UpgradeAPIInprocClient struct {
	svc svcfw.Service
}

func (c *UpgradeAPIInprocClient) UpgradeOS(ctx context.Context, in *UpgradeOSRequest, opts ...grpc.CallOption) (*UpgradeOSResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.site.UpgradeAPI")
	cah, ok := ah.(UpgradeAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *UpgradeAPISrv", ah)
	}

	var (
		rsp *UpgradeOSResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.site.UpgradeOSRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'UpgradeAPI.UpgradeOS' operation on 'site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.site.UpgradeAPI.UpgradeOS"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpgradeOS(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.site.UpgradeOSResponse", rsp)...)

	return rsp, nil
}
func (c *UpgradeAPIInprocClient) UpgradeSW(ctx context.Context, in *UpgradeSWRequest, opts ...grpc.CallOption) (*UpgradeSWResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.site.UpgradeAPI")
	cah, ok := ah.(UpgradeAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *UpgradeAPISrv", ah)
	}

	var (
		rsp *UpgradeSWResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.site.UpgradeSWRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'UpgradeAPI.UpgradeSW' operation on 'site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.site.UpgradeAPI.UpgradeSW"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.UpgradeSW(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.site.UpgradeSWResponse", rsp)...)

	return rsp, nil
}

func NewUpgradeAPIInprocClient(svc svcfw.Service) UpgradeAPIClient {
	return &UpgradeAPIInprocClient{svc: svc}
}

// RegisterGwUpgradeAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwUpgradeAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterUpgradeAPIHandlerClient(ctx, mux, NewUpgradeAPIInprocClient(s))
}

var UpgradeAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Site upgrade custom API",
        "description": "These API provide a way to upgrade site's volterra software and OS",
        "version": "version not set"
    },
    "schemes": [
        "http",
        "https"
    ],
    "consumes": [
        "application/json"
    ],
    "produces": [
        "application/json"
    ],
    "tags": null,
    "paths": {
        "/public/namespaces/{namespace}/sites/{name}/upgrade_os": {
            "post": {
                "summary": "Upgrade OS",
                "description": "Upgrade Site OS version",
                "operationId": "ves.io.schema.site.UpgradeAPI.UpgradeOS",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/siteUpgradeOSResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-required\nx-example: \"system\"\nSite namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-required\nx-example: \"ce398\"\nSite name",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/siteUpgradeOSRequest"
                        }
                    }
                ],
                "tags": [
                    "UpgradeAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-site-UpgradeAPI-UpgradeOS"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.UpgradeAPI.UpgradeOS"
            },
            "x-displayname": "Site Upgrade",
            "x-ves-proto-service": "ves.io.schema.site.UpgradeAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/sites/{name}/upgrade_sw": {
            "post": {
                "summary": "Upgrade SW",
                "description": "Upgrade Site SW version",
                "operationId": "ves.io.schema.site.UpgradeAPI.UpgradeSW",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/siteUpgradeSWResponse"
                        }
                    },
                    "401": {
                        "description": "Returned when operation is not authorized",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "403": {
                        "description": "Returned when there is no permission to access resource",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "404": {
                        "description": "Returned when resource is not found",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "409": {
                        "description": "Returned when operation on resource is conflicting with current value",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "429": {
                        "description": "Returned when operation has been rejected as it is happening too frequently",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "500": {
                        "description": "Returned when server encountered an error in processing API",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "503": {
                        "description": "Returned when service is unavailable temporarily",
                        "schema": {
                            "format": "string"
                        }
                    },
                    "504": {
                        "description": "Returned when server timed out processing request",
                        "schema": {
                            "format": "string"
                        }
                    }
                },
                "parameters": [
                    {
                        "name": "namespace",
                        "description": "Namespace\n\nx-required\nx-example: \"system\"\nSite namespace",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-required\nx-example: \"ce398\"\nSite name",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Name"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/siteUpgradeSWRequest"
                        }
                    }
                ],
                "tags": [
                    "UpgradeAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-site-UpgradeAPI-UpgradeSW"
                },
                "x-ves-proto-rpc": "ves.io.schema.site.UpgradeAPI.UpgradeSW"
            },
            "x-displayname": "Site Upgrade",
            "x-ves-proto-service": "ves.io.schema.site.UpgradeAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "siteUpgradeOSRequest": {
            "type": "object",
            "description": "Request for Upgrade OS of Site",
            "title": "Upgrade OS request",
            "x-displayname": "Upgrade OS Request",
            "x-ves-proto-message": "ves.io.schema.site.UpgradeOSRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Site name\n\nExample: - \"ce398\"-\nRequired: YES",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "ce398",
                    "x-ves-required": "true"
                },
                "namespace": {
                    "type": "string",
                    "description": " Site namespace\n\nExample: - \"system\"-\nRequired: YES",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "system",
                    "x-ves-required": "true"
                },
                "version": {
                    "type": "string",
                    "description": " Version to upgraded to\n\nExample: - \"7.2003.20\"-\nRequired: YES",
                    "title": "SW/OS Version",
                    "x-displayname": "Version",
                    "x-ves-example": "7.2003.20",
                    "x-ves-required": "true"
                }
            }
        },
        "siteUpgradeOSResponse": {
            "type": "object",
            "description": "Response for OS Upgrade Request, empty because the only returned information\nis currently error message",
            "title": "Upgrade OS Response",
            "x-displayname": "Upgrade OS Response",
            "x-ves-proto-message": "ves.io.schema.site.UpgradeOSResponse"
        },
        "siteUpgradeSWRequest": {
            "type": "object",
            "description": "Request for Upgrade SW of Site",
            "title": "Upgrade SW request",
            "x-displayname": "Upgrade SW Request",
            "x-ves-proto-message": "ves.io.schema.site.UpgradeSWRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Site name\n\nExample: - \"ce398\"-\nRequired: YES",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "ce398",
                    "x-ves-required": "true"
                },
                "namespace": {
                    "type": "string",
                    "description": " Site namespace\n\nExample: - \"system\"-\nRequired: YES",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "system",
                    "x-ves-required": "true"
                },
                "version": {
                    "type": "string",
                    "description": " Version to upgraded to\n\nExample: - \"crt-20201010-600\"-\nRequired: YES",
                    "title": "SW/OS Version",
                    "x-displayname": "Version",
                    "x-ves-example": "crt-20201010-600",
                    "x-ves-required": "true"
                }
            }
        },
        "siteUpgradeSWResponse": {
            "type": "object",
            "description": "Response for SW Upgrade Request, empty because the only returned information\nis currently error message",
            "title": "Upgrade SW Response",
            "x-displayname": "Upgrade SW Response",
            "x-ves-proto-message": "ves.io.schema.site.UpgradeSWResponse"
        }
    },
    "x-displayname": "Site",
    "x-ves-proto-file": "ves.io/schema/site/public_upgradeapi.proto"
}`
