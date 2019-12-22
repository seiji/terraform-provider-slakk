package slack

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
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
				Optional: true,
				Default:  false,
			},
			"user_ids": {
				Type: schema.TypeList,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func resourceChannelCreate(d *schema.ResourceData, m interface{}) error {
	name := d.Get("name").(string)
	isPrivate := d.Get("is_private").(bool)
	userIds := toStringSlice(d.Get("user_ids").([]interface{}))

	api := m.(*slackAPI)
	_, err := api.CreateConversation(name, isPrivate, userIds...)
	if err != nil {
		return err
	}

	return resourceChannelRead(d, m)
}

func resourceChannelRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)

	name := d.Get("name").(string)

	channel, err := api.getChannel(name)
	if err != nil {
		return err
	}

	d.SetId(channel.ID)

	d.Set("is_general", channel.IsGeneral)
	d.Set("is_archived", channel.IsArchived)
	d.Set("is_private", channel.IsPrivate)
	d.Set("is_shared", channel.IsShared)

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceChannelRead(d, m)
}

func resourceChannelDelete(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)

	name := d.Get("name").(string)

	channel, err := api.getChannel(name)
	if err != nil {
		return err
	}

	_, _, err = api.CloseConversation(channel.ID)
	if err != nil {
		return err
	}

	return nil
}
