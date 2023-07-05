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

package google

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccGKEBackupBackupPlanIamBindingGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupBackupPlanIamBinding_basicGenerated(context),
			},
			{
				ResourceName:      "google_gke_backup_backup_plan_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-basic-plan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Test Iam Binding update
				Config: testAccGKEBackupBackupPlanIamBinding_updateGenerated(context),
			},
			{
				ResourceName:      "google_gke_backup_backup_plan_iam_binding.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s roles/viewer", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-basic-plan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGKEBackupBackupPlanIamMemberGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				// Test Iam Member creation (no update for member, no need to test)
				Config: testAccGKEBackupBackupPlanIamMember_basicGenerated(context),
			},
			{
				ResourceName:      "google_gke_backup_backup_plan_iam_member.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s roles/viewer user:admin@hashicorptest.com", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-basic-plan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccGKEBackupBackupPlanIamPolicyGenerated(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": acctest.RandString(t, 10),
		"role":          "roles/viewer",
		"project":       envvar.GetTestProjectFromEnv(),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		Steps: []resource.TestStep{
			{
				Config: testAccGKEBackupBackupPlanIamPolicy_basicGenerated(context),
				Check:  resource.TestCheckResourceAttrSet("data.google_gke_backup_backup_plan_iam_policy.foo", "policy_data"),
			},
			{
				ResourceName:      "google_gke_backup_backup_plan_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-basic-plan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccGKEBackupBackupPlanIamPolicy_emptyBinding(context),
			},
			{
				ResourceName:      "google_gke_backup_backup_plan_iam_policy.foo",
				ImportStateId:     fmt.Sprintf("projects/%s/locations/%s/backupPlans/%s", envvar.GetTestProjectFromEnv(), envvar.GetTestRegionFromEnv(), fmt.Sprintf("tf-test-basic-plan%s", context["random_suffix"])),
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccGKEBackupBackupPlanIamMember_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-basic-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_backup_plan_iam_member" "foo" {
  project = google_gke_backup_backup_plan.basic.project
  location = google_gke_backup_backup_plan.basic.location
  name = google_gke_backup_backup_plan.basic.name
  role = "%{role}"
  member = "user:admin@hashicorptest.com"
}
`, context)
}

func testAccGKEBackupBackupPlanIamPolicy_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-basic-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

data "google_iam_policy" "foo" {
  binding {
    role = "%{role}"
    members = ["user:admin@hashicorptest.com"]
  }
}

resource "google_gke_backup_backup_plan_iam_policy" "foo" {
  project = google_gke_backup_backup_plan.basic.project
  location = google_gke_backup_backup_plan.basic.location
  name = google_gke_backup_backup_plan.basic.name
  policy_data = data.google_iam_policy.foo.policy_data
}

data "google_gke_backup_backup_plan_iam_policy" "foo" {
  project = google_gke_backup_backup_plan.basic.project
  location = google_gke_backup_backup_plan.basic.location
  name = google_gke_backup_backup_plan.basic.name
  depends_on = [
    google_gke_backup_backup_plan_iam_policy.foo
  ]
}
`, context)
}

func testAccGKEBackupBackupPlanIamPolicy_emptyBinding(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-basic-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

data "google_iam_policy" "foo" {
}

resource "google_gke_backup_backup_plan_iam_policy" "foo" {
  project = google_gke_backup_backup_plan.basic.project
  location = google_gke_backup_backup_plan.basic.location
  name = google_gke_backup_backup_plan.basic.name
  policy_data = data.google_iam_policy.foo.policy_data
}
`, context)
}

func testAccGKEBackupBackupPlanIamBinding_basicGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-basic-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_backup_plan_iam_binding" "foo" {
  project = google_gke_backup_backup_plan.basic.project
  location = google_gke_backup_backup_plan.basic.location
  name = google_gke_backup_backup_plan.basic.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com"]
}
`, context)
}

func testAccGKEBackupBackupPlanIamBinding_updateGenerated(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_container_cluster" "primary" {
  name               = "tf-test-basic-cluster%{random_suffix}"
  location           = "us-central1"
  initial_node_count = 1
  workload_identity_config {
    workload_pool = "%{project}.svc.id.goog"
  }
  addons_config {
    gke_backup_agent_config {
      enabled = true
    }
  }
}

resource "google_gke_backup_backup_plan" "basic" {
  name = "tf-test-basic-plan%{random_suffix}"
  cluster = google_container_cluster.primary.id
  location = "us-central1"
  backup_config {
    include_volume_data = true
    include_secrets = true
    all_namespaces = true
  }
}

resource "google_gke_backup_backup_plan_iam_binding" "foo" {
  project = google_gke_backup_backup_plan.basic.project
  location = google_gke_backup_backup_plan.basic.location
  name = google_gke_backup_backup_plan.basic.name
  role = "%{role}"
  members = ["user:admin@hashicorptest.com", "user:gterraformtest1@gmail.com"]
}
`, context)
}
