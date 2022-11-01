// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package vesenv

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

func (m *AddonServiceChoice) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AddonServiceChoice) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AddonServiceChoice) DeepCopy() *AddonServiceChoice {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AddonServiceChoice{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AddonServiceChoice) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AddonServiceChoice) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AddonServiceChoiceValidator().Validate(ctx, m, opts...)
}

type ValidateAddonServiceChoice struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAddonServiceChoice) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AddonServiceChoice)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AddonServiceChoice got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	switch m.GetChoice().(type) {
	case *AddonServiceChoice_F5XcBase:
		if fv, exists := v.FldValidators["choice.f5xc_base"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_F5XcBase).F5XcBase
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("f5xc_base"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_VesIoSupportManagement:
		if fv, exists := v.FldValidators["choice.ves_io_support_management"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_VesIoSupportManagement).VesIoSupportManagement
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ves_io_support_management"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_VesIoVolterraDefault:
		if fv, exists := v.FldValidators["choice.ves_io_volterra_default"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_VesIoVolterraDefault).VesIoVolterraDefault
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ves_io_volterra_default"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_VesIoTenantManagement:
		if fv, exists := v.FldValidators["choice.ves_io_tenant_management"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_VesIoTenantManagement).VesIoTenantManagement
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ves_io_tenant_management"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_VesIoScim:
		if fv, exists := v.FldValidators["choice.ves_io_scim"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_VesIoScim).VesIoScim
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("ves_io_scim"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_ShapeBot:
		if fv, exists := v.FldValidators["choice.shape_bot"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_ShapeBot).ShapeBot
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("shape_bot"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_ShapeRecognize:
		if fv, exists := v.FldValidators["choice.shape_recognize"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_ShapeRecognize).ShapeRecognize
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("shape_recognize"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_AidataBfdp:
		if fv, exists := v.FldValidators["choice.aidata_bfdp"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_AidataBfdp).AidataBfdp
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("aidata_bfdp"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_LilacCdn:
		if fv, exists := v.FldValidators["choice.lilac_cdn"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_LilacCdn).LilacCdn
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("lilac_cdn"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_NginxMgmtSuite:
		if fv, exists := v.FldValidators["choice.nginx_mgmt_suite"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_NginxMgmtSuite).NginxMgmtSuite
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("nginx_mgmt_suite"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_SyntheticMonitor:
		if fv, exists := v.FldValidators["choice.synthetic_monitor"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_SyntheticMonitor).SyntheticMonitor
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("synthetic_monitor"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_Safeap:
		if fv, exists := v.FldValidators["choice.safeap"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_Safeap).Safeap
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("safeap"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_AipConnector:
		if fv, exists := v.FldValidators["choice.aip_connector"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_AipConnector).AipConnector
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("aip_connector"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_ClientSideDefense:
		if fv, exists := v.FldValidators["choice.client_side_defense"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_ClientSideDefense).ClientSideDefense
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("client_side_defense"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_ShapeBotAdvanced:
		if fv, exists := v.FldValidators["choice.shape_bot_advanced"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_ShapeBotAdvanced).ShapeBotAdvanced
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("shape_bot_advanced"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AddonServiceChoice_ShapeBotPremium:
		if fv, exists := v.FldValidators["choice.shape_bot_premium"]; exists {
			val := m.GetChoice().(*AddonServiceChoice_ShapeBotPremium).ShapeBotPremium
			vOpts := append(opts,
				db.WithValidateField("choice"),
				db.WithValidateField("shape_bot_premium"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAddonServiceChoiceValidator = func() *ValidateAddonServiceChoice {
	v := &ValidateAddonServiceChoice{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func AddonServiceChoiceValidator() db.Validator {
	return DefaultAddonServiceChoiceValidator
}
