package gen

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
)

func (r *GeneratedQueryResolver) _service(ctx context.Context) (*_Service, error) {
	sdl := SchemaSDL
	return &_Service{
		Sdl: &sdl,
	}, nil
}

func getExecutionContext(ctx context.Context) executionContext {
	e := ctx.Value(KeyExecutableSchema).(*executableSchema)
	return executionContext{graphql.GetRequestContext(ctx), e}
}

func (r *GeneratedQueryResolver) _entities(ctx context.Context, representations []interface{}) (res []_Entity, err error) {
	res = []_Entity{}
	for _, repr := range representations {
		anyValue, ok := repr.(map[string]interface{})
		if !ok {
			err = fmt.Errorf("The _entities resolver received invalid representation type")
			break
		}
		typename, ok := anyValue["__typename"].(string)
		if !ok {
			err = fmt.Errorf("The _entities resolver received invalid representation type (missing __typename field)")
			break
		}

		switch typename {
		case "File":
			ec := getExecutionContext(ctx)
			f, _err := ec.unmarshalInputFileFilterType(ctx, anyValue)
			err = _err
			if err != nil {
				return
			}

			if f.IsEmpty(ctx, r.GetDB(ctx).Dialect()) {
				res = append(res, nil)
				continue
			}

			item, _err := r.File(ctx, nil, nil, &f)
			if _err != nil {
				err = _err
				return
			}
			if item == nil {
				// append nil object without reflected interface
				res = append(res, nil)
			} else {
				res = append(res, item)
			}
			break
		default:
			err = fmt.Errorf("The _entities resolver tried to load an entity for type \"%s\", but no object type of that name was found in the schema", typename)
			return
		}
	}
	return res, err
}
