package schemas

import (
	"github.com/eddycharly/terraform-provider-kops/pkg/api/resources"
	"github.com/eddycharly/terraform-provider-kops/pkg/api/utils"
	. "github.com/eddycharly/terraform-provider-kops/pkg/schemas"
	utilsschemas "github.com/eddycharly/terraform-provider-kops/pkg/schemas/utils"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var _ = Schema

func ResourceValidateOptions() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"skip":          OptionalBool(),
			"timeout":       OptionalDuration(),
			"poll_interval": OptionalDuration(),
		},
	}
}

func ExpandResourceValidateOptions(in map[string]interface{}) resources.ValidateOptions {
	if in == nil {
		panic("expand ValidateOptions failure, in is nil")
	}
	return resources.ValidateOptions{
		Skip: func(in interface{}) bool {
			return bool(ExpandBool(in))
		}(in["skip"]),
		ValidateOptions: func(in interface{}) utils.ValidateOptions {
			return utilsschemas.ExpandResourceValidateOptions(in.(map[string]interface{}))
		}(in),
	}
}

func FlattenResourceValidateOptionsInto(in resources.ValidateOptions, out map[string]interface{}) {
	out["skip"] = func(in bool) interface{} {
		return FlattenBool(bool(in))
	}(in.Skip)
	utilsschemas.FlattenResourceValidateOptionsInto(in.ValidateOptions, out)
}

func FlattenResourceValidateOptions(in resources.ValidateOptions) map[string]interface{} {
	out := map[string]interface{}{}
	FlattenResourceValidateOptionsInto(in, out)
	return out
}
