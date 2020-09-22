package gen

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/ast"
)

type GeneratedQueryResolver struct{ *GeneratedResolver }

type QueryFileHandlerOptions struct {
	ID     *string
	Q      *string
	Filter *FileFilterType
}

func (r *GeneratedQueryResolver) File(ctx context.Context, id *string, q *string, filter *FileFilterType) (*File, error) {
	opts := QueryFileHandlerOptions{
		ID:     id,
		Q:      q,
		Filter: filter,
	}
	return r.Handlers.QueryFile(ctx, r.GeneratedResolver, opts)
}
func QueryFileHandler(ctx context.Context, r *GeneratedResolver, opts QueryFileHandlerOptions) (*File, error) {
	selection := []ast.Selection{}
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		selection = append(selection, f.Field)
	}
	selectionSet := ast.SelectionSet(selection)

	query := FileQueryFilter{opts.Q}
	offset := 0
	limit := 1
	rt := &FileResultType{
		EntityResultType: EntityResultType{
			Offset:       &offset,
			Limit:        &limit,
			Query:        &query,
			Filter:       opts.Filter,
			SelectionSet: &selectionSet,
		},
	}
	qb := r.GetDB(ctx)
	if qb == nil {
		qb = r.DB.Query()
	}
	if opts.ID != nil {
		qb = qb.Where(TableName("files")+".id = ?", *opts.ID)
	}

	var items []*File
	giOpts := GetItemsOptions{
		Alias:      TableName("files"),
		Preloaders: []string{},
	}
	err := rt.GetItems(ctx, qb, giOpts, &items)
	if err != nil {
		return nil, err
	}
	if len(items) == 0 {
		return nil, nil
	}
	return items[0], err
}

type QueryFilesHandlerOptions struct {
	Offset *int
	Limit  *int
	Q      *string
	Sort   []*FileSortType
	Filter *FileFilterType
}

func (r *GeneratedQueryResolver) Files(ctx context.Context, offset *int, limit *int, q *string, sort []*FileSortType, filter *FileFilterType) (*FileResultType, error) {
	opts := QueryFilesHandlerOptions{
		Offset: offset,
		Limit:  limit,
		Q:      q,
		Sort:   sort,
		Filter: filter,
	}
	return r.Handlers.QueryFiles(ctx, r.GeneratedResolver, opts)
}
func QueryFilesHandler(ctx context.Context, r *GeneratedResolver, opts QueryFilesHandlerOptions) (*FileResultType, error) {
	query := FileQueryFilter{opts.Q}

	var selectionSet *ast.SelectionSet
	for _, f := range graphql.CollectFieldsCtx(ctx, nil) {
		if f.Field.Name == "items" {
			selectionSet = &f.Field.SelectionSet
		}
	}

	_sort := []EntitySort{}
	for _, sort := range opts.Sort {
		_sort = append(_sort, sort)
	}

	return &FileResultType{
		EntityResultType: EntityResultType{
			Offset:       opts.Offset,
			Limit:        opts.Limit,
			Query:        &query,
			Sort:         _sort,
			Filter:       opts.Filter,
			SelectionSet: selectionSet,
		},
	}, nil
}

type GeneratedFileResultTypeResolver struct{ *GeneratedResolver }

func (r *GeneratedFileResultTypeResolver) Items(ctx context.Context, obj *FileResultType) (items []*File, err error) {
	otps := GetItemsOptions{
		Alias:      TableName("files"),
		Preloaders: []string{},
	}
	err = obj.GetItems(ctx, r.DB.db, otps, &items)

	uniqueItems := []*File{}
	idMap := map[string]bool{}
	for _, item := range items {
		if _, ok := idMap[item.ID]; !ok {
			idMap[item.ID] = true
			uniqueItems = append(uniqueItems, item)
		}
	}
	items = uniqueItems
	return
}

func (r *GeneratedFileResultTypeResolver) Count(ctx context.Context, obj *FileResultType) (count int, err error) {
	opts := GetItemsOptions{
		Alias:      TableName("files"),
		Preloaders: []string{},
	}
	return obj.GetCount(ctx, r.DB.db, opts, &File{})
}
