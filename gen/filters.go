package gen

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
)

func (f *FileFilterType) IsEmpty(ctx context.Context, dialect gorm.Dialect) bool {
	wheres := []string{}
	havings := []string{}
	whereValues := []interface{}{}
	havingValues := []interface{}{}
	joins := []string{}
	err := f.ApplyWithAlias(ctx, dialect, "companies", &wheres, &whereValues, &havings, &havingValues, &joins)
	if err != nil {
		panic(err)
	}
	return len(wheres) == 0 && len(havings) == 0
}
func (f *FileFilterType) Apply(ctx context.Context, dialect gorm.Dialect, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	return f.ApplyWithAlias(ctx, dialect, TableName("files"), wheres, whereValues, havings, havingValues, joins)
}
func (f *FileFilterType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, wheres *[]string, whereValues *[]interface{}, havings *[]string, havingValues *[]interface{}, joins *[]string) error {
	if f == nil {
		return nil
	}
	aliasPrefix := dialect.Quote(alias) + "."

	_where, _whereValues := f.WhereContent(dialect, aliasPrefix)
	_having, _havingValues := f.HavingContent(dialect, aliasPrefix)
	*wheres = append(*wheres, _where...)
	*havings = append(*havings, _having...)
	*whereValues = append(*whereValues, _whereValues...)
	*havingValues = append(*havingValues, _havingValues...)

	if f.Or != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, or := range f.Or {
			_ws := []string{}
			_hs := []string{}
			err := or.ApplyWithAlias(ctx, dialect, alias, &_ws, &wvs, &_hs, &hvs, &js)
			if err != nil {
				return err
			}
			if len(_ws) > 0 {
				ws = append(ws, strings.Join(_ws, " AND "))
			}
			if len(_hs) > 0 {
				hs = append(hs, strings.Join(_hs, " AND "))
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, "("+strings.Join(ws, " OR ")+")")
		}
		if len(hs) > 0 {
			*havings = append(*havings, "("+strings.Join(hs, " OR ")+")")
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}
	if f.And != nil {
		ws := []string{}
		hs := []string{}
		wvs := []interface{}{}
		hvs := []interface{}{}
		js := []string{}
		for _, and := range f.And {
			err := and.ApplyWithAlias(ctx, dialect, alias, &ws, &wvs, &hs, &hvs, &js)
			if err != nil {
				return err
			}
		}
		if len(ws) > 0 {
			*wheres = append(*wheres, strings.Join(ws, " AND "))
		}
		if len(hs) > 0 {
			*havings = append(*havings, strings.Join(hs, " AND "))
		}
		*whereValues = append(*whereValues, wvs...)
		*havingValues = append(*havingValues, hvs...)
		*joins = append(*joins, js...)
	}

	return nil
}

func (f *FileFilterType) WhereContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.ID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" = ?")
		values = append(values, f.ID)
	}

	if f.IDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" != ?")
		values = append(values, f.IDNe)
	}

	if f.IDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" > ?")
		values = append(values, f.IDGt)
	}

	if f.IDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" < ?")
		values = append(values, f.IDLt)
	}

	if f.IDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" >= ?")
		values = append(values, f.IDGte)
	}

	if f.IDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" <= ?")
		values = append(values, f.IDLte)
	}

	if f.IDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IN (?)")
		values = append(values, f.IDIn)
	}

	if f.IDNull != nil {
		if *f.IDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("id")+" IS NOT NULL")
		}
	}

	if f.Email != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" = ?")
		values = append(values, f.Email)
	}

	if f.EmailNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" != ?")
		values = append(values, f.EmailNe)
	}

	if f.EmailGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" > ?")
		values = append(values, f.EmailGt)
	}

	if f.EmailLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" < ?")
		values = append(values, f.EmailLt)
	}

	if f.EmailGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" >= ?")
		values = append(values, f.EmailGte)
	}

	if f.EmailLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" <= ?")
		values = append(values, f.EmailLte)
	}

	if f.EmailIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IN (?)")
		values = append(values, f.EmailIn)
	}

	if f.EmailLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailPrefix))
	}

	if f.EmailSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailSuffix))
	}

	if f.EmailNull != nil {
		if *f.EmailNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("email")+" IS NOT NULL")
		}
	}

	if f.UID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" = ?")
		values = append(values, f.UID)
	}

	if f.UIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" != ?")
		values = append(values, f.UIDNe)
	}

	if f.UIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" > ?")
		values = append(values, f.UIDGt)
	}

	if f.UIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" < ?")
		values = append(values, f.UIDLt)
	}

	if f.UIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" >= ?")
		values = append(values, f.UIDGte)
	}

	if f.UIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" <= ?")
		values = append(values, f.UIDLte)
	}

	if f.UIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" IN (?)")
		values = append(values, f.UIDIn)
	}

	if f.UIDNull != nil {
		if *f.UIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("uid")+" IS NOT NULL")
		}
	}

	if f.Name != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" = ?")
		values = append(values, f.Name)
	}

	if f.NameNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" != ?")
		values = append(values, f.NameNe)
	}

	if f.NameGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" > ?")
		values = append(values, f.NameGt)
	}

	if f.NameLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" < ?")
		values = append(values, f.NameLt)
	}

	if f.NameGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" >= ?")
		values = append(values, f.NameGte)
	}

	if f.NameLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" <= ?")
		values = append(values, f.NameLte)
	}

	if f.NameIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IN (?)")
		values = append(values, f.NameIn)
	}

	if f.NameLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NamePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NamePrefix))
	}

	if f.NameSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameSuffix))
	}

	if f.NameNull != nil {
		if *f.NameNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("name")+" IS NOT NULL")
		}
	}

	if f.Size != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" = ?")
		values = append(values, f.Size)
	}

	if f.SizeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" != ?")
		values = append(values, f.SizeNe)
	}

	if f.SizeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" > ?")
		values = append(values, f.SizeGt)
	}

	if f.SizeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" < ?")
		values = append(values, f.SizeLt)
	}

	if f.SizeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" >= ?")
		values = append(values, f.SizeGte)
	}

	if f.SizeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" <= ?")
		values = append(values, f.SizeLte)
	}

	if f.SizeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" IN (?)")
		values = append(values, f.SizeIn)
	}

	if f.SizeNull != nil {
		if *f.SizeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("size")+" IS NOT NULL")
		}
	}

	if f.ContentType != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" = ?")
		values = append(values, f.ContentType)
	}

	if f.ContentTypeNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" != ?")
		values = append(values, f.ContentTypeNe)
	}

	if f.ContentTypeGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" > ?")
		values = append(values, f.ContentTypeGt)
	}

	if f.ContentTypeLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" < ?")
		values = append(values, f.ContentTypeLt)
	}

	if f.ContentTypeGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" >= ?")
		values = append(values, f.ContentTypeGte)
	}

	if f.ContentTypeLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" <= ?")
		values = append(values, f.ContentTypeLte)
	}

	if f.ContentTypeIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" IN (?)")
		values = append(values, f.ContentTypeIn)
	}

	if f.ContentTypeLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ContentTypeLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ContentTypePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentTypePrefix))
	}

	if f.ContentTypeSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentTypeSuffix))
	}

	if f.ContentTypeNull != nil {
		if *f.ContentTypeNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("contentType")+" IS NOT NULL")
		}
	}

	if f.URL != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" = ?")
		values = append(values, f.URL)
	}

	if f.URLNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" != ?")
		values = append(values, f.URLNe)
	}

	if f.URLGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" > ?")
		values = append(values, f.URLGt)
	}

	if f.URLLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" < ?")
		values = append(values, f.URLLt)
	}

	if f.URLGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" >= ?")
		values = append(values, f.URLGte)
	}

	if f.URLLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" <= ?")
		values = append(values, f.URLLte)
	}

	if f.URLIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" IN (?)")
		values = append(values, f.URLIn)
	}

	if f.URLLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.URLLike, "?", "_", -1), "*", "%", -1))
	}

	if f.URLPrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.URLPrefix))
	}

	if f.URLSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.URLSuffix))
	}

	if f.URLNull != nil {
		if *f.URLNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("url")+" IS NOT NULL")
		}
	}

	if f.Reference != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" = ?")
		values = append(values, f.Reference)
	}

	if f.ReferenceNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" != ?")
		values = append(values, f.ReferenceNe)
	}

	if f.ReferenceGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" > ?")
		values = append(values, f.ReferenceGt)
	}

	if f.ReferenceLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" < ?")
		values = append(values, f.ReferenceLt)
	}

	if f.ReferenceGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" >= ?")
		values = append(values, f.ReferenceGte)
	}

	if f.ReferenceLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" <= ?")
		values = append(values, f.ReferenceLte)
	}

	if f.ReferenceIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IN (?)")
		values = append(values, f.ReferenceIn)
	}

	if f.ReferenceLike != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferencePrefix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferencePrefix))
	}

	if f.ReferenceSuffix != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceSuffix))
	}

	if f.ReferenceNull != nil {
		if *f.ReferenceNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("reference")+" IS NOT NULL")
		}
	}

	if f.ReferenceID != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" = ?")
		values = append(values, f.ReferenceID)
	}

	if f.ReferenceIDNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" != ?")
		values = append(values, f.ReferenceIDNe)
	}

	if f.ReferenceIDGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" > ?")
		values = append(values, f.ReferenceIDGt)
	}

	if f.ReferenceIDLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" < ?")
		values = append(values, f.ReferenceIDLt)
	}

	if f.ReferenceIDGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" >= ?")
		values = append(values, f.ReferenceIDGte)
	}

	if f.ReferenceIDLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" <= ?")
		values = append(values, f.ReferenceIDLte)
	}

	if f.ReferenceIDIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IN (?)")
		values = append(values, f.ReferenceIDIn)
	}

	if f.ReferenceIDNull != nil {
		if *f.ReferenceIDNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("referenceID")+" IS NOT NULL")
		}
	}

	if f.UpdatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" = ?")
		values = append(values, f.UpdatedAt)
	}

	if f.UpdatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" != ?")
		values = append(values, f.UpdatedAtNe)
	}

	if f.UpdatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" > ?")
		values = append(values, f.UpdatedAtGt)
	}

	if f.UpdatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" < ?")
		values = append(values, f.UpdatedAtLt)
	}

	if f.UpdatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" >= ?")
		values = append(values, f.UpdatedAtGte)
	}

	if f.UpdatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" <= ?")
		values = append(values, f.UpdatedAtLte)
	}

	if f.UpdatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IN (?)")
		values = append(values, f.UpdatedAtIn)
	}

	if f.UpdatedAtNull != nil {
		if *f.UpdatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedAt")+" IS NOT NULL")
		}
	}

	if f.CreatedAt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" = ?")
		values = append(values, f.CreatedAt)
	}

	if f.CreatedAtNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" != ?")
		values = append(values, f.CreatedAtNe)
	}

	if f.CreatedAtGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" > ?")
		values = append(values, f.CreatedAtGt)
	}

	if f.CreatedAtLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" < ?")
		values = append(values, f.CreatedAtLt)
	}

	if f.CreatedAtGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" >= ?")
		values = append(values, f.CreatedAtGte)
	}

	if f.CreatedAtLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" <= ?")
		values = append(values, f.CreatedAtLte)
	}

	if f.CreatedAtIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IN (?)")
		values = append(values, f.CreatedAtIn)
	}

	if f.CreatedAtNull != nil {
		if *f.CreatedAtNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdAt")+" IS NOT NULL")
		}
	}

	if f.UpdatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" = ?")
		values = append(values, f.UpdatedBy)
	}

	if f.UpdatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" != ?")
		values = append(values, f.UpdatedByNe)
	}

	if f.UpdatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" > ?")
		values = append(values, f.UpdatedByGt)
	}

	if f.UpdatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" < ?")
		values = append(values, f.UpdatedByLt)
	}

	if f.UpdatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" >= ?")
		values = append(values, f.UpdatedByGte)
	}

	if f.UpdatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" <= ?")
		values = append(values, f.UpdatedByLte)
	}

	if f.UpdatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IN (?)")
		values = append(values, f.UpdatedByIn)
	}

	if f.UpdatedByNull != nil {
		if *f.UpdatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("updatedBy")+" IS NOT NULL")
		}
	}

	if f.CreatedBy != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" = ?")
		values = append(values, f.CreatedBy)
	}

	if f.CreatedByNe != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" != ?")
		values = append(values, f.CreatedByNe)
	}

	if f.CreatedByGt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" > ?")
		values = append(values, f.CreatedByGt)
	}

	if f.CreatedByLt != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" < ?")
		values = append(values, f.CreatedByLt)
	}

	if f.CreatedByGte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" >= ?")
		values = append(values, f.CreatedByGte)
	}

	if f.CreatedByLte != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" <= ?")
		values = append(values, f.CreatedByLte)
	}

	if f.CreatedByIn != nil {
		conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IN (?)")
		values = append(values, f.CreatedByIn)
	}

	if f.CreatedByNull != nil {
		if *f.CreatedByNull {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NULL")
		} else {
			conditions = append(conditions, aliasPrefix+dialect.Quote("createdBy")+" IS NOT NULL")
		}
	}

	return
}
func (f *FileFilterType) HavingContent(dialect gorm.Dialect, aliasPrefix string) (conditions []string, values []interface{}) {
	conditions = []string{}
	values = []interface{}{}

	if f.IDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMin)
	}

	if f.IDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") = ?")
		values = append(values, f.IDMax)
	}

	if f.IDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMinNe)
	}

	if f.IDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") != ?")
		values = append(values, f.IDMaxNe)
	}

	if f.IDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMinGt)
	}

	if f.IDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") > ?")
		values = append(values, f.IDMaxGt)
	}

	if f.IDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMinLt)
	}

	if f.IDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") < ?")
		values = append(values, f.IDMaxLt)
	}

	if f.IDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMinGte)
	}

	if f.IDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") >= ?")
		values = append(values, f.IDMaxGte)
	}

	if f.IDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMinLte)
	}

	if f.IDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") <= ?")
		values = append(values, f.IDMaxLte)
	}

	if f.IDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMinIn)
	}

	if f.IDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("id")+") IN (?)")
		values = append(values, f.IDMaxIn)
	}

	if f.EmailMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMin)
	}

	if f.EmailMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") = ?")
		values = append(values, f.EmailMax)
	}

	if f.EmailMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMinNe)
	}

	if f.EmailMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") != ?")
		values = append(values, f.EmailMaxNe)
	}

	if f.EmailMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMinGt)
	}

	if f.EmailMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") > ?")
		values = append(values, f.EmailMaxGt)
	}

	if f.EmailMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMinLt)
	}

	if f.EmailMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") < ?")
		values = append(values, f.EmailMaxLt)
	}

	if f.EmailMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMinGte)
	}

	if f.EmailMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") >= ?")
		values = append(values, f.EmailMaxGte)
	}

	if f.EmailMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMinLte)
	}

	if f.EmailMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") <= ?")
		values = append(values, f.EmailMaxLte)
	}

	if f.EmailMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMinIn)
	}

	if f.EmailMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") IN (?)")
		values = append(values, f.EmailMaxIn)
	}

	if f.EmailMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.EmailMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.EmailMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMinPrefix))
	}

	if f.EmailMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.EmailMaxPrefix))
	}

	if f.EmailMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMinSuffix))
	}

	if f.EmailMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("email")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.EmailMaxSuffix))
	}

	if f.UIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") = ?")
		values = append(values, f.UIDMin)
	}

	if f.UIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") = ?")
		values = append(values, f.UIDMax)
	}

	if f.UIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") != ?")
		values = append(values, f.UIDMinNe)
	}

	if f.UIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") != ?")
		values = append(values, f.UIDMaxNe)
	}

	if f.UIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") > ?")
		values = append(values, f.UIDMinGt)
	}

	if f.UIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") > ?")
		values = append(values, f.UIDMaxGt)
	}

	if f.UIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") < ?")
		values = append(values, f.UIDMinLt)
	}

	if f.UIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") < ?")
		values = append(values, f.UIDMaxLt)
	}

	if f.UIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") >= ?")
		values = append(values, f.UIDMinGte)
	}

	if f.UIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") >= ?")
		values = append(values, f.UIDMaxGte)
	}

	if f.UIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") <= ?")
		values = append(values, f.UIDMinLte)
	}

	if f.UIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") <= ?")
		values = append(values, f.UIDMaxLte)
	}

	if f.UIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("uid")+") IN (?)")
		values = append(values, f.UIDMinIn)
	}

	if f.UIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("uid")+") IN (?)")
		values = append(values, f.UIDMaxIn)
	}

	if f.NameMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMin)
	}

	if f.NameMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") = ?")
		values = append(values, f.NameMax)
	}

	if f.NameMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMinNe)
	}

	if f.NameMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") != ?")
		values = append(values, f.NameMaxNe)
	}

	if f.NameMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMinGt)
	}

	if f.NameMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") > ?")
		values = append(values, f.NameMaxGt)
	}

	if f.NameMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMinLt)
	}

	if f.NameMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") < ?")
		values = append(values, f.NameMaxLt)
	}

	if f.NameMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMinGte)
	}

	if f.NameMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") >= ?")
		values = append(values, f.NameMaxGte)
	}

	if f.NameMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMinLte)
	}

	if f.NameMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") <= ?")
		values = append(values, f.NameMaxLte)
	}

	if f.NameMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMinIn)
	}

	if f.NameMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") IN (?)")
		values = append(values, f.NameMaxIn)
	}

	if f.NameMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.NameMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.NameMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMinPrefix))
	}

	if f.NameMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.NameMaxPrefix))
	}

	if f.NameMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMinSuffix))
	}

	if f.NameMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("name")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.NameMaxSuffix))
	}

	if f.SizeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") = ?")
		values = append(values, f.SizeMin)
	}

	if f.SizeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") = ?")
		values = append(values, f.SizeMax)
	}

	if f.SizeAvg != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") = ?")
		values = append(values, f.SizeAvg)
	}

	if f.SizeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") != ?")
		values = append(values, f.SizeMinNe)
	}

	if f.SizeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") != ?")
		values = append(values, f.SizeMaxNe)
	}

	if f.SizeAvgNe != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") != ?")
		values = append(values, f.SizeAvgNe)
	}

	if f.SizeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") > ?")
		values = append(values, f.SizeMinGt)
	}

	if f.SizeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") > ?")
		values = append(values, f.SizeMaxGt)
	}

	if f.SizeAvgGt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") > ?")
		values = append(values, f.SizeAvgGt)
	}

	if f.SizeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") < ?")
		values = append(values, f.SizeMinLt)
	}

	if f.SizeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") < ?")
		values = append(values, f.SizeMaxLt)
	}

	if f.SizeAvgLt != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") < ?")
		values = append(values, f.SizeAvgLt)
	}

	if f.SizeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") >= ?")
		values = append(values, f.SizeMinGte)
	}

	if f.SizeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") >= ?")
		values = append(values, f.SizeMaxGte)
	}

	if f.SizeAvgGte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") >= ?")
		values = append(values, f.SizeAvgGte)
	}

	if f.SizeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") <= ?")
		values = append(values, f.SizeMinLte)
	}

	if f.SizeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") <= ?")
		values = append(values, f.SizeMaxLte)
	}

	if f.SizeAvgLte != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") <= ?")
		values = append(values, f.SizeAvgLte)
	}

	if f.SizeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("size")+") IN (?)")
		values = append(values, f.SizeMinIn)
	}

	if f.SizeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("size")+") IN (?)")
		values = append(values, f.SizeMaxIn)
	}

	if f.SizeAvgIn != nil {
		conditions = append(conditions, "Avg("+aliasPrefix+dialect.Quote("size")+") IN (?)")
		values = append(values, f.SizeAvgIn)
	}

	if f.ContentTypeMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") = ?")
		values = append(values, f.ContentTypeMin)
	}

	if f.ContentTypeMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") = ?")
		values = append(values, f.ContentTypeMax)
	}

	if f.ContentTypeMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") != ?")
		values = append(values, f.ContentTypeMinNe)
	}

	if f.ContentTypeMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") != ?")
		values = append(values, f.ContentTypeMaxNe)
	}

	if f.ContentTypeMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") > ?")
		values = append(values, f.ContentTypeMinGt)
	}

	if f.ContentTypeMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") > ?")
		values = append(values, f.ContentTypeMaxGt)
	}

	if f.ContentTypeMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") < ?")
		values = append(values, f.ContentTypeMinLt)
	}

	if f.ContentTypeMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") < ?")
		values = append(values, f.ContentTypeMaxLt)
	}

	if f.ContentTypeMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") >= ?")
		values = append(values, f.ContentTypeMinGte)
	}

	if f.ContentTypeMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") >= ?")
		values = append(values, f.ContentTypeMaxGte)
	}

	if f.ContentTypeMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") <= ?")
		values = append(values, f.ContentTypeMinLte)
	}

	if f.ContentTypeMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") <= ?")
		values = append(values, f.ContentTypeMaxLte)
	}

	if f.ContentTypeMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") IN (?)")
		values = append(values, f.ContentTypeMinIn)
	}

	if f.ContentTypeMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") IN (?)")
		values = append(values, f.ContentTypeMaxIn)
	}

	if f.ContentTypeMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ContentTypeMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ContentTypeMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ContentTypeMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ContentTypeMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentTypeMinPrefix))
	}

	if f.ContentTypeMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ContentTypeMaxPrefix))
	}

	if f.ContentTypeMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("contentType")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentTypeMinSuffix))
	}

	if f.ContentTypeMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("contentType")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ContentTypeMaxSuffix))
	}

	if f.URLMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") = ?")
		values = append(values, f.URLMin)
	}

	if f.URLMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") = ?")
		values = append(values, f.URLMax)
	}

	if f.URLMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") != ?")
		values = append(values, f.URLMinNe)
	}

	if f.URLMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") != ?")
		values = append(values, f.URLMaxNe)
	}

	if f.URLMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") > ?")
		values = append(values, f.URLMinGt)
	}

	if f.URLMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") > ?")
		values = append(values, f.URLMaxGt)
	}

	if f.URLMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") < ?")
		values = append(values, f.URLMinLt)
	}

	if f.URLMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") < ?")
		values = append(values, f.URLMaxLt)
	}

	if f.URLMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") >= ?")
		values = append(values, f.URLMinGte)
	}

	if f.URLMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") >= ?")
		values = append(values, f.URLMaxGte)
	}

	if f.URLMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") <= ?")
		values = append(values, f.URLMinLte)
	}

	if f.URLMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") <= ?")
		values = append(values, f.URLMaxLte)
	}

	if f.URLMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") IN (?)")
		values = append(values, f.URLMinIn)
	}

	if f.URLMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") IN (?)")
		values = append(values, f.URLMaxIn)
	}

	if f.URLMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.URLMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.URLMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.URLMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.URLMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.URLMinPrefix))
	}

	if f.URLMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.URLMaxPrefix))
	}

	if f.URLMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.URLMinSuffix))
	}

	if f.URLMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("url")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.URLMaxSuffix))
	}

	if f.ReferenceMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") = ?")
		values = append(values, f.ReferenceMin)
	}

	if f.ReferenceMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") = ?")
		values = append(values, f.ReferenceMax)
	}

	if f.ReferenceMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") != ?")
		values = append(values, f.ReferenceMinNe)
	}

	if f.ReferenceMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") != ?")
		values = append(values, f.ReferenceMaxNe)
	}

	if f.ReferenceMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") > ?")
		values = append(values, f.ReferenceMinGt)
	}

	if f.ReferenceMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") > ?")
		values = append(values, f.ReferenceMaxGt)
	}

	if f.ReferenceMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") < ?")
		values = append(values, f.ReferenceMinLt)
	}

	if f.ReferenceMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") < ?")
		values = append(values, f.ReferenceMaxLt)
	}

	if f.ReferenceMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") >= ?")
		values = append(values, f.ReferenceMinGte)
	}

	if f.ReferenceMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") >= ?")
		values = append(values, f.ReferenceMaxGte)
	}

	if f.ReferenceMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") <= ?")
		values = append(values, f.ReferenceMinLte)
	}

	if f.ReferenceMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") <= ?")
		values = append(values, f.ReferenceMaxLte)
	}

	if f.ReferenceMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") IN (?)")
		values = append(values, f.ReferenceMinIn)
	}

	if f.ReferenceMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") IN (?)")
		values = append(values, f.ReferenceMaxIn)
	}

	if f.ReferenceMinLike != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceMinLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceMaxLike != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, strings.Replace(strings.Replace(*f.ReferenceMaxLike, "?", "_", -1), "*", "%", -1))
	}

	if f.ReferenceMinPrefix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceMinPrefix))
	}

	if f.ReferenceMaxPrefix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%s%%", *f.ReferenceMaxPrefix))
	}

	if f.ReferenceMinSuffix != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceMinSuffix))
	}

	if f.ReferenceMaxSuffix != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("reference")+") LIKE ?")
		values = append(values, fmt.Sprintf("%%%s", *f.ReferenceMaxSuffix))
	}

	if f.ReferenceIDMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") = ?")
		values = append(values, f.ReferenceIDMin)
	}

	if f.ReferenceIDMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") = ?")
		values = append(values, f.ReferenceIDMax)
	}

	if f.ReferenceIDMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") != ?")
		values = append(values, f.ReferenceIDMinNe)
	}

	if f.ReferenceIDMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") != ?")
		values = append(values, f.ReferenceIDMaxNe)
	}

	if f.ReferenceIDMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") > ?")
		values = append(values, f.ReferenceIDMinGt)
	}

	if f.ReferenceIDMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") > ?")
		values = append(values, f.ReferenceIDMaxGt)
	}

	if f.ReferenceIDMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") < ?")
		values = append(values, f.ReferenceIDMinLt)
	}

	if f.ReferenceIDMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") < ?")
		values = append(values, f.ReferenceIDMaxLt)
	}

	if f.ReferenceIDMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") >= ?")
		values = append(values, f.ReferenceIDMinGte)
	}

	if f.ReferenceIDMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") >= ?")
		values = append(values, f.ReferenceIDMaxGte)
	}

	if f.ReferenceIDMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") <= ?")
		values = append(values, f.ReferenceIDMinLte)
	}

	if f.ReferenceIDMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") <= ?")
		values = append(values, f.ReferenceIDMaxLte)
	}

	if f.ReferenceIDMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("referenceID")+") IN (?)")
		values = append(values, f.ReferenceIDMinIn)
	}

	if f.ReferenceIDMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("referenceID")+") IN (?)")
		values = append(values, f.ReferenceIDMaxIn)
	}

	if f.UpdatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMin)
	}

	if f.UpdatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") = ?")
		values = append(values, f.UpdatedAtMax)
	}

	if f.UpdatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMinNe)
	}

	if f.UpdatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") != ?")
		values = append(values, f.UpdatedAtMaxNe)
	}

	if f.UpdatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMinGt)
	}

	if f.UpdatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") > ?")
		values = append(values, f.UpdatedAtMaxGt)
	}

	if f.UpdatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMinLt)
	}

	if f.UpdatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") < ?")
		values = append(values, f.UpdatedAtMaxLt)
	}

	if f.UpdatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMinGte)
	}

	if f.UpdatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") >= ?")
		values = append(values, f.UpdatedAtMaxGte)
	}

	if f.UpdatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMinLte)
	}

	if f.UpdatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") <= ?")
		values = append(values, f.UpdatedAtMaxLte)
	}

	if f.UpdatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMinIn)
	}

	if f.UpdatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedAt")+") IN (?)")
		values = append(values, f.UpdatedAtMaxIn)
	}

	if f.CreatedAtMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMin)
	}

	if f.CreatedAtMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") = ?")
		values = append(values, f.CreatedAtMax)
	}

	if f.CreatedAtMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMinNe)
	}

	if f.CreatedAtMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") != ?")
		values = append(values, f.CreatedAtMaxNe)
	}

	if f.CreatedAtMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMinGt)
	}

	if f.CreatedAtMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") > ?")
		values = append(values, f.CreatedAtMaxGt)
	}

	if f.CreatedAtMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMinLt)
	}

	if f.CreatedAtMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") < ?")
		values = append(values, f.CreatedAtMaxLt)
	}

	if f.CreatedAtMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMinGte)
	}

	if f.CreatedAtMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") >= ?")
		values = append(values, f.CreatedAtMaxGte)
	}

	if f.CreatedAtMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMinLte)
	}

	if f.CreatedAtMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") <= ?")
		values = append(values, f.CreatedAtMaxLte)
	}

	if f.CreatedAtMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMinIn)
	}

	if f.CreatedAtMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdAt")+") IN (?)")
		values = append(values, f.CreatedAtMaxIn)
	}

	if f.UpdatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMin)
	}

	if f.UpdatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") = ?")
		values = append(values, f.UpdatedByMax)
	}

	if f.UpdatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMinNe)
	}

	if f.UpdatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") != ?")
		values = append(values, f.UpdatedByMaxNe)
	}

	if f.UpdatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMinGt)
	}

	if f.UpdatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") > ?")
		values = append(values, f.UpdatedByMaxGt)
	}

	if f.UpdatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMinLt)
	}

	if f.UpdatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") < ?")
		values = append(values, f.UpdatedByMaxLt)
	}

	if f.UpdatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMinGte)
	}

	if f.UpdatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") >= ?")
		values = append(values, f.UpdatedByMaxGte)
	}

	if f.UpdatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMinLte)
	}

	if f.UpdatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") <= ?")
		values = append(values, f.UpdatedByMaxLte)
	}

	if f.UpdatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMinIn)
	}

	if f.UpdatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("updatedBy")+") IN (?)")
		values = append(values, f.UpdatedByMaxIn)
	}

	if f.CreatedByMin != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMin)
	}

	if f.CreatedByMax != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") = ?")
		values = append(values, f.CreatedByMax)
	}

	if f.CreatedByMinNe != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMinNe)
	}

	if f.CreatedByMaxNe != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") != ?")
		values = append(values, f.CreatedByMaxNe)
	}

	if f.CreatedByMinGt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMinGt)
	}

	if f.CreatedByMaxGt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") > ?")
		values = append(values, f.CreatedByMaxGt)
	}

	if f.CreatedByMinLt != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMinLt)
	}

	if f.CreatedByMaxLt != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") < ?")
		values = append(values, f.CreatedByMaxLt)
	}

	if f.CreatedByMinGte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMinGte)
	}

	if f.CreatedByMaxGte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") >= ?")
		values = append(values, f.CreatedByMaxGte)
	}

	if f.CreatedByMinLte != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMinLte)
	}

	if f.CreatedByMaxLte != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") <= ?")
		values = append(values, f.CreatedByMaxLte)
	}

	if f.CreatedByMinIn != nil {
		conditions = append(conditions, "Min("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMinIn)
	}

	if f.CreatedByMaxIn != nil {
		conditions = append(conditions, "Max("+aliasPrefix+dialect.Quote("createdBy")+") IN (?)")
		values = append(values, f.CreatedByMaxIn)
	}

	return
}

// AndWith convenience method for combining two or more filters with AND statement
func (f *FileFilterType) AndWith(f2 ...*FileFilterType) *FileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &FileFilterType{
		And: append(_f2, f),
	}
}

// OrWith convenience method for combining two or more filters with OR statement
func (f *FileFilterType) OrWith(f2 ...*FileFilterType) *FileFilterType {
	_f2 := f2[:0]
	for _, x := range f2 {
		if x != nil {
			_f2 = append(_f2, x)
		}
	}
	if len(_f2) == 0 {
		return f
	}
	return &FileFilterType{
		Or: append(_f2, f),
	}
}
