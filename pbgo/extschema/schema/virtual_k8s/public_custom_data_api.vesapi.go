//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//

package virtual_k8s

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

func (c *CustomDataAPIGrpcClient) doRPCPVCMetrics(ctx context.Context, yamlReq string, opts ...grpc.CallOption) (proto.Message, error) {
	req := &PVCMetricsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.virtual_k8s.PVCMetricsRequest", yamlReq)
	}
	rsp, err := c.grpcClient.PVCMetrics(ctx, req, opts...)
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
	return rsp, nil
}

func NewCustomDataAPIGrpcClient(cc *grpc.ClientConn) server.CustomClient {
	ccl := &CustomDataAPIGrpcClient{
		conn:       cc,
		grpcClient: NewCustomDataAPIClient(cc),
	}
	rpcFns := make(map[string]func(context.Context, string, ...grpc.CallOption) (proto.Message, error))
	rpcFns["PVCMetrics"] = ccl.doRPCPVCMetrics

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

func (c *CustomDataAPIRestClient) doRPCPVCMetrics(ctx context.Context, callOpts *server.CustomCallOpts) (proto.Message, error) {
	if callOpts.URI == "" {
		return nil, fmt.Errorf("Error, URI should be specified, got empty")
	}
	url := fmt.Sprintf("%s%s", c.baseURL, callOpts.URI)

	yamlReq := callOpts.YAMLReq
	req := &PVCMetricsRequest{}
	if err := codec.FromYAML(yamlReq, req); err != nil {
		return nil, fmt.Errorf("YAML Request %s is not of type *ves.io.schema.virtual_k8s.PVCMetricsRequest: %s", yamlReq, err)
	}

	var hReq *http.Request
	hm := strings.ToLower(callOpts.HTTPMethod)
	switch hm {
	case "post":
		jsn, err := req.ToJSON()
		if err != nil {
			return nil, errors.Wrap(err, "Custom RestClient converting YAML to JSON")
		}
		newReq, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsn)))
		if err != nil {
			return nil, errors.Wrap(err, "Creating new HTTP POST request for custom API")
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
		q.Add("field_selector", fmt.Sprintf("%v", req.FieldSelector))
		q.Add("filter", fmt.Sprintf("%v", req.Filter))
		q.Add("group_by", fmt.Sprintf("%v", req.GroupBy))
		q.Add("namespace", fmt.Sprintf("%v", req.Namespace))
		q.Add("range", fmt.Sprintf("%v", req.Range))
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

	if rsp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(rsp.Body)
		return nil, fmt.Errorf("Unsuccessful custom API %s on %s, status code %d, body %s, err %s", callOpts.HTTPMethod, callOpts.URI, rsp.StatusCode, body, err)
	}

	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "Custom API RestClient read body")
	}
	pbRsp := &PVCMetricsResponse{}
	if err := codec.FromJSON(string(body), pbRsp); err != nil {
		return nil, fmt.Errorf("JSON Response %s is not of type *ves.io.schema.virtual_k8s.PVCMetricsResponse", body)
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
	rpcFns["PVCMetrics"] = ccl.doRPCPVCMetrics

	ccl.rpcFns = rpcFns

	return ccl
}

// Create CustomDataAPIInprocClient

// INPROC Client (satisfying CustomDataAPIClient interface)
type CustomDataAPIInprocClient struct {
	svc svcfw.Service
}

func (c *CustomDataAPIInprocClient) PVCMetrics(ctx context.Context, in *PVCMetricsRequest, opts ...grpc.CallOption) (*PVCMetricsResponse, error) {
	ah := c.svc.GetAPIHandler("ves.io.schema.virtual_k8s.CustomDataAPI")
	cah, ok := ah.(CustomDataAPIServer)
	if !ok {
		return nil, fmt.Errorf("ah %v is not of type *CustomDataAPISrv", ah)
	}

	var (
		rsp *PVCMetricsResponse
		err error
	)

	bodyFields := svcfw.GenAuditReqBodyFields(ctx, c.svc, "ves.io.schema.virtual_k8s.PVCMetricsRequest", in)
	defer func() {
		if len(bodyFields) > 0 {
			server.ExtendAPIAudit(ctx, svcfw.PublicAPIBodyLog.Uid, bodyFields)
		}
		userMsg := "The 'CustomDataAPI.PVCMetrics' operation on 'virtual_k8s'"
		if err == nil {
			userMsg += " was successfully performed."
		} else {
			userMsg += " failed to be performed."
		}
		server.AddUserMsgToAPIAudit(ctx, userMsg)
	}()

	if c.svc.Config().EnableAPIValidation {
		if rvFn := c.svc.GetRPCValidator("ves.io.schema.virtual_k8s.CustomDataAPI.PVCMetrics"); rvFn != nil {
			if verr := rvFn(ctx, in); verr != nil {
				err = server.MaybePublicRestError(ctx, errors.Wrapf(verr, "Validating Request"))
				return nil, server.GRPCStatusFromError(err).Err()
			}
		}
	}

	rsp, err = cah.PVCMetrics(ctx, in)
	if err != nil {
		return rsp, server.GRPCStatusFromError(server.MaybePublicRestError(ctx, err)).Err()
	}

	bodyFields = append(bodyFields, svcfw.GenAuditRspBodyFields(ctx, c.svc, "ves.io.schema.virtual_k8s.PVCMetricsResponse", rsp)...)

	return rsp, nil
}

func NewCustomDataAPIInprocClient(svc svcfw.Service) CustomDataAPIClient {
	return &CustomDataAPIInprocClient{svc: svc}
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

var CustomDataAPISwaggerJSON string = `{
    "swagger": "2.0",
    "info": {
        "title": "Virtual K8s Object",
        "description": "APIs to get monitoring data for Virtual Kubernetes object.",
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
        "/public/namespaces/{namespace}/virtual_k8s/pvc/metrics": {
            "post": {
                "summary": "PVC Metrics",
                "description": "API to get PVC capacity/usage.",
                "operationId": "ves.io.schema.virtual_k8s.CustomDataAPI.PVCMetrics",
                "responses": {
                    "200": {
                        "description": "",
                        "schema": {
                            "$ref": "#/definitions/virtual_k8sPVCMetricsResponse"
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
                        "in": "path",
                        "required": true,
                        "type": "string"
                    },
                    {
                        "name": "body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/virtual_k8sPVCMetricsRequest"
                        }
                    }
                ],
                "tags": [
                    "CustomDataAPI"
                ],
                "externalDocs": {
                    "description": "Examples of this operation",
                    "url": "https://www.volterra.io/docs/reference/api-ref/ves-io-schema-virtual_k8s-CustomDataAPI-PVCMetrics"
                },
                "x-ves-proto-rpc": "ves.io.schema.virtual_k8s.CustomDataAPI.PVCMetrics"
            },
            "x-displayname": "Custom Data API",
            "x-ves-proto-service": "ves.io.schema.virtual_k8s.CustomDataAPI",
            "x-ves-proto-service-type": "CUSTOM_PUBLIC"
        }
    },
    "definitions": {
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
        },
        "virtual_k8sPVCMetricData": {
            "type": "object",
            "description": "PVC Metric data contains the metric type and the corresponding metric value",
            "title": "PVC Metric Data",
            "x-displayname": "PVC Metric Data",
            "x-ves-proto-message": "ves.io.schema.virtual_k8s.PVCMetricData",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " List of metric data ",
                    "title": "Data",
                    "items": {
                        "$ref": "#/definitions/virtual_k8sPVCMetricTypeData"
                    },
                    "x-displayname": "Data"
                },
                "type": {
                    "description": " PVC Metric Type",
                    "title": "Type",
                    "$ref": "#/definitions/virtual_k8sPVCMetricType",
                    "x-displayname": "Type"
                }
            }
        },
        "virtual_k8sPVCMetricLabel": {
            "type": "string",
            "description": "Labels in PVC metrics\n",
            "title": "PVC Metric Label",
            "enum": [
                "NAMESPACE",
                "PVC",
                "SITE"
            ],
            "default": "NAMESPACE",
            "x-displayname": "PVC Metric Label",
            "x-ves-proto-enum": "ves.io.schema.virtual_k8s.PVCMetricLabel"
        },
        "virtual_k8sPVCMetricType": {
            "type": "string",
            "description": "List of PVC metric types\n\nx-unit: \"bytes\"\nx-unit: \"bytes\"",
            "title": "PVC Metric Type",
            "enum": [
                "PVC_USAGE",
                "PVC_CAPACITY"
            ],
            "default": "PVC_USAGE",
            "x-displayname": "PVC Metric Type",
            "x-ves-proto-enum": "ves.io.schema.virtual_k8s.PVCMetricType"
        },
        "virtual_k8sPVCMetricTypeData": {
            "type": "object",
            "description": "PVC Metric Type Data contains key/value pair that uniquely identifies the group_by labels specified in the request.",
            "title": "PVC Metric Type Data",
            "x-displayname": "PVC Metric Type Data",
            "x-ves-proto-message": "ves.io.schema.virtual_k8s.PVCMetricTypeData",
            "properties": {
                "key": {
                    "type": "object",
                    "description": " Key contains the name/value pair.\n \"name\" is the label name defined in \"PVCMetricLabel\"",
                    "title": "Key",
                    "x-displayname": "Key"
                },
                "value": {
                    "type": "array",
                    "description": " List of metric values.",
                    "title": "Value",
                    "items": {
                        "$ref": "#/definitions/schemaMetricValue"
                    },
                    "x-displayname": "Value"
                }
            }
        },
        "virtual_k8sPVCMetricsRequest": {
            "type": "object",
            "description": "Request to get the metrics for PVC",
            "title": "PVC Metrics Request",
            "x-displayname": "PVC Metrics Request",
            "x-ves-proto-message": "ves.io.schema.virtual_k8s.PVCMetricsRequest",
            "properties": {
                "end_time": {
                    "type": "string",
                    "description": " end time of metric collection from which data will be considered to build graph.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the end_time will be evaluated to start_time+10m\n           If start_time is not specified, then the end_time will be evaluated to \u003ccurrent time\u003e\n\nExample: - \"1570197600\"-",
                    "title": "End time",
                    "x-displayname": "End Time",
                    "x-ves-example": "1570197600"
                },
                "field_selector": {
                    "type": "array",
                    "description": " Select fields to be returned in the response.\n One or more fields in {PVC_USAGE, PVC_CAPACITY} can be specified.\nRequired: YES",
                    "title": "Field Selector",
                    "items": {
                        "$ref": "#/definitions/virtual_k8sPVCMetricType"
                    },
                    "x-displayname": "Field Selector",
                    "x-ves-required": "true"
                },
                "filter": {
                    "type": "string",
                    "description": " filter is used to specify the list of matchers\n syntax for filter := {[\u003cmatcher\u003e]}\n \u003cmatcher\u003e := \u003clabel\u003e\u003coperator\u003e\"\u003cvalue\u003e\"\n   \u003clabel\u003e := string\n     One or more labels defined in MetricLabel {PVC, SITE} can be specified in the filter.\n   \u003cvalue\u003e := string\n   \u003coperator\u003e := [\"=\"|\"!=\"|\"=~\"|\"!~\"]\n     =  : equal to\n     != : not equal to\n     =~ : regex match\n     !~ : not regex match\n\n Optional: If not specified, metric will be aggregated based on the group_by labels. \n\nExample: - \"{PVC=\\\"pvc-1\\\"}\"-",
                    "title": "Label Filter",
                    "x-displayname": "Filter",
                    "x-ves-example": "{PVC=\\\"pvc-1\\\"}"
                },
                "group_by": {
                    "type": "array",
                    "description": " Aggregate data by zero or more labels {NAMESPACE, PVC, SITE}",
                    "title": "Group by",
                    "items": {
                        "$ref": "#/definitions/virtual_k8sPVCMetricLabel"
                    },
                    "x-displayname": "Group By"
                },
                "namespace": {
                    "type": "string",
                    "description": "\n\nExample: - \"ns1\"-",
                    "title": "Namespace",
                    "x-displayname": "Namespace",
                    "x-ves-example": "ns1"
                },
                "range": {
                    "type": "string",
                    "description": " range decides how far to go back in time to fetch values for each step.\n For example, if the range is 5m, then for step t1, query will be evaluated for t1-5m and for\n t2, query will be evaluated for t2-5m and so on.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\n Note: For non-timeseries query, i.e., for step=end_time-start_time, range should be set to end_time-start_time\n\n Optional: If not specified, range is set to 5m\n\nExample: - \"10m\"-",
                    "title": "Range",
                    "x-displayname": "Range",
                    "x-ves-example": "10m"
                },
                "start_time": {
                    "type": "string",
                    "description": "\n start time of metric collection from which data will be considered to build graph.\n Format: unix_timestamp|rfc 3339\n\n Optional: If not specified, then the start_time will be evaluated to end_time-10m\n           If end_time is not specified, then the start_time will be evaluated to \u003ccurrent time\u003e-10m\n\nExample: - \"1570194000\"-",
                    "title": "Start time",
                    "x-displayname": "Start Time",
                    "x-ves-example": "1570194000"
                },
                "step": {
                    "type": "string",
                    "description": " step is the resolution width, which determines the number of the data points [x-axis (time)] to be returned in the response.\n The timestamps in the response will be t1=start_time, t2=t1+step, ... tn=tn-1+step, where tn \u003c= end_time.\n Format: [0-9][smhd], where s - seconds, m - minutes, h - hours, d - days\n\n Optional: If not specified, then step size is evaluated to \u003cend_time - start_time\u003e\n\nExample: - \"5m\"-",
                    "title": "Step",
                    "x-displayname": "Step",
                    "x-ves-example": "5m"
                }
            }
        },
        "virtual_k8sPVCMetricsResponse": {
            "type": "object",
            "description": "Metrics response for PVC",
            "title": "PVC Metrics Response",
            "x-displayname": "PVC Metrics Response",
            "x-ves-proto-message": "ves.io.schema.virtual_k8s.PVCMetricsResponse",
            "properties": {
                "data": {
                    "type": "array",
                    "description": " Data contains time-series metric data for PVC",
                    "title": "Data",
                    "items": {
                        "$ref": "#/definitions/virtual_k8sPVCMetricData"
                    },
                    "x-displayname": "Data"
                }
            }
        }
    },
    "x-displayname": "Virtual Kubernetes",
    "x-ves-proto-file": "ves.io/schema/virtual_k8s/public_custom_data_api.proto"
}`