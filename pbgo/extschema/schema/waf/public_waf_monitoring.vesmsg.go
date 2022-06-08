//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package waf

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

func (m *RuleHitsCountRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RuleHitsCountRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RuleHitsCountRequest) DeepCopy() *RuleHitsCountRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RuleHitsCountRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RuleHitsCountRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RuleHitsCountRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RuleHitsCountRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRuleHitsCountRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRuleHitsCountRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RuleHitsCountRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RuleHitsCountRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["group_by"]; exists {

		vOpts := append(opts, db.WithValidateField("group_by"))
		for idx, item := range m.GetGroupBy() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		for idx, item := range m.GetLabelFilter() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
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

	if fv, exists := v.FldValidators["range"]; exists {

		vOpts := append(opts, db.WithValidateField("range"))
		if err := fv(ctx, m.GetRange(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRuleHitsCountRequestValidator = func() *ValidateRuleHitsCountRequest {
	v := &ValidateRuleHitsCountRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RuleHitsCountRequestValidator() db.Validator {
	return DefaultRuleHitsCountRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *RuleHitsCountResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RuleHitsCountResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RuleHitsCountResponse) DeepCopy() *RuleHitsCountResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RuleHitsCountResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RuleHitsCountResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RuleHitsCountResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RuleHitsCountResponseValidator().Validate(ctx, m, opts...)
}

type ValidateRuleHitsCountResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRuleHitsCountResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RuleHitsCountResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RuleHitsCountResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRuleHitsCountResponseValidator = func() *ValidateRuleHitsCountResponse {
	v := &ValidateRuleHitsCountResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RuleHitsCountResponseValidator() db.Validator {
	return DefaultRuleHitsCountResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SecurityEventsCountRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SecurityEventsCountRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SecurityEventsCountRequest) DeepCopy() *SecurityEventsCountRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SecurityEventsCountRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SecurityEventsCountRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SecurityEventsCountRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SecurityEventsCountRequestValidator().Validate(ctx, m, opts...)
}

type ValidateSecurityEventsCountRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSecurityEventsCountRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SecurityEventsCountRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SecurityEventsCountRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["end_time"]; exists {

		vOpts := append(opts, db.WithValidateField("end_time"))
		if err := fv(ctx, m.GetEndTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["group_by"]; exists {

		vOpts := append(opts, db.WithValidateField("group_by"))
		for idx, item := range m.GetGroupBy() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["label_filter"]; exists {

		vOpts := append(opts, db.WithValidateField("label_filter"))
		for idx, item := range m.GetLabelFilter() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
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

	if fv, exists := v.FldValidators["range"]; exists {

		vOpts := append(opts, db.WithValidateField("range"))
		if err := fv(ctx, m.GetRange(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["step"]; exists {

		vOpts := append(opts, db.WithValidateField("step"))
		if err := fv(ctx, m.GetStep(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSecurityEventsCountRequestValidator = func() *ValidateSecurityEventsCountRequest {
	v := &ValidateSecurityEventsCountRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SecurityEventsCountRequestValidator() db.Validator {
	return DefaultSecurityEventsCountRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *SecurityEventsCountResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SecurityEventsCountResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SecurityEventsCountResponse) DeepCopy() *SecurityEventsCountResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SecurityEventsCountResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SecurityEventsCountResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SecurityEventsCountResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SecurityEventsCountResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSecurityEventsCountResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSecurityEventsCountResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SecurityEventsCountResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SecurityEventsCountResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["data"]; exists {

		vOpts := append(opts, db.WithValidateField("data"))
		for idx, item := range m.GetData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSecurityEventsCountResponseValidator = func() *ValidateSecurityEventsCountResponse {
	v := &ValidateSecurityEventsCountResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SecurityEventsCountResponseValidator() db.Validator {
	return DefaultSecurityEventsCountResponseValidator
}
