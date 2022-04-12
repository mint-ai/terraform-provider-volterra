//
// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package virtual_k8s

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
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

func (m *CreateSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetDefaultFlavorRefDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetDefaultFlavorRefDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetVsiteRefsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetVsiteRefsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *CreateSpecType) GetDefaultFlavorRefDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("workload_flavor.Object")
	dri := db.DRefInfo{
		RefdType:   "workload_flavor.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "default_flavor_ref",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetDefaultFlavorRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CreateSpecType) GetDefaultFlavorRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "workload_flavor.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: workload_flavor")
	}

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "workload_flavor.Object",
		Tenant:    vref.Tenant,
		Namespace: vref.Namespace,
		Name:      vref.Name,
	}
	refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
	if err != nil {
		return nil, errors.Wrap(err, "Getting referred entry")
	}
	if refdEnt != nil {
		entries = append(entries, refdEnt)
	}

	return entries, nil
}

func (m *CreateSpecType) GetVsiteRefsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetVsiteRefs()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("CreateSpecType.vsite_refs[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "virtual_site.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "vsite_refs",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetVsiteRefsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *CreateSpecType) GetVsiteRefsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "virtual_site.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: virtual_site")
	}
	for _, ref := range m.GetVsiteRefs() {
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

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) VsiteRefsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vsite_refs")
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
			return errors.Wrap(err, "repeated vsite_refs")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vsite_refs")
		}
		return nil
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

	if fv, exists := v.FldValidators["default_flavor_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("default_flavor_ref"))
		if err := fv(ctx, m.GetDefaultFlavorRef(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetServiceIsolationChoice().(type) {
	case *CreateSpecType_Isolated:
		if fv, exists := v.FldValidators["service_isolation_choice.isolated"]; exists {
			val := m.GetServiceIsolationChoice().(*CreateSpecType_Isolated).Isolated
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("isolated"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *CreateSpecType_Disabled:
		if fv, exists := v.FldValidators["service_isolation_choice.disabled"]; exists {
			val := m.GetServiceIsolationChoice().(*CreateSpecType_Disabled).Disabled
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("disabled"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["vsite_refs"]; exists {
		vOpts := append(opts, db.WithValidateField("vsite_refs"))
		if err := fv(ctx, m.GetVsiteRefs(), vOpts...); err != nil {
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

	vrhVsiteRefs := v.VsiteRefsValidationRuleHandler
	rulesVsiteRefs := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "32",
	}
	vFn, err = vrhVsiteRefs(rulesVsiteRefs)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.vsite_refs: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vsite_refs"] = vFn

	v.FldValidators["default_flavor_ref"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

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

func (m *GetSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetDefaultFlavorRefDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetDefaultFlavorRefDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetVsiteRefsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetVsiteRefsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *GetSpecType) GetDefaultFlavorRefDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("workload_flavor.Object")
	dri := db.DRefInfo{
		RefdType:   "workload_flavor.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "default_flavor_ref",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetDefaultFlavorRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetSpecType) GetDefaultFlavorRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "workload_flavor.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: workload_flavor")
	}

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "workload_flavor.Object",
		Tenant:    vref.Tenant,
		Namespace: vref.Namespace,
		Name:      vref.Name,
	}
	refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
	if err != nil {
		return nil, errors.Wrap(err, "Getting referred entry")
	}
	if refdEnt != nil {
		entries = append(entries, refdEnt)
	}

	return entries, nil
}

func (m *GetSpecType) GetVsiteRefsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetVsiteRefs()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GetSpecType.vsite_refs[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "virtual_site.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "vsite_refs",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetVsiteRefsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetSpecType) GetVsiteRefsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "virtual_site.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: virtual_site")
	}
	for _, ref := range m.GetVsiteRefs() {
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

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) VsiteRefsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vsite_refs")
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
			return errors.Wrap(err, "repeated vsite_refs")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vsite_refs")
		}
		return nil
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

	if fv, exists := v.FldValidators["default_flavor_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("default_flavor_ref"))
		if err := fv(ctx, m.GetDefaultFlavorRef(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetServiceIsolationChoice().(type) {
	case *GetSpecType_Isolated:
		if fv, exists := v.FldValidators["service_isolation_choice.isolated"]; exists {
			val := m.GetServiceIsolationChoice().(*GetSpecType_Isolated).Isolated
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("isolated"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GetSpecType_Disabled:
		if fv, exists := v.FldValidators["service_isolation_choice.disabled"]; exists {
			val := m.GetServiceIsolationChoice().(*GetSpecType_Disabled).Disabled
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("disabled"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["vsite_refs"]; exists {
		vOpts := append(opts, db.WithValidateField("vsite_refs"))
		if err := fv(ctx, m.GetVsiteRefs(), vOpts...); err != nil {
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

	vrhVsiteRefs := v.VsiteRefsValidationRuleHandler
	rulesVsiteRefs := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "32",
	}
	vFn, err = vrhVsiteRefs(rulesVsiteRefs)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.vsite_refs: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vsite_refs"] = vFn

	v.FldValidators["default_flavor_ref"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

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

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetDefaultFlavorRefDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetDefaultFlavorRefDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetVsiteRefsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetVsiteRefsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *GlobalSpecType) GetDefaultFlavorRefDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("workload_flavor.Object")
	dri := db.DRefInfo{
		RefdType:   "workload_flavor.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "default_flavor_ref",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetDefaultFlavorRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetDefaultFlavorRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "workload_flavor.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: workload_flavor")
	}

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "workload_flavor.Object",
		Tenant:    vref.Tenant,
		Namespace: vref.Namespace,
		Name:      vref.Name,
	}
	refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
	if err != nil {
		return nil, errors.Wrap(err, "Getting referred entry")
	}
	if refdEnt != nil {
		entries = append(entries, refdEnt)
	}

	return entries, nil
}

func (m *GlobalSpecType) GetVsiteRefsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetVsiteRefs()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.vsite_refs[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "virtual_site.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "vsite_refs",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetVsiteRefsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetVsiteRefsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "virtual_site.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: virtual_site")
	}
	for _, ref := range m.GetVsiteRefs() {
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

func (v *ValidateGlobalSpecType) VsiteRefsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vsite_refs")
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
			return errors.Wrap(err, "repeated vsite_refs")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vsite_refs")
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

	if fv, exists := v.FldValidators["default_flavor_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("default_flavor_ref"))
		if err := fv(ctx, m.GetDefaultFlavorRef(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetServiceIsolationChoice().(type) {
	case *GlobalSpecType_Isolated:
		if fv, exists := v.FldValidators["service_isolation_choice.isolated"]; exists {
			val := m.GetServiceIsolationChoice().(*GlobalSpecType_Isolated).Isolated
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("isolated"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_Disabled:
		if fv, exists := v.FldValidators["service_isolation_choice.disabled"]; exists {
			val := m.GetServiceIsolationChoice().(*GlobalSpecType_Disabled).Disabled
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("disabled"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["vsite_refs"]; exists {
		vOpts := append(opts, db.WithValidateField("vsite_refs"))
		if err := fv(ctx, m.GetVsiteRefs(), vOpts...); err != nil {
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

	vrhVsiteRefs := v.VsiteRefsValidationRuleHandler
	rulesVsiteRefs := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "32",
	}
	vFn, err = vrhVsiteRefs(rulesVsiteRefs)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.vsite_refs: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vsite_refs"] = vFn

	v.FldValidators["default_flavor_ref"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

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

func (m *ReplaceSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetDefaultFlavorRefDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetDefaultFlavorRefDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	if fdrInfos, err := m.GetVsiteRefsDRefInfo(); err != nil {
		return nil, errors.Wrap(err, "GetVsiteRefsDRefInfo() FAILED")
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil

}

func (m *ReplaceSpecType) GetDefaultFlavorRefDRefInfo() ([]db.DRefInfo, error) {

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	vdRef := db.NewDirectRefForView(vref)
	vdRef.SetKind("workload_flavor.Object")
	dri := db.DRefInfo{
		RefdType:   "workload_flavor.Object",
		RefdTenant: vref.Tenant,
		RefdNS:     vref.Namespace,
		RefdName:   vref.Name,
		DRField:    "default_flavor_ref",
		Ref:        vdRef,
	}
	return []db.DRefInfo{dri}, nil

}

// GetDefaultFlavorRefDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *ReplaceSpecType) GetDefaultFlavorRefDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "workload_flavor.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: workload_flavor")
	}

	vref := m.GetDefaultFlavorRef()
	if vref == nil {
		return nil, nil
	}
	ref := &ves_io_schema.ObjectRefType{
		Kind:      "workload_flavor.Object",
		Tenant:    vref.Tenant,
		Namespace: vref.Namespace,
		Name:      vref.Name,
	}
	refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
	if err != nil {
		return nil, errors.Wrap(err, "Getting referred entry")
	}
	if refdEnt != nil {
		entries = append(entries, refdEnt)
	}

	return entries, nil
}

func (m *ReplaceSpecType) GetVsiteRefsDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetVsiteRefs()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("ReplaceSpecType.vsite_refs[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "virtual_site.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "vsite_refs",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetVsiteRefsDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *ReplaceSpecType) GetVsiteRefsDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "virtual_site.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: virtual_site")
	}
	for _, ref := range m.GetVsiteRefs() {
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

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) VsiteRefsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for vsite_refs")
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
			return errors.Wrap(err, "repeated vsite_refs")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items vsite_refs")
		}
		return nil
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

	if fv, exists := v.FldValidators["default_flavor_ref"]; exists {

		vOpts := append(opts, db.WithValidateField("default_flavor_ref"))
		if err := fv(ctx, m.GetDefaultFlavorRef(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetServiceIsolationChoice().(type) {
	case *ReplaceSpecType_Isolated:
		if fv, exists := v.FldValidators["service_isolation_choice.isolated"]; exists {
			val := m.GetServiceIsolationChoice().(*ReplaceSpecType_Isolated).Isolated
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("isolated"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *ReplaceSpecType_Disabled:
		if fv, exists := v.FldValidators["service_isolation_choice.disabled"]; exists {
			val := m.GetServiceIsolationChoice().(*ReplaceSpecType_Disabled).Disabled
			vOpts := append(opts,
				db.WithValidateField("service_isolation_choice"),
				db.WithValidateField("disabled"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["vsite_refs"]; exists {
		vOpts := append(opts, db.WithValidateField("vsite_refs"))
		if err := fv(ctx, m.GetVsiteRefs(), vOpts...); err != nil {
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

	vrhVsiteRefs := v.VsiteRefsValidationRuleHandler
	rulesVsiteRefs := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "32",
	}
	vFn, err = vrhVsiteRefs(rulesVsiteRefs)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.vsite_refs: %s", err)
		panic(errMsg)
	}
	v.FldValidators["vsite_refs"] = vFn

	v.FldValidators["default_flavor_ref"] = ves_io_schema_views.ObjectRefTypeValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// create setters in CreateSpecType from GlobalSpecType for oneof fields
func (r *CreateSpecType) SetServiceIsolationChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.ServiceIsolationChoice.(type) {
	case nil:
		o.ServiceIsolationChoice = nil

	case *CreateSpecType_Disabled:
		o.ServiceIsolationChoice = &GlobalSpecType_Disabled{Disabled: of.Disabled}

	case *CreateSpecType_Isolated:
		o.ServiceIsolationChoice = &GlobalSpecType_Isolated{Isolated: of.Isolated}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *CreateSpecType) GetServiceIsolationChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.ServiceIsolationChoice.(type) {
	case nil:
		r.ServiceIsolationChoice = nil

	case *GlobalSpecType_Disabled:
		r.ServiceIsolationChoice = &CreateSpecType_Disabled{Disabled: of.Disabled}

	case *GlobalSpecType_Isolated:
		r.ServiceIsolationChoice = &CreateSpecType_Isolated{Isolated: of.Isolated}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.DefaultFlavorRef = f.GetDefaultFlavorRef()
	m.GetServiceIsolationChoiceFromGlobalSpecType(f)
	m.VsiteRefs = f.GetVsiteRefs()
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.DefaultFlavorRef = m1.DefaultFlavorRef
	m1.SetServiceIsolationChoiceToGlobalSpecType(f)
	f.VsiteRefs = m1.VsiteRefs
}

// create setters in GetSpecType from GlobalSpecType for oneof fields
func (r *GetSpecType) SetServiceIsolationChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.ServiceIsolationChoice.(type) {
	case nil:
		o.ServiceIsolationChoice = nil

	case *GetSpecType_Disabled:
		o.ServiceIsolationChoice = &GlobalSpecType_Disabled{Disabled: of.Disabled}

	case *GetSpecType_Isolated:
		o.ServiceIsolationChoice = &GlobalSpecType_Isolated{Isolated: of.Isolated}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *GetSpecType) GetServiceIsolationChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.ServiceIsolationChoice.(type) {
	case nil:
		r.ServiceIsolationChoice = nil

	case *GlobalSpecType_Disabled:
		r.ServiceIsolationChoice = &GetSpecType_Disabled{Disabled: of.Disabled}

	case *GlobalSpecType_Isolated:
		r.ServiceIsolationChoice = &GetSpecType_Isolated{Isolated: of.Isolated}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.DefaultFlavorRef = f.GetDefaultFlavorRef()
	m.GetServiceIsolationChoiceFromGlobalSpecType(f)
	m.VsiteRefs = f.GetVsiteRefs()
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.DefaultFlavorRef = m1.DefaultFlavorRef
	m1.SetServiceIsolationChoiceToGlobalSpecType(f)
	f.VsiteRefs = m1.VsiteRefs
}

// create setters in ReplaceSpecType from GlobalSpecType for oneof fields
func (r *ReplaceSpecType) SetServiceIsolationChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.ServiceIsolationChoice.(type) {
	case nil:
		o.ServiceIsolationChoice = nil

	case *ReplaceSpecType_Disabled:
		o.ServiceIsolationChoice = &GlobalSpecType_Disabled{Disabled: of.Disabled}

	case *ReplaceSpecType_Isolated:
		o.ServiceIsolationChoice = &GlobalSpecType_Isolated{Isolated: of.Isolated}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *ReplaceSpecType) GetServiceIsolationChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.ServiceIsolationChoice.(type) {
	case nil:
		r.ServiceIsolationChoice = nil

	case *GlobalSpecType_Disabled:
		r.ServiceIsolationChoice = &ReplaceSpecType_Disabled{Disabled: of.Disabled}

	case *GlobalSpecType_Isolated:
		r.ServiceIsolationChoice = &ReplaceSpecType_Isolated{Isolated: of.Isolated}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	if f == nil {
		return
	}
	m.DefaultFlavorRef = f.GetDefaultFlavorRef()
	m.GetServiceIsolationChoiceFromGlobalSpecType(f)
	m.VsiteRefs = f.GetVsiteRefs()
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m1 := m.DeepCopy()
	_ = m1
	if f == nil {
		return
	}
	f.DefaultFlavorRef = m1.DefaultFlavorRef
	m1.SetServiceIsolationChoiceToGlobalSpecType(f)
	f.VsiteRefs = m1.VsiteRefs
}
