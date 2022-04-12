//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package tenant

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

func (m *CA) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CA) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CA) DeepCopy() *CA {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CA{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CA) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CA) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CAValidator().Validate(ctx, m, opts...)
}

type ValidateCA struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCA) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CA)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CA got type %s", t)
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

	if fv, exists := v.FldValidators["password"]; exists {

		vOpts := append(opts, db.WithValidateField("password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["pem"]; exists {

		vOpts := append(opts, db.WithValidateField("pem"))
		if err := fv(ctx, m.GetPem(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["version"]; exists {

		vOpts := append(opts, db.WithValidateField("version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCAValidator = func() *ValidateCA {
	v := &ValidateCA{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["password"] = EncryptedPasswordValidator().Validate

	return v
}()

func CAValidator() db.Validator {
	return DefaultCAValidator
}

// augmented methods on protoc/std generated struct

func (m *EncryptedPassword) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EncryptedPassword) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EncryptedPassword) DeepCopy() *EncryptedPassword {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EncryptedPassword{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EncryptedPassword) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EncryptedPassword) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EncryptedPasswordValidator().Validate(ctx, m, opts...)
}

type ValidateEncryptedPassword struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEncryptedPassword) PasswordValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for Password")
	}

	return validatorFn, nil
}

func (v *ValidateEncryptedPassword) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EncryptedPassword)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EncryptedPassword got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["Password"]; exists {

		vOpts := append(opts, db.WithValidateField("Password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["Version"]; exists {

		vOpts := append(opts, db.WithValidateField("Version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEncryptedPasswordValidator = func() *ValidateEncryptedPassword {
	v := &ValidateEncryptedPassword{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPassword := v.PasswordValidationRuleHandler
	rulesPassword := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhPassword(rulesPassword)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EncryptedPassword.Password: %s", err)
		panic(errMsg)
	}
	v.FldValidators["Password"] = vFn

	return v
}()

func EncryptedPasswordValidator() db.Validator {
	return DefaultEncryptedPasswordValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetShapeSharedInstanceAuthKey().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.shape_shared_instance_auth_key")
	}

	return nil
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

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetSharedPublicVipDRefInfo()

}

func (m *GlobalSpecType) GetSharedPublicVipDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetSharedPublicVip()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.shared_public_vip[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "public_ip.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "shared_public_vip",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetSharedPublicVipDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetSharedPublicVipDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "public_ip.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: public_ip")
	}
	for _, ref := range m.GetSharedPublicVip() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) PublicVipValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for public_vip")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TgwAsnOffsetValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tgw_asn_offset")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) SharedPublicVipValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for shared_public_vip")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
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
			return errors.Wrap(err, "repeated shared_public_vip")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items shared_public_vip")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) CompanyNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for company_name")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TenantFqdnValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tenant_fqdn")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) AddonServicesSubscribedValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for addon_services_subscribed")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for addon_services_subscribed")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated addon_services_subscribed")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items addon_services_subscribed")
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

	if fv, exists := v.FldValidators["addon_services_subscribed"]; exists {
		vOpts := append(opts, db.WithValidateField("addon_services_subscribed"))
		if err := fv(ctx, m.GetAddonServicesSubscribed(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company_name"]; exists {

		vOpts := append(opts, db.WithValidateField("company_name"))
		if err := fv(ctx, m.GetCompanyName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["default_disable_public_ap"]; exists {

		vOpts := append(opts, db.WithValidateField("default_disable_public_ap"))
		if err := fv(ctx, m.GetDefaultDisablePublicAp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["deleted"]; exists {

		vOpts := append(opts, db.WithValidateField("deleted"))
		if err := fv(ctx, m.GetDeleted(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["k8s_server_sub_ca_latest_version"]; exists {

		vOpts := append(opts, db.WithValidateField("k8s_server_sub_ca_latest_version"))
		if err := fv(ctx, m.GetK8SServerSubCaLatestVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["k8s_server_sub_cas"]; exists {

		vOpts := append(opts, db.WithValidateField("k8s_server_sub_cas"))
		for idx, item := range m.GetK8SServerSubCas() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["proxy_root_ca_latest_version"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_root_ca_latest_version"))
		if err := fv(ctx, m.GetProxyRootCaLatestVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["proxy_root_cas"]; exists {

		vOpts := append(opts, db.WithValidateField("proxy_root_cas"))
		for idx, item := range m.GetProxyRootCas() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["public_vip"]; exists {

		vOpts := append(opts, db.WithValidateField("public_vip"))
		if err := fv(ctx, m.GetPublicVip(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["shape_shared_instance_auth_key"]; exists {

		vOpts := append(opts, db.WithValidateField("shape_shared_instance_auth_key"))
		if err := fv(ctx, m.GetShapeSharedInstanceAuthKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["shared_public_vip"]; exists {
		vOpts := append(opts, db.WithValidateField("shared_public_vip"))
		if err := fv(ctx, m.GetSharedPublicVip(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_fqdn"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_fqdn"))
		if err := fv(ctx, m.GetTenantFqdn(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tgw_asn_offset"]; exists {

		vOpts := append(opts, db.WithValidateField("tgw_asn_offset"))
		if err := fv(ctx, m.GetTgwAsnOffset(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["use_global_ain_vrf"]; exists {

		vOpts := append(opts, db.WithValidateField("use_global_ain_vrf"))
		if err := fv(ctx, m.GetUseGlobalAinVrf(), vOpts...); err != nil {
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

	vrhPublicVip := v.PublicVipValidationRuleHandler
	rulesPublicVip := map[string]string{
		"ves.io.schema.rules.string.ip": "true",
	}
	vFn, err = vrhPublicVip(rulesPublicVip)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.public_vip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["public_vip"] = vFn

	vrhTgwAsnOffset := v.TgwAsnOffsetValidationRuleHandler
	rulesTgwAsnOffset := map[string]string{
		"ves.io.schema.rules.uint32.lte": "4294963198",
	}
	vFn, err = vrhTgwAsnOffset(rulesTgwAsnOffset)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.tgw_asn_offset: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tgw_asn_offset"] = vFn

	vrhSharedPublicVip := v.SharedPublicVipValidationRuleHandler
	rulesSharedPublicVip := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhSharedPublicVip(rulesSharedPublicVip)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.shared_public_vip: %s", err)
		panic(errMsg)
	}
	v.FldValidators["shared_public_vip"] = vFn

	vrhCompanyName := v.CompanyNameValidationRuleHandler
	rulesCompanyName := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhCompanyName(rulesCompanyName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.company_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["company_name"] = vFn

	vrhTenantFqdn := v.TenantFqdnValidationRuleHandler
	rulesTenantFqdn := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhTenantFqdn(rulesTenantFqdn)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.tenant_fqdn: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tenant_fqdn"] = vFn

	vrhAddonServicesSubscribed := v.AddonServicesSubscribedValidationRuleHandler
	rulesAddonServicesSubscribed := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "32",
	}
	vFn, err = vrhAddonServicesSubscribed(rulesAddonServicesSubscribed)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.addon_services_subscribed: %s", err)
		panic(errMsg)
	}
	v.FldValidators["addon_services_subscribed"] = vFn

	v.FldValidators["k8s_server_sub_cas"] = SubCAValidator().Validate

	v.FldValidators["proxy_root_cas"] = CAValidator().Validate

	v.FldValidators["shape_shared_instance_auth_key"] = ves_io_schema.SecretTypeValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *SubCA) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SubCA) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SubCA) DeepCopy() *SubCA {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SubCA{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SubCA) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SubCA) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SubCAValidator().Validate(ctx, m, opts...)
}

type ValidateSubCA struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSubCA) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SubCA)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SubCA got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["CaName"]; exists {

		vOpts := append(opts, db.WithValidateField("CaName"))
		if err := fv(ctx, m.GetCaName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["Password"]; exists {

		vOpts := append(opts, db.WithValidateField("Password"))
		if err := fv(ctx, m.GetPassword(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["SubCAPEM"]; exists {

		vOpts := append(opts, db.WithValidateField("SubCAPEM"))
		if err := fv(ctx, m.GetSubCAPEM(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["Version"]; exists {

		vOpts := append(opts, db.WithValidateField("Version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSubCAValidator = func() *ValidateSubCA {
	v := &ValidateSubCA{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["Password"] = EncryptedPasswordValidator().Validate

	return v
}()

func SubCAValidator() db.Validator {
	return DefaultSubCAValidator
}
