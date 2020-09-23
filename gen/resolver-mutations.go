package gen

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/novacloudcz/graphql-orm/events"
)

type GeneratedMutationResolver struct{ *GeneratedResolver }

type MutationEvents struct {
	Events []events.Event
}

func EnrichContextWithMutations(ctx context.Context, r *GeneratedResolver) context.Context {
	_ctx := context.WithValue(ctx, KeyMutationTransaction, r.DB.db.Begin())
	_ctx = context.WithValue(_ctx, KeyMutationEvents, &MutationEvents{})
	return _ctx
}
func FinishMutationContext(ctx context.Context, r *GeneratedResolver) (err error) {
	s := GetMutationEventStore(ctx)

	for _, event := range s.Events {
		err = r.Handlers.OnEvent(ctx, r, &event)
		if err != nil {
			return
		}
	}

	tx := r.GetDB(ctx)
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return
	}

	for _, event := range s.Events {
		err = r.EventController.SendEvent(ctx, &event)
	}

	return
}
func RollbackMutationContext(ctx context.Context, r *GeneratedResolver) error {
	tx := r.GetDB(ctx)
	return tx.Rollback().Error
}
func GetMutationEventStore(ctx context.Context) *MutationEvents {
	return ctx.Value(KeyMutationEvents).(*MutationEvents)
}
func AddMutationEvent(ctx context.Context, e events.Event) {
	s := GetMutationEventStore(ctx)
	s.Events = append(s.Events, e)
}

func (r *GeneratedMutationResolver) CreateFile(ctx context.Context, input map[string]interface{}) (item *File, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.CreateFile(ctx, r.GeneratedResolver, input)
	if err != nil {
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func CreateFileHandler(ctx context.Context, r *GeneratedResolver, input map[string]interface{}) (item *File, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	now := time.Now()
	item = &File{ID: uuid.Must(uuid.NewV4()).String(), CreatedAt: now, CreatedBy: principalID}
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeCreated,
		Entity:      "File",
		EntityID:    item.ID,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes FileChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	if _, ok := input["id"]; ok && (item.ID != changes.ID) {
		item.ID = changes.ID
		event.EntityID = item.ID
		event.AddNewValue("id", changes.ID)
	}

	if _, ok := input["uid"]; ok && (item.UID != changes.UID) && (item.UID == nil || changes.UID == nil || *item.UID != *changes.UID) {
		item.UID = changes.UID

		event.AddNewValue("uid", changes.UID)
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		item.Name = changes.Name

		event.AddNewValue("name", changes.Name)
	}

	if _, ok := input["size"]; ok && (item.Size != changes.Size) && (item.Size == nil || changes.Size == nil || *item.Size != *changes.Size) {
		item.Size = changes.Size

		event.AddNewValue("size", changes.Size)
	}

	if _, ok := input["contentType"]; ok && (item.ContentType != changes.ContentType) && (item.ContentType == nil || changes.ContentType == nil || *item.ContentType != *changes.ContentType) {
		item.ContentType = changes.ContentType

		event.AddNewValue("contentType", changes.ContentType)
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) && (item.Reference == nil || changes.Reference == nil || *item.Reference != *changes.Reference) {
		item.Reference = changes.Reference

		event.AddNewValue("reference", changes.Reference)
	}

	err = tx.Create(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}
func (r *GeneratedMutationResolver) UpdateFile(ctx context.Context, id string, input map[string]interface{}) (item *File, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.UpdateFile(ctx, r.GeneratedResolver, id, input)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func UpdateFileHandler(ctx context.Context, r *GeneratedResolver, id string, input map[string]interface{}) (item *File, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &File{}
	now := time.Now()
	tx := r.GetDB(ctx)

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeUpdated,
		Entity:      "File",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	var changes FileChanges
	err = ApplyChanges(input, &changes)
	if err != nil {
		tx.Rollback()
		return
	}

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	item.UpdatedBy = principalID

	if _, ok := input["uid"]; ok && (item.UID != changes.UID) && (item.UID == nil || changes.UID == nil || *item.UID != *changes.UID) {
		event.AddOldValue("uid", item.UID)
		event.AddNewValue("uid", changes.UID)
		item.UID = changes.UID
	}

	if _, ok := input["name"]; ok && (item.Name != changes.Name) && (item.Name == nil || changes.Name == nil || *item.Name != *changes.Name) {
		event.AddOldValue("name", item.Name)
		event.AddNewValue("name", changes.Name)
		item.Name = changes.Name
	}

	if _, ok := input["size"]; ok && (item.Size != changes.Size) && (item.Size == nil || changes.Size == nil || *item.Size != *changes.Size) {
		event.AddOldValue("size", item.Size)
		event.AddNewValue("size", changes.Size)
		item.Size = changes.Size
	}

	if _, ok := input["contentType"]; ok && (item.ContentType != changes.ContentType) && (item.ContentType == nil || changes.ContentType == nil || *item.ContentType != *changes.ContentType) {
		event.AddOldValue("contentType", item.ContentType)
		event.AddNewValue("contentType", changes.ContentType)
		item.ContentType = changes.ContentType
	}

	if _, ok := input["reference"]; ok && (item.Reference != changes.Reference) && (item.Reference == nil || changes.Reference == nil || *item.Reference != *changes.Reference) {
		event.AddOldValue("reference", item.Reference)
		event.AddNewValue("reference", changes.Reference)
		item.Reference = changes.Reference
	}

	err = tx.Save(item).Error
	if err != nil {
		tx.Rollback()
		return
	}

	if len(event.Changes) > 0 {
		AddMutationEvent(ctx, event)
	}

	return
}
func (r *GeneratedMutationResolver) DeleteFile(ctx context.Context, id string) (item *File, err error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	item, err = r.Handlers.DeleteFile(ctx, r.GeneratedResolver, id)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return
}
func DeleteFileHandler(ctx context.Context, r *GeneratedResolver, id string) (item *File, err error) {
	principalID := GetPrincipalIDFromContext(ctx)
	item = &File{}
	now := time.Now()
	tx := r.GetDB(ctx)

	err = GetItem(ctx, tx, item, &id)
	if err != nil {
		tx.Rollback()
		return
	}

	event := events.NewEvent(events.EventMetadata{
		Type:        events.EventTypeDeleted,
		Entity:      "File",
		EntityID:    id,
		Date:        now,
		PrincipalID: principalID,
	})

	err = tx.Delete(item, TableName("files")+".id = ?", id).Error
	if err != nil {
		tx.Rollback()
		return
	}

	AddMutationEvent(ctx, event)

	return
}
func (r *GeneratedMutationResolver) DeleteAllFiles(ctx context.Context) (bool, error) {
	ctx = EnrichContextWithMutations(ctx, r.GeneratedResolver)
	done, err := r.Handlers.DeleteAllFiles(ctx, r.GeneratedResolver)
	if err != nil {
		RollbackMutationContext(ctx, r.GeneratedResolver)
		return done, err
	}
	err = FinishMutationContext(ctx, r.GeneratedResolver)
	return done, err
}
func DeleteAllFilesHandler(ctx context.Context, r *GeneratedResolver) (bool, error) {
	tx := r.GetDB(ctx)
	err := tx.Delete(&File{}).Error
	if err != nil {
		tx.Rollback()
		return false, err
	}
	return true, err
}
