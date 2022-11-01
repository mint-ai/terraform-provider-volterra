// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package aws_tgw_site

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *SetTGWInfoRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetTGWInfoRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetTGWInfoRequest) DeepCopy() *SetTGWInfoRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetTGWInfoRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetTGWInfoRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetTGWInfoRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetTGWInfoRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetTGWInfoRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetTGWInfoRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetTGWInfoRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetTGWInfoRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["direct_connect_info"]; exists {

		vOpts := append(opts, db.WithValidateField("direct_connect_info"))
		if err := fv(ctx, m.GetDirectConnectInfo(), vOpts...); err != nil {
			return err
		}

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

	if fv, exists := v.FldValidators["tgw_info"]; exists {

		vOpts := append(opts, db.WithValidateField("tgw_info"))
		if err := fv(ctx, m.GetTgwInfo(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetTGWInfoRequestValidator = func() *ValidateSetTGWInfoRequest {
	v := &ValidateSetTGWInfoRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tgw_info"] = AWSTGWInfoConfigTypeValidator().Validate

	v.FldValidators["direct_connect_info"] = ves_io_schema_views.DirectConnectInfoValidator().Validate

	return v
}()

func SetTGWInfoRequestValidator() db.Validator {
	return DefaultSetTGWInfoRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetTGWInfoResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetTGWInfoResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetTGWInfoResponse) DeepCopy() *SetTGWInfoResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetTGWInfoResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetTGWInfoResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetTGWInfoResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetTGWInfoResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetTGWInfoResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetTGWInfoResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetTGWInfoResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetTGWInfoResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetTGWInfoResponseValidator = func() *ValidateSetTGWInfoResponse {
	v := &ValidateSetTGWInfoResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetTGWInfoResponseValidator() db.Validator {
	return DefaultSetTGWInfoResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVIPInfoRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVIPInfoRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVIPInfoRequest) DeepCopy() *SetVIPInfoRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVIPInfoRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVIPInfoRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVIPInfoRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVIPInfoRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetVIPInfoRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVIPInfoRequest) VipParamsPerAzValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for vip_params_per_az")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_site.PublishVIPParamsPerAz, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema_site.PublishVIPParamsPerAzValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vip_params_per_az")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_site.PublishVIPParamsPerAz)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_site.PublishVIPParamsPerAz, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated vip_params_per_az")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vip_params_per_az")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateSetVIPInfoRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVIPInfoRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVIPInfoRequest got type %s", t)
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

	if fv, exists := v.FldValidators["vip_params_per_az"]; exists {
		vOpts := append(opts, db.WithValidateField("vip_params_per_az"))
		if err := fv(ctx, m.GetVipParamsPerAz(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVIPInfoRequestValidator = func() *ValidateSetVIPInfoRequest {
	v := &ValidateSetVIPInfoRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhVipParamsPerAz := v.VipParamsPerAzValidationRuleHandler
	rulesVipParamsPerAz := map[string]string{
		"ves.io.schema.rules.repeated.num_items": "1,2,3",
	}
	vFn, err = vrhVipParamsPerAz(rulesVipParamsPerAz)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SetVIPInfoRequest.vip_params_per_az: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vip_params_per_az"] = vFn

	return v
}()

func SetVIPInfoRequestValidator() db.Validator {
	return DefaultSetVIPInfoRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVIPInfoResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVIPInfoResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVIPInfoResponse) DeepCopy() *SetVIPInfoResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVIPInfoResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVIPInfoResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVIPInfoResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVIPInfoResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetVIPInfoResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVIPInfoResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVIPInfoResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVIPInfoResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVIPInfoResponseValidator = func() *ValidateSetVIPInfoResponse {
	v := &ValidateSetVIPInfoResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVIPInfoResponseValidator() db.Validator {
	return DefaultSetVIPInfoResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPCIpPrefixesRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPCIpPrefixesRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPCIpPrefixesRequest) DeepCopy() *SetVPCIpPrefixesRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPCIpPrefixesRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPCIpPrefixesRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPCIpPrefixesRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPCIpPrefixesRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPCIpPrefixesRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPCIpPrefixesRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPCIpPrefixesRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPCIpPrefixesRequest got type %s", t)
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

	if fv, exists := v.FldValidators["vpc_ip_prefixes"]; exists {

		vOpts := append(opts, db.WithValidateField("vpc_ip_prefixes"))
		for key, value := range m.GetVpcIpPrefixes() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPCIpPrefixesRequestValidator = func() *ValidateSetVPCIpPrefixesRequest {
	v := &ValidateSetVPCIpPrefixesRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["vpc_ip_prefixes"] = VPCIpPrefixesTypeValidator().Validate

	return v
}()

func SetVPCIpPrefixesRequestValidator() db.Validator {
	return DefaultSetVPCIpPrefixesRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPCIpPrefixesResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPCIpPrefixesResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPCIpPrefixesResponse) DeepCopy() *SetVPCIpPrefixesResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPCIpPrefixesResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPCIpPrefixesResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPCIpPrefixesResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPCIpPrefixesResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPCIpPrefixesResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPCIpPrefixesResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPCIpPrefixesResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPCIpPrefixesResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPCIpPrefixesResponseValidator = func() *ValidateSetVPCIpPrefixesResponse {
	v := &ValidateSetVPCIpPrefixesResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVPCIpPrefixesResponseValidator() db.Validator {
	return DefaultSetVPCIpPrefixesResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPNTunnelsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPNTunnelsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPNTunnelsRequest) DeepCopy() *SetVPNTunnelsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPNTunnelsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPNTunnelsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPNTunnelsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPNTunnelsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPNTunnelsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPNTunnelsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPNTunnelsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPNTunnelsRequest got type %s", t)
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

	if fv, exists := v.FldValidators["tunnels"]; exists {

		vOpts := append(opts, db.WithValidateField("tunnels"))
		for idx, item := range m.GetTunnels() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPNTunnelsRequestValidator = func() *ValidateSetVPNTunnelsRequest {
	v := &ValidateSetVPNTunnelsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tunnels"] = AWSVPNTunnelConfigTypeValidator().Validate

	return v
}()

func SetVPNTunnelsRequestValidator() db.Validator {
	return DefaultSetVPNTunnelsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SetVPNTunnelsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetVPNTunnelsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetVPNTunnelsResponse) DeepCopy() *SetVPNTunnelsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetVPNTunnelsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetVPNTunnelsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetVPNTunnelsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetVPNTunnelsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSetVPNTunnelsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetVPNTunnelsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetVPNTunnelsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetVPNTunnelsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetVPNTunnelsResponseValidator = func() *ValidateSetVPNTunnelsResponse {
	v := &ValidateSetVPNTunnelsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetVPNTunnelsResponseValidator() db.Validator {
	return DefaultSetVPNTunnelsResponseValidator
}
