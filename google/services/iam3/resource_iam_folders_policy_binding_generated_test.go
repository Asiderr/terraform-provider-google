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

package iam3_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccIAM3FoldersPolicyBinding_iamFoldersPolicyBindingExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"org_id":        envvar.GetTestOrgFromEnv(t),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {},
		},
		CheckDestroy: testAccCheckIAM3FoldersPolicyBindingDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccIAM3FoldersPolicyBinding_iamFoldersPolicyBindingExample(context),
			},
			{
				ResourceName:            "google_iam_folders_policy_binding.my-folder-binding",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"annotations", "folder", "location", "policy_binding_id"},
			},
		},
	})
}

func testAccIAM3FoldersPolicyBinding_iamFoldersPolicyBindingExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_iam_principal_access_boundary_policy" "pab_policy" {
  organization   = "%{org_id}"
  location       = "global"
  display_name   = "test folder binding%{random_suffix}"
  principal_access_boundary_policy_id = "tf-test-my-pab-policy%{random_suffix}"
}

resource "google_folder" "folder" {
  display_name        = "test folder%{random_suffix}"
  parent              = "organizations/%{org_id}"
  deletion_protection = false
}

resource "time_sleep" "wait_120s" {
  depends_on      = [google_folder.folder]
  create_duration = "120s"
}

resource "google_iam_folders_policy_binding" "my-folder-binding" {
  folder         = google_folder.folder.folder_id
  location       = "global"
  display_name   = "test folder binding%{random_suffix}"
  policy_kind    = "PRINCIPAL_ACCESS_BOUNDARY"
  policy_binding_id = "tf-test-test-folder-binding%{random_suffix}"
  policy         = "organizations/%{org_id}/locations/global/principalAccessBoundaryPolicies/${google_iam_principal_access_boundary_policy.pab_policy.principal_access_boundary_policy_id}"
  target {
    principal_set = "//cloudresourcemanager.googleapis.com/folders/${google_folder.folder.folder_id}"
  }
  depends_on = [time_sleep.wait_120s]
}
`, context)
}

func testAccCheckIAM3FoldersPolicyBindingDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_iam_folders_policy_binding" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{IAM3BasePath}}folders/{{folder}}/locations/{{location}}/policyBindings/{{policy_binding_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				Project:   billingProject,
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err == nil {
				return fmt.Errorf("IAM3FoldersPolicyBinding still exists at %s", url)
			}
		}

		return nil
	}
}
