//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package api_inventory

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_views_api_definition "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/api_definition"
	ves_io_schema_views_app_api_group "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/app_api_group"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *EvaluateApiGroupReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateApiGroupReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateApiGroupReq) DeepCopy() *EvaluateApiGroupReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateApiGroupReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateApiGroupReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateApiGroupReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateApiGroupReqValidator().Validate(ctx, m, opts...)
}

type ValidateEvaluateApiGroupReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateApiGroupReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateApiGroupReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateApiGroupReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_group_builder"]; exists {

		vOpts := append(opts, db.WithValidateField("api_group_builder"))
		if err := fv(ctx, m.GetApiGroupBuilder(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["group_name"]; exists {

		vOpts := append(opts, db.WithValidateField("group_name"))
		if err := fv(ctx, m.GetGroupName(), vOpts...); err != nil {
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateApiGroupReqValidator = func() *ValidateEvaluateApiGroupReq {
	v := &ValidateEvaluateApiGroupReq{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_group_builder"] = ves_io_schema_views_api_definition.ApiGroupBuilderValidator().Validate

	return v
}()

func EvaluateApiGroupReqValidator() db.Validator {
	return DefaultEvaluateApiGroupReqValidator
}

// augmented methods on protoc/std generated struct

func (m *EvaluateApiGroupRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EvaluateApiGroupRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EvaluateApiGroupRsp) DeepCopy() *EvaluateApiGroupRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EvaluateApiGroupRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EvaluateApiGroupRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EvaluateApiGroupRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EvaluateApiGroupRspValidator().Validate(ctx, m, opts...)
}

func (m *EvaluateApiGroupRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetApiGroupDRefInfo()

}

// GetDRefInfo for the field's type
func (m *EvaluateApiGroupRsp) GetApiGroupDRefInfo() ([]db.DRefInfo, error) {
	if m.GetApiGroup() == nil {
		return nil, nil
	}

	drInfos, err := m.GetApiGroup().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetApiGroup().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "api_group." + dri.DRField
	}
	return drInfos, err

}

type ValidateEvaluateApiGroupRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEvaluateApiGroupRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EvaluateApiGroupRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EvaluateApiGroupRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_group"]; exists {

		vOpts := append(opts, db.WithValidateField("api_group"))
		if err := fv(ctx, m.GetApiGroup(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEvaluateApiGroupRspValidator = func() *ValidateEvaluateApiGroupRsp {
	v := &ValidateEvaluateApiGroupRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_group"] = ves_io_schema_views_app_api_group.GlobalSpecTypeValidator().Validate

	return v
}()

func EvaluateApiGroupRspValidator() db.Validator {
	return DefaultEvaluateApiGroupRspValidator
}

// augmented methods on protoc/std generated struct

func (m *GetContentReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetContentReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetContentReq) DeepCopy() *GetContentReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetContentReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetContentReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetContentReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetContentReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetContentReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetContentReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetContentReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetContentReq got type %s", t)
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
var DefaultGetContentReqValidator = func() *ValidateGetContentReq {
	v := &ValidateGetContentReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetContentReqValidator() db.Validator {
	return DefaultGetContentReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetContentRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetContentRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetContentRsp) DeepCopy() *GetContentRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetContentRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetContentRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetContentRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetContentRspValidator().Validate(ctx, m, opts...)
}

type ValidateGetContentRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetContentRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetContentRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetContentRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_endpoints"]; exists {

		vOpts := append(opts, db.WithValidateField("api_endpoints"))
		for idx, item := range m.GetApiEndpoints() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetContentRspValidator = func() *ValidateGetContentRsp {
	v := &ValidateGetContentRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_endpoints"] = ApiEndpointValidator().Validate

	return v
}()

func GetContentRspValidator() db.Validator {
	return DefaultGetContentRspValidator
}
