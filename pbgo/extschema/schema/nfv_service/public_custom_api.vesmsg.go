// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package nfv_service

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

func (m *ForceDeleteNFVServiceRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForceDeleteNFVServiceRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForceDeleteNFVServiceRequest) DeepCopy() *ForceDeleteNFVServiceRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForceDeleteNFVServiceRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForceDeleteNFVServiceRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForceDeleteNFVServiceRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForceDeleteNFVServiceRequestValidator().Validate(ctx, m, opts...)
}

type ValidateForceDeleteNFVServiceRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForceDeleteNFVServiceRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForceDeleteNFVServiceRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForceDeleteNFVServiceRequest got type %s", t)
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

	return nil
}

// Well-known symbol for default validator implementation
var DefaultForceDeleteNFVServiceRequestValidator = func() *ValidateForceDeleteNFVServiceRequest {
	v := &ValidateForceDeleteNFVServiceRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForceDeleteNFVServiceRequestValidator() db.Validator {
	return DefaultForceDeleteNFVServiceRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ForceDeleteNFVServiceResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForceDeleteNFVServiceResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForceDeleteNFVServiceResponse) DeepCopy() *ForceDeleteNFVServiceResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForceDeleteNFVServiceResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForceDeleteNFVServiceResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForceDeleteNFVServiceResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForceDeleteNFVServiceResponseValidator().Validate(ctx, m, opts...)
}

type ValidateForceDeleteNFVServiceResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForceDeleteNFVServiceResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForceDeleteNFVServiceResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForceDeleteNFVServiceResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultForceDeleteNFVServiceResponseValidator = func() *ValidateForceDeleteNFVServiceResponse {
	v := &ValidateForceDeleteNFVServiceResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForceDeleteNFVServiceResponseValidator() db.Validator {
	return DefaultForceDeleteNFVServiceResponseValidator
}
