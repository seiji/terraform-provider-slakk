package slack

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/slack-go/slack"
)

func dataSourceSlackChannel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackChannelRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_general": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_shared": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSlackChannelRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)

	name := d.Get("name").(string)

	channels, err := api.GetChannels(false)
	if err != nil {
		return err
	}

	var channel *slack.Channel
	for _, c := range channels {
		if strings.ToLower(c.Name) == strings.ToLower(name) {
			channel = &c
			break
		}
	}

	if channel == nil {
		return fmt.Errorf("channel '%s' is not found", name)
	}

	d.SetId(channel.ID)
	d.Set("is_general", channel.IsGeneral)
	d.Set("is_archived", channel.IsArchived)
	d.Set("is_private", channel.IsPrivate)
	d.Set("is_shared", channel.IsShared)

	return nil
}
