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
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GetApiGroupReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetApiGroupReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetApiGroupReq) DeepCopy() *GetApiGroupReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetApiGroupReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetApiGroupReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetApiGroupReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetApiGroupReqValidator().Validate(ctx, m, opts...)
}

type ValidateGetApiGroupReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetApiGroupReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetApiGroupReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetApiGroupReq got type %s", t)
		}
	}
	if m == nil {
		return nil
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
var DefaultGetApiGroupReqValidator = func() *ValidateGetApiGroupReq {
	v := &ValidateGetApiGroupReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetApiGroupReqValidator() db.Validator {
	return DefaultGetApiGroupReqValidator
}

// augmented methods on protoc/std generated struct

func (m *GetApiGroupRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetApiGroupRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetApiGroupRsp) DeepCopy() *GetApiGroupRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetApiGroupRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetApiGroupRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetApiGroupRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetApiGroupRspValidator().Validate(ctx, m, opts...)
}

func (m *GetApiGroupRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetApiGroupDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetApiGroupRsp) GetApiGroupDRefInfo() ([]db.DRefInfo, error) {
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

type ValidateGetApiGroupRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetApiGroupRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetApiGroupRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetApiGroupRsp got type %s", t)
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
var DefaultGetApiGroupRspValidator = func() *ValidateGetApiGroupRsp {
	v := &ValidateGetApiGroupRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_group"] = CustomApiGroupValidator().Validate

	return v
}()

func GetApiGroupRspValidator() db.Validator {
	return DefaultGetApiGroupRspValidator
}

// augmented methods on protoc/std generated struct

func (m *ListApiGroupsReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListApiGroupsReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListApiGroupsReq) DeepCopy() *ListApiGroupsReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListApiGroupsReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListApiGroupsReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListApiGroupsReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListApiGroupsReqValidator().Validate(ctx, m, opts...)
}

type ValidateListApiGroupsReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListApiGroupsReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListApiGroupsReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListApiGroupsReq got type %s", t)
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

	if fv, exists := v.FldValidators["with_contents"]; exists {

		vOpts := append(opts, db.WithValidateField("with_contents"))
		if err := fv(ctx, m.GetWithContents(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListApiGroupsReqValidator = func() *ValidateListApiGroupsReq {
	v := &ValidateListApiGroupsReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListApiGroupsReqValidator() db.Validator {
	return DefaultListApiGroupsReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ListApiGroupsRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListApiGroupsRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListApiGroupsRsp) DeepCopy() *ListApiGroupsRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListApiGroupsRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListApiGroupsRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListApiGroupsRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListApiGroupsRspValidator().Validate(ctx, m, opts...)
}

func (m *ListApiGroupsRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetApiGroupsContentsDRefInfo()

}

// GetDRefInfo for the field's type
func (m *ListApiGroupsRsp) GetApiGroupsContentsDRefInfo() ([]db.DRefInfo, error) {
	if m.GetApiGroupsContents() == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	for idx, e := range m.GetApiGroupsContents() {
		driSet, err := e.GetDRefInfo()
		if err != nil {
			return nil, errors.Wrap(err, "GetApiGroupsContents() GetDRefInfo() FAILED")
		}
		for i := range driSet {
			dri := &driSet[i]
			dri.DRField = fmt.Sprintf("api_groups_contents[%v].%s", idx, dri.DRField)
		}
		drInfos = append(drInfos, driSet...)
	}
	return drInfos, nil

}

type ValidateListApiGroupsRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListApiGroupsRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListApiGroupsRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListApiGroupsRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["api_groups_contents"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups_contents"))
		for idx, item := range m.GetApiGroupsContents() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["api_groups_names"]; exists {

		vOpts := append(opts, db.WithValidateField("api_groups_names"))
		for idx, item := range m.GetApiGroupsNames() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListApiGroupsRspValidator = func() *ValidateListApiGroupsRsp {
	v := &ValidateListApiGroupsRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_groups_contents"] = CustomApiGroupValidator().Validate

	return v
}()

func ListApiGroupsRspValidator() db.Validator {
	return DefaultListApiGroupsRspValidator
}

// augmented methods on protoc/std generated struct

func (m *UpdateApiGroupReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateApiGroupReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateApiGroupReq) DeepCopy() *UpdateApiGroupReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateApiGroupReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateApiGroupReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateApiGroupReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateApiGroupReqValidator().Validate(ctx, m, opts...)
}

func (m *UpdateApiGroupReq) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetApiGroupDRefInfo()

}

// GetDRefInfo for the field's type
func (m *UpdateApiGroupReq) GetApiGroupDRefInfo() ([]db.DRefInfo, error) {
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

type ValidateUpdateApiGroupReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateApiGroupReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateApiGroupReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateApiGroupReq got type %s", t)
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
var DefaultUpdateApiGroupReqValidator = func() *ValidateUpdateApiGroupReq {
	v := &ValidateUpdateApiGroupReq{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["api_group"] = CustomApiGroupValidator().Validate

	return v
}()

func UpdateApiGroupReqValidator() db.Validator {
	return DefaultUpdateApiGroupReqValidator
}

// augmented methods on protoc/std generated struct

func (m *UpdateApiGroupRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateApiGroupRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateApiGroupRsp) DeepCopy() *UpdateApiGroupRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateApiGroupRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateApiGroupRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateApiGroupRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateApiGroupRspValidator().Validate(ctx, m, opts...)
}

type ValidateUpdateApiGroupRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateApiGroupRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateApiGroupRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateApiGroupRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpdateApiGroupRspValidator = func() *ValidateUpdateApiGroupRsp {
	v := &ValidateUpdateApiGroupRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UpdateApiGroupRspValidator() db.Validator {
	return DefaultUpdateApiGroupRspValidator
}
