package main

import (
	"bytes"
	"encoding/json"
	grafana "github.com/lbrlabs/pulumi-grafana/provider"
	"github.com/lbrlabs/pulumi-grafana/provider/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfgen"
	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	//"regexp"

	"github.com/pulumi/pulumi/sdk/v3/go/common/util/contract"
)

func main() {

	info := grafana.Provider()

	info.SchemaPostProcessor = func(spec *schema.PackageSpec) {

		// Temporary workaround until this is exposed in tfbridge.
		if val, ok := spec.Language["go"]; ok {
			var options map[string]any
			err := json.Unmarshal(val, &options)
			contract.AssertNoErrorf(err, "unexpected error unmarshalling go options")
			options["generics"] = "side-by-side"
			spec.Language["go"] = rawMessage(options)
		}

		//removeExamples(spec)
	}

	// Modify the path to point to the new provider
	tfgen.Main("grafana", version.Version, info)
}

func rawMessage(v any) schema.RawMessage {
	var out bytes.Buffer
	encoder := json.NewEncoder(&out)
	encoder.SetEscapeHTML(false)
	err := encoder.Encode(v)
	contract.AssertNoErrorf(err, "unexpected error while encoding JSON")
	return out.Bytes()
}

// var re = regexp.MustCompile(`(?ms)\n?\n?## Example Usage.*`)

// func removeExamples(spec *schema.PackageSpec) {
// 	for tok, t := range spec.Types {
// 		t.Description = re.ReplaceAllString(t.Description, "")
// 		spec.Types[tok] = t
// 	}
// 	for tok, r := range spec.Resources {
// 		r.Description = re.ReplaceAllString(r.Description, "")
// 		spec.Resources[tok] = r
// 	}
// 	for tok, f := range spec.Functions {
// 		f.Description = re.ReplaceAllString(f.Description, "")
// 		spec.Functions[tok] = f
// 	}
// }
