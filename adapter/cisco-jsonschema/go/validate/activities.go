package validate

import (
	"context"
	"fmt"

	"cisco.com/cwm/adapters/cisco/jsonschema/common"
	"github.com/xeipuuv/gojsonschema"
)

func (adp *Adapter) Validate(ctx context.Context, req *Request, cfg *common.Config) (*Response, error) {
	var response string
	var str1 string

	schemaLoader := gojsonschema.NewGoLoader(req.GetJsonSchema())
	documentLoader := gojsonschema.NewGoLoader(req.GetJsonInput())

	result, err := gojsonschema.Validate(schemaLoader, documentLoader)
	if err != nil {
		message := fmt.Sprintf("Received error while validating JSON Input against JSON Schema %s", err.Error())
		response = message
	} else {
		if result.Valid() {
			response = "The document is valid"
			common.GetLogger(ctx).Info("The document is valid")
		} else {
			response = "The document is not valid. see errors :\n"
			for _, desc := range result.Errors() {
				str1 = fmt.Sprintf("- %s\n", desc)
				response = response + str1
			}
			common.GetLogger(ctx).Info(response)
		}
	}

	out := &Response{
		Output: response,
	}
	return out, nil
}
