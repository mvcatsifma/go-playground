package main

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"os"
)

func main() {
	ctx := context.Background()
	ctx, u := NewUnitOfWork(ctx)
	a := NewAlbum(ctx, "Living on the edge")
	if len(u.newObjects) != 1 {
		fmt.Printf("ERROR: expected len(u.newObjects) to be %d, was %d\n", 1, len(u.newObjects))
		os.Exit(1)
	}
	if !cmp.Equal(a, u.newObjects[0], cmpopts.IgnoreUnexported(Album{})) {
		fmt.Printf("ERROR: expected objects to be equal\n")
		os.Exit(1)
	}
}

type Album struct {
	doHelper
	Name  string
	Title string
}

func (a *Album) GetID() string {
	return ""
}

func NewAlbum(ctx context.Context, name string) *Album {
	a := &Album{Name: name}
	a.markNew(ctx, a)
	return a
}

func (a *Album) setTitle(ctx context.Context, title string) {
	a.Title = title
	a.markDirty(ctx, a)
}

type DomainObject interface {
	GetID() string
}

// doHelper is embedded in domain objects to give them access to the UnitOfWork via context.
type doHelper struct {
}

func (d *doHelper) markDeleted(ctx context.Context, o DomainObject) {
	GetCurrentUnitOfWork(ctx).registerDeleted(o)
}

func (d *doHelper) markNew(ctx context.Context, o DomainObject) {
	GetCurrentUnitOfWork(ctx).registerNew(o)
}

func (d *doHelper) markDirty(ctx context.Context, o DomainObject) {
	GetCurrentUnitOfWork(ctx).registerDirty(o)
}

const ctxKeyUnitOfWork = "unitOfWork"

// NewUnitOfWork creates a UnitOfWork and stores it in the context; typically called in an HTTP filter.
func NewUnitOfWork(ctx context.Context) (context.Context, *UnitOfWork) {
	uow := &UnitOfWork{}
	ctx = context.WithValue(ctx, ctxKeyUnitOfWork, uow)
	return ctx, uow
}

// GetCurrentUnitOfWork retrieves the active UnitOfWork from the context; panics if none was set.
func GetCurrentUnitOfWork(ctx context.Context) *UnitOfWork {
	return ctx.Value(ctxKeyUnitOfWork).(*UnitOfWork)
}

type UnitOfWork struct {
	newObjects     []DomainObject
	dirtyObjects   []DomainObject
	deletedObjects []DomainObject
}

func (u *UnitOfWork) registerDeleted(d DomainObject) {
	u.deletedObjects = append(u.deletedObjects, d)
}

func (u *UnitOfWork) registerNew(d DomainObject) {
	u.newObjects = append(u.newObjects, d)
}

func (u *UnitOfWork) registerDirty(d DomainObject) {
	u.dirtyObjects = append(u.dirtyObjects, d)
}
