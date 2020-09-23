package gen

import (
	"context"

	"github.com/jinzhu/gorm"
)

func (s FileSortType) Apply(ctx context.Context, dialect gorm.Dialect, sorts *[]SortInfo, joins *[]string) error {
	return s.ApplyWithAlias(ctx, dialect, TableName("files"), sorts, joins)
}
func (s FileSortType) ApplyWithAlias(ctx context.Context, dialect gorm.Dialect, alias string, sorts *[]SortInfo, joins *[]string) error {
	aliasPrefix := dialect.Quote(alias) + "."

	if s.ID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("id"), Direction: s.ID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.IDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.IDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("id") + ")", Direction: s.IDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Email != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("email"), Direction: s.Email.String()}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.EmailMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("email") + ")", Direction: s.EmailMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("uid"), Direction: s.UID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("uid") + ")", Direction: s.UIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("uid") + ")", Direction: s.UIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Name != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("name"), Direction: s.Name.String()}
		*sorts = append(*sorts, sort)
	}

	if s.NameMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.NameMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("name") + ")", Direction: s.NameMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Size != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("size"), Direction: s.Size.String()}
		*sorts = append(*sorts, sort)
	}

	if s.SizeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("size") + ")", Direction: s.SizeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SizeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("size") + ")", Direction: s.SizeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.SizeAvg != nil {
		sort := SortInfo{Field: "Avg(" + aliasPrefix + dialect.Quote("size") + ")", Direction: s.SizeAvg.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ContentType != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("contentType"), Direction: s.ContentType.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ContentTypeMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("contentType") + ")", Direction: s.ContentTypeMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ContentTypeMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("contentType") + ")", Direction: s.ContentTypeMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.URL != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("url"), Direction: s.URL.String()}
		*sorts = append(*sorts, sort)
	}

	if s.URLMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("url") + ")", Direction: s.URLMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.URLMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("url") + ")", Direction: s.URLMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.Reference != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("reference"), Direction: s.Reference.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("reference") + ")", Direction: s.ReferenceMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("reference") + ")", Direction: s.ReferenceMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceID != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("referenceID"), Direction: s.ReferenceID.String()}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceIDMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("referenceID") + ")", Direction: s.ReferenceIDMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.ReferenceIDMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("referenceID") + ")", Direction: s.ReferenceIDMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedAt"), Direction: s.UpdatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedAt") + ")", Direction: s.UpdatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAt != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdAt"), Direction: s.CreatedAt.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedAtMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdAt") + ")", Direction: s.CreatedAtMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("updatedBy"), Direction: s.UpdatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.UpdatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("updatedBy") + ")", Direction: s.UpdatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedBy != nil {
		sort := SortInfo{Field: aliasPrefix + dialect.Quote("createdBy"), Direction: s.CreatedBy.String()}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMin != nil {
		sort := SortInfo{Field: "Min(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMin.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	if s.CreatedByMax != nil {
		sort := SortInfo{Field: "Max(" + aliasPrefix + dialect.Quote("createdBy") + ")", Direction: s.CreatedByMax.String(), IsAggregation: true}
		*sorts = append(*sorts, sort)
	}

	return nil
}
