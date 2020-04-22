package slack

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceSlackEmojis() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackEmojisRead,
		Schema: map[string]*schema.Schema{
			"emojis": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
			"aliases": {
				Type:     schema.TypeMap,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Computed: true,
			},
		},
	}
}

func dataSourceSlackEmojisRead(d *schema.ResourceData, m interface{}) error {
	api := m.(*slackAPI)

	log.Println("[DEBUG] Fetching emojis")
	emojiMap, err := api.GetEmoji()
	if err != nil {
		return fmt.Errorf("Error fetching emojis: %s", err)
	}

	pp(emojiMap)
	emojis := map[string]string{}
	aliases := map[string]string{}
	for k, v := range emojiMap {
		if strings.HasPrefix(v, "alias:") {
			aliases[k] = strings.Replace(v, "alias:", "", 1)
		} else {
			emojis[k] = v
		}
	}

	d.SetId(time.Now().UTC().String())
	if err := d.Set("emojis", emojis); err != nil {
		return fmt.Errorf("Error setting emojis: %s", err)
	}
	if err := d.Set("aliases", aliases); err != nil {
		return fmt.Errorf("Error setting aliases: %s", err)
	}

	return nil
}
