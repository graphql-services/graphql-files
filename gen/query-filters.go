package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/vektah/gqlparser/ast"
)

type FileQueryFilter struct {
	Query *string
}

func (qf *FileQueryFilter) Apply(ctx context.Context, dialect gorm.Dialect, selectionSet *ast.SelectionSet, wheres *[]string, values *[]interface{}, joins *[]string) error {
	if qf.Query == nil {
		return nil
	}

	fields := []*ast.Field{}
	if selectionSet != nil {
		for _, s := range *selectionSet {
			if f, ok := s.(*ast.Field); ok {
				fields = append(fields, f)
			}
		}
	} else {
		return fmt.Errorf("Cannot query with 'q' attribute without items field.")
	}

	queryParts := strings.Split(*qf.Query, " ")
	for _, part := range queryParts {
		ors := []string{}
		if err := qf.applyQueryWithFields(dialect, fields, part, TableName("files"), &ors, values, joins); err != nil {
			return err
		}
		*wheres = append(*wheres, "("+strings.Join(ors, " OR ")+")")
	}
	return nil
}

func (qf *FileQueryFilter) applyQueryWithFields(dialect gorm.Dialect, fields []*ast.Field, query, alias string, ors *[]string, values *[]interface{}, joins *[]string) error {
	if len(fields) == 0 {
		return nil
	}

	fieldsMap := map[string][]*ast.Field{}
	for _, f := range fields {
		fieldsMap[f.Name] = append(fieldsMap[f.Name], f)
	}

	if _, ok := fieldsMap["email"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("email")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["Size"]; ok {

		cast := "TEXT"
		if dialect.GetName() == "mysql" {
			cast = "CHAR"
		}
		column := fmt.Sprintf("CAST(%s"+dialect.Quote("Size")+" AS %s)", dialect.Quote(alias)+".", cast)

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["ContentType"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("ContentType")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["URL"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("URL")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["Name"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("Name")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	if _, ok := fieldsMap["Reference"]; ok {

		column := dialect.Quote(alias) + "." + dialect.Quote("Reference")

		*ors = append(*ors, fmt.Sprintf("%[1]s LIKE ? OR %[1]s LIKE ?", column))
		*values = append(*values, query+"%", "% "+query+"%")
	}

	return nil
}
