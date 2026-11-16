
Action items
============

Next
----

- [ ] Extend parsing to support ISO 8601 date/time standard this would increase flexibility in notating a log file

Bugs
----

None known.

## Someday Maybe

+ Output iCal/RSS feed of content for a given time range
+ Added Excel .xlsx import/export
+ Add JavaScript support


## Modernization (2026-11-16)

To match the workspace standard (knowledge and opml are the models; Go 1.27.3 across repositories). Do one repository at a time; run the tests and `make -n release` after each step.

- [x] Go 1.27.3: change `go 1.26.4` in go.mod, the `Go >=` lines in codemeta.json, README.md and INSTALL.md; run `go mod tidy`, `go vet ./...` and the tests.
- [x] Makefile: add `check-go-version` (go.mod, codemeta.json and the Go in use agree); `check-release` calls it.
- [x] Drop the armv7l target and 32-bit Pi wording (dropped 2026-10-27).
- [x] release.ps1: take knowledge's (existing-draft check, checksums).
