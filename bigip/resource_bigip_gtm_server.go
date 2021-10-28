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

func resourceBigipGtmServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceBigipGtmServerCreate,
		Read:   resourceBigipGtmServerRead,
		// Update: resourceBigipGtmServerUpdate,
		Delete: resourceBigipGtmServerDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "Name of the GTM Server",
				ForceNew:     true,
				ValidateFunc: validateF5Name,
			},
		},
	}
}

func resourceBigipGtmServerCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)

	gtmServer := bigip.Server{
		Name: d.Get("name").(string),
	}

	log.Printf("[INFO] Creating GTM Server %s", gtmServer.Name)

	err := client.CreateGtmserver(&gtmServer)
	if err != nil {
		return fmt.Errorf("Error creating GTM Server %s: %v", gtmServer.Name, err)
	}

	d.SetId(gtmServer.Name)

	return resourceBigipGtmServerRead(d, meta)
}

func resourceBigipGtmServerRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*bigip.BigIP)

	name := d.Id()
	log.Printf("[INFO] Retrieving GTM Server %s", name)

	server, err := client.GetGtmserver(name)
	if err != nil {
		return fmt.Errorf("Error retrieving GTM Server %s: %v", name, err)
	}

	d.Set("name", server.Name)

	return nil
}

// func resourceBigipGtmServerUpdate(d *schema.ResourceData, meta interface{}) error {
// 	client := meta.(*bigip.BigIP)

// 	name := d.Id()

// 	gtmServer := bigip.Server{
// 		Name: d.Get("name").(string),
// 	}

// 	err := client.UpdateGtmserver(name, &gtmServer)
// 	if err != nil {
// 		return fmt.Errorf("Error modifying GTM Server %s: %v", name, err)
// 	}
// 	return resourceBigipGtmServerRead(d, meta)
// }

func resourceBigipGtmServerDelete(d *schema.ResourceData, meta interface{}) error {

	client := meta.(*bigip.BigIP)

	name := d.Id()

	if err := client.DeleteGtmserver(name); err != nil {
		return fmt.Errorf("Error deleting GTM Server %s: %v", name, err)
	}

	d.SetId("")
	return nil
}
