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

Brought up to current workspace practice. The Go package is now named `stngo`, matching its import path and the repository, so code that imports `github.com/rsdoiel/stngo` without an alias must now refer to it as `stngo` (or alias it as `stn`). The generated installers now download from `rsdoiel/stngo` instead of the non-existent `rsdoiel/stn`. `stnparse`, `stnfilter` and `stnreport` accept `-h`, `-l` and `-v` beside the long options and exit with the workspace codes: 2 for a bad command line or date, 65 for an unparseable input line, 66 for a missing input file, 73 for an output file that cannot be created. Requires Go 1.26.4. The Makefile gains `generate` and `check-release` targets, and the publish and release scripts stop at the first failed step.

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


