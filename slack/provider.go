package slack

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.com/nlopes/slack"
)

// Provider is for slack.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ConfigureFunc: func(d *schema.ResourceData) (interface{}, error) {
			accessToken := d.Get("access_token").(string)
			return &slackAPI{slack.New(accessToken)}, nil
		},
		DataSourcesMap: map[string]*schema.Resource{
			"slack_channel": dataSourceSlackChannel(),
			"slack_user":    dataSourceSlackUser(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"slack_channel": resourceChannel(),
		},
		Schema: map[string]*schema.Schema{
			"access_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_ACCESS_TOKEN", nil),
				Description: "The Slack Oauth Access Token",
			},
		},
	}
}

type slackAPI struct {
	*slack.Client
}

func (api *slackAPI) getUserInfo(name string) (user *slack.User, err error) {
	var users []slack.User
	if users, err = api.GetUsers(); err != nil {
		return
	}

	for _, u := range users {
		if strings.ToLower(u.Name) == strings.ToLower(name) {
			user = &u
			return
		}
	}

	err = fmt.Errorf("user '%s' is not found", name)

	return
}

func (api *slackAPI) getChannel(name string) (channel *slack.Channel, err error) {
	var channels []slack.Channel
	channels, err = api.GetChannels(false)
	if err != nil {
		return
	}

	for _, c := range channels {
		if strings.ToLower(c.Name) == strings.ToLower(name) {
			channel = &c
			return
		}
	}

	err = fmt.Errorf("channel '%s' is not found", name)

	return
}
