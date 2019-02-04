package infoblox

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

//Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USERNAME", nil),
				Description: "Infoblox Username",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_PASSWORD", nil),
				Description: "Infoblox User Password",
			},
			"host": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_HOST", nil),
				Description: "Infoblox Base Url(defaults to testing)",
			},
			"sslverify": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_SSLVERIFY", true),
				Description: "Enable ssl",
			},
			"usecookies": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("INFOBLOX_USECOOKIES", false),
				Description: "Use cookies",
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"infoblox_record": resourceInfobloxRecord(),
			"infoblox_ip":     resourceInfobloxIP(),
			"infoblox_record_a":     infobloxRecordA(),
			"infoblox_record_aaaa":  infobloxRecordAAAA(),
			"infoblox_record_cname": infobloxRecordCNAME(),
			"infoblox_record_ptr":   infobloxRecordPTR(),
			"infoblox_record_host":  infobloxRecordHost(),
			"infoblox_record_txt":   infobloxRecordTXT(),
			"infoblox_record_mx":    infobloxRecordMX(),
			"infoblox_record_srv":   infobloxRecordSRV(),
		},

		ConfigureFunc: provideConfigure,
	}
}

func provideConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		Username:   d.Get("username").(string),
		Password:   d.Get("password").(string),
		Host:       d.Get("host").(string),
		SSLVerify:  d.Get("sslverify").(bool),
		UseCookies: d.Get("usecookies").(bool),
	}

	return config.Client()
}
