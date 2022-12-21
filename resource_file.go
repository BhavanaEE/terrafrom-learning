package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
	fmt.Println("Reading a file")
	fileName := d.Get("path").(string)

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
		return err
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)
	return nil
}

func resourceFileCreate(d *schema.ResourceData, m interface{}) error {
	filepath := d.Get("path").(string)
	fmt.Println("Writing to a file")

	file, err := os.Create("test.txt")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
		return err
	}

	defer file.Close()
	
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
		return err
	}

	fmt.Printf("\nSize: %d bytes", len(data))
	fmt.Printf("\nData: %s", data)


	len, err := file.WriteString(string(data))
	if err != nil {
		log.Fatalf("failed in writing content to file: %s", err)
		return err
	}
	fmt.Printf("\nwrote %d bytes\n", len)
	return nil
}
