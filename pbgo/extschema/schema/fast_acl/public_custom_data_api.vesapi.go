// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package fast_acl

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

// Create CustomDataAPI GRPC Client satisfying server.CustomClient
type CustomDataAPIGrpcClient struct {
	conn       *grpc.ClientConn
	grpcClient CustomDataAPIClient
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error)
}

func (c *CustomDataAPIGrpcClient) doRPCFastACLHits(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &FastACLHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.fast_acl.FastACLHitsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.FastACLHits(ctx, req, opts...)
	return rsp, err
}

func (c *CustomDataAPIGrpcClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomDataAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomDataAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["FastACLHits"] = ccl.doRPCFastACLHits

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomDataAPI REST Client satisfying server.CustomClient
type CustomDataAPIRestClient struct {
	baseURL string
	client  http.Client
	// map of rpc name to its invocation
	rpcFns map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error)
}

func (c *CustomDataAPIRestClient) doRPCFastACLHits(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &FastACLHitsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.fast_acl.FastACLHitsRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post", "put":
		jsn, err := codec.ToJSON(req, codec.ToWithUseProtoFieldName())
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
		q.Add("end_time", fmt.Sprintf("%v", req.EndTime))
		q.Add("group_by", fmt.Sprintf("%v", req.GroupBy))
		q.Add("label_filter", fmt.Sprintf("%v", req.LabelFilter))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("start_time", fmt.Sprintf("%v", req.StartTime))
		q.Add("step", fmt.Sprintf("%v", req.Step))

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

	// checking whether the status code is a successful status code (2xx series)
	if rsp.StatusCode < 200 || rsp.StatusCode > 299 {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &FastACLHitsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, errors.Wrapf(err, "JSON Response %s is not of type *ves.io.schema.fast_acl.FastACLHitsResponse", body)

	}
	if callOpts.OutCallResponse != nil {
		callOpts.OutCallResponse.ProtoMsg = pbRsp
		callOpts.OutCallResponse.JSON = string(body)
	}
	return pbRsp, nil
}

func (c *CustomDataAPIRestClient) DoRPC(ctx context.Context, rpc string, opts ...server.CustomCallOpt) (proto.Message, error) {
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

func NewCustomDataAPIRestClient(baseURL string, hc http.Client) server.CustomClient {
	ccl := &CustomDataAPIRestClient{
		baseURL: baseURL,
		client:  hc,
	}

	rpcFns := make(map[string]func(context.Context, *server.CustomCallOpts) (proto.Message, error))
	rpcFns["FastACLHits"] = ccl.doRPCFastACLHits

	ccl.rpcFns = rpcFns

	return ccl
}

// Create customDataAPIInprocClient

// INPROC Client (satisfying CustomDataAPIClient interface)
type customDataAPIInprocClient struct {
	CustomDataAPIServer
}

func (c *customDataAPIInprocClient) FastACLHits(ctx context.Context, in *FastACLHitsRequest, opts ...grpc.CallOption) (*FastACLHitsResponse, error) {
	return c.CustomDataAPIServer.FastACLHits(ctx, in)
}

func NewCustomDataAPIInprocClient(svc svcfw.Service) CustomDataAPIClient {
	return &customDataAPIInprocClient{CustomDataAPIServer: NewCustomDataAPIServer(svc)}
}

// RegisterGwCustomDataAPIHandler registers with grpc-gw with an inproc-client backing so that
// rest to grpc happens without a grpc.Dial (thus avoiding additional certs for mTLS)
func RegisterGwCustomDataAPIHandler(ctx context.Context, mux *runtime.ServeMux, svc interface{}) error {
	s, ok := svc.(svcfw.Service)
	if !ok {
		return fmt.Errorf("svc is not svcfw.Service")
	}
	return RegisterCustomDataAPIHandlerClient(ctx, mux, NewCustomDataAPIInprocClient(s))
}

// Create customDataAPISrv

// SERVER (satisfying CustomDataAPIServer interface)
type customDataAPISrv struct {
	svc svcfw.Service
}

func (s *customDataAPISrv) FastACLHits(ctx context.Context, in *FastACLHitsRequest) (*FastACLHitsResponse, error) {
	ah := s.svc.GetAPIHandler("ves.io.schema.fast_acl.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPIServer", ah)
	}

	var (
		rsp *FastACLHitsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, s.svc, "ves.io.schema.fast_acl.FastACLHitsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.FastACLHits' operation on 'fast_acl'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if err := svcfw.FillOneofDefaultChoice(ctx, s.svc, in); err != nil {
		err = server.MaybePublicRestError(ctx, errors.Wrapf(err, "Filling oneof default choice"))
		return nil, server.GRPCStatusFromError(err).Err()
	}

	if s.svc.Config().EnableAPIValidation {
		if rvFn := s.svc.GetRPCValidator("ves.io.schema.fast_acl.CustomDataAPI.FastACLHits"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.FastACLHits(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, s.svc, "ves.io.schema.fast_acl.FastACLHitsResponse", rsp)...)

	return rsp, nil
}

func NewCustomDataAPIServer(svc svcfw.Service) CustomDataAPIServer {
	return &customDataAPISrv{svc: svc}
}

var CustomDataAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Fast ACL",
        "description": "Monitoring APIs Fast ACL",
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
    "tags": [],
    "paths": {
        "/public/namespaces/{namespace}/fast_acl/hits": {
            "post": {
                "summary": "Fast ACL Hits",
                "description": "Get the counter for Fast ACL hits for a given namespace.",
                "operationId": "ves.io.schema.fast_acl.CustomDataAPI.FastACLHits",
                "responses": {
                    "200": {
                        "description": "A successful response.",
                        "schema": {
                            "$ref": "#/definitions/fast_aclFastACLHitsResponse"
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
                        "description": "Namespace\n\nx-example: \"ns1\"\nNamespace is used to scope Fast ACL hits for the given namespace.",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "x-displayname": "Namespace"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/fast_aclFastACLHitsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-fast_acl-customdataapi-fastaclhits"
                },
                "x-ves-proto-rpc": "ves.io.schema.fast_acl.CustomDataAPI.FastACLHits"
            },
            "x-displayname": "Fast ACL Monitoring APIs",
            "x-ves-proto-service": "ves.io.schema.fast_acl.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
        "fast_aclFastACLHits": {
            "type": "object",
            "description": "FastACLHits contains the timeseries data of Fast ACL hits",
            "title": "Fast ACL Hits",
            "x-displayname": "Fast ACL Hits",
            "x-ves-proto-message": "ves.io.schema.fast_acl.FastACLHits",
            "properties": {
                "id": {
                    "description": " ID identifies the unique combination of group_by label values in the response",
                    "title": "ID",
                    "$ref": "#/definitions/fast_aclFastACLHitsId",
                    "x-displayname": "ID"
                },
                "metric": {
                    "type": "array",
                    "description": " x-unit: \"count\"\n List of metric values",
                    "title": "Metric",
                    "items": {
                        "$ref": "#/definitions/schemaMetricValue"
                    },
                    "x-displayname": "Metric"
                }
            }
        },
        "fast_aclFastACLHitsId": {
            "type": "object",
            "description": "FastACLHitsId uniquely identifies an entry in the response to Fast ACL hits request.\nFast ACL hits counter is aggregated based on the labels specified in the group_by field in the request.\nTherefore, only the feields that corresponds to the MetricLabel in the group_by field will have non-empty\nvalue in the response.",
            "title": "Fast ACL Hits ID",
            "x-displayname": "Fast ACL Hits ID",
            "x-ves-proto-message": "ves.io.schema.fast_acl.FastACLHitsId",
            "properties": {
                "fast_acl": {
                    "type": "string",
                    "description": " Fast ACL name\n\nExample: - \"facl-1\"-",
                    "title": "Fast ACL",
                    "x-displayname": "Fast ACL",
                    "x-ves-example": "facl-1"
                },
                "fast_acl_rule": {
                    "type": "string",
                    "description": " Fast ACL Rule name\n\nExample: - \"rule1\"-",
                    "title": "Fast ACL Rule",
                    "x-displayname": "Fast ACL Rule",
                    "x-ves-example": "rule1"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace in which the policy rule was hit\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "site": {
                    "type": "string",
                    "description": " Site name\n\nExample: - \"ce1\"-",
                    "title": "Site",
                    "x-displayname": "Site",
                    "x-ves-example": "ce1"
                }
            }
        },
        "fast_aclFastACLHitsRequest": {
            "type": "object",
            "description": "Request to get the Fast ACL hits counter.",
            "title": "Fast ACL Hits Request",
            "x-displayname": "Fast ACL Hits Request",
            "x-ves-proto-message": "ves.io.schema.fast_acl.FastACLHitsRequest",
            "properties": {
                "end_time": {
                    "type": "string",
                    "description": " end time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the end_time will be evaluated to start_time+10m\n           If start_time is not specified, then the end_time will be evaluated to \u003ccurrent time\u003e\n\nExample: - \"1570007981\"-",
                    "title": "End time",
                    "x-displayname": "End Time",
                    "x-ves-example": "1570007981"
                },
                "group_by": {
                    "type": "array",
                    "description": " Aggregate data by one of more labels specified in -group_by-.\n\n Optional: If not specified, then the rule hits are aggregated/grouped by -FAST_ACL-.",
                    "title": "Group by",
                    "items": {
                        "$ref": "#/definitions/fast_aclFastACLMetricLabel"
                    },
                    "x-displayname": "Group By"
                },
                "label_filter": {
                    "type": "array",
                    "description": " List of label filter expressions of the form \"label\" -Op- \"value\".\n Response will only contain data that matches all the conditions specified in the -label_filter-.\n\n Optional: If not specified, then the metrics will be filtered only based on the -namespace- in the request.",
                    "title": "Label Filter",
                    "items": {
                        "$ref": "#/definitions/fast_aclFastACLMetricLabelFilter"
                    },
                    "x-displayname": "Label Filter"
                },
                "namespace": {
                    "type": "string",
                    "description": " Namespace is used to scope Fast ACL hits for the given namespace.\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "start_time": {
                    "type": "string",
                    "description": " start time of metric collection from which data will be considered.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the start_time will be evaluated to end_time-10m\n           If end_time is not specified, then the start_time will be evaluated to \u003ccurrent time\u003e-10m\n\nExample: - \"1570007981\"-",
                    "title": "Start time",
                    "x-displayname": "Start Time",
                    "x-ves-example": "1570007981"
                },
                "step": {
                    "type": "string",
                    "description": " step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.\n The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn \u003c= end_time.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\n Optional: If not specified, then step size is evaluated to \u003cend_time - start_time\u003e\n\nExample: - \"5m\"-\n\nValidation Rules:\n  ves.io.schema.rules.string.time_interval: true\n",
                    "title": "Step",
                    "x-displayname": "Step",
                    "x-ves-example": "5m",
                    "x-ves-validation-rules": {
                        "ves.io.schema.rules.string.time_interval": "true"
                    }
                }
            }
        },
        "fast_aclFastACLHitsResponse": {
            "type": "object",
            "description": "Number of Fast ACL rule hits for each unique combination of group_by labels in the request.",
            "title": "Fast ACL Hits Response",
            "x-displayname": "Fast ACL Hits Response",
            "x-ves-proto-message": "ves.io.schema.fast_acl.FastACLHitsResponse",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " List of Fast ACL hits data",
                    "title": "Fast ACL Hits",
                    "items": {
                        "$ref": "#/definitions/fast_aclFastACLHits"
                    },
                    "x-displayname": "Fast ACL Hits"
                }
            }
        },
        "fast_aclFastACLMetricLabel": {
            "type": "string",
            "description": "Fast ACL hits can be sliced and diced based on one or more labels listed below.\n",
            "title": "Fast ACL Metric Labels",
            "enum": [
                "NAMESPACE",
                "FAST_ACL",
                "FAST_ACL_RULE",
                "SITE"
            ],
            "default": "NAMESPACE",
            "x-displayname": "Fast ACL Metric Labels",
            "x-ves-proto-enum": "ves.io.schema.fast_acl.FastACLMetricLabel"
        },
        "fast_aclFastACLMetricLabelFilter": {
            "type": "object",
            "description": "Label filter can be specified to filter metrics based on label match",
            "title": "Fast ACL Metric Label Filter",
            "x-displayname": "Fast ACL Metric Label Filter",
            "x-ves-proto-message": "ves.io.schema.fast_acl.FastACLMetricLabelFilter",
            "properties": {
                "label": {
                    "description": " Label associated with Fast ACL hits",
                    "title": "Label",
                    "$ref": "#/definitions/fast_aclFastACLMetricLabel",
                    "x-displayname": "Label"
                },
                "op": {
                    "description": " Operator to evaluate the label ",
                    "title": "Operator",
                    "$ref": "#/definitions/schemaMetricLabelOp",
                    "x-displayname": "Operator"
                },
                "value": {
                    "type": "string",
                    "description": " Value to be compared with\n\nExample: - \"policy1\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "policy1"
                }
            }
        },
        "schemaMetricLabelOp": {
            "type": "string",
            "description": "The operator to use when filtering metrics based on label values.\n",
            "title": "Metric Label Operator",
            "enum": [
                "EQ",
                "NEQ"
            ],
            "default": "EQ",
            "x-displayname": "Metric Label Operator",
            "x-ves-proto-enum": "ves.io.schema.MetricLabelOp"
        },
        "schemaMetricValue": {
            "type": "object",
            "description": "Metric data contains timestamp and the value.",
            "title": "Metric Value",
            "x-displayname": "Metric Value",
            "x-ves-proto-message": "ves.io.schema.MetricValue",
            "properties": {
                "timestamp": {
                    "type": "number",
                    "description": " timestamp\n\nExample: - \"1570007981\"-",
                    "title": "Timestamp",
                    "format": "double",
                    "x-displayname": "Timestamp",
                    "x-ves-example": "1570007981"
                },
                "value": {
                    "type": "string",
                    "description": "\n\nExample: - \"15\"-",
                    "title": "Value",
                    "x-displayname": "Value",
                    "x-ves-example": "15"
                }
            }
        }
    },
    "x-displayname": "Fast ACL",
    "x-ves-proto-file": "ves.io/schema/fast_acl/public_custom_data_api.proto"
}`
