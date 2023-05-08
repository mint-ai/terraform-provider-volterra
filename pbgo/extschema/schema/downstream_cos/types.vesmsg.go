// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package downstream_cos

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cos_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("cos_limit"))
		if err := fv(ctx, m.GetCosLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_limit"))
		if err := fv(ctx, m.GetTenantLimit(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tenant_limit"] = PerCpuUtilizationLimitValidator().Validate

	v.FldValidators["cos_limit"] = PerCpuUtilizationLimitValidator().Validate

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cos_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("cos_limit"))
		if err := fv(ctx, m.GetCosLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_limit"))
		if err := fv(ctx, m.GetTenantLimit(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tenant_limit"] = PerCpuUtilizationLimitValidator().Validate

	v.FldValidators["cos_limit"] = PerCpuUtilizationLimitValidator().Validate

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cos_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("cos_limit"))
		if err := fv(ctx, m.GetCosLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_limit"))
		if err := fv(ctx, m.GetTenantLimit(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tenant_limit"] = PerCpuUtilizationLimitValidator().Validate

	v.FldValidators["cos_limit"] = PerCpuUtilizationLimitValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *PerCpuUtilizationLimit) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *PerCpuUtilizationLimit) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *PerCpuUtilizationLimit) DeepCopy() *PerCpuUtilizationLimit {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &PerCpuUtilizationLimit{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *PerCpuUtilizationLimit) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *PerCpuUtilizationLimit) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return PerCpuUtilizationLimitValidator().Validate(ctx, m, opts...)
}

type ValidatePerCpuUtilizationLimit struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidatePerCpuUtilizationLimit) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*PerCpuUtilizationLimit)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *PerCpuUtilizationLimit got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["hard_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("hard_limit"))
		if err := fv(ctx, m.GetHardLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["soft_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("soft_limit"))
		if err := fv(ctx, m.GetSoftLimit(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultPerCpuUtilizationLimitValidator = func() *ValidatePerCpuUtilizationLimit {
	v := &ValidatePerCpuUtilizationLimit{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["soft_limit"] = ves_io_schema.FractionalPercentValidator().Validate

	v.FldValidators["hard_limit"] = ves_io_schema.FractionalPercentValidator().Validate

	return v
}()

func PerCpuUtilizationLimitValidator() db.Validator {
	return DefaultPerCpuUtilizationLimitValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cos_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("cos_limit"))
		if err := fv(ctx, m.GetCosLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_limit"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_limit"))
		if err := fv(ctx, m.GetTenantLimit(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["tenant_limit"] = PerCpuUtilizationLimitValidator().Validate

	v.FldValidators["cos_limit"] = PerCpuUtilizationLimitValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.CosLimit = f.GetCosLimit()
	m.TenantLimit = f.GetTenantLimit()
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *CreateSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *CreateSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.CosLimit = m1.CosLimit
	f.TenantLimit = m1.TenantLimit
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *CreateSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.CosLimit = f.GetCosLimit()
	m.TenantLimit = f.GetTenantLimit()
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *GetSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *GetSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.CosLimit = m1.CosLimit
	f.TenantLimit = m1.TenantLimit
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.CosLimit = f.GetCosLimit()
	m.TenantLimit = f.GetTenantLimit()
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.CosLimit = m1.CosLimit
	f.TenantLimit = m1.TenantLimit
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
