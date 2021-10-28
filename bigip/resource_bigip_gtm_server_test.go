/*
Original work from https://github.com/DealerDotCom/terraform-provider-bigip
Modifications Copyright 2019 F5 Networks Inc.
This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file,You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package bigip

import (
	"fmt"
	"testing"

	"github.com/f5devcentral/go-bigip"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

var TEST_GTM_SERVER_NAME = "/" + TEST_PARTITION + "/test-server_1"

var TEST_GTM_SERVER_RESOURCE = `
	resource "bigip_gtm_server" "test-gtm-server" {
		name = "` + TEST_GTM_SERVER_NAME + `"
	}`

func TestAccBigipGtmServer_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers:    testAccProviders,
		CheckDestroy: testCheckGtmServerDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_GTM_SERVER_RESOURCE,
				Check:  resource.ComposeTestCheckFunc(
				// testCheckIRuleExists(TEST_IRULE_NAME),
				),
			},
		},
	})
}

func testCheckGtmServerDestroyed(s *terraform.State) error {
	client := testAccProvider.Meta().(*bigip.BigIP)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "bigip_gtm_server" {
			continue
		}

		name := rs.Primary.ID
		server, err := client.GetGtmserver(name)

		if err != nil {
			return nil
		}
		if server != nil {
			return fmt.Errorf("GTM Server %s not destroyed.", name)
		}
	}
	return nil
}
