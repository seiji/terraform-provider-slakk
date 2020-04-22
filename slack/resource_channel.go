package slack

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceChannelCreate,
		Read:   resourceChannelRead,
		Update: resourceChannelUpdate,
		Delete: resourceChannelDelete,
		Importer: &schema.ResourceImporter{
			State: resourceChannelImport,
		},

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
				Optional: true,
				Default:  false,
			},
			// "user_ids": {
			// 	Type: schema.TypeList,
			// 	Elem: &schema.Schema{
			// 		Type: schema.TypeString,
			// 	},
			// 	Optional: true,
			// },
		},
	}
}

func resourceChannelCreate(d *schema.ResourceData, meta interface{}) error {
	name := d.Get("name").(string)
	api := meta.(*slackAPI)
	channel, err := api.CreateChannel(name)
	if err != nil {
		return err
	}

	d.SetId(channel.ID)
	return resourceChannelRead(d, meta)
}

func resourceChannelRead(d *schema.ResourceData, meta interface{}) error {
	channelID := d.Id()
	api := meta.(*slackAPI)

	channel, err := api.GetChannelInfo(channelID)
	if err != nil {
		return err
	}

	d.Set("name", channel.Name)
	d.Set("is_general", channel.IsGeneral)
	d.Set("is_archived", channel.IsArchived)
	d.Set("is_private", channel.IsPrivate)
	d.Set("is_shared", channel.IsShared)

	return nil
}

func resourceChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	channelID := d.Id()
	api := meta.(*slackAPI)

	_, err := api.GetChannelInfo(channelID)
	if err != nil {
		return err
	}

	if d.HasChange("name") {
		name := d.Get("name").(string)
		if _, err = api.RenameChannel(channelID, name); err != nil {
			return err
		}
	}
	return resourceChannelRead(d, meta)
}

func resourceChannelDelete(d *schema.ResourceData, meta interface{}) error {
	log.Printf("[WARN] delete is not yet implemented. ")
	return nil
}

func resourceChannelImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	return []*schema.ResourceData{d}, nil
}
