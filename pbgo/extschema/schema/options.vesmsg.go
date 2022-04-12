//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package schema

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

func (m *Dependencies) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Dependencies) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Dependencies) DeepCopy() *Dependencies {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Dependencies{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Dependencies) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Dependencies) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DependenciesValidator().Validate(ctx, m, opts...)
}

type ValidateDependencies struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDependencies) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Dependencies)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Dependencies got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["on"]; exists {

		vOpts := append(opts, db.WithValidateField("on"))
		for idx, item := range m.GetOn() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDependenciesValidator = func() *ValidateDependencies {
	v := &ValidateDependencies{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DependenciesValidator() db.Validator {
	return DefaultDependenciesValidator
}

// augmented methods on protoc/std generated struct

func (m *Key) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Key) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Key) DeepCopy() *Key {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Key{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Key) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Key) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return KeyValidator().Validate(ctx, m, opts...)
}

type ValidateKey struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateKey) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Key)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Key got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["field_path"]; exists {

		vOpts := append(opts, db.WithValidateField("field_path"))
		if err := fv(ctx, m.GetFieldPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultKeyValidator = func() *ValidateKey {
	v := &ValidateKey{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func KeyValidator() db.Validator {
	return DefaultKeyValidator
}

// augmented methods on protoc/std generated struct

func (m *Keys) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Keys) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Keys) DeepCopy() *Keys {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Keys{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Keys) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Keys) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return KeysValidator().Validate(ctx, m, opts...)
}

type ValidateKeys struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateKeys) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Keys)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Keys got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["keys"]; exists {

		vOpts := append(opts, db.WithValidateField("keys"))
		for idx, item := range m.GetKeys() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultKeysValidator = func() *ValidateKeys {
	v := &ValidateKeys{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func KeysValidator() db.Validator {
	return DefaultKeysValidator
}

// augmented methods on protoc/std generated struct

func (m *MetricDef) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *MetricDef) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *MetricDef) DeepCopy() *MetricDef {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &MetricDef{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *MetricDef) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *MetricDef) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return MetricDefValidator().Validate(ctx, m, opts...)
}

type ValidateMetricDef struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateMetricDef) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*MetricDef)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *MetricDef got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["buckets"]; exists {

		vOpts := append(opts, db.WithValidateField("buckets"))
		for idx, item := range m.GetBuckets() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["labels"]; exists {

		vOpts := append(opts, db.WithValidateField("labels"))
		for idx, item := range m.GetLabels() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultMetricDefValidator = func() *ValidateMetricDef {
	v := &ValidateMetricDef{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func MetricDefValidator() db.Validator {
	return DefaultMetricDefValidator
}

// augmented methods on protoc/std generated struct

func (m *On) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *On) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *On) DeepCopy() *On {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &On{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *On) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *On) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OnValidator().Validate(ctx, m, opts...)
}

type ValidateOn struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOn) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*On)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *On got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["via"]; exists {

		vOpts := append(opts, db.WithValidateField("via"))
		for idx, item := range m.GetVia() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOnValidator = func() *ValidateOn {
	v := &ValidateOn{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func OnValidator() db.Validator {
	return DefaultOnValidator
}

// augmented methods on protoc/std generated struct

func (m *Via) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Via) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Via) DeepCopy() *Via {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Via{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Via) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Via) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ViaValidator().Validate(ctx, m, opts...)
}

type ValidateVia struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateVia) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Via)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Via got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["field_path"]; exists {

		vOpts := append(opts, db.WithValidateField("field_path"))
		if err := fv(ctx, m.GetFieldPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["gen_field_path"]; exists {

		vOpts := append(opts, db.WithValidateField("gen_field_path"))
		if err := fv(ctx, m.GetGenFieldPath(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ref_type"]; exists {

		vOpts := append(opts, db.WithValidateField("ref_type"))
		if err := fv(ctx, m.GetRefType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultViaValidator = func() *ValidateVia {
	v := &ValidateVia{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ViaValidator() db.Validator {
	return DefaultViaValidator
}
