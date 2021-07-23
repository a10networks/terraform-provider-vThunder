package thunder

//Thunder resource IpPrefixList

import (
	"context"
	"fmt"
	"util"

	go_thunder "github.com/go_thunder/thunder"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIpPrefixList() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIpPrefixListCreate,
		UpdateContext: resourceIpPrefixListUpdate,
		ReadContext:   resourceIpPrefixListRead,
		DeleteContext: resourceIpPrefixListDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
			"rules": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"le": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"action": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"ipaddr": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "",
						},
						"any": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"seq": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
						"ge": {
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "",
						},
					},
				},
			},
			"uuid": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "",
			},
		},
	}
}

func resourceIpPrefixListCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics

	if client.Host != "" {
		logger.Println("[INFO] Creating IpPrefixList (Inside resourceIpPrefixListCreate) ")
		name := d.Get("name").(string)
		data := dataToIpPrefixList(d)
		logger.Println("[INFO] received V from method data to IpPrefixList --")
		d.SetId(name)
		err := go_thunder.PostIpPrefixList(client.Token, data, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}

		return resourceIpPrefixListRead(ctx, d, meta)

	}
	return diags
}

func resourceIpPrefixListRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logger := util.GetLoggerInstance()
	client := meta.(Thunder)

	var diags diag.Diagnostics
	logger.Println("[INFO] Reading IpPrefixList (Inside resourceIpPrefixListRead)")

	if client.Host != "" {
		name := d.Id()
		logger.Println("[INFO] Fetching service Read" + name)
		data, err := go_thunder.GetIpPrefixList(client.Token, name, client.Host)
		if err != nil {
			return diag.FromErr(err)
		}
		if data == nil {
			d.SetId("")
			return nil
		}
		return diags
	}
	return nil
}

func resourceIpPrefixListUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpPrefixListRead(ctx, d, meta)
}

func resourceIpPrefixListDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return resourceIpPrefixListRead(ctx, d, meta)
}
func dataToIpPrefixList(d *schema.ResourceData) go_thunder.PrefixList {
	var vc go_thunder.PrefixList
	var c go_thunder.PrefixListInstance
	c.Name = d.Get("name").(string)

	RulesCount := d.Get("rules.#").(int)
	c.Le = make([]go_thunder.Rules, 0, RulesCount)

	for i := 0; i < RulesCount; i++ {
		var obj1 go_thunder.Rules
		prefix := fmt.Sprintf("rules.%d.", i)
		obj1.Seq = d.Get(prefix + "seq").(int)
		obj1.Ipaddr = d.Get(prefix + "ipaddr").(string)
		obj1.Ge = d.Get(prefix + "ge").(int)
		obj1.Any = d.Get(prefix + "any").(int)
		obj1.Description = d.Get(prefix + "description").(string)
		obj1.Action = d.Get(prefix + "action").(string)
		obj1.Le = d.Get(prefix + "le").(int)
		c.Le = append(c.Le, obj1)
	}

	vc.UUID = c
	return vc
}
