// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package api_credential

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

func (m *ValidateTokenRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ValidateTokenRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ValidateTokenRequest) DeepCopy() *ValidateTokenRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ValidateTokenRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ValidateTokenRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ValidateTokenRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ValidateTokenRequestValidator().Validate(ctx, m, opts...)
}

type ValidateValidateTokenRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateValidateTokenRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ValidateTokenRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ValidateTokenRequest got type %s", t)
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

	if fv, exists := v.FldValidators["token"]; exists {

		vOpts := append(opts, db.WithValidateField("token"))
		if err := fv(ctx, m.GetToken(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultValidateTokenRequestValidator = func() *ValidateValidateTokenRequest {
	v := &ValidateValidateTokenRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ValidateTokenRequestValidator() db.Validator {
	return DefaultValidateTokenRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ValidateTokenResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ValidateTokenResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ValidateTokenResponse) DeepCopy() *ValidateTokenResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ValidateTokenResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ValidateTokenResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ValidateTokenResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ValidateTokenResponseValidator().Validate(ctx, m, opts...)
}

type ValidateValidateTokenResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateValidateTokenResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ValidateTokenResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ValidateTokenResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["type"]; exists {

		vOpts := append(opts, db.WithValidateField("type"))
		if err := fv(ctx, m.GetType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["user"]; exists {

		vOpts := append(opts, db.WithValidateField("user"))
		if err := fv(ctx, m.GetUser(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["valid"]; exists {

		vOpts := append(opts, db.WithValidateField("valid"))
		if err := fv(ctx, m.GetValid(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultValidateTokenResponseValidator = func() *ValidateValidateTokenResponse {
	v := &ValidateValidateTokenResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ValidateTokenResponseValidator() db.Validator {
	return DefaultValidateTokenResponseValidator
}
