# unitofwork

Unit of Work pattern: domain objects register state changes via a UoW stored in the context.

## What's here

`main.go` — `Album` embeds `doHelper` which proxies `markNew`/`markDirty`/`markDeleted` to the `UnitOfWork` in the context.

## Todo

- [ ] Add `Commit(ctx context.Context)` to `UnitOfWork` that logs what would be inserted, updated, and deleted — make the pattern actually do something.
- [ ] Call `a.setTitle(ctx, "new title")` and verify the album moves from `newObjects` to `dirtyObjects` in a test.
- [ ] Add a `Track` struct; verify the UoW tracks both `Album` and `Track` correctly under the same context.
- [ ] Replace the string context key `"unitOfWork"` with a typed key `type ctxKey struct{}`; understand why typed keys prevent collisions.
