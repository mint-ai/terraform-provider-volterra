// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package http_loadbalancer

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_virtual_host_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GetDnsInfoRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDnsInfoRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDnsInfoRequest) DeepCopy() *GetDnsInfoRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDnsInfoRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDnsInfoRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDnsInfoRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDnsInfoRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetDnsInfoRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDnsInfoRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDnsInfoRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDnsInfoRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetDnsInfoRequestValidator = func() *ValidateGetDnsInfoRequest {
	v := &ValidateGetDnsInfoRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetDnsInfoRequestValidator() db.Validator {
	return DefaultGetDnsInfoRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetDnsInfoResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDnsInfoResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDnsInfoResponse) DeepCopy() *GetDnsInfoResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDnsInfoResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDnsInfoResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDnsInfoResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDnsInfoResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetDnsInfoResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetDnsInfoDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetDnsInfoResponse) GetDnsInfoDRefInfo() ([]db.DRefInfo, error) {
	if m.GetDnsInfo() == nil {
		return nil, nil
	}

	drInfos, err := m.GetDnsInfo().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetDnsInfo().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "dns_info." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetDnsInfoResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDnsInfoResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDnsInfoResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDnsInfoResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_info"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_info"))
		if err := fv(ctx, m.GetDnsInfo(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetDnsInfoResponseValidator = func() *ValidateGetDnsInfoResponse {
	v := &ValidateGetDnsInfoResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["dns_info"] = ves_io_schema_virtual_host_dns_info.GlobalSpecTypeValidator().Validate

	return v
}()

func GetDnsInfoResponseValidator() db.Validator {
	return DefaultGetDnsInfoResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSecurityConfigReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSecurityConfigReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSecurityConfigReq) DeepCopy() *GetSecurityConfigReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSecurityConfigReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSecurityConfigReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSecurityConfigReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSecurityConfigReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetSecurityConfigReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSecurityConfigReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSecurityConfigReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSecurityConfigReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetLoadbalancerChoice().(type) {
	case *GetSecurityConfigReq_AllHttpLoadbalancers:
		if fv, exists := v.FldValidators["loadbalancer_choice.all_http_loadbalancers"]; exists {
			val := m.GetLoadbalancerChoice().(*GetSecurityConfigReq_AllHttpLoadbalancers).AllHttpLoadbalancers
			vOpts := append(opts,
				db.WithValidateField("loadbalancer_choice"),
				db.WithValidateField("all_http_loadbalancers"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GetSecurityConfigReq_HttpLoadbalancersList:
		if fv, exists := v.FldValidators["loadbalancer_choice.http_loadbalancers_list"]; exists {
			val := m.GetLoadbalancerChoice().(*GetSecurityConfigReq_HttpLoadbalancersList).HttpLoadbalancersList
			vOpts := append(opts,
				db.WithValidateField("loadbalancer_choice"),
				db.WithValidateField("http_loadbalancers_list"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSecurityConfigReqValidator = func() *ValidateGetSecurityConfigReq {
	v := &ValidateGetSecurityConfigReq{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["loadbalancer_choice.http_loadbalancers_list"] = HTTPLoadBalancerListValidator().Validate

	return v
}()

func GetSecurityConfigReqValidator() db.Validator {
	return DefaultGetSecurityConfigReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSecurityConfigRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSecurityConfigRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSecurityConfigRsp) DeepCopy() *GetSecurityConfigRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSecurityConfigRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSecurityConfigRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSecurityConfigRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSecurityConfigRspValidator().Validate(ctx, m, opts...)
}

type ValidateGetSecurityConfigRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSecurityConfigRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSecurityConfigRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSecurityConfigRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_protection"]; exists {

		vOpts := append(opts, db.WithValidateField("api_protection"))
		for idx, item := range m.GetApiProtection() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["app_firewall"]; exists {

		vOpts := append(opts, db.WithValidateField("app_firewall"))
		for idx, item := range m.GetAppFirewall() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["bot_defense"]; exists {

		vOpts := append(opts, db.WithValidateField("bot_defense"))
		for idx, item := range m.GetBotDefense() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["ddos_detection"]; exists {

		vOpts := append(opts, db.WithValidateField("ddos_detection"))
		for idx, item := range m.GetDdosDetection() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["protected"]; exists {

		vOpts := append(opts, db.WithValidateField("protected"))
		for idx, item := range m.GetProtected() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSecurityConfigRspValidator = func() *ValidateGetSecurityConfigRsp {
	v := &ValidateGetSecurityConfigRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetSecurityConfigRspValidator() db.Validator {
	return DefaultGetSecurityConfigRspValidator
}

// augmented methods on protoc/std generated struct

func (m *HTTPLoadBalancerList) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *HTTPLoadBalancerList) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *HTTPLoadBalancerList) DeepCopy() *HTTPLoadBalancerList {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &HTTPLoadBalancerList{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *HTTPLoadBalancerList) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *HTTPLoadBalancerList) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return HTTPLoadBalancerListValidator().Validate(ctx, m, opts...)
}

type ValidateHTTPLoadBalancerList struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateHTTPLoadBalancerList) HttpLoadbalancerValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for http_loadbalancer")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for http_loadbalancer")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated http_loadbalancer")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items http_loadbalancer")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateHTTPLoadBalancerList) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*HTTPLoadBalancerList)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *HTTPLoadBalancerList got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["http_loadbalancer"]; exists {
		vOpts := append(opts, db.WithValidateField("http_loadbalancer"))
		if err := fv(ctx, m.GetHttpLoadbalancer(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultHTTPLoadBalancerListValidator = func() *ValidateHTTPLoadBalancerList {
	v := &ValidateHTTPLoadBalancerList{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhHttpLoadbalancer := v.HttpLoadbalancerValidationRuleHandler
	rulesHttpLoadbalancer := map[string]string{
		"ves.io.schema.rules.message.required":              "true",
		"ves.io.schema.rules.repeated.items.string.max_len": "256",
		"ves.io.schema.rules.repeated.items.string.min_len": "1",
		"ves.io.schema.rules.repeated.max_items":            "5",
		"ves.io.schema.rules.repeated.min_items":            "1",
	}
	vFn, err = vrhHttpLoadbalancer(rulesHttpLoadbalancer)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for HTTPLoadBalancerList.http_loadbalancer: %s", err)
		panic(errMsg)
	}
	v.FldValidators["http_loadbalancer"] = vFn

	return v
}()

func HTTPLoadBalancerListValidator() db.Validator {
	return DefaultHTTPLoadBalancerListValidator
}
