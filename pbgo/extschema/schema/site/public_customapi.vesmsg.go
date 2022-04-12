//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package site

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

func (m *SetStateReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetStateReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetStateReq) DeepCopy() *SetStateReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetStateReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetStateReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetStateReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetStateReqValidator().Validate(ctx, m, opts...)
}

type ValidateSetStateReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetStateReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateSetStateReq) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateSetStateReq) StateValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(SiteState)
		return int32(i)
	}
	// SiteState_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, SiteState_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for state")
	}

	return validatorFn, nil
}

func (v *ValidateSetStateReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetStateReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetStateReq got type %s", t)
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

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetStateReqValidator = func() *ValidateSetStateReq {
	v := &ValidateSetStateReq{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SetStateReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SetStateReq.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhState := v.StateValidationRuleHandler
	rulesState := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhState(rulesState)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for SetStateReq.state: %s", err)
		panic(errMsg)
	}
	v.FldValidators["state"] = vFn

	return v
}()

func SetStateReqValidator() db.Validator {
	return DefaultSetStateReqValidator
}

// augmented methods on protoc/std generated struct

func (m *SetStateResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SetStateResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SetStateResp) DeepCopy() *SetStateResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SetStateResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SetStateResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SetStateResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SetStateRespValidator().Validate(ctx, m, opts...)
}

type ValidateSetStateResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSetStateResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SetStateResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SetStateResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSetStateRespValidator = func() *ValidateSetStateResp {
	v := &ValidateSetStateResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SetStateRespValidator() db.Validator {
	return DefaultSetStateRespValidator
}
