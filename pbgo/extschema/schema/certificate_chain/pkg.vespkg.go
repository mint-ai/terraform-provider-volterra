// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package certificate_chain

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.certificate_chain.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.certificate_chain.Object"] = ObjectValidator()
	vr["ves.io.schema.certificate_chain.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.certificate_chain.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.certificate_chain.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.certificate_chain.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.certificate_chain.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.certificate_chain.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.certificate_chain.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.certificate_chain.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.certificate_chain.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.certificate_chain.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.certificate_chain.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.certificate_chain.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.certificate_chain.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.certificate_chain.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.certificate_chain.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.certificate_chain.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.certificate_chain.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.certificate_chain.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.certificate_chain.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.certificate_chain.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.certificate_chain.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.certificate_chain.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.certificate_chain.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.certificate_chain.API.Get"] = []string{
		"object",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.certificate_chain.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.certificate_chain.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.certificate_chain.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.certificate_chain.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.certificate_chain.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.certificate_chain.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.certificate_chain.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.certificate_chain.Object"] = NewCRUDAPIServer

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
