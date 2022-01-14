//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package virtual_host

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_virtual_host_dns_info "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host_dns_info"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GetDnsInfoRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDnsInfoRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDnsInfoRequest) DeepCopy() *GetDnsInfoRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDnsInfoRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDnsInfoRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDnsInfoRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDnsInfoRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetDnsInfoRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDnsInfoRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDnsInfoRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDnsInfoRequest got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetDnsInfoRequestValidator = func() *ValidateGetDnsInfoRequest {
	v := &ValidateGetDnsInfoRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetDnsInfoRequestValidator() db.Validator {
	return DefaultGetDnsInfoRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetDnsInfoResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetDnsInfoResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetDnsInfoResponse) DeepCopy() *GetDnsInfoResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetDnsInfoResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetDnsInfoResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetDnsInfoResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetDnsInfoResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetDnsInfoResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetDnsInfoDRefInfo()

}

// GetDRefInfo for the field's type
func (m *GetDnsInfoResponse) GetDnsInfoDRefInfo() ([]db.DRefInfo, error) {
	if m.GetDnsInfo() == nil {
		return nil, nil
	}

	drInfos, err := m.GetDnsInfo().GetDRefInfo()
	if err != nil {
		return nil, errors.Wrap(err, "GetDnsInfo().GetDRefInfo() FAILED")
	}
	for i := range drInfos {
		dri := &drInfos[i]
		dri.DRField = "dns_info." + dri.DRField
	}
	return drInfos, err

}

type ValidateGetDnsInfoResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetDnsInfoResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetDnsInfoResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetDnsInfoResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["dns_info"]; exists {

		vOpts := append(opts, db.WithValidateField("dns_info"))
		if err := fv(ctx, m.GetDnsInfo(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetDnsInfoResponseValidator = func() *ValidateGetDnsInfoResponse {
	v := &ValidateGetDnsInfoResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["dns_info"] = ves_io_schema_virtual_host_dns_info.GlobalSpecTypeValidator().Validate

	return v
}()

func GetDnsInfoResponseValidator() db.Validator {
	return DefaultGetDnsInfoResponseValidator
}
