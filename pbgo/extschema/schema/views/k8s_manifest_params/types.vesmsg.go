// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package k8s_manifest_params

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

func (m *BigIpBareMetalDeviceK8SParams) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *BigIpBareMetalDeviceK8SParams) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *BigIpBareMetalDeviceK8SParams) DeepCopy() *BigIpBareMetalDeviceK8SParams {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &BigIpBareMetalDeviceK8SParams{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *BigIpBareMetalDeviceK8SParams) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *BigIpBareMetalDeviceK8SParams) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return BigIpBareMetalDeviceK8SParamsValidator().Validate(ctx, m, opts...)
}

type ValidateBigIpBareMetalDeviceK8SParams struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) NodeNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for node_name")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) LicenseKeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for license_key")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) InternalNetworkSelfipValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for internal_network_selfip")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) InternalNetworkGatewayValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for internal_network_gateway")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) ExternalNetworkSelfipValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for external_network_selfip")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) ExternalNetworkGatewayValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for external_network_gateway")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) NadForExternalInterfaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for nad_for_external_interface")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) NadForInternalInterfaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for nad_for_internal_interface")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalDeviceK8SParams) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*BigIpBareMetalDeviceK8SParams)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *BigIpBareMetalDeviceK8SParams got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["external_network_gateway"]; exists {

		vOpts := append(opts, db.WithValidateField("external_network_gateway"))
		if err := fv(ctx, m.GetExternalNetworkGateway(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["external_network_selfip"]; exists {

		vOpts := append(opts, db.WithValidateField("external_network_selfip"))
		if err := fv(ctx, m.GetExternalNetworkSelfip(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["internal_network_gateway"]; exists {

		vOpts := append(opts, db.WithValidateField("internal_network_gateway"))
		if err := fv(ctx, m.GetInternalNetworkGateway(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["internal_network_selfip"]; exists {

		vOpts := append(opts, db.WithValidateField("internal_network_selfip"))
		if err := fv(ctx, m.GetInternalNetworkSelfip(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["license_key"]; exists {

		vOpts := append(opts, db.WithValidateField("license_key"))
		if err := fv(ctx, m.GetLicenseKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["nad_for_external_interface"]; exists {

		vOpts := append(opts, db.WithValidateField("nad_for_external_interface"))
		if err := fv(ctx, m.GetNadForExternalInterface(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["nad_for_internal_interface"]; exists {

		vOpts := append(opts, db.WithValidateField("nad_for_internal_interface"))
		if err := fv(ctx, m.GetNadForInternalInterface(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["node_name"]; exists {

		vOpts := append(opts, db.WithValidateField("node_name"))
		if err := fv(ctx, m.GetNodeName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultBigIpBareMetalDeviceK8SParamsValidator = func() *ValidateBigIpBareMetalDeviceK8SParams {
	v := &ValidateBigIpBareMetalDeviceK8SParams{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNodeName := v.NodeNameValidationRuleHandler
	rulesNodeName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNodeName(rulesNodeName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.node_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["node_name"] = vFn

	vrhLicenseKey := v.LicenseKeyValidationRuleHandler
	rulesLicenseKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhLicenseKey(rulesLicenseKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.license_key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["license_key"] = vFn

	vrhInternalNetworkSelfip := v.InternalNetworkSelfipValidationRuleHandler
	rulesInternalNetworkSelfip := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInternalNetworkSelfip(rulesInternalNetworkSelfip)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.internal_network_selfip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["internal_network_selfip"] = vFn

	vrhInternalNetworkGateway := v.InternalNetworkGatewayValidationRuleHandler
	rulesInternalNetworkGateway := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhInternalNetworkGateway(rulesInternalNetworkGateway)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.internal_network_gateway: %s", err)
		panic(errMsg)
	}
	v.FldValidators["internal_network_gateway"] = vFn

	vrhExternalNetworkSelfip := v.ExternalNetworkSelfipValidationRuleHandler
	rulesExternalNetworkSelfip := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhExternalNetworkSelfip(rulesExternalNetworkSelfip)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.external_network_selfip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["external_network_selfip"] = vFn

	vrhExternalNetworkGateway := v.ExternalNetworkGatewayValidationRuleHandler
	rulesExternalNetworkGateway := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhExternalNetworkGateway(rulesExternalNetworkGateway)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.external_network_gateway: %s", err)
		panic(errMsg)
	}
	v.FldValidators["external_network_gateway"] = vFn

	vrhNadForExternalInterface := v.NadForExternalInterfaceValidationRuleHandler
	rulesNadForExternalInterface := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNadForExternalInterface(rulesNadForExternalInterface)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.nad_for_external_interface: %s", err)
		panic(errMsg)
	}
	v.FldValidators["nad_for_external_interface"] = vFn

	vrhNadForInternalInterface := v.NadForInternalInterfaceValidationRuleHandler
	rulesNadForInternalInterface := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNadForInternalInterface(rulesNadForInternalInterface)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalDeviceK8SParams.nad_for_internal_interface: %s", err)
		panic(errMsg)
	}
	v.FldValidators["nad_for_internal_interface"] = vFn

	return v
}()

func BigIpBareMetalDeviceK8SParamsValidator() db.Validator {
	return DefaultBigIpBareMetalDeviceK8SParamsValidator
}

// augmented methods on protoc/std generated struct

func (m *BigIpBareMetalK8SType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *BigIpBareMetalK8SType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *BigIpBareMetalK8SType) DeepCopy() *BigIpBareMetalK8SType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &BigIpBareMetalK8SType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *BigIpBareMetalK8SType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *BigIpBareMetalK8SType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return BigIpBareMetalK8STypeValidator().Validate(ctx, m, opts...)
}

type ValidateBigIpBareMetalK8SType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateBigIpBareMetalK8SType) AdminUsernameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for admin_username")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalK8SType) SshKeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for ssh_key")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalK8SType) DevicesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for devices")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*BigIpBareMetalDeviceK8SParams, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := BigIpBareMetalDeviceK8SParamsValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for devices")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*BigIpBareMetalDeviceK8SParams)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*BigIpBareMetalDeviceK8SParams, got %T", val)
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
			return errors.Wrap(err, "repeated devices")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items devices")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalK8SType) PublicDownloadUrlValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for public_download_url")
	}

	return validatorFn, nil
}

func (v *ValidateBigIpBareMetalK8SType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*BigIpBareMetalK8SType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *BigIpBareMetalK8SType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["admin_username"]; exists {

		vOpts := append(opts, db.WithValidateField("admin_username"))
		if err := fv(ctx, m.GetAdminUsername(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["devices"]; exists {
		vOpts := append(opts, db.WithValidateField("devices"))
		if err := fv(ctx, m.GetDevices(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["public_download_url"]; exists {

		vOpts := append(opts, db.WithValidateField("public_download_url"))
		if err := fv(ctx, m.GetPublicDownloadUrl(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ssh_key"]; exists {

		vOpts := append(opts, db.WithValidateField("ssh_key"))
		if err := fv(ctx, m.GetSshKey(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultBigIpBareMetalK8STypeValidator = func() *ValidateBigIpBareMetalK8SType {
	v := &ValidateBigIpBareMetalK8SType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAdminUsername := v.AdminUsernameValidationRuleHandler
	rulesAdminUsername := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhAdminUsername(rulesAdminUsername)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalK8SType.admin_username: %s", err)
		panic(errMsg)
	}
	v.FldValidators["admin_username"] = vFn

	vrhSshKey := v.SshKeyValidationRuleHandler
	rulesSshKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhSshKey(rulesSshKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalK8SType.ssh_key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ssh_key"] = vFn

	vrhDevices := v.DevicesValidationRuleHandler
	rulesDevices := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhDevices(rulesDevices)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalK8SType.devices: %s", err)
		panic(errMsg)
	}
	v.FldValidators["devices"] = vFn

	vrhPublicDownloadUrl := v.PublicDownloadUrlValidationRuleHandler
	rulesPublicDownloadUrl := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPublicDownloadUrl(rulesPublicDownloadUrl)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for BigIpBareMetalK8SType.public_download_url: %s", err)
		panic(errMsg)
	}
	v.FldValidators["public_download_url"] = vFn

	return v
}()

func BigIpBareMetalK8STypeValidator() db.Validator {
	return DefaultBigIpBareMetalK8STypeValidator
}

// augmented methods on protoc/std generated struct

func (m *DeploymentStatusType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DeploymentStatusType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DeploymentStatusType) DeepCopy() *DeploymentStatusType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DeploymentStatusType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DeploymentStatusType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DeploymentStatusType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DeploymentStatusTypeValidator().Validate(ctx, m, opts...)
}

type ValidateDeploymentStatusType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDeploymentStatusType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DeploymentStatusType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DeploymentStatusType got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDeploymentStatusTypeValidator = func() *ValidateDeploymentStatusType {
	v := &ValidateDeploymentStatusType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DeploymentStatusTypeValidator() db.Validator {
	return DefaultDeploymentStatusTypeValidator
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

	switch m.GetF5K8SType().(type) {
	case *GlobalSpecType_F5BareMetalK8SParams:
		if fv, exists := v.FldValidators["f5_k8s_type.f5_bare_metal_k8s_params"]; exists {
			val := m.GetF5K8SType().(*GlobalSpecType_F5BareMetalK8SParams).F5BareMetalK8SParams
			vOpts := append(opts,
				db.WithValidateField("f5_k8s_type"),
				db.WithValidateField("f5_bare_metal_k8s_params"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["f5_k8s_type.f5_bare_metal_k8s_params"] = BigIpBareMetalK8STypeValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}
