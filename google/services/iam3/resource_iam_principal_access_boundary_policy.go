// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package iam3

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func ResourceIAM3PrincipalAccessBoundaryPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceIAM3PrincipalAccessBoundaryPolicyCreate,
		Read:   resourceIAM3PrincipalAccessBoundaryPolicyRead,
		Update: resourceIAM3PrincipalAccessBoundaryPolicyUpdate,
		Delete: resourceIAM3PrincipalAccessBoundaryPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceIAM3PrincipalAccessBoundaryPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetAnnotationsDiff,
		),

		Schema: map[string]*schema.Schema{
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location the principal access boundary policy is in.`,
			},
			"organization": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The parent organization of the principal access boundary policy.`,
			},
			"principal_access_boundary_policy_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `The ID to use to create the principal access boundary policy.
This value must start with a lowercase letter followed by up to 62 lowercase letters, numbers, hyphens, or dots. Pattern, /a-z{2,62}/.`,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User defined annotations. See https://google.aip.dev/148#annotations
for more details such as format and size limitations


**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"details": {
				Type:        schema.TypeList,
				Computed:    true,
				Optional:    true,
				Description: `Principal access boundary policy details`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rules": {
							Type:        schema.TypeList,
							Required:    true,
							Description: `A list of principal access boundary policy rules. The number of rules in a policy is limited to 500.`,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"effect": {
										Type:     schema.TypeString,
										Required: true,
										Description: `The access relationship of principals to the resources in this rule.
Possible values: ALLOW`,
									},
									"resources": {
										Type:     schema.TypeList,
										Required: true,
										Description: `A list of Cloud Resource Manager resources. The resource
and all the descendants are included. The number of resources in a policy
is limited to 500 across all rules.
The following resource types are supported:
* Organizations, such as '//cloudresourcemanager.googleapis.com/organizations/123'.
* Folders, such as '//cloudresourcemanager.googleapis.com/folders/123'.
* Projects, such as '//cloudresourcemanager.googleapis.com/projects/123'
or '//cloudresourcemanager.googleapis.com/projects/my-project-id'.`,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"description": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: `The description of the principal access boundary policy rule. Must be less than or equal to 256 characters.`,
									},
								},
							},
						},
						"enforcement_version": {
							Type:     schema.TypeString,
							Computed: true,
							Optional: true,
							Description: `The version number that indicates which Google Cloud services
are included in the enforcement (e.g. \"latest\", \"1\", ...). If empty, the
PAB policy version will be set to the current latest version, and this version
won't get updated when new versions are released.`,
						},
					},
				},
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `The description of the principal access boundary policy. Must be less than or equal to 63 characters.`,
			},
			"create_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when the principal access boundary policy was created.`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `The etag for the principal access boundary. If this is provided on update, it must match the server's etag.`,
			},
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Identifier. The resource name of the principal access boundary policy.  The following format is supported:
 'organizations/{organization_id}/locations/{location}/principalAccessBoundaryPolicies/{policy_id}'`,
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The globally unique ID of the principal access boundary policy.`,
			},
			"update_time": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The time when the principal access boundary policy was most recently updated.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceIAM3PrincipalAccessBoundaryPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM3PrincipalAccessBoundaryPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	detailsProp, err := expandIAM3PrincipalAccessBoundaryPolicyDetails(d.Get("details"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("details"); !tpgresource.IsEmptyValue(reflect.ValueOf(detailsProp)) && (ok || !reflect.DeepEqual(v, detailsProp)) {
		obj["details"] = detailsProp
	}
	annotationsProp, err := expandIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies?principalAccessBoundaryPolicyId={{principal_access_boundary_policy_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new PrincipalAccessBoundaryPolicy: %#v", obj)
	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating PrincipalAccessBoundaryPolicy: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	// Use the resource in the operation response to populate
	// identity fields and d.Id() before read
	var opRes map[string]interface{}
	err = IAM3OperationWaitTimeWithResponse(
		config, res, &opRes, project, "Creating PrincipalAccessBoundaryPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))
	if err != nil {
		// The resource didn't actually create
		d.SetId("")

		return fmt.Errorf("Error waiting to create PrincipalAccessBoundaryPolicy: %s", err)
	}

	if err := d.Set("name", flattenIAM3PrincipalAccessBoundaryPolicyName(opRes["name"], d, config)); err != nil {
		return err
	}

	// This may have caused the ID to update - update it if so.
	id, err = tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	log.Printf("[DEBUG] Finished creating PrincipalAccessBoundaryPolicy %q: %#v", d.Id(), res)

	return resourceIAM3PrincipalAccessBoundaryPolicyRead(d, meta)
}

func resourceIAM3PrincipalAccessBoundaryPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("IAM3PrincipalAccessBoundaryPolicy %q", d.Id()))
	}

	if err := d.Set("name", flattenIAM3PrincipalAccessBoundaryPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("uid", flattenIAM3PrincipalAccessBoundaryPolicyUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("etag", flattenIAM3PrincipalAccessBoundaryPolicyEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("display_name", flattenIAM3PrincipalAccessBoundaryPolicyDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("annotations", flattenIAM3PrincipalAccessBoundaryPolicyAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("create_time", flattenIAM3PrincipalAccessBoundaryPolicyCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("update_time", flattenIAM3PrincipalAccessBoundaryPolicyUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("details", flattenIAM3PrincipalAccessBoundaryPolicyDetails(res["details"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}
	if err := d.Set("effective_annotations", flattenIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading PrincipalAccessBoundaryPolicy: %s", err)
	}

	return nil
}

func resourceIAM3PrincipalAccessBoundaryPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	obj := make(map[string]interface{})
	displayNameProp, err := expandIAM3PrincipalAccessBoundaryPolicyDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	detailsProp, err := expandIAM3PrincipalAccessBoundaryPolicyDetails(d.Get("details"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("details"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, detailsProp)) {
		obj["details"] = detailsProp
	}
	annotationsProp, err := expandIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating PrincipalAccessBoundaryPolicy %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("details") {
		updateMask = append(updateMask, "details")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating PrincipalAccessBoundaryPolicy %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating PrincipalAccessBoundaryPolicy %q: %#v", d.Id(), res)
		}

		err = IAM3OperationWaitTime(
			config, res, project, "Updating PrincipalAccessBoundaryPolicy", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceIAM3PrincipalAccessBoundaryPolicyRead(d, meta)
}

func resourceIAM3PrincipalAccessBoundaryPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var project string
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	url, err := tpgresource.ReplaceVars(d, config, "{{IAM3BasePath}}organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting PrincipalAccessBoundaryPolicy %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "PrincipalAccessBoundaryPolicy")
	}

	err = IAM3OperationWaitTime(
		config, res, project, "Deleting PrincipalAccessBoundaryPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting PrincipalAccessBoundaryPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceIAM3PrincipalAccessBoundaryPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^organizations/(?P<organization>[^/]+)/locations/(?P<location>[^/]+)/principalAccessBoundaryPolicies/(?P<principal_access_boundary_policy_id>[^/]+)$",
		"^(?P<organization>[^/]+)/(?P<location>[^/]+)/(?P<principal_access_boundary_policy_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "organizations/{{organization}}/locations/{{location}}/principalAccessBoundaryPolicies/{{principal_access_boundary_policy_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenIAM3PrincipalAccessBoundaryPolicyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenIAM3PrincipalAccessBoundaryPolicyCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyDetails(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["rules"] =
		flattenIAM3PrincipalAccessBoundaryPolicyDetailsRules(original["rules"], d, config)
	transformed["enforcement_version"] =
		flattenIAM3PrincipalAccessBoundaryPolicyDetailsEnforcementVersion(original["enforcementVersion"], d, config)
	return []interface{}{transformed}
}
func flattenIAM3PrincipalAccessBoundaryPolicyDetailsRules(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}
	l := v.([]interface{})
	transformed := make([]interface{}, 0, len(l))
	for _, raw := range l {
		original := raw.(map[string]interface{})
		if len(original) < 1 {
			// Do not include empty json objects coming back from the api
			continue
		}
		transformed = append(transformed, map[string]interface{}{
			"description": flattenIAM3PrincipalAccessBoundaryPolicyDetailsRulesDescription(original["description"], d, config),
			"resources":   flattenIAM3PrincipalAccessBoundaryPolicyDetailsRulesResources(original["resources"], d, config),
			"effect":      flattenIAM3PrincipalAccessBoundaryPolicyDetailsRulesEffect(original["effect"], d, config),
		})
	}
	return transformed
}
func flattenIAM3PrincipalAccessBoundaryPolicyDetailsRulesDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyDetailsRulesResources(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyDetailsRulesEffect(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyDetailsEnforcementVersion(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandIAM3PrincipalAccessBoundaryPolicyDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetails(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedRules, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRules(original["rules"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedRules); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["rules"] = transformedRules
	}

	transformedEnforcementVersion, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsEnforcementVersion(original["enforcement_version"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnforcementVersion); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["enforcementVersion"] = transformedEnforcementVersion
	}

	return transformed, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRules(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	req := make([]interface{}, 0, len(l))
	for _, raw := range l {
		if raw == nil {
			continue
		}
		original := raw.(map[string]interface{})
		transformed := make(map[string]interface{})

		transformedDescription, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesDescription(original["description"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedDescription); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["description"] = transformedDescription
		}

		transformedResources, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesResources(original["resources"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedResources); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["resources"] = transformedResources
		}

		transformedEffect, err := expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesEffect(original["effect"], d, config)
		if err != nil {
			return nil, err
		} else if val := reflect.ValueOf(transformedEffect); val.IsValid() && !tpgresource.IsEmptyValue(val) {
			transformed["effect"] = transformedEffect
		}

		req = append(req, transformed)
	}
	return req, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesResources(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsRulesEffect(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyDetailsEnforcementVersion(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandIAM3PrincipalAccessBoundaryPolicyEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
