package slack

import (
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSlackUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackUserRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceSlackUserRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)
	name := d.Get("name").(string)

	log.Println("[DEBUG] Fetching user info")
	user, err := api.getUserInfo(name)
	if err != nil {
		return fmt.Errorf("Error fetching user info: %s", err)
	}

	d.SetId(user.ID)
	if err := d.Set("id", user.ID); err != nil {
		return fmt.Errorf("Error setting id: %s", err)
	}

	return nil
}
