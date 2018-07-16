package blueprint

import (
	"github.com/go-openapi/spec"
)

type ExtendedMethods struct {
	Properties AbstractPropertiesMethodType
	Path       string
	Method     string
}

func (b BlueprintType) GetMethodMap() map[string]ExtendedMethods {
	type extendedOps struct {
		Ops    spec.Operation
		Path   string
		Method string
	}

	var ops map[string]extendedOps

	ops = make(map[string]extendedOps)

	addToOps := func(method string, path string, ops *spec.Operation, data map[string]extendedOps) {
		data[ops.ID] = extendedOps{Ops: *ops, Path: path, Method: method}
	}

	//TODO: nil check
	if b.API.Paths != nil {
		for k, v := range b.API.Paths.Paths {
			if v.Get != nil {
				addToOps("GET", k, v.Get, ops)
			}
			if v.Post != nil {
				addToOps("POST", k, v.Post, ops)
			}
			if v.Put != nil {
				addToOps("PUT", k, v.Put, ops)
			}
			if v.Delete != nil {
				addToOps("DELETE", k, v.Delete, ops)
			}
			if v.Head != nil {
				addToOps("HEAD", k, v.Head, ops)
			}
			if v.Options != nil {
				addToOps("OPTIONS", k, v.Options, ops)
			}
			if v.Patch != nil {
				addToOps("PATCH", k, v.Patch, ops)
			}
		}
	}

	methods := make(map[string]AbstractPropertiesMethodType)

	for _, v := range b.AbstractProperties {
		if v.MethodId == nil {
			continue
		}

		methods[*v.MethodId] = v
	}

	result := make(map[string]ExtendedMethods)

	for k, v := range methods {
		exOps := ops[k]
		result[k] = ExtendedMethods{
			Properties: v,
			Path:       exOps.Path,
			Method:     exOps.Method,
		}
	}

	return result
}
