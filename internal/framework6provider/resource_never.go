package framework

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type resourceNeverType struct{}

var _ resource.Resource = (*resourceNever)(nil)

func NewNeverResource() resource.Resource {
	return &resourceNever{}
}

func (r resourceNeverType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"name": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (r resourceNeverType) NewResource(_ context.Context, p provider.Provider) (resource.Resource, diag.Diagnostics) {
	return resourceNever{
		p: *(p.(*testProvider)),
	}, nil
}

type resourceNever struct {
	p testProvider
}

type never struct {
	Name string `tfsdk:"name"`
}

func (r resourceNever) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	for {
		time.Sleep(time.Second)
		log.Printf("[DEBUG] Zzz")
	}
}

func (r resourceNever) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state never
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Name = "Hello, World!"

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceNever) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NOP
}

func (r resourceNever) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
