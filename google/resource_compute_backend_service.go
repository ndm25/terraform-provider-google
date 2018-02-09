// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    AUTO GENERATED CODE     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by terraform-codegen and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in README.md and
//     CONTRIBUTING.md located at the root of this package.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	compute "google.golang.org/api/compute/v1"
)

func resourceComputeBackendService() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeBackendServiceCreate,
		Read:   resourceComputeBackendServiceRead,
		Update: resourceComputeBackendServiceUpdate,
		Delete: resourceComputeBackendServiceDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeBackendServiceImport,
		},

		Schema: map[string]*schema.Schema{
			"affinity_cookie_ttl_sec": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"backends": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"balancing_mode": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"capacity_scaler": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"description": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"group": {
							Type:             schema.TypeString,
							Optional:         true,
							DiffSuppressFunc: compareSelfLinkOrResourceName,
						},
						"max_connections": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_connections_per_instance": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_rate": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"max_rate_per_instance": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
						"max_utilization": {
							Type:     schema.TypeFloat,
							Optional: true,
						},
					},
				},
			},
			"cdn_policy": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cache_key_policy": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"include_host": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"include_protocol": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"include_query_string": {
										Type:     schema.TypeBool,
										Optional: true,
									},
									"query_string_blacklist": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"query_string_whitelist": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
			"connection_draining": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"draining_timeout_sec": {
							Type:     schema.TypeInt,
							Optional: true,
						},
					},
				},
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_cdn": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"health_checks": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"port_name": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"protocol": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"region": {
				Type:             schema.TypeString,
				Optional:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
			},
			"session_affinity": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"timeout_sec": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"creation_timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"id": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"project": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceComputeBackendServiceCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"affinityCookieTtlSec": expandComputeBackendServiceAffinityCookieTtlSec(d.Get("affinity_cookie_ttl_sec")),
		"backends":             expandComputeBackendServiceBackends(d.Get("backends")),
		"cdnPolicy":            expandComputeBackendServiceCdnPolicy(d.Get("cdn_policy")),
		"connectionDraining":   expandComputeBackendServiceConnectionDraining(d.Get("connection_draining")),
		"description":          expandComputeBackendServiceDescription(d.Get("description")),
		"enableCDN":            expandComputeBackendServiceEnableCDN(d.Get("enable_cdn")),
		"healthChecks":         expandComputeBackendServiceHealthChecks(d.Get("health_checks")),
		"name":                 expandComputeBackendServiceName(d.Get("name")),
		"portName":             expandComputeBackendServicePortName(d.Get("port_name")),
		"protocol":             expandComputeBackendServiceProtocol(d.Get("protocol")),
		"region":               expandComputeBackendServiceRegion(d.Get("region")),
		"sessionAffinity":      expandComputeBackendServiceSessionAffinity(d.Get("session_affinity")),
		"timeoutSec":           expandComputeBackendServiceTimeoutSec(d.Get("timeout_sec")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendServices")
	if err != nil {
		return err
	}
	res, err := Post(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error creating BackendService: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	waitErr := computeOperationWait(config.clientCompute, op, project, "Creating BackendService")
	if waitErr != nil {
		// The resource didn't actually create
		d.SetId("")
		return waitErr
	}

	return resourceComputeBackendServiceRead(d, meta)
}

func resourceComputeBackendServiceRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendServices/{{name}}")
	if err != nil {
		return err
	}
	res, err := Get(config, url)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeBackendService %q", d.Id()))
	}

	d.Set("affinity_cookie_ttl_sec", flattenComputeBackendServiceAffinityCookieTtlSec(res["affinityCookieTtlSec"]))
	d.Set("backends", flattenComputeBackendServiceBackends(res["backends"]))
	d.Set("cdn_policy", flattenComputeBackendServiceCdnPolicy(res["cdnPolicy"]))
	d.Set("connection_draining", flattenComputeBackendServiceConnectionDraining(res["connectionDraining"]))
	d.Set("creation_timestamp", flattenComputeBackendServiceCreationTimestamp(res["creationTimestamp"]))
	d.Set("description", flattenComputeBackendServiceDescription(res["description"]))
	d.Set("enable_cdn", flattenComputeBackendServiceEnableCDN(res["enableCDN"]))
	d.Set("health_checks", flattenComputeBackendServiceHealthChecks(res["healthChecks"]))
	d.Set("id", flattenComputeBackendServiceId(res["id"]))
	d.Set("name", flattenComputeBackendServiceName(res["name"]))
	d.Set("port_name", flattenComputeBackendServicePortName(res["portName"]))
	d.Set("protocol", flattenComputeBackendServiceProtocol(res["protocol"]))
	d.Set("region", flattenComputeBackendServiceRegion(res["region"]))
	d.Set("session_affinity", flattenComputeBackendServiceSessionAffinity(res["sessionAffinity"]))
	d.Set("timeout_sec", flattenComputeBackendServiceTimeoutSec(res["timeoutSec"]))
	d.Set("self_link", res["selfLink"])
	d.Set("project", project)

	return nil
}

func resourceComputeBackendServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	obj := map[string]interface{}{
		"affinityCookieTtlSec": expandComputeBackendServiceAffinityCookieTtlSec(d.Get("affinity_cookie_ttl_sec")),
		"backends":             expandComputeBackendServiceBackends(d.Get("backends")),
		"cdnPolicy":            expandComputeBackendServiceCdnPolicy(d.Get("cdn_policy")),
		"connectionDraining":   expandComputeBackendServiceConnectionDraining(d.Get("connection_draining")),
		"description":          expandComputeBackendServiceDescription(d.Get("description")),
		"enableCDN":            expandComputeBackendServiceEnableCDN(d.Get("enable_cdn")),
		"healthChecks":         expandComputeBackendServiceHealthChecks(d.Get("health_checks")),
		"name":                 expandComputeBackendServiceName(d.Get("name")),
		"portName":             expandComputeBackendServicePortName(d.Get("port_name")),
		"protocol":             expandComputeBackendServiceProtocol(d.Get("protocol")),
		"region":               expandComputeBackendServiceRegion(d.Get("region")),
		"sessionAffinity":      expandComputeBackendServiceSessionAffinity(d.Get("session_affinity")),
		"timeoutSec":           expandComputeBackendServiceTimeoutSec(d.Get("timeout_sec")),
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendServices/{{name}}")
	if err != nil {
		return err
	}
	res, err := Put(config, url, obj)
	if err != nil {
		return fmt.Errorf("Error updating BackendService %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating BackendService")
	if err != nil {
		return err
	}

	return resourceComputeBackendServiceRead(d, meta)
}

func resourceComputeBackendServiceDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)

	project, err := getProject(d, config)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "https://www.googleapis.com/compute/v1/projects/{{project}}/global/backendServices/{{name}}")
	if err != nil {
		return err
	}
	res, err := Delete(config, url)
	if err != nil {
		return fmt.Errorf("Error deleting BackendService %q: %s", d.Id(), err)
	}

	op := &compute.Operation{}
	err = Convert(res, op)
	if err != nil {
		return err
	}

	err = computeOperationWait(config.clientCompute, op, project, "Updating BackendService")
	if err != nil {
		return err
	}

	return nil
}

func resourceComputeBackendServiceImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	d.Set("", d.Id())
	return []*schema.ResourceData{d}, nil
}

func flattenComputeBackendServiceAffinityCookieTtlSec(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceBackends(v interface{}) interface{} {
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed = append(transformed, map[string]interface{}{
			"balancing_mode":               flattenComputeBackendServiceBackendsBalancingMode(original["balancingMode"]),
			"capacity_scaler":              flattenComputeBackendServiceBackendsCapacityScaler(original["capacityScaler"]),
			"description":                  flattenComputeBackendServiceBackendsDescription(original["description"]),
			"group":                        flattenComputeBackendServiceBackendsGroup(original["group"]),
			"max_connections":              flattenComputeBackendServiceBackendsMaxConnections(original["maxConnections"]),
			"max_connections_per_instance": flattenComputeBackendServiceBackendsMaxConnectionsPerInstance(original["maxConnectionsPerInstance"]),
			"max_rate":                     flattenComputeBackendServiceBackendsMaxRate(original["maxRate"]),
			"max_rate_per_instance":        flattenComputeBackendServiceBackendsMaxRatePerInstance(original["maxRatePerInstance"]),
			"max_utilization":              flattenComputeBackendServiceBackendsMaxUtilization(original["maxUtilization"]),
		})
	}
	return transformed
}
func flattenComputeBackendServiceBackendsBalancingMode(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsCapacityScaler(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsDescription(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsGroup(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsMaxConnections(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsMaxConnectionsPerInstance(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsMaxRate(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsMaxRatePerInstance(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceBackendsMaxUtilization(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceCdnPolicy(v interface{}) interface{} {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["cache_key_policy"] =
		flattenComputeBackendServiceCdnPolicyCacheKeyPolicy(original["cacheKeyPolicy"])
	return []interface{}{transformed}
}
func flattenComputeBackendServiceCdnPolicyCacheKeyPolicy(v interface{}) interface{} {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["include_host"] =
		flattenComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeHost(original["includeHost"])
	transformed["include_protocol"] =
		flattenComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeProtocol(original["includeProtocol"])
	transformed["include_query_string"] =
		flattenComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeQueryString(original["includeQueryString"])
	transformed["query_string_blacklist"] =
		flattenComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringBlacklist(original["queryStringBlacklist"])
	transformed["query_string_whitelist"] =
		flattenComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringWhitelist(original["queryStringWhitelist"])
	return []interface{}{transformed}
}
func flattenComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeHost(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeProtocol(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeQueryString(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringBlacklist(v interface{}) interface{} {
	return v
}
func flattenComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringWhitelist(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceConnectionDraining(v interface{}) interface{} {
	original := v.(map[string]interface{})
	transformed := make(map[string]interface{})
	transformed["draining_timeout_sec"] =
		flattenComputeBackendServiceConnectionDrainingDrainingTimeoutSec(original["drainingTimeoutSec"])
	return []interface{}{transformed}
}
func flattenComputeBackendServiceConnectionDrainingDrainingTimeoutSec(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceCreationTimestamp(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceDescription(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceEnableCDN(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceHealthChecks(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceId(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceName(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServicePortName(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceProtocol(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceRegion(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceSessionAffinity(v interface{}) interface{} {
	return v
}

func flattenComputeBackendServiceTimeoutSec(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceAffinityCookieTtlSec(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceBackends(v interface{}) interface{} {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformed["balancingMode"] =
			expandComputeBackendServiceBackendsBalancingMode(original["balancing_mode"])
		transformed["capacityScaler"] =
			expandComputeBackendServiceBackendsCapacityScaler(original["capacity_scaler"])
		transformed["description"] =
			expandComputeBackendServiceBackendsDescription(original["description"])
		transformed["group"] =
			expandComputeBackendServiceBackendsGroup(original["group"])
		transformed["maxConnections"] =
			expandComputeBackendServiceBackendsMaxConnections(original["max_connections"])
		transformed["maxConnectionsPerInstance"] =
			expandComputeBackendServiceBackendsMaxConnectionsPerInstance(original["max_connections_per_instance"])
		transformed["maxRate"] =
			expandComputeBackendServiceBackendsMaxRate(original["max_rate"])
		transformed["maxRatePerInstance"] =
			expandComputeBackendServiceBackendsMaxRatePerInstance(original["max_rate_per_instance"])
		transformed["maxUtilization"] =
			expandComputeBackendServiceBackendsMaxUtilization(original["max_utilization"])

		req = append(req, transformed)
	}
	return req
}
func expandComputeBackendServiceBackendsBalancingMode(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsCapacityScaler(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsDescription(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsGroup(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsMaxConnections(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsMaxConnectionsPerInstance(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsMaxRate(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsMaxRatePerInstance(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceBackendsMaxUtilization(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceCdnPolicy(v interface{}) interface{} {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformed["cacheKeyPolicy"] =
			expandComputeBackendServiceCdnPolicyCacheKeyPolicy(original["cache_key_policy"])

		req = append(req, transformed)
	}
	return req
}
func expandComputeBackendServiceCdnPolicyCacheKeyPolicy(v interface{}) interface{} {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformed["includeHost"] =
			expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeHost(original["include_host"])
		transformed["includeProtocol"] =
			expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeProtocol(original["include_protocol"])
		transformed["includeQueryString"] =
			expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeQueryString(original["include_query_string"])
		transformed["queryStringBlacklist"] =
			expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringBlacklist(original["query_string_blacklist"])
		transformed["queryStringWhitelist"] =
			expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringWhitelist(original["query_string_whitelist"])

		req = append(req, transformed)
	}
	return req
}
func expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeHost(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeProtocol(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceCdnPolicyCacheKeyPolicyIncludeQueryString(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringBlacklist(v interface{}) interface{} {
	return v
}
func expandComputeBackendServiceCdnPolicyCacheKeyPolicyQueryStringWhitelist(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceConnectionDraining(v interface{}) interface{} {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformed["drainingTimeoutSec"] =
			expandComputeBackendServiceConnectionDrainingDrainingTimeoutSec(original["draining_timeout_sec"])

		req = append(req, transformed)
	}
	return req
}
func expandComputeBackendServiceConnectionDrainingDrainingTimeoutSec(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceDescription(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceEnableCDN(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceHealthChecks(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceName(v interface{}) interface{} {
	return v
}

func expandComputeBackendServicePortName(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceProtocol(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceRegion(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceSessionAffinity(v interface{}) interface{} {
	return v
}

func expandComputeBackendServiceTimeoutSec(v interface{}) interface{} {
	return v
}
