/*
Original work from https://github.com/DealerDotCom/terraform-provider-bigip
Modifications Copyright 2019 F5 Networks Inc.
This Source Code Form is subject to the terms of the Mozilla Public License, v. 2.0.
If a copy of the MPL was not distributed with this file,You can obtain one at https://mozilla.org/MPL/2.0/.
*/
package bigip

import (
	"fmt"
	"log"

	"github.com/f5devcentral/go-bigip"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceBigipGtmPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigipGtmPoolCreate,
		Read:   resourceBigipGtmPoolRead,
		// Update: resourceBigipGtmPoolUpdate,
		Delete: resourceBigipGtmPoolDelete,
		// Exists: resourceBigipGtmPoolExists,
		// Importer: &schema.ResourceImporter{
		// 	State: schema.ImportStatePassthrough,
		// },

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the pool",
				ForceNew:     true,
				ValidateFunc: validateF5Name,
			},
		},
	}
}

func resourceBigipGtmPoolCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)

	name := d.Get("name").(string)
	monitor := ""
	loadBalancingMode := ""
	maxAnswersReturned := 0
	alternativeMode := ""
	fallbackIP := ""
	fallbackMode := ""
	members := []string{}

	log.Printf("[INFO] Creating GTM Pool %s", name)

	err := client.CreatePool_a(name, monitor, loadBalancingMode, maxAnswersReturned, alternativeMode, fallbackIP, fallbackMode, members)
	if err != nil {
		return fmt.Errorf("Error creating GTM Pool %s: %v", name, err)
	}

	d.SetId(name)

	return resourceBigipGtmPoolRead(d, meta)
}

func resourceBigipGtmPoolRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)

	name := d.Id()
	log.Printf("[INFO] Retrieving GTM Pool %s", name)

	// poolQuery := bigip.Pool_a{
	// 	Name: name,
	// }

	pool, err := client.Pool_as()
	if err != nil {
		return fmt.Errorf("Error retrieving iRule %s: %v", name, err)
	}

	d.Set("name", pool.Name)

	return nil
}

// func resourceBigipGtmPoolExists(d *schema.ResourceData, meta interface{}) (bool, error) {
// 	client := meta.(*bigip.BigIP)

// 	name := d.Id()
// 	log.Printf("[INFO] Checking if iRule (%s) exists", name)

// 	irule, err := client.IRule(name)
// 	if err != nil {
// 		return false, fmt.Errorf("Error retrieving iRule %s: %v", name, err)
// 	}
// 	if irule == nil {
// 		log.Printf("[DEBUG] iRule (%s) not found, removing from state", name)
// 		d.SetId("")
// 		return false, nil
// 	}

// 	return irule != nil, nil
// }

// func resourceBigipGtmPoolUpdate(d *schema.ResourceData, meta interface{}) error {
// 	client := meta.(*bigip.BigIP)

// 	name := d.Id()

// 	r := &bigip.IRule{
// 		FullPath: name,
// 		Rule:     d.Get("irule").(string),
// 	}

// 	err := client.ModifyIRule(name, r)
// 	if err != nil {
// 		return fmt.Errorf("Error modifying iRule %s: %v", name, err)
// 	}
// 	return resourceBigipGtmPoolRead(d, meta)
// }

func resourceBigipGtmPoolDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
