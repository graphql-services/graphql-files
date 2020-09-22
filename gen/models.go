package gen

import (
	"fmt"
	"reflect"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
)

type FileResultType struct {
	EntityResultType
}

type File struct {
	ID          string     `json:"id" gorm:"column:id;primary_key"`
	Email       *string    `json:"email" gorm:"column:email"`
	UID         *string    `json:"UID" gorm:"column:UID"`
	Size        *int       `json:"Size" gorm:"column:Size"`
	ContentType *string    `json:"ContentType" gorm:"column:ContentType"`
	URL         *string    `json:"URL" gorm:"column:URL"`
	Name        *string    `json:"Name" gorm:"column:Name"`
	Reference   *string    `json:"Reference" gorm:"column:Reference;index:findref"`
	ReferenceID *string    `json:"ReferenceID" gorm:"column:ReferenceID;index:findref"`
	UpdatedAt   *time.Time `json:"updatedAt" gorm:"column:updatedAt"`
	CreatedAt   time.Time  `json:"createdAt" gorm:"column:createdAt"`
	UpdatedBy   *string    `json:"updatedBy" gorm:"column:updatedBy"`
	CreatedBy   *string    `json:"createdBy" gorm:"column:createdBy"`
}

func (m *File) Is_Entity() {}

type FileChanges struct {
	ID          string
	Email       *string
	UID         *string
	Size        *int
	ContentType *string
	URL         *string
	Name        *string
	Reference   *string
	ReferenceID *string
	UpdatedAt   *time.Time
	CreatedAt   time.Time
	UpdatedBy   *string
	CreatedBy   *string
}

// used to convert map[string]interface{} to EntityChanges struct
func ApplyChanges(changes map[string]interface{}, to interface{}) error {
	dec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		ErrorUnused: true,
		TagName:     "json",
		Result:      to,
		ZeroFields:  true,
		// This is needed to get mapstructure to call the gqlgen unmarshaler func for custom scalars (eg Date)
		DecodeHook: func(a reflect.Type, b reflect.Type, v interface{}) (interface{}, error) {

			if b == reflect.TypeOf(time.Time{}) {
				switch a.Kind() {
				case reflect.String:
					return time.Parse(time.RFC3339, v.(string))
				case reflect.Float64:
					return time.Unix(0, int64(v.(float64))*int64(time.Millisecond)), nil
				case reflect.Int64:
					return time.Unix(0, v.(int64)*int64(time.Millisecond)), nil
				default:
					return v, fmt.Errorf("Unable to parse date from %v", v)
				}
			}

			if reflect.PtrTo(b).Implements(reflect.TypeOf((*graphql.Unmarshaler)(nil)).Elem()) {
				resultType := reflect.New(b)
				result := resultType.MethodByName("UnmarshalGQL").Call([]reflect.Value{reflect.ValueOf(v)})
				err, _ := result[0].Interface().(error)
				return resultType.Elem().Interface(), err
			}

			return v, nil
		},
	})

	if err != nil {
		return err
	}

	return dec.Decode(changes)
}
