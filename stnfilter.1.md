%stnfilter(1) user manual | version 0.0.15 9c12c25
% R. S. Doiel
% 2026-10-20

# NAME

stnfilter — filter the output of stnparse by date or text

# SYNOPSIS

stnfilter [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

stnfilter reads the lines written by stnparse and writes the ones that match.
An entry is kept when it starts inside the date range, and, if -match is
given, when any of its cells contains the text. Lines are passed through
unchanged.

A date range is whole days in local time, the zone the timesheet was written
in. -start 2026-10-12 -end 2026-10-12 is the day from 00:00 to 24:00 on the
wall clock, whatever the time zone offset was that day. If only -start is given
the range runs to the present moment, and if only -end is given it runs from
the beginning.

The input and output are standard input and output unless a file is named. A
file name of - means standard input or output.

# OPTIONS

-h, -help
: display this help page

-l, -license
: display the license

-v, -version
: display the program name, version and release hash

-i FILENAME
: read from FILENAME

-o FILENAME
: write to FILENAME

-start DATE
: first day to keep, as YYYY-MM-DD

-end DATE
: last day to keep, as YYYY-MM-DD, inclusive

-match TEXT
: keep entries with TEXT in any cell

-json
: accepted and ignored; the lines are passed through unchanged

# EXIT STATUS

0
: success

2
: bad command line, or a -start or -end that is not a date

65
: an input line is not output of stnparse

66
: the input file cannot be opened

73
: the output file cannot be created

74
: reading or writing failed part way

# EXAMPLES

One day:

~~~shell
stnparse TimeSheet.txt | stnfilter -start=2026-10-16 -end=2026-10-16
~~~

A week, Sunday through Saturday, mentioning one project:

~~~shell
stnparse TimeSheet.txt | stnfilter -start 2026-10-11 -end 2026-10-17 -match Widgets
~~~

# SEE ALSO

stnparse(1), stnreport(1), stn(5)

