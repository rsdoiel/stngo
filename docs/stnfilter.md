%stnfilter(1) stnfilter user manual
% R. S. Doiel
% August 14, 2022

# NAME

stnfilter

# SYNOPSIS

stnfilter [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

stnfilter is a standard timesheet notation filter.
stnfilter will filter the output from stnparse based on date
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
    stnfilter -start 2015-07-04 -end 2015-07-14 -json TimeSheet.tab
~~~

To render the same in a tab delimited output

~~~shell
    stnfilter -start 2015-07-04 -end 2015-07-14 TimeSheet.tab
~~~

Typical usage would be in a pipeline with Unix stnparse

~~~shell
   stnparse Time_Sheet.txt | stnfilter -start 2015-07-06 -end 2015-07-010
~~~

Matching a project name "Fred" for the same week would look like

~~~shell
    stnparse Time_Sheet.txt | stnfilter s -start 2015-07-06 -end 2015-07-010 -match Fred
~~~


