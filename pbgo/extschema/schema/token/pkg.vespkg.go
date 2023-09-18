// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package token

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.token.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.token.Object"] = ObjectValidator()
	vr["ves.io.schema.token.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.token.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.token.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.token.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.token.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.token.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.token.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.token.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.token.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.token.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.token.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.token.ObjectChangeResp"] = ObjectChangeRespValidator()
	vr["ves.io.schema.token.StateReq"] = StateReqValidator()

	vr["ves.io.schema.token.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.token.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.token.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.token.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.token.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.token.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.token.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.token.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.token.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.token.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.token.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.token.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.token.API.Get"] = []string{
		"object",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.token.API"] = "register"
	sm["ves.io.schema.token.CustomAPI"] = "register"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

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
		csr.CRUDSwaggerRegistry["ves.io.schema.token.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.token.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.token.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.token.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.token.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.token.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.token.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.token.Object"] = CustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.token.CustomAPI"] = NewCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.token.CustomAPI"] = NewCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.token.CustomAPI"] = RegisterCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.token.CustomAPI"] = RegisterGwCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.token.CustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomAPIServer(svc)
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