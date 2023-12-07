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

type resourceNeverCreatesType struct{}

var _ resource.Resource = (*resourceNeverCreates)(nil)

func (r resourceNeverCreatesType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"foo": {
				Type:     types.StringType,
				Required: true,
			},
		},
	}, nil
}

func (r resourceNeverCreatesType) NewResource(_ context.Context, p provider.Provider) (resource.Resource, diag.Diagnostics) {
	return resourceNeverCreates{
		p: *(p.(*testProvider)),
	}, nil
}

type resourceNeverCreates struct {
	p testProvider
}

type neverCreates struct {
	Foo string `tfsdk:"foo"`
}

func (r resourceNeverCreates) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	for {
		time.Sleep(time.Second)
		log.Printf("[DEBUG] Zzz")
	}
}

func (r resourceNeverCreates) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state neverCreates
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	state.Foo = "Bar"

	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

func (r resourceNeverCreates) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// NOP
}

func (r resourceNeverCreates) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	resp.State.RemoveResource(ctx)
}
