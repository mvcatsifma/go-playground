# godirwalk2

Directory walking with symlink skipping and tiered error recovery.

## What's here

`main.go` — skips symlinks, halts on `ErrFatal`, skips all other erroring nodes.

## Todo

- [ ] Add a file extension filter (e.g. only count `.go` files) and report the total at the end.
- [ ] Pass a `context.Context` through the callback and return an error when it is done — practice wiring cancellation into a walk.
- [ ] Rewrite the same walk with stdlib `fs.WalkDir` and compare: when does godirwalk add value?
- [ ] Collect all visited paths into `[]string` and write a test using a temp directory with known contents.
