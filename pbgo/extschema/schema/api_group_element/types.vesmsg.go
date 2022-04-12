//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package api_group_element

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

func (v *ValidateCreateSpecType) MethodsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for methods")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema.HttpMethod, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for methods")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema.HttpMethod)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema.HttpMethod, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated methods")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items methods")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) PathRegexValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path_regex")
	}

	return validatorFn, nil
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

	if fv, exists := v.FldValidators["methods"]; exists {
		vOpts := append(opts, db.WithValidateField("methods"))
		if err := fv(ctx, m.GetMethods(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path_regex"]; exists {

		vOpts := append(opts, db.WithValidateField("path_regex"))
		if err := fv(ctx, m.GetPathRegex(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMethods := v.MethodsValidationRuleHandler
	rulesMethods := map[string]string{
		"ves.io.schema.rules.message.required":                 "true",
		"ves.io.schema.rules.repeated.items.enum.defined_only": "true",
		"ves.io.schema.rules.repeated.items.enum.not_in":       "0",
		"ves.io.schema.rules.repeated.min_items":               "1",
		"ves.io.schema.rules.repeated.unique":                  "true",
	}
	vFn, err = vrhMethods(rulesMethods)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.methods: %s", err)
		panic(errMsg)
	}
	v.FldValidators["methods"] = vFn

	vrhPathRegex := v.PathRegexValidationRuleHandler
	rulesPathRegex := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_bytes": "1024",
		"ves.io.schema.rules.string.min_bytes": "1",
		"ves.io.schema.rules.string.regex":     "true",
	}
	vFn, err = vrhPathRegex(rulesPathRegex)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.path_regex: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path_regex"] = vFn

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

func (v *ValidateGetSpecType) MethodsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for methods")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema.HttpMethod, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for methods")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema.HttpMethod)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema.HttpMethod, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated methods")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items methods")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) PathRegexValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path_regex")
	}

	return validatorFn, nil
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

	if fv, exists := v.FldValidators["methods"]; exists {
		vOpts := append(opts, db.WithValidateField("methods"))
		if err := fv(ctx, m.GetMethods(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path_regex"]; exists {

		vOpts := append(opts, db.WithValidateField("path_regex"))
		if err := fv(ctx, m.GetPathRegex(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMethods := v.MethodsValidationRuleHandler
	rulesMethods := map[string]string{
		"ves.io.schema.rules.message.required":                 "true",
		"ves.io.schema.rules.repeated.items.enum.defined_only": "true",
		"ves.io.schema.rules.repeated.items.enum.not_in":       "0",
		"ves.io.schema.rules.repeated.min_items":               "1",
		"ves.io.schema.rules.repeated.unique":                  "true",
	}
	vFn, err = vrhMethods(rulesMethods)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.methods: %s", err)
		panic(errMsg)
	}
	v.FldValidators["methods"] = vFn

	vrhPathRegex := v.PathRegexValidationRuleHandler
	rulesPathRegex := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_bytes": "1024",
		"ves.io.schema.rules.string.min_bytes": "1",
		"ves.io.schema.rules.string.regex":     "true",
	}
	vFn, err = vrhPathRegex(rulesPathRegex)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.path_regex: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path_regex"] = vFn

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

func (v *ValidateGlobalSpecType) MethodsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for methods")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema.HttpMethod, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for methods")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema.HttpMethod)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema.HttpMethod, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated methods")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items methods")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) PathRegexValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path_regex")
	}

	return validatorFn, nil
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

	if fv, exists := v.FldValidators["methods"]; exists {
		vOpts := append(opts, db.WithValidateField("methods"))
		if err := fv(ctx, m.GetMethods(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path_regex"]; exists {

		vOpts := append(opts, db.WithValidateField("path_regex"))
		if err := fv(ctx, m.GetPathRegex(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMethods := v.MethodsValidationRuleHandler
	rulesMethods := map[string]string{
		"ves.io.schema.rules.message.required":                 "true",
		"ves.io.schema.rules.repeated.items.enum.defined_only": "true",
		"ves.io.schema.rules.repeated.items.enum.not_in":       "0",
		"ves.io.schema.rules.repeated.min_items":               "1",
		"ves.io.schema.rules.repeated.unique":                  "true",
	}
	vFn, err = vrhMethods(rulesMethods)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.methods: %s", err)
		panic(errMsg)
	}
	v.FldValidators["methods"] = vFn

	vrhPathRegex := v.PathRegexValidationRuleHandler
	rulesPathRegex := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_bytes": "1024",
		"ves.io.schema.rules.string.min_bytes": "1",
		"ves.io.schema.rules.string.regex":     "true",
	}
	vFn, err = vrhPathRegex(rulesPathRegex)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.path_regex: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path_regex"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
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

func (v *ValidateReplaceSpecType) MethodsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.HttpMethod)
		return int32(i)
	}
	// ves_io_schema.HttpMethod_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, ves_io_schema.HttpMethod_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for methods")
	}
	itemsValidatorFn := func(ctx context.Context, elems []ves_io_schema.HttpMethod, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for methods")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]ves_io_schema.HttpMethod)
		if !ok {
			return fmt.Errorf("Repeated validation expected []ves_io_schema.HttpMethod, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated methods")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items methods")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) PathRegexValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path_regex")
	}

	return validatorFn, nil
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

	if fv, exists := v.FldValidators["methods"]; exists {
		vOpts := append(opts, db.WithValidateField("methods"))
		if err := fv(ctx, m.GetMethods(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["path_regex"]; exists {

		vOpts := append(opts, db.WithValidateField("path_regex"))
		if err := fv(ctx, m.GetPathRegex(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhMethods := v.MethodsValidationRuleHandler
	rulesMethods := map[string]string{
		"ves.io.schema.rules.message.required":                 "true",
		"ves.io.schema.rules.repeated.items.enum.defined_only": "true",
		"ves.io.schema.rules.repeated.items.enum.not_in":       "0",
		"ves.io.schema.rules.repeated.min_items":               "1",
		"ves.io.schema.rules.repeated.unique":                  "true",
	}
	vFn, err = vrhMethods(rulesMethods)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.methods: %s", err)
		panic(errMsg)
	}
	v.FldValidators["methods"] = vFn

	vrhPathRegex := v.PathRegexValidationRuleHandler
	rulesPathRegex := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_bytes": "1024",
		"ves.io.schema.rules.string.min_bytes": "1",
		"ves.io.schema.rules.string.regex":     "true",
	}
	vFn, err = vrhPathRegex(rulesPathRegex)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.path_regex: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path_regex"] = vFn

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Methods = f.GetMethods()
	m.PathRegex = f.GetPathRegex()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Methods = m1.Methods
	f.PathRegex = m1.PathRegex
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Methods = f.GetMethods()
	m.PathRegex = f.GetPathRegex()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Methods = m1.Methods
	f.PathRegex = m1.PathRegex
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.Methods = f.GetMethods()
	m.PathRegex = f.GetPathRegex()
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.Methods = m1.Methods
	f.PathRegex = m1.PathRegex
}
