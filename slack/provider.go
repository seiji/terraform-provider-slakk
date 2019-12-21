package slack

import (
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
