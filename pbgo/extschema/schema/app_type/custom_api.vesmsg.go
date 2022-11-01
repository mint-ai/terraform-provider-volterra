// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package app_type

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *APIEndpointLearntSchemaReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIEndpointLearntSchemaReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIEndpointLearntSchemaReq) DeepCopy() *APIEndpointLearntSchemaReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIEndpointLearntSchemaReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIEndpointLearntSchemaReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIEndpointLearntSchemaReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIEndpointLearntSchemaReqValidator().Validate(ctx, m, opts...)
}

type ValidateAPIEndpointLearntSchemaReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIEndpointLearntSchemaReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIEndpointLearntSchemaReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIEndpointLearntSchemaReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["collapsed_url"]; exists {

		vOpts := append(opts, db.WithValidateField("collapsed_url"))
		if err := fv(ctx, m.GetCollapsedUrl(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
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
var DefaultAPIEndpointLearntSchemaReqValidator = func() *ValidateAPIEndpointLearntSchemaReq {
	v := &ValidateAPIEndpointLearntSchemaReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIEndpointLearntSchemaReqValidator() db.Validator {
	return DefaultAPIEndpointLearntSchemaReqValidator
}

// augmented methods on protoc/std generated struct

func (m *APIEndpointLearntSchemaRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIEndpointLearntSchemaRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIEndpointLearntSchemaRsp) DeepCopy() *APIEndpointLearntSchemaRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIEndpointLearntSchemaRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIEndpointLearntSchemaRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIEndpointLearntSchemaRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIEndpointLearntSchemaRspValidator().Validate(ctx, m, opts...)
}

type ValidateAPIEndpointLearntSchemaRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIEndpointLearntSchemaRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIEndpointLearntSchemaRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIEndpointLearntSchemaRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["last_updated_time"]; exists {

		vOpts := append(opts, db.WithValidateField("last_updated_time"))
		if err := fv(ctx, m.GetLastUpdatedTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["learnt_schema"]; exists {

		vOpts := append(opts, db.WithValidateField("learnt_schema"))
		if err := fv(ctx, m.GetLearntSchema(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["swagger_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("swagger_spec"))
		if err := fv(ctx, m.GetSwaggerSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIEndpointLearntSchemaRspValidator = func() *ValidateAPIEndpointLearntSchemaRsp {
	v := &ValidateAPIEndpointLearntSchemaRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIEndpointLearntSchemaRspValidator() db.Validator {
	return DefaultAPIEndpointLearntSchemaRspValidator
}

// augmented methods on protoc/std generated struct

func (m *APIEndpointPDFReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIEndpointPDFReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIEndpointPDFReq) DeepCopy() *APIEndpointPDFReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIEndpointPDFReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIEndpointPDFReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIEndpointPDFReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIEndpointPDFReqValidator().Validate(ctx, m, opts...)
}

type ValidateAPIEndpointPDFReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIEndpointPDFReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIEndpointPDFReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIEndpointPDFReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["collapsed_url"]; exists {

		vOpts := append(opts, db.WithValidateField("collapsed_url"))
		if err := fv(ctx, m.GetCollapsedUrl(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
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
var DefaultAPIEndpointPDFReqValidator = func() *ValidateAPIEndpointPDFReq {
	v := &ValidateAPIEndpointPDFReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIEndpointPDFReqValidator() db.Validator {
	return DefaultAPIEndpointPDFReqValidator
}

// augmented methods on protoc/std generated struct

func (m *APIEndpointPDFRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIEndpointPDFRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIEndpointPDFRsp) DeepCopy() *APIEndpointPDFRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIEndpointPDFRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIEndpointPDFRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIEndpointPDFRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIEndpointPDFRspValidator().Validate(ctx, m, opts...)
}

type ValidateAPIEndpointPDFRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIEndpointPDFRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIEndpointPDFRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIEndpointPDFRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["pdf_info"]; exists {

		vOpts := append(opts, db.WithValidateField("pdf_info"))
		if err := fv(ctx, m.GetPdfInfo(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIEndpointPDFRspValidator = func() *ValidateAPIEndpointPDFRsp {
	v := &ValidateAPIEndpointPDFRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIEndpointPDFRspValidator() db.Validator {
	return DefaultAPIEndpointPDFRspValidator
}

// augmented methods on protoc/std generated struct

func (m *APIEndpointsReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIEndpointsReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIEndpointsReq) DeepCopy() *APIEndpointsReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIEndpointsReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIEndpointsReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIEndpointsReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIEndpointsReqValidator().Validate(ctx, m, opts...)
}

type ValidateAPIEndpointsReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIEndpointsReq) ApiEndpointInfoRequestValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ApiEndpointInfoRequest)
		return int32(i)
	}
	// ApiEndpointInfoRequest_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ApiEndpointInfoRequest_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for api_endpoint_info_request")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ApiEndpointInfoRequest, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for api_endpoint_info_request")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ApiEndpointInfoRequest)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ApiEndpointInfoRequest, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated api_endpoint_info_request")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items api_endpoint_info_request")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateAPIEndpointsReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIEndpointsReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIEndpointsReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoint_info_request"]; exists {
		vOpts := append(opts, db.WithValidateField("api_endpoint_info_request"))
		if err := fv(ctx, m.GetApiEndpointInfoRequest(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
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
var DefaultAPIEndpointsReqValidator = func() *ValidateAPIEndpointsReq {
	v := &ValidateAPIEndpointsReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhApiEndpointInfoRequest := v.ApiEndpointInfoRequestValidationRuleHandler
	rulesApiEndpointInfoRequest := map[string]string{
		"ves.io.schema.rules.repeated.unique": "true",
	}
	vFn, err = vrhApiEndpointInfoRequest(rulesApiEndpointInfoRequest)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for APIEndpointsReq.api_endpoint_info_request: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_endpoint_info_request"] = vFn

	return v
}()

func APIEndpointsReqValidator() db.Validator {
	return DefaultAPIEndpointsReqValidator
}

// augmented methods on protoc/std generated struct

func (m *APIEndpointsRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *APIEndpointsRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *APIEndpointsRsp) DeepCopy() *APIEndpointsRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &APIEndpointsRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *APIEndpointsRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *APIEndpointsRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return APIEndpointsRspValidator().Validate(ctx, m, opts...)
}

type ValidateAPIEndpointsRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAPIEndpointsRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*APIEndpointsRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *APIEndpointsRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["apiep_list"]; exists {

		vOpts := append(opts, db.WithValidateField("apiep_list"))
		for idx, item := range m.GetApiepList() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAPIEndpointsRspValidator = func() *ValidateAPIEndpointsRsp {
	v := &ValidateAPIEndpointsRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func APIEndpointsRspValidator() db.Validator {
	return DefaultAPIEndpointsRspValidator
}

// augmented methods on protoc/std generated struct

func (m *OverridePopReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OverridePopReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OverridePopReq) DeepCopy() *OverridePopReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OverridePopReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OverridePopReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OverridePopReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OverridePopReqValidator().Validate(ctx, m, opts...)
}

type ValidateOverridePopReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOverridePopReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OverridePopReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OverridePopReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
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
var DefaultOverridePopReqValidator = func() *ValidateOverridePopReq {
	v := &ValidateOverridePopReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OverridePopReqValidator() db.Validator {
	return DefaultOverridePopReqValidator
}

// augmented methods on protoc/std generated struct

func (m *OverridePopRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OverridePopRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OverridePopRsp) DeepCopy() *OverridePopRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OverridePopRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OverridePopRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OverridePopRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OverridePopRspValidator().Validate(ctx, m, opts...)
}

type ValidateOverridePopRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOverridePopRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OverridePopRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OverridePopRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status_msg"]; exists {

		vOpts := append(opts, db.WithValidateField("status_msg"))
		if err := fv(ctx, m.GetStatusMsg(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOverridePopRspValidator = func() *ValidateOverridePopRsp {
	v := &ValidateOverridePopRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OverridePopRspValidator() db.Validator {
	return DefaultOverridePopRspValidator
}

// augmented methods on protoc/std generated struct

func (m *OverridePushReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OverridePushReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OverridePushReq) DeepCopy() *OverridePushReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OverridePushReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OverridePushReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OverridePushReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OverridePushReqValidator().Validate(ctx, m, opts...)
}

type ValidateOverridePushReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOverridePushReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OverridePushReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OverridePushReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["override_info"]; exists {

		vOpts := append(opts, db.WithValidateField("override_info"))
		if err := fv(ctx, m.GetOverrideInfo(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOverridePushReqValidator = func() *ValidateOverridePushReq {
	v := &ValidateOverridePushReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OverridePushReqValidator() db.Validator {
	return DefaultOverridePushReqValidator
}

// augmented methods on protoc/std generated struct

func (m *OverridePushRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OverridePushRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OverridePushRsp) DeepCopy() *OverridePushRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OverridePushRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OverridePushRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OverridePushRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OverridePushRspValidator().Validate(ctx, m, opts...)
}

type ValidateOverridePushRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOverridePushRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OverridePushRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OverridePushRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status_msg"]; exists {

		vOpts := append(opts, db.WithValidateField("status_msg"))
		if err := fv(ctx, m.GetStatusMsg(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOverridePushRspValidator = func() *ValidateOverridePushRsp {
	v := &ValidateOverridePushRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OverridePushRspValidator() db.Validator {
	return DefaultOverridePushRspValidator
}

// augmented methods on protoc/std generated struct

func (m *OverridesReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OverridesReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OverridesReq) DeepCopy() *OverridesReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OverridesReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OverridesReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OverridesReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OverridesReqValidator().Validate(ctx, m, opts...)
}

type ValidateOverridesReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOverridesReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OverridesReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OverridesReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
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
var DefaultOverridesReqValidator = func() *ValidateOverridesReq {
	v := &ValidateOverridesReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OverridesReqValidator() db.Validator {
	return DefaultOverridesReqValidator
}

// augmented methods on protoc/std generated struct

func (m *OverridesRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *OverridesRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *OverridesRsp) DeepCopy() *OverridesRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &OverridesRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *OverridesRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *OverridesRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OverridesRspValidator().Validate(ctx, m, opts...)
}

type ValidateOverridesRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOverridesRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*OverridesRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *OverridesRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["override_list"]; exists {

		vOpts := append(opts, db.WithValidateField("override_list"))
		for idx, item := range m.GetOverrideList() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOverridesRspValidator = func() *ValidateOverridesRsp {
	v := &ValidateOverridesRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OverridesRspValidator() db.Validator {
	return DefaultOverridesRspValidator
}

// augmented methods on protoc/std generated struct

func (m *ServiceAPIEndpointPDFReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ServiceAPIEndpointPDFReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ServiceAPIEndpointPDFReq) DeepCopy() *ServiceAPIEndpointPDFReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ServiceAPIEndpointPDFReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ServiceAPIEndpointPDFReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ServiceAPIEndpointPDFReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ServiceAPIEndpointPDFReqValidator().Validate(ctx, m, opts...)
}

type ValidateServiceAPIEndpointPDFReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateServiceAPIEndpointPDFReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ServiceAPIEndpointPDFReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ServiceAPIEndpointPDFReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["collapsed_url"]; exists {

		vOpts := append(opts, db.WithValidateField("collapsed_url"))
		if err := fv(ctx, m.GetCollapsedUrl(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["method"]; exists {

		vOpts := append(opts, db.WithValidateField("method"))
		if err := fv(ctx, m.GetMethod(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service_name"]; exists {

		vOpts := append(opts, db.WithValidateField("service_name"))
		if err := fv(ctx, m.GetServiceName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultServiceAPIEndpointPDFReqValidator = func() *ValidateServiceAPIEndpointPDFReq {
	v := &ValidateServiceAPIEndpointPDFReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ServiceAPIEndpointPDFReqValidator() db.Validator {
	return DefaultServiceAPIEndpointPDFReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ServiceAPIEndpointsReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ServiceAPIEndpointsReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ServiceAPIEndpointsReq) DeepCopy() *ServiceAPIEndpointsReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ServiceAPIEndpointsReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ServiceAPIEndpointsReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ServiceAPIEndpointsReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ServiceAPIEndpointsReqValidator().Validate(ctx, m, opts...)
}

type ValidateServiceAPIEndpointsReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateServiceAPIEndpointsReq) ApiEndpointInfoRequestValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ApiEndpointInfoRequest)
		return int32(i)
	}
	// ApiEndpointInfoRequest_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ApiEndpointInfoRequest_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for api_endpoint_info_request")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ApiEndpointInfoRequest, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for api_endpoint_info_request")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ApiEndpointInfoRequest)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ApiEndpointInfoRequest, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated api_endpoint_info_request")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items api_endpoint_info_request")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateServiceAPIEndpointsReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ServiceAPIEndpointsReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ServiceAPIEndpointsReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoint_info_request"]; exists {
		vOpts := append(opts, db.WithValidateField("api_endpoint_info_request"))
		if err := fv(ctx, m.GetApiEndpointInfoRequest(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service_name"]; exists {

		vOpts := append(opts, db.WithValidateField("service_name"))
		if err := fv(ctx, m.GetServiceName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultServiceAPIEndpointsReqValidator = func() *ValidateServiceAPIEndpointsReq {
	v := &ValidateServiceAPIEndpointsReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhApiEndpointInfoRequest := v.ApiEndpointInfoRequestValidationRuleHandler
	rulesApiEndpointInfoRequest := map[string]string{
		"ves.io.schema.rules.repeated.unique": "true",
	}
	vFn, err = vrhApiEndpointInfoRequest(rulesApiEndpointInfoRequest)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ServiceAPIEndpointsReq.api_endpoint_info_request: %s", err)
		panic(errMsg)
	}
	v.FldValidators["api_endpoint_info_request"] = vFn

	return v
}()

func ServiceAPIEndpointsReqValidator() db.Validator {
	return DefaultServiceAPIEndpointsReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SwaggerSpecReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SwaggerSpecReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SwaggerSpecReq) DeepCopy() *SwaggerSpecReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SwaggerSpecReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SwaggerSpecReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SwaggerSpecReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SwaggerSpecReqValidator().Validate(ctx, m, opts...)
}

type ValidateSwaggerSpecReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSwaggerSpecReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SwaggerSpecReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SwaggerSpecReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["app_type_name"]; exists {

		vOpts := append(opts, db.WithValidateField("app_type_name"))
		if err := fv(ctx, m.GetAppTypeName(), vOpts...); err != nil {
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
var DefaultSwaggerSpecReqValidator = func() *ValidateSwaggerSpecReq {
	v := &ValidateSwaggerSpecReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SwaggerSpecReqValidator() db.Validator {
	return DefaultSwaggerSpecReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SwaggerSpecRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SwaggerSpecRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SwaggerSpecRsp) DeepCopy() *SwaggerSpecRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SwaggerSpecRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SwaggerSpecRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SwaggerSpecRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SwaggerSpecRspValidator().Validate(ctx, m, opts...)
}

type ValidateSwaggerSpecRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSwaggerSpecRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SwaggerSpecRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SwaggerSpecRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["swagger_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("swagger_spec"))
		if err := fv(ctx, m.GetSwaggerSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSwaggerSpecRspValidator = func() *ValidateSwaggerSpecRsp {
	v := &ValidateSwaggerSpecRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SwaggerSpecRspValidator() db.Validator {
	return DefaultSwaggerSpecRspValidator
}
