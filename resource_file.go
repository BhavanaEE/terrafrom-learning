package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceFileCreate,
        Read:   resourceFileRead,
        Update: resourceFileUpdate,
        Delete: resourceFileDelete,

		Schema: map[string]*schema.Schema{
            "path": {
                Type:     schema.TypeString,
                Required: true,
            },
        },
	}
}

func resourceFileDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceFileUpdate(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceFileRead(d *schema.ResourceData, m interface{}) error {
	return nil
}
func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	return nil
}
