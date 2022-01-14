//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package registration

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

func (m *ApprovalReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ApprovalReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ApprovalReq) DeepCopy() *ApprovalReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ApprovalReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ApprovalReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ApprovalReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ApprovalReqValidator().Validate(ctx, m, opts...)
}

type ValidateApprovalReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateApprovalReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateApprovalReq) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateApprovalReq) StateValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ObjectState)
		return int32(i)
	}
	// ObjectState_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ObjectState_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for state")
	}

	return validatorFn, nil
}

func (v *ValidateApprovalReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ApprovalReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ApprovalReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["annotations"]; exists {

		vOpts := append(opts, db.WithValidateField("annotations"))
		for key, value := range m.GetAnnotations() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["connected_region"]; exists {

		vOpts := append(opts, db.WithValidateField("connected_region"))
		if err := fv(ctx, m.GetConnectedRegion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["labels"]; exists {

		vOpts := append(opts, db.WithValidateField("labels"))
		for key, value := range m.GetLabels() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
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

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["passport"]; exists {

		vOpts := append(opts, db.WithValidateField("passport"))
		if err := fv(ctx, m.GetPassport(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["state"]; exists {

		vOpts := append(opts, db.WithValidateField("state"))
		if err := fv(ctx, m.GetState(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tunnel_type"]; exists {

		vOpts := append(opts, db.WithValidateField("tunnel_type"))
		if err := fv(ctx, m.GetTunnelType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultApprovalReqValidator = func() *ValidateApprovalReq {
	v := &ValidateApprovalReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApprovalReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApprovalReq.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhState := v.StateValidationRuleHandler
	rulesState := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhState(rulesState)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ApprovalReq.state: %s", err)
		panic(errMsg)
	}
	v.FldValidators["state"] = vFn

	v.FldValidators["passport"] = PassportValidator().Validate

	return v
}()

func ApprovalReqValidator() db.Validator {
	return DefaultApprovalReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ConfigReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ConfigReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ConfigReq) DeepCopy() *ConfigReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ConfigReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ConfigReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ConfigReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ConfigReqValidator().Validate(ctx, m, opts...)
}

type ValidateConfigReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateConfigReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateConfigReq) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateConfigReq) TokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for token")
	}

	return validatorFn, nil
}

func (v *ValidateConfigReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ConfigReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ConfigReq got type %s", t)
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

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultConfigReqValidator = func() *ValidateConfigReq {
	v := &ValidateConfigReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for ConfigReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ConfigReq.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhToken := v.TokenValidationRuleHandler
	rulesToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhToken(rulesToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ConfigReq.token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["token"] = vFn

	return v
}()

func ConfigReqValidator() db.Validator {
	return DefaultConfigReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ConfigResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ConfigResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ConfigResp) DeepCopy() *ConfigResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ConfigResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ConfigResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ConfigResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ConfigRespValidator().Validate(ctx, m, opts...)
}

type ValidateConfigResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateConfigResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ConfigResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ConfigResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["hash"]; exists {

		vOpts := append(opts, db.WithValidateField("hash"))
		if err := fv(ctx, m.GetHash(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["workload"]; exists {

		vOpts := append(opts, db.WithValidateField("workload"))
		for key, value := range m.GetWorkload() {
			vOpts := append(vOpts, db.WithValidateMapKey(key))
			if err := fv(ctx, value, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultConfigRespValidator = func() *ValidateConfigResp {
	v := &ValidateConfigResp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ConfigRespValidator() db.Validator {
	return DefaultConfigRespValidator
}

// augmented methods on protoc/std generated struct

func (m *ListBySiteReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListBySiteReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListBySiteReq) DeepCopy() *ListBySiteReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListBySiteReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListBySiteReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListBySiteReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListBySiteReqValidator().Validate(ctx, m, opts...)
}

type ValidateListBySiteReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListBySiteReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateListBySiteReq) SiteNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for site_name")
	}

	return validatorFn, nil
}

func (v *ValidateListBySiteReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListBySiteReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListBySiteReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["site_name"]; exists {

		vOpts := append(opts, db.WithValidateField("site_name"))
		if err := fv(ctx, m.GetSiteName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListBySiteReqValidator = func() *ValidateListBySiteReq {
	v := &ValidateListBySiteReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for ListBySiteReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhSiteName := v.SiteNameValidationRuleHandler
	rulesSiteName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSiteName(rulesSiteName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ListBySiteReq.site_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["site_name"] = vFn

	return v
}()

func ListBySiteReqValidator() db.Validator {
	return DefaultListBySiteReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ListStateReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListStateReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListStateReq) DeepCopy() *ListStateReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListStateReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListStateReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListStateReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListStateReqValidator().Validate(ctx, m, opts...)
}

type ValidateListStateReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListStateReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateListStateReq) StateValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ObjectState)
		return int32(i)
	}
	// ObjectState_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ObjectState_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for state")
	}

	return validatorFn, nil
}

func (v *ValidateListStateReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListStateReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListStateReq got type %s", t)
		}
	}
	if m == nil {
		return nil
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
var DefaultListStateReqValidator = func() *ValidateListStateReq {
	v := &ValidateListStateReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for ListStateReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhState := v.StateValidationRuleHandler
	rulesState := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhState(rulesState)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ListStateReq.state: %s", err)
		panic(errMsg)
	}
	v.FldValidators["state"] = vFn

	return v
}()

func ListStateReqValidator() db.Validator {
	return DefaultListStateReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ObjectChangeResp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ObjectChangeResp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ObjectChangeResp) DeepCopy() *ObjectChangeResp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ObjectChangeResp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ObjectChangeResp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ObjectChangeResp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ObjectChangeRespValidator().Validate(ctx, m, opts...)
}

func (m *ObjectChangeResp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return nil, nil

}

type ValidateObjectChangeResp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateObjectChangeResp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ObjectChangeResp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ObjectChangeResp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["obj"]; exists {

		vOpts := append(opts, db.WithValidateField("obj"))
		if err := fv(ctx, m.GetObj(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultObjectChangeRespValidator = func() *ValidateObjectChangeResp {
	v := &ValidateObjectChangeResp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["obj"] = ObjectValidator().Validate

	return v
}()

func ObjectChangeRespValidator() db.Validator {
	return DefaultObjectChangeRespValidator
}

// augmented methods on protoc/std generated struct

func (m *RegistrationCreateRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RegistrationCreateRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RegistrationCreateRequest) DeepCopy() *RegistrationCreateRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RegistrationCreateRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RegistrationCreateRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RegistrationCreateRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RegistrationCreateRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRegistrationCreateRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRegistrationCreateRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RegistrationCreateRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RegistrationCreateRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("metadata"))
		if err := fv(ctx, m.GetMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["spec"]; exists {

		vOpts := append(opts, db.WithValidateField("spec"))
		if err := fv(ctx, m.GetSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRegistrationCreateRequestValidator = func() *ValidateRegistrationCreateRequest {
	v := &ValidateRegistrationCreateRequest{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["metadata"] = ves_io_schema.ObjectCreateMetaTypeValidator().Validate

	v.FldValidators["spec"] = CreateSpecTypeValidator().Validate

	return v
}()

func RegistrationCreateRequestValidator() db.Validator {
	return DefaultRegistrationCreateRequestValidator
}

func (m *RegistrationCreateRequest) FromObject(e db.Entry) {
	f := e.DeepCopy().(*DBObject)
	_ = f

	if m.Metadata == nil {
		m.Metadata = &ves_io_schema.ObjectCreateMetaType{}
	}
	m.Metadata.FromObjectMetaType(f.GetMetadata())

	if m.Spec == nil {
		m.Spec = &CreateSpecType{}
	}
	m.Spec.FromGlobalSpecType(f.GetSpec().GetGcSpec())

}

func (m *RegistrationCreateRequest) ToObject(e db.Entry) {
	m1 := m.DeepCopy()
	_ = m1
	f := e.(*DBObject)
	_ = f

	if m1.Metadata != nil {
		if f.Metadata == nil {
			f.Metadata = &ves_io_schema.ObjectMetaType{}
		}
	} else if f.Metadata != nil {
		f.Metadata = nil
	}

	if m1.Metadata != nil {
		m1.Metadata.ToObjectMetaType(f.Metadata)
	}

	if m1.Spec != nil {
		if f.Spec == nil {
			f.Spec = &SpecType{}
		}
	} else if f.Spec != nil {
		f.Spec = nil
	}

	if m1.Spec != nil {
		if f.Spec.GcSpec == nil {
			f.Spec.GcSpec = &GlobalSpecType{}
		}
	} else if f.Spec != nil {
		f.Spec.GcSpec = nil
	}

	if m1.Spec != nil {
		m1.Spec.ToGlobalSpecType(f.Spec.GcSpec)
	}

}
