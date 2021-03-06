package google

import (
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"account_file": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("GOOGLE_ACCOUNT_FILE"),
			},

			"project": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("GOOGLE_PROJECT"),
			},

			"region": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: envDefaultFunc("GOOGLE_REGION"),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"google_compute_address":  resourceComputeAddress(),
			"google_compute_disk":     resourceComputeDisk(),
			"google_compute_firewall": resourceComputeFirewall(),
			"google_compute_instance": resourceComputeInstance(),
			"google_compute_network":  resourceComputeNetwork(),
			"google_compute_route":    resourceComputeRoute(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func envDefaultFunc(k string) schema.SchemaDefaultFunc {
	return func() (interface{}, error) {
		if v := os.Getenv(k); v != "" {
			return v, nil
		}

		return nil, nil
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		AccountFile: d.Get("account_file").(string),
		Project:     d.Get("project").(string),
		Region:      d.Get("region").(string),
	}

	if err := config.loadAndValidate(); err != nil {
		return nil, err
	}

	return &config, nil
}
