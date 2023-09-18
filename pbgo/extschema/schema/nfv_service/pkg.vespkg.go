// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package nfv_service

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.nfv_service.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.nfv_service.Object"] = ObjectValidator()
	vr["ves.io.schema.nfv_service.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.nfv_service.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.nfv_service.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.nfv_service.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.nfv_service.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.nfv_service.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.nfv_service.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.nfv_service.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.nfv_service.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.nfv_service.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.nfv_service.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.nfv_service.ForceDeleteNFVServiceRequest"] = ForceDeleteNFVServiceRequestValidator()
	vr["ves.io.schema.nfv_service.ForceDeleteNFVServiceResponse"] = ForceDeleteNFVServiceResponseValidator()

	vr["ves.io.schema.nfv_service.MetricData"] = MetricDataValidator()
	vr["ves.io.schema.nfv_service.MetricTypeData"] = MetricTypeDataValidator()
	vr["ves.io.schema.nfv_service.MetricsRequest"] = MetricsRequestValidator()
	vr["ves.io.schema.nfv_service.MetricsResponse"] = MetricsResponseValidator()

	vr["ves.io.schema.nfv_service.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.nfv_service.EndpointRefType"] = EndpointRefTypeValidator()
	vr["ves.io.schema.nfv_service.EndpointServiceReplaceType"] = EndpointServiceReplaceTypeValidator()
	vr["ves.io.schema.nfv_service.EndpointServiceType"] = EndpointServiceTypeValidator()
	vr["ves.io.schema.nfv_service.ExternalNLBInfo"] = ExternalNLBInfoValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAWSBYOLImageType"] = F5BigIpAWSBYOLImageTypeValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAWSMarketPlaceImageType"] = F5BigIpAWSMarketPlaceImageTypeValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAWSReplaceType"] = F5BigIpAWSReplaceTypeValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAWSTGWSiteType"] = F5BigIpAWSTGWSiteTypeValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAWSType"] = F5BigIpAWSTypeValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAWSVPCSiteType"] = F5BigIpAWSVPCSiteTypeValidator()
	vr["ves.io.schema.nfv_service.F5BigIpAppStackBareMetalType"] = F5BigIpAppStackBareMetalTypeValidator()
	vr["ves.io.schema.nfv_service.ForwardingServiceType"] = ForwardingServiceTypeValidator()
	vr["ves.io.schema.nfv_service.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.nfv_service.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.nfv_service.InterfaceDetails"] = InterfaceDetailsValidator()
	vr["ves.io.schema.nfv_service.PANAWSAutoSetupType"] = PANAWSAutoSetupTypeValidator()
	vr["ves.io.schema.nfv_service.PaloAltoAzNodesAWSType"] = PaloAltoAzNodesAWSTypeValidator()
	vr["ves.io.schema.nfv_service.PaloAltoFWAWSReplaceType"] = PaloAltoFWAWSReplaceTypeValidator()
	vr["ves.io.schema.nfv_service.PaloAltoFWAWSType"] = PaloAltoFWAWSTypeValidator()
	vr["ves.io.schema.nfv_service.PaloAltoServiceNodesAWSType"] = PaloAltoServiceNodesAWSTypeValidator()
	vr["ves.io.schema.nfv_service.PanoramaServerType"] = PanoramaServerTypeValidator()
	vr["ves.io.schema.nfv_service.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.nfv_service.SSHKeyType"] = SSHKeyTypeValidator()
	vr["ves.io.schema.nfv_service.SSHManagementNodePorts"] = SSHManagementNodePortsValidator()
	vr["ves.io.schema.nfv_service.SSHManagementType"] = SSHManagementTypeValidator()
	vr["ves.io.schema.nfv_service.ServiceHttpsManagementType"] = ServiceHttpsManagementTypeValidator()
	vr["ves.io.schema.nfv_service.ServiceNodesAWSType"] = ServiceNodesAWSTypeValidator()
	vr["ves.io.schema.nfv_service.ServiceNodesBareMetalType"] = ServiceNodesBareMetalTypeValidator()
	vr["ves.io.schema.nfv_service.SuggestedCommands"] = SuggestedCommandsValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.nfv_service.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.nfv_service.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.nfv_service.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.nfv_service.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.nfv_service.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.nfv_service.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.nfv_service.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.nfv_service.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.nfv_service.API.Create"] = []string{
		"spec.enabled_ssh_access.advertise_on_public",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.nfv_service.API.Create"] = []string{
		"spec.enabled_ssh_access.advertise_on_public",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.nfv_service.API.Create"] = []string{
		"spec.f5_big_ip_aws_service.admin_password.blindfold_secret_info_internal",
		"spec.f5_big_ip_aws_service.admin_password.secret_encoding_type",
		"spec.f5_big_ip_aws_service.admin_password.vault_secret_info",
		"spec.f5_big_ip_aws_service.admin_password.wingman_secret_info",
		"spec.f5_big_ip_aws_service.byol_image.license.blindfold_secret_info_internal",
		"spec.f5_big_ip_aws_service.byol_image.license.secret_encoding_type",
		"spec.f5_big_ip_aws_service.byol_image.license.vault_secret_info",
		"spec.f5_big_ip_aws_service.byol_image.license.wingman_secret_info",
		"spec.https_management.advertise_on_public",
		"spec.https_management.advertise_on_public_default_vip",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.disable_local",
		"spec.https_management.do_not_advertise",
		"spec.https_management.do_not_advertise_on_internet",
		"spec.palo_alto_fw_service.auto_setup.admin_password.blindfold_secret_info_internal",
		"spec.palo_alto_fw_service.auto_setup.admin_password.secret_encoding_type",
		"spec.palo_alto_fw_service.auto_setup.admin_password.vault_secret_info",
		"spec.palo_alto_fw_service.auto_setup.admin_password.wingman_secret_info",
		"spec.palo_alto_fw_service.auto_setup.autogenerated_ssh_keys",
		"spec.palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.blindfold_secret_info_internal",
		"spec.palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.secret_encoding_type",
		"spec.palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.vault_secret_info",
		"spec.palo_alto_fw_service.auto_setup.manual_ssh_keys.private_key.wingman_secret_info",
		"spec.palo_alto_fw_service.panorama_server.authorization_key.blindfold_secret_info_internal",
		"spec.palo_alto_fw_service.panorama_server.authorization_key.secret_encoding_type",
		"spec.palo_alto_fw_service.panorama_server.authorization_key.vault_secret_info",
		"spec.palo_alto_fw_service.panorama_server.authorization_key.wingman_secret_info",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.nfv_service.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.nfv_service.API.Create"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.nfv_service.API.Create"] = "ves.io.schema.nfv_service.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.nfv_service.API.Get"] = []string{
		"create_form.spec.enabled_ssh_access.advertise_on_public",
		"object",
		"replace_form.spec.enabled_ssh_access.advertise_on_public",
		"spec.enabled_ssh_access.advertise_on_public",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.nfv_service.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "create_form.spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
		{
			FieldPath:           "object.spec.gc_spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
		{
			FieldPath:           "replace_form.spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
		{
			FieldPath:           "spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.nfv_service.API.List"] = []string{
		"items.#.get_spec.enabled_ssh_access.advertise_on_public",
		"items.#.object.spec.gc_spec.enabled_ssh_access.advertise_on_public",
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.nfv_service.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
		{
			FieldPath:           "items.#.object.spec.gc_spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
	}

	mdr.RPCDeprecatedRequestFieldsRegistry["ves.io.schema.nfv_service.API.Replace"] = []string{
		"spec.enabled_ssh_access.advertise_on_public",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.nfv_service.API.Replace"] = []string{
		"spec.https_management.advertise_on_public",
		"spec.https_management.advertise_on_public_default_vip",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_sli_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_internet_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_sli.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.blindfold_secret_info_internal",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.secret_encoding_type",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.vault_secret_info",
		"spec.https_management.advertise_on_slo_vip.tls_certificates.#.private_key.wingman_secret_info",
		"spec.https_management.disable_local",
		"spec.https_management.do_not_advertise",
		"spec.https_management.do_not_advertise_on_internet",
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.nfv_service.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.ssh_management_choice",
			AllowedEnvironments: []string{"test"},
		},
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.nfv_service.API.Replace"] = "ves.io.schema.nfv_service.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.nfv_service.API"] = "config"
	sm["ves.io.schema.nfv_service.CustomAPI"] = "config"
	sm["ves.io.schema.nfv_service.CustomDataAPI"] = "data"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

	sm["config"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-config",
		ServiceSelector: "akar\\.gc.*\\",
	}

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.nfv_service.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.nfv_service.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.nfv_service.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.nfv_service.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.nfv_service.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.nfv_service.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.nfv_service.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.nfv_service.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.nfv_service.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.nfv_service.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.nfv_service.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.nfv_service.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.nfv_service.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.nfv_service.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.nfv_service.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.nfv_service.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.nfv_service.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.nfv_service.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.nfv_service.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
		}

	}()

}

func InitializeMDRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	initializeEntryRegistry(mdr)
	initializeValidatorRegistry(mdr.ValidatorRegistry)

	initializeCRUDServiceRegistry(mdr, isExternal)
	if isExternal {
		return
	}

	initializeRPCRegistry(mdr)
	initializeAPIGwServiceSlugsRegistry(mdr.APIGwServiceSlugs)
	initializeP0PolicyRegistry(mdr.P0PolicyRegistry)

}