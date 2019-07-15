package bigip

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.com/pirotrav/go-bigip"
)

var TEST_ServerSsl_NAME = fmt.Sprintf("/%s/test-ServerSsl", TEST_PARTITION)

var TEST_ServerSsl_RESOURCE = `
resource "bigip_ltm_profile_server_ssl" "profile_mutualssl" {
	alert_timeout                   = "indefinite"
	authenticate                    = "once"
	authenticate_depth              = 9
	ca_file                         = "none"
	cache_size                      = 262144
	chain                           = "none"
	ciphers                         = "DEFAULT"
	defaults_from                   = "/Common/serverssl"
	expire_cert_response_control    = "drop"
	handshake_timeout               = "10"
	id                              = "terraform_test"
	key                             = "none"
	mod_ssl_methods                 = "disabled"
	mode                            = "enabled"
	name                            = "terraform_test"
    partition                       = "Common"
	proxy_ssl                       = "disabled"
	renegotiate_period              = "indefinite"
	renegotiate_size                = "indefinite"
	renegotiation                   = "enabled"
	retain_certificate              = "true"
	secure_renegotiation            = "require-strict"
	server_name                     = "none"
	session_mirroring               = "disabled"
	session_ticket                  = "disabled"
	sni_default                     = "false"
	sni_require                     = "false"
	ssl_forward_proxy               = "disabled"
	ssl_forward_proxy_bypass        = "disabled"
	ssl_sign_hash                   = "any"
	strict_resume                   = "disabled"
    tm_options                      = [
	  - "dont-insert-empty-fragments",
	]
	unclean_shutdown                = "enabled"
	untrusted_cert_response_control = "drop"
}
`

func TestAccBigipLtmProfileServerSsl_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckServerSslsDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_ServerSsl_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckServerSslExists(TEST_ServerSsl_NAME, true),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "name", "/Common/test-ServerSsl"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "defaults_from", "/Common/serverssl"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "alert_timeout", "indefinite"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "authenticate", "once"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "authenticate_depth", "9"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "ca_file", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "cache_size", "262144"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "chain", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "ciphers", "DEFAULT"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "expire_cert_response_control", "drop"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "handshake_timeout", "10"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "id", "terraform_test"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "key", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "mod_ssl_methods", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "mode", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "name", "terraform_test"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "partition", "Common"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "proxy_ssl", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "renegotiate_period", "indefinite"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "renegotiate_size", "indefinite"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "renegotiation", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "retain_certificate", "true"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "secure_renegotiation", "require-strict"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "server_name", "none"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "session_mirroring", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "session_ticket", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "sni_default", "false"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "sni_require", "false"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "ssl_forward_proxy", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "ssl_forward_proxy_bypass", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "ssl_sign_hash", "any"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "strict_resume", "disabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "tm_options", "[dont-insert-empty-fragments,]"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "unclean_shutdown", "enabled"),
					resource.TestCheckResourceAttr("bigip_ltm_profile_server_ssl.test-ServerSsl", "untrusted_cert_response_control", "drop"),
				),
			},
		},
	})
}

func TestAccBigipLtmProfileServerSsl_import(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckServerSslsDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_ServerSsl_RESOURCE,
				Check: resource.ComposeTestCheckFunc(
					testCheckServerSslExists(TEST_ServerSsl_NAME, true),
				),
				ResourceName:      TEST_ServerSsl_NAME,
				ImportState:       false,
				ImportStateVerify: true,
			},
		},
	})
}

func testCheckServerSslExists(name string, exists bool) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client := testAccProvider.Meta().(*bigip.BigIP)
		p, err := client.GetServerSsl(name)
		if err != nil {
			return err
		}
		if exists && p == nil {
			return fmt.Errorf("ServerSsl %s was not created.", name)
		}
		if !exists && p == nil {
			return fmt.Errorf("ServerSsl %s still exists.", name)
		}
		return nil
	}
}

func testCheckServerSslsDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(*bigip.BigIP)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "bigip_ltm_profile_ServerSsl" {
			continue
		}

		name := rs.Primary.ID
		ServerSsl, err := client.GetServerSsl(name)
		if err != nil {
			return err
		}
		if ServerSsl != nil {
			return fmt.Errorf("ServerSsl %s not destroyed.", name)
		}
	}
	return nil
}