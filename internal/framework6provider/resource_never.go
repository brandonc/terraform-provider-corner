package framework

import (
	"context"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

var _ resource.Resource = (*resourceNever)(nil)

func NewNeverResource() resource.Resource {
	return &resourceNever{}
}

func (r *resourceNever) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name": schema.StringAttribute{
				Required: true,
			},
		},
	}
}

func (r resourceNever) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_never"
}

type resourceNever struct{}

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
