package gen

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/novacloudcz/graphql-orm/events"
)

type ResolutionHandlers struct {
	OnEvent func(ctx context.Context, r *GeneratedResolver, e *events.Event) error

	CreateFile     func(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *File, err error)
	UpdateFile     func(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *File, err error)
	DeleteFile     func(ctx context.Context, r *GeneratedResolver, id string) (item *File, err error)
	DeleteAllFiles func(ctx context.Context, r *GeneratedResolver) (bool, error)
	QueryFile      func(ctx context.Context, r *GeneratedResolver, opts QueryFileHandlerOptions) (*File, error)
	QueryFiles     func(ctx context.Context, r *GeneratedResolver, opts QueryFilesHandlerOptions) (*FileResultType, error)
}

func DefaultResolutionHandlers() ResolutionHandlers {
	handlers := ResolutionHandlers{
		OnEvent: func(ctx context.Context, r *GeneratedResolver, e *events.Event) error { return nil },

		CreateFile:     CreateFileHandler,
		UpdateFile:     UpdateFileHandler,
		DeleteFile:     DeleteFileHandler,
		DeleteAllFiles: DeleteAllFilesHandler,
		QueryFile:      QueryFileHandler,
		QueryFiles:     QueryFilesHandler,
	}
	return handlers
}

type GeneratedResolver struct {
	Handlers        ResolutionHandlers
	DB              *DB
	EventController *EventController
}

// GetDB returns database connection or transaction for given context (if exists)
func (r *GeneratedResolver) GetDB(ctx context.Context) *gorm.DB {
	db, _ := ctx.Value(KeyMutationTransaction).(*gorm.DB)
	if db == nil {
		db = r.DB.Query()
	}
	return db
}
