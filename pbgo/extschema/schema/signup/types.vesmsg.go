// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package signup

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_infraprotect_information "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/infraprotect_information"
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

func (m *CreateSpecType) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.PaymentProviderToken = ""

	return copy.string()
}

func (m *CreateSpecType) GoString() string {
	copy := m.DeepCopy()
	copy.PaymentProviderToken = ""

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *CreateSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.PaymentProviderToken = ""

	if err := m.GetInfraprotectInfo().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting CreateSpecType.infraprotect_info")
	}

	return nil
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

func (m *CreateSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetCompanyDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCompanyDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetCustomerDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCustomerDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetCompanyDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCompany() == nil {
		return nil, nil
	}

	drInfos, err := m.GetCompany().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetCompany().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "company." + dri.DRField
	}
	return drInfos, err

}

// GetDRefInfo for the field's type
func (m *CreateSpecType) GetCustomerDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCustomer() == nil {
		return nil, nil
	}

	drInfos, err := m.GetCustomer().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetCustomer().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "customer." + dri.DRField
	}
	return drInfos, err

}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) LocaleValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for locale")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) TypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.TenantType)
		return int32(i)
	}
	// ves_io_schema.TenantType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.TenantType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for type")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) TokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for token")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) SfdcSubscriptionIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for sfdc_subscription_id")
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

	if fv, exists := v.FldValidators["billing_address"]; exists {

		vOpts := append(opts, db.WithValidateField("billing_address"))
		if err := fv(ctx, m.GetBillingAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company"]; exists {

		vOpts := append(opts, db.WithValidateField("company"))
		if err := fv(ctx, m.GetCompany(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company_contact"]; exists {

		vOpts := append(opts, db.WithValidateField("company_contact"))
		if err := fv(ctx, m.GetCompanyContact(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company_name"]; exists {

		vOpts := append(opts, db.WithValidateField("company_name"))
		if err := fv(ctx, m.GetCompanyName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["contact_number"]; exists {

		vOpts := append(opts, db.WithValidateField("contact_number"))
		if err := fv(ctx, m.GetContactNumber(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["crm_info"]; exists {

		vOpts := append(opts, db.WithValidateField("crm_info"))
		if err := fv(ctx, m.GetCrmInfo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["currency"]; exists {

		vOpts := append(opts, db.WithValidateField("currency"))
		if err := fv(ctx, m.GetCurrency(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["customer"]; exists {

		vOpts := append(opts, db.WithValidateField("customer"))
		if err := fv(ctx, m.GetCustomer(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["customer_contact"]; exists {

		vOpts := append(opts, db.WithValidateField("customer_contact"))
		if err := fv(ctx, m.GetCustomerContact(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["domain"]; exists {

		vOpts := append(opts, db.WithValidateField("domain"))
		if err := fv(ctx, m.GetDomain(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["infraprotect_info"]; exists {

		vOpts := append(opts, db.WithValidateField("infraprotect_info"))
		if err := fv(ctx, m.GetInfraprotectInfo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locale"]; exists {

		vOpts := append(opts, db.WithValidateField("locale"))
		if err := fv(ctx, m.GetLocale(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["payment_provider_token"]; exists {

		vOpts := append(opts, db.WithValidateField("payment_provider_token"))
		if err := fv(ctx, m.GetPaymentProviderToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sfdc_subscription_id"]; exists {

		vOpts := append(opts, db.WithValidateField("sfdc_subscription_id"))
		if err := fv(ctx, m.GetSfdcSubscriptionId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["support_plan_name"]; exists {

		vOpts := append(opts, db.WithValidateField("support_plan_name"))
		if err := fv(ctx, m.GetSupportPlanName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tax_exempt"]; exists {

		vOpts := append(opts, db.WithValidateField("tax_exempt"))
		if err := fv(ctx, m.GetTaxExempt(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_accepted"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_accepted"))
		if err := fv(ctx, m.GetTosAccepted(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_accepted_at"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_accepted_at"))
		if err := fv(ctx, m.GetTosAcceptedAt(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_version"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_version"))
		if err := fv(ctx, m.GetTosVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["usage_plan_name"]; exists {

		vOpts := append(opts, db.WithValidateField("usage_plan_name"))
		if err := fv(ctx, m.GetUsagePlanName(), vOpts...); err != nil {
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

	vrhLocale := v.LocaleValidationRuleHandler
	rulesLocale := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhLocale(rulesLocale)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.locale: %s", err)
		panic(errMsg)
	}
	v.FldValidators["locale"] = vFn

	vrhType := v.TypeValidationRuleHandler
	rulesType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhType(rulesType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["type"] = vFn

	vrhToken := v.TokenValidationRuleHandler
	rulesToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhToken(rulesToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["token"] = vFn

	vrhSfdcSubscriptionId := v.SfdcSubscriptionIdValidationRuleHandler
	rulesSfdcSubscriptionId := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhSfdcSubscriptionId(rulesSfdcSubscriptionId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.sfdc_subscription_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["sfdc_subscription_id"] = vFn

	v.FldValidators["infraprotect_info"] = ves_io_schema_infraprotect_information.GlobalSpecTypeValidator().Validate

	v.FldValidators["crm_info"] = CrmInfoValidator().Validate

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *CrmInfo) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CrmInfo) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CrmInfo) DeepCopy() *CrmInfo {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CrmInfo{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CrmInfo) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CrmInfo) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CrmInfoValidator().Validate(ctx, m, opts...)
}

type ValidateCrmInfo struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCrmInfo) AccountIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for account_id")
	}

	return validatorFn, nil
}

func (v *ValidateCrmInfo) EntitlementIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for entitlement_id")
	}

	return validatorFn, nil
}

func (v *ValidateCrmInfo) SubscriptionIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for subscription_id")
	}

	return validatorFn, nil
}

func (v *ValidateCrmInfo) OrderTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for order_type")
	}

	return validatorFn, nil
}

func (v *ValidateCrmInfo) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CrmInfo)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CrmInfo got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["account_id"]; exists {

		vOpts := append(opts, db.WithValidateField("account_id"))
		if err := fv(ctx, m.GetAccountId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["entitled_skus"]; exists {

		vOpts := append(opts, db.WithValidateField("entitled_skus"))
		for idx, item := range m.GetEntitledSkus() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["entitlement_id"]; exists {

		vOpts := append(opts, db.WithValidateField("entitlement_id"))
		if err := fv(ctx, m.GetEntitlementId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["order_type"]; exists {

		vOpts := append(opts, db.WithValidateField("order_type"))
		if err := fv(ctx, m.GetOrderType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subscription_id"]; exists {

		vOpts := append(opts, db.WithValidateField("subscription_id"))
		if err := fv(ctx, m.GetSubscriptionId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCrmInfoValidator = func() *ValidateCrmInfo {
	v := &ValidateCrmInfo{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhAccountId := v.AccountIdValidationRuleHandler
	rulesAccountId := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhAccountId(rulesAccountId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CrmInfo.account_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["account_id"] = vFn

	vrhEntitlementId := v.EntitlementIdValidationRuleHandler
	rulesEntitlementId := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhEntitlementId(rulesEntitlementId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CrmInfo.entitlement_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["entitlement_id"] = vFn

	vrhSubscriptionId := v.SubscriptionIdValidationRuleHandler
	rulesSubscriptionId := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhSubscriptionId(rulesSubscriptionId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CrmInfo.subscription_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["subscription_id"] = vFn

	vrhOrderType := v.OrderTypeValidationRuleHandler
	rulesOrderType := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhOrderType(rulesOrderType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CrmInfo.order_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["order_type"] = vFn

	return v
}()

func CrmInfoValidator() db.Validator {
	return DefaultCrmInfoValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *GetSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

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

func (m *GlobalSpecType) String() string {
	if m == nil {
		return ""
	}
	copy := m.DeepCopy()
	copy.PaymentProviderToken = ""

	return copy.string()
}

func (m *GlobalSpecType) GoString() string {
	copy := m.DeepCopy()
	copy.PaymentProviderToken = ""

	return copy.goString()
}

// Redact squashes sensitive info in m (in-place)
func (m *GlobalSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	m.PaymentProviderToken = ""

	if err := m.GetInfraprotectInfo().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting GlobalSpecType.infraprotect_info")
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

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetCompanyDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCompanyDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetCustomerDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetCustomerDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetCompanyDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCompany() == nil {
		return nil, nil
	}

	drInfos, err := m.GetCompany().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetCompany().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "company." + dri.DRField
	}
	return drInfos, err

}

// GetDRefInfo for the field's type
func (m *GlobalSpecType) GetCustomerDRefInfo() ([]db.DRefInfo, error) {
	if m.GetCustomer() == nil {
		return nil, nil
	}

	drInfos, err := m.GetCustomer().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetCustomer().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "customer." + dri.DRField
	}
	return drInfos, err

}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) LocaleValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for locale")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ves_io_schema.TenantType)
		return int32(i)
	}
	// ves_io_schema.TenantType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ves_io_schema.TenantType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for type")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TokenValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for token")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) SfdcSubscriptionIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for sfdc_subscription_id")
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

	if fv, exists := v.FldValidators["billing_address"]; exists {

		vOpts := append(opts, db.WithValidateField("billing_address"))
		if err := fv(ctx, m.GetBillingAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["billing_provider_account_id"]; exists {

		vOpts := append(opts, db.WithValidateField("billing_provider_account_id"))
		if err := fv(ctx, m.GetBillingProviderAccountId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company"]; exists {

		vOpts := append(opts, db.WithValidateField("company"))
		if err := fv(ctx, m.GetCompany(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company_contact"]; exists {

		vOpts := append(opts, db.WithValidateField("company_contact"))
		if err := fv(ctx, m.GetCompanyContact(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["company_name"]; exists {

		vOpts := append(opts, db.WithValidateField("company_name"))
		if err := fv(ctx, m.GetCompanyName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["contact_number"]; exists {

		vOpts := append(opts, db.WithValidateField("contact_number"))
		if err := fv(ctx, m.GetContactNumber(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["crm_info"]; exists {

		vOpts := append(opts, db.WithValidateField("crm_info"))
		if err := fv(ctx, m.GetCrmInfo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["currency"]; exists {

		vOpts := append(opts, db.WithValidateField("currency"))
		if err := fv(ctx, m.GetCurrency(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["customer"]; exists {

		vOpts := append(opts, db.WithValidateField("customer"))
		if err := fv(ctx, m.GetCustomer(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["customer_contact"]; exists {

		vOpts := append(opts, db.WithValidateField("customer_contact"))
		if err := fv(ctx, m.GetCustomerContact(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["domain"]; exists {

		vOpts := append(opts, db.WithValidateField("domain"))
		if err := fv(ctx, m.GetDomain(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["infraprotect_info"]; exists {

		vOpts := append(opts, db.WithValidateField("infraprotect_info"))
		if err := fv(ctx, m.GetInfraprotectInfo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["locale"]; exists {

		vOpts := append(opts, db.WithValidateField("locale"))
		if err := fv(ctx, m.GetLocale(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["payment_provider_token"]; exists {

		vOpts := append(opts, db.WithValidateField("payment_provider_token"))
		if err := fv(ctx, m.GetPaymentProviderToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["sfdc_subscription_id"]; exists {

		vOpts := append(opts, db.WithValidateField("sfdc_subscription_id"))
		if err := fv(ctx, m.GetSfdcSubscriptionId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["support_plan_name"]; exists {

		vOpts := append(opts, db.WithValidateField("support_plan_name"))
		if err := fv(ctx, m.GetSupportPlanName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tax_exempt"]; exists {

		vOpts := append(opts, db.WithValidateField("tax_exempt"))
		if err := fv(ctx, m.GetTaxExempt(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_accepted"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_accepted"))
		if err := fv(ctx, m.GetTosAccepted(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_accepted_at"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_accepted_at"))
		if err := fv(ctx, m.GetTosAcceptedAt(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tos_version"]; exists {

		vOpts := append(opts, db.WithValidateField("tos_version"))
		if err := fv(ctx, m.GetTosVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["usage_plan_name"]; exists {

		vOpts := append(opts, db.WithValidateField("usage_plan_name"))
		if err := fv(ctx, m.GetUsagePlanName(), vOpts...); err != nil {
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

	vrhLocale := v.LocaleValidationRuleHandler
	rulesLocale := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhLocale(rulesLocale)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.locale: %s", err)
		panic(errMsg)
	}
	v.FldValidators["locale"] = vFn

	vrhType := v.TypeValidationRuleHandler
	rulesType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhType(rulesType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["type"] = vFn

	vrhToken := v.TokenValidationRuleHandler
	rulesToken := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhToken(rulesToken)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.token: %s", err)
		panic(errMsg)
	}
	v.FldValidators["token"] = vFn

	vrhSfdcSubscriptionId := v.SfdcSubscriptionIdValidationRuleHandler
	rulesSfdcSubscriptionId := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
	}
	vFn, err = vrhSfdcSubscriptionId(rulesSfdcSubscriptionId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.sfdc_subscription_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["sfdc_subscription_id"] = vFn

	v.FldValidators["infraprotect_info"] = ves_io_schema_infraprotect_information.GlobalSpecTypeValidator().Validate

	v.FldValidators["crm_info"] = CrmInfoValidator().Validate

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

// Redact squashes sensitive info in m (in-place)
func (m *ReplaceSpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	return nil
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.BillingAddress = f.GetBillingAddress()
	m.Company = f.GetCompany()
	m.CompanyContact = f.GetCompanyContact()
	m.CompanyName = f.GetCompanyName()
	m.ContactNumber = f.GetContactNumber()
	m.CrmInfo = f.GetCrmInfo()
	m.Currency = f.GetCurrency()
	m.Customer = f.GetCustomer()
	m.CustomerContact = f.GetCustomerContact()
	m.Domain = f.GetDomain()
	m.Email = f.GetEmail()
	m.FirstName = f.GetFirstName()
	m.InfraprotectInfo = f.GetInfraprotectInfo()
	m.LastName = f.GetLastName()
	m.Locale = f.GetLocale()
	m.PaymentProviderToken = f.GetPaymentProviderToken()
	m.SfdcSubscriptionId = f.GetSfdcSubscriptionId()
	m.SupportPlanName = f.GetSupportPlanName()
	m.TaxExempt = f.GetTaxExempt()
	m.Token = f.GetToken()
	m.TosAccepted = f.GetTosAccepted()
	m.TosAcceptedAt = f.GetTosAcceptedAt()
	m.TosVersion = f.GetTosVersion()
	m.Type = f.GetType()
	m.UsagePlanName = f.GetUsagePlanName()
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

	f.BillingAddress = m1.BillingAddress
	f.Company = m1.Company
	f.CompanyContact = m1.CompanyContact
	f.CompanyName = m1.CompanyName
	f.ContactNumber = m1.ContactNumber
	f.CrmInfo = m1.CrmInfo
	f.Currency = m1.Currency
	f.Customer = m1.Customer
	f.CustomerContact = m1.CustomerContact
	f.Domain = m1.Domain
	f.Email = m1.Email
	f.FirstName = m1.FirstName
	f.InfraprotectInfo = m1.InfraprotectInfo
	f.LastName = m1.LastName
	f.Locale = m1.Locale
	f.PaymentProviderToken = m1.PaymentProviderToken
	f.SfdcSubscriptionId = m1.SfdcSubscriptionId
	f.SupportPlanName = m1.SupportPlanName
	f.TaxExempt = m1.TaxExempt
	f.Token = m1.Token
	f.TosAccepted = m1.TosAccepted
	f.TosAcceptedAt = m1.TosAcceptedAt
	f.TosVersion = m1.TosVersion
	f.Type = m1.Type
	f.UsagePlanName = m1.UsagePlanName
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

}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}