package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": {
				Type:     schema.TypeString,
				Required: true,
			},
			"token": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"cloud_order": resourceOrder(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"cloud_order": dataSourceOrder(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}
func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	var diags diag.Diagnostics
	username := d.Get("username").(string)
	token := d.Get("token").(string)
	if (username == "shakirmengrani") && (token == "shakirmengrani") {
		// construct api client
		return "Login successfully", diags
	}
	diags = append(diags, diag.Diagnostic{
		Severity: diag.Error,
		Summary:  "Unable to authenticate user for authenticated",
		Detail:   "Unable to authenticate user for authenticated",
	})
	return nil, diags
}
