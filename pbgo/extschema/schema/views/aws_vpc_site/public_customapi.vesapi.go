//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package aws_vpc_site

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

// Create CustomAPI GRPC Client satisfying server.CustomClient
type CustomAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomAPIGrpcClient) doRPCSetVIPInfo(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetVIPInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_vpc_site.SetVIPInfoRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetVIPInfo(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) doRPCSetVPCK8SHostnames(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &SetVPCK8SHostnamesRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesRequest", yamlReq)
	}
	rsp, err := c.grpcClient.SetVPCK8SHostnames(ctx, req, opts...)
	return rsp, err
}

func (c *CustomAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["SetVIPInfo"] = ccl.doRPCSetVIPInfo

	rpcFns["SetVPCK8SHostnames"] = ccl.doRPCSetVPCK8SHostnames

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPI REST Client satisfying server.CustomClient
type CustomAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomAPIRestClient) doRPCSetVIPInfo(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetVIPInfoRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_vpc_site.SetVIPInfoRequest: %s", yamlReq, err)
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
		q.Add("vip_params_per_az", fmt.Sprintf("%v", req.VipParamsPerAz))

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
	pbRsp := &SetVIPInfoResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.views.aws_vpc_site.SetVIPInfoResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) doRPCSetVPCK8SHostnames(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &SetVPCK8SHostnamesRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesRequest: %s", yamlReq, err)
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
		q.Add("node_names", fmt.Sprintf("%v", req.NodeNames))

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
	pbRsp := &SetVPCK8SHostnamesResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["SetVIPInfo"] = ccl.doRPCSetVIPInfo

	rpcFns["SetVPCK8SHostnames"] = ccl.doRPCSetVPCK8SHostnames

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomAPIInprocClient

// INPROC Client (satisfying CustomAPIClient interface)
type CustomAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomAPIInprocClient) SetVIPInfo(ctx context.Context, in *SetVIPInfoRequest, opts ...grpc.CallOption) (*SetVIPInfoResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.aws_vpc_site.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *SetVIPInfoResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.aws_vpc_site.SetVIPInfoRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SetVIPInfo' operation on 'aws_vpc_site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.aws_vpc_site.CustomAPI.SetVIPInfo"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetVIPInfo(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.views.aws_vpc_site.SetVIPInfoResponse", rsp)...)

	return rsp, nil
}
func (c *CustomAPIInprocClient) SetVPCK8SHostnames(ctx context.Context, in *SetVPCK8SHostnamesRequest, opts ...grpc.CallOption) (*SetVPCK8SHostnamesResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.views.aws_vpc_site.CustomAPI")
	cah, ok := ah.(CustomAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomAPISrv", ah)
	}

	var (
		rsp *SetVPCK8SHostnamesResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomAPI.SetVPCK8SHostnames' operation on 'aws_vpc_site'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.views.aws_vpc_site.CustomAPI.SetVPCK8SHostnames"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.SetVPCK8SHostnames(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesResponse", rsp)...)

	return rsp, nil
}

func NewCustomAPIInprocClient(svc svcfw.Service) CustomAPIClient {
	return &CustomAPIInprocClient{svc: svc}
}

// RegisterGwCustomAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomAPIHandlerClient(ctx, mux, NewCustomAPIInprocClient(s))
}

var CustomAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "AWS VPC site",
        "description": "AWS VPC site view defines a required parameters that can be used in CRUD, to create and manage a volterra site in AWS VPC.\nIt can be used to either automatically create or Manually assisted site creation in AWS VPC.\n\nView will create following child objects.\n\n* Site",
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
        "/public/namespaces/{namespace}/aws_vpc_site/{name}/set_vip_info": {
            "post": {
                "summary": "Configure AWS VPC Site VIP Information",
                "description": "Configure AWS VPC Site VIP Information",
                "operationId": "ves.io.schema.views.aws_vpc_site.CustomAPI.SetVIPInfo",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/aws_vpc_siteSetVIPInfoResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"aws-vpc-site-1\"\nName of the object to be configured",
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
                            "$ref": "#/definitions/aws_vpc_siteSetVIPInfoRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-aws_vpc_site-CustomAPI-SetVIPInfo"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.aws_vpc_site.CustomAPI.SetVIPInfo"
            },
            "x-displayname": "Custom API for AWS VPC site",
            "x-ves-proto-service": "ves.io.schema.views.aws_vpc_site.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        },
        "/public/namespaces/{namespace}/aws_vpc_site/{name}/storage/set_vpc_k8s_hostnames": {
            "post": {
                "summary": "Configure VPC k8s hostnames",
                "description": "Configure VPC k8s node hostname set",
                "operationId": "ves.io.schema.views.aws_vpc_site.CustomAPI.SetVPCK8SHostnames",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/aws_vpc_siteSetVPCK8SHostnamesResponse"
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
                        "description": "Namespace\n\nx-example: \"default\"\nNamespace for the object to be configured",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "name",
                        "description": "Name\n\nx-example: \"aws-vpc-site-1\"\nName of the object to be configured",
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
                            "$ref": "#/definitions/aws_vpc_siteSetVPCK8SHostnamesRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-views-aws_vpc_site-CustomAPI-SetVPCK8SHostnames"
                },
                "x-ves-proto-rpc": "ves.io.schema.views.aws_vpc_site.CustomAPI.SetVPCK8SHostnames"
            },
            "x-displayname": "Custom API for AWS VPC site",
            "x-ves-proto-service": "ves.io.schema.views.aws_vpc_site.CustomAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "aws_vpc_siteSetVIPInfoRequest": {
            "type": "object",
            "description": "Request to configure AWS VPC Site VIP information",
            "title": "Request to configure AWS VPC Site VIP information",
            "x-displayname": "Request to configure AWS VPC Site VIP information",
            "x-ves-proto-message": "ves.io.schema.views.aws_vpc_site.SetVIPInfoRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"aws-vpc-site-1\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-vpc-site-1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "vip_params_per_az": {
                    "type": "array",
                    "description": " VIP Parameters per AZ\n\nExample: - \"master-0\"-",
                    "title": "VIP Params Per AZ",
                    "items": {
                        "$ref": "#/definitions/sitePublishVIPParamsPerAz"
                    },
                    "x-displayname": "VIP Params Per AZ",
                    "x-ves-example": "master-0"
                }
            }
        },
        "aws_vpc_siteSetVIPInfoResponse": {
            "type": "object",
            "title": "Response to configure AWS VPC Site VIP Information",
            "x-displayname": "Response to configure AWS VPC Site VIP Information",
            "x-ves-proto-message": "ves.io.schema.views.aws_vpc_site.SetVIPInfoResponse"
        },
        "aws_vpc_siteSetVPCK8SHostnamesRequest": {
            "type": "object",
            "description": "Request to configure VPC k8s node hostname set",
            "title": "Request to configure VPC k8s node hostname set",
            "x-displayname": "Request to configure VPC k8s node hostname set",
            "x-ves-proto-message": "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesRequest",
            "properties": {
                "name": {
                    "type": "string",
                    "description": " Name of the object to be configured\n\nExample: - \"aws-vpc-site-1\"-",
                    "title": "Name",
                    "x-displayname": "Name",
                    "x-ves-example": "aws-vpc-site-1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace for the object to be configured\n\nExample: - \"default\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "default"
                },
                "node_names": {
                    "type": "array",
                    "description": " Hostname of K8S nodes for deinition of Mayastore pools\n\nExample: - \"master-0\"-",
                    "title": "Hostname of K8S nodes",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Hostname of K8S nodes",
                    "x-ves-example": "master-0"
                }
            }
        },
        "aws_vpc_siteSetVPCK8SHostnamesResponse": {
            "type": "object",
            "title": "Response to configure VPC k8s node hostname set",
            "x-displayname": "Response to configure VPC k8s node hostname set",
            "x-ves-proto-message": "ves.io.schema.views.aws_vpc_site.SetVPCK8SHostnamesResponse"
        },
        "sitePublishVIPParamsPerAz": {
            "type": "object",
            "description": "Per AZ parameters needed to publish VIP for public cloud sites",
            "title": "Publish VIP Params Per AZ",
            "x-displayname": "Publish VIP Params Per AZ",
            "x-ves-proto-message": "ves.io.schema.site.PublishVIPParamsPerAz",
            "properties": {
                "az_name": {
                    "type": "string",
                    "description": " Name of the Availability zone\n\nExample: - \"us-east-2a\"-",
                    "title": "AZ Name",
                    "x-displayname": "AZ Name",
                    "x-ves-example": "us-east-2a"
                },
                "inside_vip": {
                    "type": "array",
                    "description": " List of Inside VIPs for an AZ\n\nExample: - \"192.168.0.156\"-",
                    "title": "Inside VIP(s)",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Inside VIP(s)",
                    "x-ves-example": "192.168.0.156"
                },
                "inside_vip_cname": {
                    "type": "string",
                    "description": " CNAME value for the inside VIP,\n These are usually public cloud generated CNAME\n\nExample: - \"test.56670-387196482.useast2.ves.io\"-",
                    "title": "Inside VIP CNAME",
                    "x-displayname": "Inside VIP CNAME",
                    "x-ves-example": "test.56670-387196482.useast2.ves.io"
                },
                "outside_vip": {
                    "type": "array",
                    "description": " List of Outside VIPs for an AZ\n\nExample: - \"192.168.0.156\"-\nRequired: YES",
                    "title": "Outside VIP(s)",
                    "items": {
                        "type": "string"
                    },
                    "x-displayname": "Outside VIP(s)",
                    "x-ves-example": "192.168.0.156",
                    "x-ves-required": "true"
                },
                "outside_vip_cname": {
                    "type": "string",
                    "description": " CNAME value for the outside VIP\n These are usually public cloud generated CNAME\n\nExample: - \"test.56670-387196482.useast2.ves.io\"-",
                    "title": "Outside VIP CNAME",
                    "x-displayname": "Outside VIP CNAME",
                    "x-ves-example": "test.56670-387196482.useast2.ves.io"
                }
            }
        }
    },
    "x-displayname": "Configure AWS VPC Site",
    "x-ves-proto-file": "ves.io/schema/views/aws_vpc_site/public_customapi.proto"
}`
