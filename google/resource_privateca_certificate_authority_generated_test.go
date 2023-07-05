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
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
)

func TestAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"pool_name":           acctest.BootstrapSharedCaPoolInLocation(t, "us-central1"),
		"pool_location":       "us-central1",
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityBasicExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate_authority.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pem_ca_certificate", "ignore_active_certificates_on_deletion", "skip_grace_period", "location", "certificate_authority_id", "pool", "deletion_protection"},
			},
		},
	})
}

func testAccPrivatecaCertificateAuthority_privatecaCertificateAuthorityBasicExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_certificate_authority" "default" {
  // This example assumes this pool already exists.
  // Pools cannot be deleted in normal test circumstances, so we depend on static pools
  pool = "%{pool_name}"
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}"
  location = "%{pool_location}"
  deletion_protection = "%{deletion_protection}"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        is_ca = true
        max_issuer_path_length = 10
      }
      key_usage {
        base_key_usage {
          digital_signature = true
          content_commitment = true
          key_encipherment = false
          data_encipherment = true
          key_agreement = true
          cert_sign = true
          crl_sign = true
          decipher_only = true
        }
        extended_key_usage {
          server_auth = true
          client_auth = false
          email_protection = true
          code_signing = true
          time_stamping = true
        }
      }
    }
  }
  lifetime = "86400s"
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
}
`, context)
}

func TestAccPrivatecaCertificateAuthority_privatecaCertificateAuthoritySubordinateExample(t *testing.T) {
	acctest.SkipIfVcr(t)
	t.Parallel()

	context := map[string]interface{}{
		"pool_name":           acctest.BootstrapSharedCaPoolInLocation(t, "us-central1"),
		"pool_location":       "us-central1",
		"deletion_protection": false,
		"random_suffix":       acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccPrivatecaCertificateAuthority_privatecaCertificateAuthoritySubordinateExample(context),
			},
			{
				ResourceName:            "google_privateca_certificate_authority.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"pem_ca_certificate", "ignore_active_certificates_on_deletion", "skip_grace_period", "location", "certificate_authority_id", "pool", "deletion_protection"},
			},
		},
	})
}

func testAccPrivatecaCertificateAuthority_privatecaCertificateAuthoritySubordinateExample(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_privateca_certificate_authority" "root-ca" {
  pool = "%{pool_name}"
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}-root"
  location = "us-central1"
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-certificate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        # is_ca *MUST* be true for certificate authorities
        is_ca = true
      }
      key_usage {
        base_key_usage {
          # cert_sign and crl_sign *MUST* be true for certificate authorities
          cert_sign = true
          crl_sign = true
        }
        extended_key_usage {
          server_auth = false
        }
      }
    }
  }
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }

  // Disable CA deletion related safe checks for easier cleanup.
  deletion_protection                    = false
  skip_grace_period                      = true
  ignore_active_certificates_on_deletion = true
}

resource "google_privateca_certificate_authority" "default" {
  // This example assumes this pool already exists.
  // Pools cannot be deleted in normal test circumstances, so we depend on static pools
  pool = "%{pool_name}"
  certificate_authority_id = "tf-test-my-certificate-authority%{random_suffix}-sub"
  location = "%{pool_location}"
  deletion_protection = "%{deletion_protection}"
  subordinate_config {
    certificate_authority = google_privateca_certificate_authority.root-ca.name
  }
  config {
    subject_config {
      subject {
        organization = "HashiCorp"
        common_name = "my-subordinate-authority"
      }
      subject_alt_name {
        dns_names = ["hashicorp.com"]
      }
    }
    x509_config {
      ca_options {
        is_ca = true
        # Force the sub CA to only issue leaf certs
        max_issuer_path_length = 0
      }
      key_usage {
        base_key_usage {
          digital_signature = true
          content_commitment = true
          key_encipherment = false
          data_encipherment = true
          key_agreement = true
          cert_sign = true
          crl_sign = true
          decipher_only = true
        }
        extended_key_usage {
          server_auth = true
          client_auth = false
          email_protection = true
          code_signing = true
          time_stamping = true
        }
      }
    }
  }
  lifetime = "86400s"
  key_spec {
    algorithm = "RSA_PKCS1_4096_SHA256"
  }
  type = "SUBORDINATE"
}
`, context)
}

func testAccCheckPrivatecaCertificateAuthorityDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_privateca_certificate_authority" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := acctest.GoogleProviderConfig(t)

			url, err := tpgresource.ReplaceVarsForTest(config, rs, "{{PrivatecaBasePath}}projects/{{project}}/locations/{{location}}/caPools/{{pool}}/certificateAuthorities/{{certificate_authority_id}}")
			if err != nil {
				return err
			}

			res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
				Config:    config,
				Method:    "GET",
				RawURL:    url,
				UserAgent: config.UserAgent,
			})
			if err != nil {
				return nil
			}

			if s := res["state"]; s != "DELETED" {
				return fmt.Errorf("CertificateAuthority %s got %s, want DELETED", url, s)
			}
		}

		return nil
	}
}
