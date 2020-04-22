package slack

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/slack-go/slack"
)

func dataSourceSlackConversations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackConversationsList,
		Schema: map[string]*schema.Schema{
			"ids": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"exclude_archived": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"names": {
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"types": {
				Type:     schema.TypeList,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
			},
		},
	}
}

func dataSourceSlackConversationsList(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)

	types := []string{"public_channel"}
	excludeArchived := !d.Get("exclude_archived").(bool)
	if v, ok := d.GetOk("types"); ok {
		types = toStringSlice(v.([]interface{}))
	}

	params := slack.GetConversationsParameters{
		ExcludeArchived: strconv.FormatBool(excludeArchived),
		Limit:           100,
		Types:           types,
	}

	log.Printf("[DEBUG] Fetching conversations with params: %v", params)
	cs, _, err := api.GetConversations(&params)
	if err != nil {
		return fmt.Errorf("Error fetching conversations: %s", err)
	}

	names := []string{}
	ids := map[string]string{}
	for _, c := range cs {
		names = append(names, c.Name)
		ids[c.Name] = c.ID
	}

	d.SetId(time.Now().UTC().String())
	if err := d.Set("ids", ids); err != nil {
		return fmt.Errorf("Error setting ids: %s", err)
	}
	if err := d.Set("names", names); err != nil {
		return fmt.Errorf("Error setting names: %s", err)
	}

	return nil
}
