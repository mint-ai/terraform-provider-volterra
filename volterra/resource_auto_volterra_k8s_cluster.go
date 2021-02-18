//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_k8s_cluster "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/k8s_cluster"
	ves_io_schema_views "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
)

// resourceVolterraK8SCluster is implementation of Volterra's K8SCluster resources
func resourceVolterraK8SCluster() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraK8SClusterCreate,
		Read:   resourceVolterraK8SClusterRead,
		Update: resourceVolterraK8SClusterUpdate,
		Delete: resourceVolterraK8SClusterDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"use_custom_cluster_role_bindings": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cluster_role_bindings": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"use_default_cluster_role_bindings": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_custom_cluster_role_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"cluster_roles": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"use_default_cluster_roles": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"global_access_enable": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"no_global_access": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"local_access_config": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"local_domain": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"default_port": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"port": {

							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},

			"no_local_access": {

				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_custom_psp_list": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"pod_security_policies": {

							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"namespace": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"tenant": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
					},
				},
			},

			"use_default_psp": {

				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// resourceVolterraK8SClusterCreate creates K8SCluster resource
func resourceVolterraK8SClusterCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_k8s_cluster.CreateSpecType{}
	createReq := &ves_io_schema_k8s_cluster.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	//cluster_role_bindings_choice

	clusterRoleBindingsChoiceTypeFound := false

	if v, ok := d.GetOk("use_custom_cluster_role_bindings"); ok && !clusterRoleBindingsChoiceTypeFound {

		clusterRoleBindingsChoiceTypeFound = true
		clusterRoleBindingsChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_UseCustomClusterRoleBindings{}
		clusterRoleBindingsChoiceInt.UseCustomClusterRoleBindings = &ves_io_schema_k8s_cluster.ClusterRoleBindingListType{}
		createSpec.ClusterRoleBindingsChoice = clusterRoleBindingsChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cluster_role_bindings"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				clusterRoleBindingsIntNew := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				clusterRoleBindingsChoiceInt.UseCustomClusterRoleBindings.ClusterRoleBindings = clusterRoleBindingsIntNew
				for i, ps := range sl {

					crbMapToStrVal := ps.(map[string]interface{})
					clusterRoleBindingsIntNew[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := crbMapToStrVal["name"]; ok && !isIntfNil(v) {
						clusterRoleBindingsIntNew[i].Name = v.(string)
					}

					if v, ok := crbMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						clusterRoleBindingsIntNew[i].Namespace = v.(string)
					}

					if v, ok := crbMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						clusterRoleBindingsIntNew[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("use_default_cluster_role_bindings"); ok && !clusterRoleBindingsChoiceTypeFound {

		clusterRoleBindingsChoiceTypeFound = true

		if v.(bool) {
			clusterRoleBindingsChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_UseDefaultClusterRoleBindings{}
			clusterRoleBindingsChoiceInt.UseDefaultClusterRoleBindings = &ves_io_schema.Empty{}
			createSpec.ClusterRoleBindingsChoice = clusterRoleBindingsChoiceInt
		}

	}

	//cluster_role_choice

	clusterRoleChoiceTypeFound := false

	if v, ok := d.GetOk("use_custom_cluster_role_list"); ok && !clusterRoleChoiceTypeFound {

		clusterRoleChoiceTypeFound = true
		clusterRoleChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_UseCustomClusterRoleList{}
		clusterRoleChoiceInt.UseCustomClusterRoleList = &ves_io_schema_k8s_cluster.ClusterRoleListType{}
		createSpec.ClusterRoleChoice = clusterRoleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cluster_roles"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				clusterRolesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				clusterRoleChoiceInt.UseCustomClusterRoleList.ClusterRoles = clusterRolesInt
				for i, ps := range sl {

					crMapToStrVal := ps.(map[string]interface{})
					clusterRolesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := crMapToStrVal["name"]; ok && !isIntfNil(v) {
						clusterRolesInt[i].Name = v.(string)
					}

					if v, ok := crMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						clusterRolesInt[i].Namespace = v.(string)
					}

					if v, ok := crMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						clusterRolesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("use_default_cluster_roles"); ok && !clusterRoleChoiceTypeFound {

		clusterRoleChoiceTypeFound = true

		if v.(bool) {
			clusterRoleChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_UseDefaultClusterRoles{}
			clusterRoleChoiceInt.UseDefaultClusterRoles = &ves_io_schema.Empty{}
			createSpec.ClusterRoleChoice = clusterRoleChoiceInt
		}

	}

	//global_access_choice

	globalAccessChoiceTypeFound := false

	if v, ok := d.GetOk("global_access_enable"); ok && !globalAccessChoiceTypeFound {

		globalAccessChoiceTypeFound = true

		if v.(bool) {
			globalAccessChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_GlobalAccessEnable{}
			globalAccessChoiceInt.GlobalAccessEnable = &ves_io_schema.Empty{}
			createSpec.GlobalAccessChoice = globalAccessChoiceInt
		}

	}

	if v, ok := d.GetOk("no_global_access"); ok && !globalAccessChoiceTypeFound {

		globalAccessChoiceTypeFound = true

		if v.(bool) {
			globalAccessChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_NoGlobalAccess{}
			globalAccessChoiceInt.NoGlobalAccess = &ves_io_schema.Empty{}
			createSpec.GlobalAccessChoice = globalAccessChoiceInt
		}

	}

	//local_access_choice

	localAccessChoiceTypeFound := false

	if v, ok := d.GetOk("local_access_config"); ok && !localAccessChoiceTypeFound {

		localAccessChoiceTypeFound = true
		localAccessChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_LocalAccessConfig{}
		localAccessChoiceInt.LocalAccessConfig = &ves_io_schema_k8s_cluster.LocalAccessConfigType{}
		createSpec.LocalAccessChoice = localAccessChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["local_domain"]; ok && !isIntfNil(v) {

				localAccessChoiceInt.LocalAccessConfig.LocalDomain = v.(string)
			}

			portChoiceTypeFound := false

			if v, ok := cs["default_port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

				portChoiceTypeFound = true

				if v.(bool) {
					portChoiceInt := &ves_io_schema_k8s_cluster.LocalAccessConfigType_DefaultPort{}
					portChoiceInt.DefaultPort = &ves_io_schema.Empty{}
					localAccessChoiceInt.LocalAccessConfig.PortChoice = portChoiceInt
				}

			}

			if v, ok := cs["port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

				portChoiceTypeFound = true
				portChoiceInt := &ves_io_schema_k8s_cluster.LocalAccessConfigType_Port{}

				localAccessChoiceInt.LocalAccessConfig.PortChoice = portChoiceInt

				portChoiceInt.Port =
					uint32(v.(int))

			}

		}

	}

	if v, ok := d.GetOk("no_local_access"); ok && !localAccessChoiceTypeFound {

		localAccessChoiceTypeFound = true

		if v.(bool) {
			localAccessChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_NoLocalAccess{}
			localAccessChoiceInt.NoLocalAccess = &ves_io_schema.Empty{}
			createSpec.LocalAccessChoice = localAccessChoiceInt
		}

	}

	//pod_security_policy_choice

	podSecurityPolicyChoiceTypeFound := false

	if v, ok := d.GetOk("use_custom_psp_list"); ok && !podSecurityPolicyChoiceTypeFound {

		podSecurityPolicyChoiceTypeFound = true
		podSecurityPolicyChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_UseCustomPspList{}
		podSecurityPolicyChoiceInt.UseCustomPspList = &ves_io_schema_k8s_cluster.PodSecurityPolicyListType{}
		createSpec.PodSecurityPolicyChoice = podSecurityPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["pod_security_policies"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				podSecurityPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				podSecurityPolicyChoiceInt.UseCustomPspList.PodSecurityPolicies = podSecurityPoliciesInt
				for i, ps := range sl {

					pspMapToStrVal := ps.(map[string]interface{})
					podSecurityPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := pspMapToStrVal["name"]; ok && !isIntfNil(v) {
						podSecurityPoliciesInt[i].Name = v.(string)
					}

					if v, ok := pspMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						podSecurityPoliciesInt[i].Namespace = v.(string)
					}

					if v, ok := pspMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						podSecurityPoliciesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("use_default_psp"); ok && !podSecurityPolicyChoiceTypeFound {

		podSecurityPolicyChoiceTypeFound = true

		if v.(bool) {
			podSecurityPolicyChoiceInt := &ves_io_schema_k8s_cluster.CreateSpecType_UseDefaultPsp{}
			podSecurityPolicyChoiceInt.UseDefaultPsp = &ves_io_schema.Empty{}
			createSpec.PodSecurityPolicyChoice = podSecurityPolicyChoiceInt
		}

	}

	log.Printf("[DEBUG] Creating Volterra K8SCluster object with struct: %+v", createReq)

	createK8SClusterResp, err := client.CreateObject(context.Background(), ves_io_schema_k8s_cluster.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating K8SCluster: %s", err)
	}
	d.SetId(createK8SClusterResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraK8SClusterRead(d, meta)
}

func resourceVolterraK8SClusterRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_k8s_cluster.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] K8SCluster %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra K8SCluster %q: %s", d.Id(), err)
	}
	return setK8SClusterFields(client, d, resp)
}

func setK8SClusterFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraK8SClusterUpdate updates K8SCluster resource
func resourceVolterraK8SClusterUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_k8s_cluster.ReplaceSpecType{}
	updateReq := &ves_io_schema_k8s_cluster.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}
	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	clusterRoleBindingsChoiceTypeFound := false

	if v, ok := d.GetOk("use_custom_cluster_role_bindings"); ok && !clusterRoleBindingsChoiceTypeFound {

		clusterRoleBindingsChoiceTypeFound = true
		clusterRoleBindingsChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_UseCustomClusterRoleBindings{}
		clusterRoleBindingsChoiceInt.UseCustomClusterRoleBindings = &ves_io_schema_k8s_cluster.ClusterRoleBindingListType{}
		updateSpec.ClusterRoleBindingsChoice = clusterRoleBindingsChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cluster_role_bindings"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				clusterRoleBindingsIntNew := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				clusterRoleBindingsChoiceInt.UseCustomClusterRoleBindings.ClusterRoleBindings = clusterRoleBindingsIntNew
				for i, ps := range sl {

					crbMapToStrVal := ps.(map[string]interface{})
					clusterRoleBindingsIntNew[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := crbMapToStrVal["name"]; ok && !isIntfNil(v) {
						clusterRoleBindingsIntNew[i].Name = v.(string)
					}

					if v, ok := crbMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						clusterRoleBindingsIntNew[i].Namespace = v.(string)
					}

					if v, ok := crbMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						clusterRoleBindingsIntNew[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("use_default_cluster_role_bindings"); ok && !clusterRoleBindingsChoiceTypeFound {

		clusterRoleBindingsChoiceTypeFound = true

		if v.(bool) {
			clusterRoleBindingsChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_UseDefaultClusterRoleBindings{}
			clusterRoleBindingsChoiceInt.UseDefaultClusterRoleBindings = &ves_io_schema.Empty{}
			updateSpec.ClusterRoleBindingsChoice = clusterRoleBindingsChoiceInt
		}

	}

	clusterRoleChoiceTypeFound := false

	if v, ok := d.GetOk("use_custom_cluster_role_list"); ok && !clusterRoleChoiceTypeFound {

		clusterRoleChoiceTypeFound = true
		clusterRoleChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_UseCustomClusterRoleList{}
		clusterRoleChoiceInt.UseCustomClusterRoleList = &ves_io_schema_k8s_cluster.ClusterRoleListType{}
		updateSpec.ClusterRoleChoice = clusterRoleChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["cluster_roles"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				clusterRolesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				clusterRoleChoiceInt.UseCustomClusterRoleList.ClusterRoles = clusterRolesInt
				for i, ps := range sl {

					crMapToStrVal := ps.(map[string]interface{})
					clusterRolesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := crMapToStrVal["name"]; ok && !isIntfNil(v) {
						clusterRolesInt[i].Name = v.(string)
					}

					if v, ok := crMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						clusterRolesInt[i].Namespace = v.(string)
					}

					if v, ok := crMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						clusterRolesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("use_default_cluster_roles"); ok && !clusterRoleChoiceTypeFound {

		clusterRoleChoiceTypeFound = true

		if v.(bool) {
			clusterRoleChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_UseDefaultClusterRoles{}
			clusterRoleChoiceInt.UseDefaultClusterRoles = &ves_io_schema.Empty{}
			updateSpec.ClusterRoleChoice = clusterRoleChoiceInt
		}

	}

	globalAccessChoiceTypeFound := false

	if v, ok := d.GetOk("global_access_enable"); ok && !globalAccessChoiceTypeFound {

		globalAccessChoiceTypeFound = true

		if v.(bool) {
			globalAccessChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_GlobalAccessEnable{}
			globalAccessChoiceInt.GlobalAccessEnable = &ves_io_schema.Empty{}
			updateSpec.GlobalAccessChoice = globalAccessChoiceInt
		}

	}

	if v, ok := d.GetOk("no_global_access"); ok && !globalAccessChoiceTypeFound {

		globalAccessChoiceTypeFound = true

		if v.(bool) {
			globalAccessChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_NoGlobalAccess{}
			globalAccessChoiceInt.NoGlobalAccess = &ves_io_schema.Empty{}
			updateSpec.GlobalAccessChoice = globalAccessChoiceInt
		}

	}

	localAccessChoiceTypeFound := false

	if v, ok := d.GetOk("local_access_config"); ok && !localAccessChoiceTypeFound {

		localAccessChoiceTypeFound = true
		localAccessChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_LocalAccessConfig{}
		localAccessChoiceInt.LocalAccessConfig = &ves_io_schema_k8s_cluster.LocalAccessConfigType{}
		updateSpec.LocalAccessChoice = localAccessChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["local_domain"]; ok && !isIntfNil(v) {

				localAccessChoiceInt.LocalAccessConfig.LocalDomain = v.(string)
			}

			portChoiceTypeFound := false

			if v, ok := cs["default_port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

				portChoiceTypeFound = true

				if v.(bool) {
					portChoiceInt := &ves_io_schema_k8s_cluster.LocalAccessConfigType_DefaultPort{}
					portChoiceInt.DefaultPort = &ves_io_schema.Empty{}
					localAccessChoiceInt.LocalAccessConfig.PortChoice = portChoiceInt
				}

			}

			if v, ok := cs["port"]; ok && !isIntfNil(v) && !portChoiceTypeFound {

				portChoiceTypeFound = true
				portChoiceInt := &ves_io_schema_k8s_cluster.LocalAccessConfigType_Port{}

				localAccessChoiceInt.LocalAccessConfig.PortChoice = portChoiceInt

				portChoiceInt.Port =
					uint32(v.(int))

			}

		}

	}

	if v, ok := d.GetOk("no_local_access"); ok && !localAccessChoiceTypeFound {

		localAccessChoiceTypeFound = true

		if v.(bool) {
			localAccessChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_NoLocalAccess{}
			localAccessChoiceInt.NoLocalAccess = &ves_io_schema.Empty{}
			updateSpec.LocalAccessChoice = localAccessChoiceInt
		}

	}

	podSecurityPolicyChoiceTypeFound := false

	if v, ok := d.GetOk("use_custom_psp_list"); ok && !podSecurityPolicyChoiceTypeFound {

		podSecurityPolicyChoiceTypeFound = true
		podSecurityPolicyChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_UseCustomPspList{}
		podSecurityPolicyChoiceInt.UseCustomPspList = &ves_io_schema_k8s_cluster.PodSecurityPolicyListType{}
		updateSpec.PodSecurityPolicyChoice = podSecurityPolicyChoiceInt

		sl := v.(*schema.Set).List()
		for _, set := range sl {
			cs := set.(map[string]interface{})

			if v, ok := cs["pod_security_policies"]; ok && !isIntfNil(v) {

				sl := v.([]interface{})
				podSecurityPoliciesInt := make([]*ves_io_schema_views.ObjectRefType, len(sl))
				podSecurityPolicyChoiceInt.UseCustomPspList.PodSecurityPolicies = podSecurityPoliciesInt
				for i, ps := range sl {

					pspMapToStrVal := ps.(map[string]interface{})
					podSecurityPoliciesInt[i] = &ves_io_schema_views.ObjectRefType{}

					if v, ok := pspMapToStrVal["name"]; ok && !isIntfNil(v) {
						podSecurityPoliciesInt[i].Name = v.(string)
					}

					if v, ok := pspMapToStrVal["namespace"]; ok && !isIntfNil(v) {
						podSecurityPoliciesInt[i].Namespace = v.(string)
					}

					if v, ok := pspMapToStrVal["tenant"]; ok && !isIntfNil(v) {
						podSecurityPoliciesInt[i].Tenant = v.(string)
					}

				}

			}

		}

	}

	if v, ok := d.GetOk("use_default_psp"); ok && !podSecurityPolicyChoiceTypeFound {

		podSecurityPolicyChoiceTypeFound = true

		if v.(bool) {
			podSecurityPolicyChoiceInt := &ves_io_schema_k8s_cluster.ReplaceSpecType_UseDefaultPsp{}
			podSecurityPolicyChoiceInt.UseDefaultPsp = &ves_io_schema.Empty{}
			updateSpec.PodSecurityPolicyChoice = podSecurityPolicyChoiceInt
		}

	}

	log.Printf("[DEBUG] Updating Volterra K8SCluster obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_k8s_cluster.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating K8SCluster: %s", err)
	}

	return resourceVolterraK8SClusterRead(d, meta)
}

func resourceVolterraK8SClusterDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_k8s_cluster.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] K8SCluster %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra K8SCluster before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra K8SCluster obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_k8s_cluster.ObjectType, namespace, name)
}