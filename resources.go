package main

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var schemaOrder = map[string]*schema.Schema{
	"items": {
		Type:     schema.TypeList,
		Required: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"name": {
					Type:     schema.TypeString,
					Required: true,
				},
				"rate": {
					Type:     schema.TypeFloat,
					Required: true,
				},
				"qty": {
					Type:     schema.TypeFloat,
					Required: true,
				},
			},
		},
	},
}

func resourceOrder() *schema.Resource {
	return &schema.Resource{
		CreateContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			return nil
		},
		ReadContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			return nil
		},
		UpdateContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			return nil
		},
		DeleteContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			return nil
		},
		Schema: schemaOrder,
	}
}

func dataSourceOrder() *schema.Resource {
	return &schema.Resource{
		ReadContext: func(ctx context.Context, rd *schema.ResourceData, i interface{}) diag.Diagnostics {
			var diags diag.Diagnostics

			return diags
		},
		Schema: schemaOrder,
	}
}
