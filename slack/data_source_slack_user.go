package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSlackUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackUserRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
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

	user, err := api.getUserInfo(name)
	if err != nil {
		return err
	}

	d.SetId(user.ID)
	d.Set("id", user.ID)

	return nil
}
