package slack

import (
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/nlopes/slack"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceChannelCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	// isPrivate := d.Get("is_private").(bool)

	api := m.(*slackAPI)
	channel, err := api.CreateChannel(name)
	if err != nil {
		return err
	}

	d.SetId(channel.ID)
	return resourceChannelRead(d, m)
}

func resourceChannelRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)
	channels, err := api.GetChannels(false)
	if err != nil {
		return err
	}

	var channel *slack.Channel
	name := d.Get("name").(string)
	for _, c := range channels {
		if strings.ToLower(c.Name) == strings.ToLower(name) {
			channel = &c
			break
		}
	}

	d.SetId(channel.ID)
	d.Set("is_private", channel.IsPrivate)

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceChannelRead(d, m)
}

func resourceChannelDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
