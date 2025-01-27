// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package k8s_cluster_role

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.k8s_cluster_role.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.k8s_cluster_role.Object"] = ObjectValidator()
	vr["ves.io.schema.k8s_cluster_role.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.k8s_cluster_role.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.k8s_cluster_role.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.k8s_cluster_role.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.k8s_cluster_role.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.k8s_cluster_role.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.k8s_cluster_role.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.k8s_cluster_role.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.k8s_cluster_role.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.k8s_cluster_role.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.k8s_cluster_role.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.k8s_cluster_role.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.NonResourceURLListType"] = NonResourceURLListTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.PolicyRuleListType"] = PolicyRuleListTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.PolicyRuleType"] = PolicyRuleTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.k8s_cluster_role.ResourceListType"] = ResourceListTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.k8s_cluster_role.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.k8s_cluster_role.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.k8s_cluster_role.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.k8s_cluster_role.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.k8s_cluster_role.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.k8s_cluster_role.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.k8s_cluster_role.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.k8s_cluster_role.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.k8s_cluster_role.API.Get"] = []string{
		"object",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.k8s_cluster_role.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.k8s_cluster_role.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.k8s_cluster_role.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.k8s_cluster_role.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.k8s_cluster_role.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.k8s_cluster_role.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.k8s_cluster_role.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.k8s_cluster_role.Object"] = NewCRUDAPIServer

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
