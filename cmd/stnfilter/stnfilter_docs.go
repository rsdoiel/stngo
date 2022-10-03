package main

import (
	"strings"
)

const (
	helpText = `%{app_name}(1) {app_name} user manual
% R. S. Doiel
% August 14, 2022

# NAME

{app_name}

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

{app_name} is a standard timesheet notation filter.
{app_name} will filter the output from stnparse based on date
or matching text.

# OPTIONS

-i
: Set input filename

-o
: Set output filename

-start
: filter by start date

-end
: filter by end date

-match
: Match a text against entries

-json
: output as JSON

# EXAMPLE

Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

~~~shell
    {app_name} -start 2015-07-04 -end 2015-07-14 -json TimeSheet.tab
~~~

To render the same in a tab delimited output

~~~shell
    {app_name} -start 2015-07-04 -end 2015-07-14 TimeSheet.tab
~~~

Typical usage would be in a pipeline with Unix stnparse

~~~shell
   stnparse Time_Sheet.txt | {app_name} -start 2015-07-06 -end 2015-07-010
~~~

Matching a project name "Fred" for the same week would look like

~~~shell
    stnparse Time_Sheet.txt | {app_name} s -start 2015-07-06 -end 2015-07-010 -match Fred
~~~

`

	licenseText = `
{app_name} {version}

Copyright (c) 2021, R. S. Doiel
All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions are met:

* Redistributions of source code must retain the above copyright notice, this
  list of conditions and the following disclaimer.

* Redistributions in binary form must reproduce the above copyright notice,
  this list of conditions and the following disclaimer in the documentation
  and/or other materials provided with the distribution.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

`
)

func fmtText(txt string, appName string, version string) string {
	return strings.ReplaceAll(strings.ReplaceAll(txt, "{app_name}", appName), "{version}", version)
}
