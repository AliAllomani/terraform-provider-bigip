/*
Original work from https://github.com/DealerDotCom/terraform-provider-bigip
Modifications Copyright 2019 F5 Networks Inc.
This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file,You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package bigip

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var TEST_GTM_POOL_NAME = "/" + TEST_PARTITION + "/test-pool_1"

var TEST_GTM_POOL_RESOURCE = `
	resource "bigip_gtm_pool" "test-rule" {
		name = "` + TEST_GTM_POOL_NAME + `"
	}`

func TestAccBigipGtmPool_create(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck: func() {
			testAcctPreCheck(t)
		},
		Providers: testAccProviders,
		// CheckDestroy: testCheckIRulesDestroyed,
		Steps: []resource.TestStep{
			{
				Config: TEST_GTM_POOL_RESOURCE,
				Check:  resource.ComposeTestCheckFunc(
				// testCheckIRuleExists(TEST_IRULE_NAME),
				),
			},
		},
	})
}

// func testCheckIRuleExists(name string) resource.TestCheckFunc {
// 	return func(s *terraform.State) error {
// 		client := testAccProvider.Meta().(*bigip.BigIP)

// 		irule, err := client.IRule(name)
// 		if err != nil {
// 			return fmt.Errorf("Error while fetching irule: %v", err)

// 		}

// 		body := s.RootModule().Resources["bigip_ltm_irule.test-rule"].Primary.Attributes["irule"]
// 		if irule.Rule != body {
// 			return fmt.Errorf("IRule body does not match. Expecting %s got %s.", body, irule.Rule)
// 		}

// 		irule_name := fmt.Sprintf("/%s/%s", irule.Partition, irule.Name)
// 		if irule_name != name {
// 			return fmt.Errorf("IRule name does not match. Expecting %s got %s.", name, irule_name)
// 		}
// 		return nil
// 	}
// }

// func testCheckIRulesDestroyed(s *terraform.State) error {
// 	client := testAccProvider.Meta().(*bigip.BigIP)

// 	for _, rs := range s.RootModule().Resources {
// 		if rs.Type != "bigip_ltm_irule" {
// 			continue
// 		}

// 		name := rs.Primary.ID
// 		irule, err := client.IRule(name)

// 		if err != nil {
// 			return nil
// 		}
// 		if irule != nil {
// 			return fmt.Errorf("IRule %s not destroyed.", name)
// 		}
// 	}
// 	return nil
// }
