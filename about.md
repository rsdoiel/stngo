---
title: stngo
abstract: |-
  A set of command line tools and Go package for working with Standard Timesheet Notation.
authors:
  - family_name: Doiel
    given_name: R. S.
    id: https://orcid.org/0000-0003-0900-6903



repository_code: https://github.com/rsdoiel/stngo
version: 0.0.15
license_url: https://opensource.org/licenses/BSD-2-Clause

programming_language:
  - Go >= 1.26.4

keywords:
  - text markup
  - golang
  - timesheet

date_released: 2026-10-20
---

About this software
===================

## stngo 0.0.15

Maintenance release. Brought up to current workspace practice. Times are now handled as local time, the clock the timesheet was written on: an entry gets the UTC offset in force on its own date, so winter entries are no longer shifted an hour when the tools run in daylight time (a log line of 8:00 printed as 07:00), and the `stnfilter` day range is whole local days, so an entry after 5pm is no longer reported on the next day. `stnparse` warns on standard error about a line that starts like an entry but has a typo (a missing or wrong semicolon, an extra space) instead of silently dropping its hours. `stnparse` now works out morning and afternoon start times from the order of the day: the start times of a day ascend, so `1:00 - 4:00` after an entry that started at 8:00 is 13:00 to 16:00 (on the author's logs this corrects about 28% of the displayed start times and changes no duration or report), while `8:00 - 4:00` is still 8:00 to 16:00. Fixed: the `stnparse` segfault on a `24:00` end time (`24:00` is now midnight at the end of the day, and a line that cannot be parsed is reported and skipped, with exit 65 at the end); an entry with no notes left a newline in its last cell, which printed a stray blank line in `stnreport` and split one activity into two rows; `stnreport -itemize` dropped the first entry; a last line with no newline was dropped; `stnfilter -end` without `-start` was ignored. The Go package is now named `stngo`, matching its import path, so code that imports `github.com/rsdoiel/stngo` without an alias must refer to it as `stngo` (or alias it as `stn`), and the generated installers now download from `rsdoiel/stngo` instead of the non-existent `rsdoiel/stn`. The programs take `-h`, `-l` and `-v` beside the long options, exit with the workspace codes (2 usage, 65 data, 66 missing input, 73 cannot create output, 74 I/O), and their command line code moved to `internal/cli` with help text as constants in `helptext.go` that now document every option. Snap packaging is dropped; use the curl installer. Requires Go 1.26.4.

## Authors

- [R. S. Doiel](https://orcid.org/0000-0003-0900-6903)






A set of command line tools and Go package for working with Standard Timesheet Notation.

- [License](https://opensource.org/licenses/BSD-2-Clause)
- [Code Repository](https://github.com/rsdoiel/stngo)
  - [Issue Tracker](https://github.com/rsdoiel/stngo/issues)

## Programming languages

- Go >= 1.26.4




## Software Requirements

- Go >= 1.26.4


## Software Suggestions

- CMTools >= 0.0.45b
- Pandoc >= 3.9
- GNU Make >= 3


