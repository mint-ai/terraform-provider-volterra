// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package flow_exporter

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

func (m *FlowCollector) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FlowCollector) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FlowCollector) DeepCopy() *FlowCollector {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FlowCollector{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FlowCollector) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FlowCollector) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FlowCollectorValidator().Validate(ctx, m, opts...)
}

type ValidateFlowCollector struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFlowCollector) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FlowCollector)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FlowCollector got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["collector_end_point"]; exists {

		vOpts := append(opts, db.WithValidateField("collector_end_point"))
		if err := fv(ctx, m.GetCollectorEndPoint(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["collector_type"]; exists {

		vOpts := append(opts, db.WithValidateField("collector_type"))
		if err := fv(ctx, m.GetCollectorType(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetCollectorTypeParameters().(type) {
	case *FlowCollector_IpfixCollectorParameters:
		if fv, exists := v.FldValidators["collector_type_parameters.ipfix_collector_parameters"]; exists {
			val := m.GetCollectorTypeParameters().(*FlowCollector_IpfixCollectorParameters).IpfixCollectorParameters
			vOpts := append(opts,
				db.WithValidateField("collector_type_parameters"),
				db.WithValidateField("ipfix_collector_parameters"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFlowCollectorValidator = func() *ValidateFlowCollector {
	v := &ValidateFlowCollector{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["collector_type_parameters.ipfix_collector_parameters"] = IpfixParametersValidator().Validate

	v.FldValidators["collector_end_point"] = FlowCollectorEndPointValidator().Validate

	return v
}()

func FlowCollectorValidator() db.Validator {
	return DefaultFlowCollectorValidator
}

// augmented methods on protoc/std generated struct

func (m *FlowCollectorEndPoint) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FlowCollectorEndPoint) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FlowCollectorEndPoint) DeepCopy() *FlowCollectorEndPoint {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FlowCollectorEndPoint{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FlowCollectorEndPoint) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FlowCollectorEndPoint) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FlowCollectorEndPointValidator().Validate(ctx, m, opts...)
}

type ValidateFlowCollectorEndPoint struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFlowCollectorEndPoint) CollectorIpInformationValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {
	validatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for collector_ip_information")
	}
	return validatorFn, nil
}

func (v *ValidateFlowCollectorEndPoint) CollectorProtoValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for collector_proto")
	}

	return validatorFn, nil
}

func (v *ValidateFlowCollectorEndPoint) CollectorPortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for collector_port")
	}

	return validatorFn, nil
}

func (v *ValidateFlowCollectorEndPoint) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FlowCollectorEndPoint)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FlowCollectorEndPoint got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["collector_ip_information"]; exists {
		val := m.GetCollectorIpInformation()
		vOpts := append(opts,
			db.WithValidateField("collector_ip_information"),
		)
		if err := fv(ctx, val, vOpts...); err != nil {
			return err
		}
	}

	switch m.GetCollectorIpInformation().(type) {
	case *FlowCollectorEndPoint_CollectorIp:
		if fv, exists := v.FldValidators["collector_ip_information.collector_ip"]; exists {
			val := m.GetCollectorIpInformation().(*FlowCollectorEndPoint_CollectorIp).CollectorIp
			vOpts := append(opts,
				db.WithValidateField("collector_ip_information"),
				db.WithValidateField("collector_ip"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *FlowCollectorEndPoint_CollectorServiceName:
		if fv, exists := v.FldValidators["collector_ip_information.collector_service_name"]; exists {
			val := m.GetCollectorIpInformation().(*FlowCollectorEndPoint_CollectorServiceName).CollectorServiceName
			vOpts := append(opts,
				db.WithValidateField("collector_ip_information"),
				db.WithValidateField("collector_service_name"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	switch m.GetCollectorNetworkChoice().(type) {
	case *FlowCollectorEndPoint_SiteLocalNetwork:
		if fv, exists := v.FldValidators["collector_network_choice.site_local_network"]; exists {
			val := m.GetCollectorNetworkChoice().(*FlowCollectorEndPoint_SiteLocalNetwork).SiteLocalNetwork
			vOpts := append(opts,
				db.WithValidateField("collector_network_choice"),
				db.WithValidateField("site_local_network"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *FlowCollectorEndPoint_SiteLocalInsideNetwork:
		if fv, exists := v.FldValidators["collector_network_choice.site_local_inside_network"]; exists {
			val := m.GetCollectorNetworkChoice().(*FlowCollectorEndPoint_SiteLocalInsideNetwork).SiteLocalInsideNetwork
			vOpts := append(opts,
				db.WithValidateField("collector_network_choice"),
				db.WithValidateField("site_local_inside_network"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["collector_port"]; exists {

		vOpts := append(opts, db.WithValidateField("collector_port"))
		if err := fv(ctx, m.GetCollectorPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["collector_proto"]; exists {

		vOpts := append(opts, db.WithValidateField("collector_proto"))
		if err := fv(ctx, m.GetCollectorProto(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFlowCollectorEndPointValidator = func() *ValidateFlowCollectorEndPoint {
	v := &ValidateFlowCollectorEndPoint{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhCollectorIpInformation := v.CollectorIpInformationValidationRuleHandler
	rulesCollectorIpInformation := map[string]string{
		"ves.io.schema.rules.message.required_oneof": "true",
	}
	vFn, err = vrhCollectorIpInformation(rulesCollectorIpInformation)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FlowCollectorEndPoint.collector_ip_information: %s", err)
		panic(errMsg)
	}
	v.FldValidators["collector_ip_information"] = vFn

	vrhCollectorProto := v.CollectorProtoValidationRuleHandler
	rulesCollectorProto := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.in":        "[\"UDP\"]",
	}
	vFn, err = vrhCollectorProto(rulesCollectorProto)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FlowCollectorEndPoint.collector_proto: %s", err)
		panic(errMsg)
	}
	v.FldValidators["collector_proto"] = vFn

	vrhCollectorPort := v.CollectorPortValidationRuleHandler
	rulesCollectorPort := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhCollectorPort(rulesCollectorPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FlowCollectorEndPoint.collector_port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["collector_port"] = vFn

	v.FldValidators["collector_ip_information.collector_ip"] = ves_io_schema.IpAddressTypeValidator().Validate

	return v
}()

func FlowCollectorEndPointValidator() db.Validator {
	return DefaultFlowCollectorEndPointValidator
}

// augmented methods on protoc/std generated struct

func (m *FlowExportTimeout) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FlowExportTimeout) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FlowExportTimeout) DeepCopy() *FlowExportTimeout {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FlowExportTimeout{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FlowExportTimeout) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FlowExportTimeout) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FlowExportTimeoutValidator().Validate(ctx, m, opts...)
}

type ValidateFlowExportTimeout struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFlowExportTimeout) TimeoutPacketsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for timeout_packets")
	}

	return validatorFn, nil
}

func (v *ValidateFlowExportTimeout) TimeoutSecondsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for timeout_seconds")
	}

	return validatorFn, nil
}

func (v *ValidateFlowExportTimeout) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FlowExportTimeout)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FlowExportTimeout got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["timeout_packets"]; exists {

		vOpts := append(opts, db.WithValidateField("timeout_packets"))
		if err := fv(ctx, m.GetTimeoutPackets(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["timeout_seconds"]; exists {

		vOpts := append(opts, db.WithValidateField("timeout_seconds"))
		if err := fv(ctx, m.GetTimeoutSeconds(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFlowExportTimeoutValidator = func() *ValidateFlowExportTimeout {
	v := &ValidateFlowExportTimeout{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhTimeoutPackets := v.TimeoutPacketsValidationRuleHandler
	rulesTimeoutPackets := map[string]string{
		"ves.io.schema.rules.uint32.gt": "500",
	}
	vFn, err = vrhTimeoutPackets(rulesTimeoutPackets)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FlowExportTimeout.timeout_packets: %s", err)
		panic(errMsg)
	}
	v.FldValidators["timeout_packets"] = vFn

	vrhTimeoutSeconds := v.TimeoutSecondsValidationRuleHandler
	rulesTimeoutSeconds := map[string]string{
		"ves.io.schema.rules.uint32.gt": "0",
	}
	vFn, err = vrhTimeoutSeconds(rulesTimeoutSeconds)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FlowExportTimeout.timeout_seconds: %s", err)
		panic(errMsg)
	}
	v.FldValidators["timeout_seconds"] = vFn

	return v
}()

func FlowExportTimeoutValidator() db.Validator {
	return DefaultFlowExportTimeoutValidator
}

// augmented methods on protoc/std generated struct

func (m *FlowSampler) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FlowSampler) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FlowSampler) DeepCopy() *FlowSampler {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FlowSampler{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FlowSampler) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FlowSampler) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FlowSamplerValidator().Validate(ctx, m, opts...)
}

type ValidateFlowSampler struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFlowSampler) SamplerRateValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for sampler_rate")
	}

	return validatorFn, nil
}

func (v *ValidateFlowSampler) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FlowSampler)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FlowSampler got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["sampler_rate"]; exists {

		vOpts := append(opts, db.WithValidateField("sampler_rate"))
		if err := fv(ctx, m.GetSamplerRate(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFlowSamplerValidator = func() *ValidateFlowSampler {
	v := &ValidateFlowSampler{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhSamplerRate := v.SamplerRateValidationRuleHandler
	rulesSamplerRate := map[string]string{
		"ves.io.schema.rules.uint32.gt":  "500",
		"ves.io.schema.rules.uint32.lte": "8000",
	}
	vFn, err = vrhSamplerRate(rulesSamplerRate)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FlowSampler.sampler_rate: %s", err)
		panic(errMsg)
	}
	v.FldValidators["sampler_rate"] = vFn

	return v
}()

func FlowSamplerValidator() db.Validator {
	return DefaultFlowSamplerValidator
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

func (v *ValidateGlobalSpecType) FlowCollectorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for flow_collector")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*FlowCollector, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := FlowCollectorValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for flow_collector")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*FlowCollector)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*FlowCollector, got %T", val)
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
			return errors.Wrap(err, "repeated flow_collector")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items flow_collector")
		}
		return nil
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

	if fv, exists := v.FldValidators["flow_active_timeout"]; exists {

		vOpts := append(opts, db.WithValidateField("flow_active_timeout"))
		if err := fv(ctx, m.GetFlowActiveTimeout(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["flow_collector"]; exists {
		vOpts := append(opts, db.WithValidateField("flow_collector"))
		if err := fv(ctx, m.GetFlowCollector(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["flow_sampler"]; exists {

		vOpts := append(opts, db.WithValidateField("flow_sampler"))
		if err := fv(ctx, m.GetFlowSampler(), vOpts...); err != nil {
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

	vrhFlowCollector := v.FlowCollectorValidationRuleHandler
	rulesFlowCollector := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "2",
	}
	vFn, err = vrhFlowCollector(rulesFlowCollector)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.flow_collector: %s", err)
		panic(errMsg)
	}
	v.FldValidators["flow_collector"] = vFn

	v.FldValidators["flow_sampler"] = FlowSamplerValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *IpfixParameters) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *IpfixParameters) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *IpfixParameters) DeepCopy() *IpfixParameters {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &IpfixParameters{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *IpfixParameters) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *IpfixParameters) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return IpfixParametersValidator().Validate(ctx, m, opts...)
}

type ValidateIpfixParameters struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateIpfixParameters) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*IpfixParameters)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *IpfixParameters got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["ipfix_template_refresh_time"]; exists {

		vOpts := append(opts, db.WithValidateField("ipfix_template_refresh_time"))
		if err := fv(ctx, m.GetIpfixTemplateRefreshTime(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultIpfixParametersValidator = func() *ValidateIpfixParameters {
	v := &ValidateIpfixParameters{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["ipfix_template_refresh_time"] = FlowExportTimeoutValidator().Validate

	return v
}()

func IpfixParametersValidator() db.Validator {
	return DefaultIpfixParametersValidator
}
