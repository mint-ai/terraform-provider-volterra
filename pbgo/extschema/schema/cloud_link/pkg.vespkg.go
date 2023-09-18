// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cloud_link

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.cloud_link.AWSStatusType"] = AWSStatusTypeValidator()
	vr["ves.io.schema.cloud_link.AzureStatusType"] = AzureStatusTypeValidator()
	vr["ves.io.schema.cloud_link.BGPPeerType"] = BGPPeerTypeValidator()
	vr["ves.io.schema.cloud_link.DirectConnectConnectionStatusType"] = DirectConnectConnectionStatusTypeValidator()
	vr["ves.io.schema.cloud_link.DirectConnectGatewayStatusType"] = DirectConnectGatewayStatusTypeValidator()
	vr["ves.io.schema.cloud_link.SpecType"] = SpecTypeValidator()
	vr["ves.io.schema.cloud_link.VirtualInterfaceStatusType"] = VirtualInterfaceStatusTypeValidator()

	vr["ves.io.schema.cloud_link.Object"] = ObjectValidator()
	vr["ves.io.schema.cloud_link.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.cloud_link.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.cloud_link.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.cloud_link.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.cloud_link.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.cloud_link.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.cloud_link.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.cloud_link.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.cloud_link.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.cloud_link.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.cloud_link.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.cloud_link.AWSBYOCListType"] = AWSBYOCListTypeValidator()
	vr["ves.io.schema.cloud_link.AWSBYOCType"] = AWSBYOCTypeValidator()
	vr["ves.io.schema.cloud_link.AWSF5XCManagedType"] = AWSF5XCManagedTypeValidator()
	vr["ves.io.schema.cloud_link.AWSType"] = AWSTypeValidator()
	vr["ves.io.schema.cloud_link.AzureType"] = AzureTypeValidator()
	vr["ves.io.schema.cloud_link.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.cloud_link.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.cloud_link.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.cloud_link.Ipv4Type"] = Ipv4TypeValidator()
	vr["ves.io.schema.cloud_link.Ipv6Type"] = Ipv6TypeValidator()
	vr["ves.io.schema.cloud_link.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.cloud_link.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.cloud_link.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cloud_link.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.cloud_link.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.cloud_link.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.cloud_link.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.cloud_link.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.cloud_link.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.cloud_link.API.Create"] = []string{
		"spec.aws.byoc.connections.#.auth_key.blindfold_secret_info_internal",
		"spec.aws.byoc.connections.#.auth_key.secret_encoding_type",
		"spec.aws.byoc.connections.#.auth_key.vault_secret_info",
		"spec.aws.byoc.connections.#.auth_key.wingman_secret_info",
		"spec.aws.byoc.connections.#.ipv6",
		"spec.aws.f5xc_managed",
		"spec.azure",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.cloud_link.API.Create"] = "ves.io.schema.cloud_link.CreateRequest"

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.cloud_link.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.cloud_link.API.Replace"] = []string{
		"spec.aws.byoc.connections.#.auth_key.blindfold_secret_info_internal",
		"spec.aws.byoc.connections.#.auth_key.secret_encoding_type",
		"spec.aws.byoc.connections.#.auth_key.vault_secret_info",
		"spec.aws.byoc.connections.#.auth_key.wingman_secret_info",
		"spec.aws.byoc.connections.#.ipv6",
		"spec.aws.f5xc_managed",
		"spec.azure",
	}

	mdr.RPCConfidentialRequestRegistry["ves.io.schema.cloud_link.API.Replace"] = "ves.io.schema.cloud_link.ReplaceRequest"

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.cloud_link.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.cloud_link.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.cloud_link.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.cloud_link.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.cloud_link.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.cloud_link.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.cloud_link.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.cloud_link.Object"] = NewCRUDAPIServer

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