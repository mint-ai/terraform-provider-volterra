// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package service_policy_rule

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.service_policy_rule.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.service_policy_rule.Object"] = ObjectValidator()
	vr["ves.io.schema.service_policy_rule.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.service_policy_rule.CreateRequest"] = CreateRequestValidator()
	vr["ves.io.schema.service_policy_rule.CreateResponse"] = CreateResponseValidator()
	vr["ves.io.schema.service_policy_rule.DeleteRequest"] = DeleteRequestValidator()
	vr["ves.io.schema.service_policy_rule.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.service_policy_rule.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.service_policy_rule.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.service_policy_rule.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.service_policy_rule.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.service_policy_rule.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.service_policy_rule.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.service_policy_rule.ChallengeRuleSpec"] = ChallengeRuleSpecValidator()
	vr["ves.io.schema.service_policy_rule.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.service_policy_rule.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.service_policy_rule.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.service_policy_rule.IPThreatCategoryListType"] = IPThreatCategoryListTypeValidator()
	vr["ves.io.schema.service_policy_rule.RateLimiterRuleSpec"] = RateLimiterRuleSpecValidator()
	vr["ves.io.schema.service_policy_rule.ReplaceSpecType"] = ReplaceSpecTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.service_policy_rule.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.service_policy_rule.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.service_policy_rule.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.service_policy_rule.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.service_policy_rule.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.service_policy_rule.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.service_policy_rule.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.service_policy_rule.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.service_policy_rule.API.Create"] = []string{
		"spec.any_dst_asn",
		"spec.any_dst_ip",
		"spec.challenge_action",
		"spec.client_role",
		"spec.content_rewrite_action",
		"spec.dst_asn_list",
		"spec.dst_asn_matcher",
		"spec.dst_ip_matcher",
		"spec.dst_ip_prefix_list",
		"spec.goto_policy.#",
		"spec.ip_prefix_list.ipv6_prefixes.#",
		"spec.ip_reputation_action",
		"spec.l4_dest_matcher",
		"spec.origin_server_subsets_action",
		"spec.rate_limiter.#",
		"spec.scheme.#",
		"spec.server_selector",
		"spec.shape_protected_endpoint_action",
		"spec.url_matcher",
		"spec.virtual_host_matcher",
		"spec.waf_action.data_guard_control",
		"spec.waf_action.waf_in_monitoring_mode",
	}

	mdr.RPCDeprecatedResponseFieldsRegistry["ves.io.schema.service_policy_rule.API.Get"] = []string{
		"object",
	}

	mdr.RPCHiddenInternalFieldsRegistry["ves.io.schema.service_policy_rule.API.Replace"] = []string{
		"spec.any_dst_asn",
		"spec.any_dst_ip",
		"spec.challenge_action",
		"spec.client_role",
		"spec.content_rewrite_action",
		"spec.dst_asn_list",
		"spec.dst_asn_matcher",
		"spec.dst_ip_matcher",
		"spec.dst_ip_prefix_list",
		"spec.goto_policy.#",
		"spec.ip_prefix_list.ipv6_prefixes.#",
		"spec.ip_reputation_action",
		"spec.l4_dest_matcher",
		"spec.origin_server_subsets_action",
		"spec.rate_limiter.#",
		"spec.scheme.#",
		"spec.server_selector",
		"spec.shape_protected_endpoint_action",
		"spec.url_matcher",
		"spec.virtual_host_matcher",
		"spec.waf_action.data_guard_control",
		"spec.waf_action.waf_in_monitoring_mode",
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.service_policy_rule.API"] = "config"

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
		csr.CRUDSwaggerRegistry["ves.io.schema.service_policy_rule.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.service_policy_rule.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.service_policy_rule.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.service_policy_rule.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.service_policy_rule.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.service_policy_rule.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.service_policy_rule.Object"] = NewCRUDAPIServer

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