package cli

/** ParseHelpText is the stnparse(1) page, shown by `stnparse -help`. The
 * {app_name}, {version}, {release_hash} and {release_date} tokens are
 * replaced by stngo.FmtHelp. The Makefile writes the page to stnparse.1.md.
 *
 * Example:
 *   fmt.Println(stngo.FmtHelp(cli.ParseHelpText, "stnparse", stngo.Version, stngo.ReleaseDate, stngo.ReleaseHash))
 */
const ParseHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — parse Simple Timesheet Notation

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads a timesheet written in Simple Timesheet Notation and writes one
line for each entry: the start and end time as RFC 3339 timestamps, then the
cells of the entry, separated by tabs. With -json it writes a JSON array
instead. Pipe the result to stnfilter and stnreport.

Times are local time. The hours in a timesheet are the hours on the clock where
it was written, so 8:00 is 8:00 whatever the date, and each entry gets the UTC
offset in force on its own day, summer or winter.

Times are written without am or pm, so {app_name} works them out from the order
of the day. The entries of a day are in order, so their start times ascend: if a
start such as 1:00 would be earlier than the entry before it, it is 13:00. After
an entry that started at 8:00, "1:00 - 4:00" is 13:00 to 16:00. The first entry of
a day, and a start that is not earlier than the one before it, is taken as
written, so two entries that both start at 10:00 overlap. An end time earlier
than the start is the afternoon, so "8:00 - 4:00" is 8:00 to 16:00 and
"8:00 - 1:00" is 8:00 to 13:00. 24:00 is midnight at the end of the day. A time
of 12:00 or later is never moved.

Lines that are not entries or dates (headings, notes, blank lines) are ignored.
A line that starts like an entry, a time and a dash, but is not one is
skipped with a warning on standard error, because its hours would otherwise go
uncounted. The usual causes are a missing semicolon after the end time, a
colon where the semicolon belongs and an extra space.

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
: read the timesheet from FILENAME

-o FILENAME
: write the result to FILENAME

-json
: write a JSON array of entries instead of tab delimited lines

# EXIT STATUS

0
: success, including when lines were skipped with a warning

2
: bad command line

65
: an entry line could not be parsed, for example an hour of 25. It is reported
  and skipped, the rest of the timesheet is parsed, and the exit status is 65

66
: the input file cannot be opened

73
: the output file cannot be created

74
: reading or writing failed part way

# EXAMPLES

Parse a timesheet into a table:

~~~shell
{app_name} TimeSheet.txt
~~~

Parse it into JSON:

~~~shell
{app_name} -json TimeSheet.txt
~~~

Report the hours for one day:

~~~shell
{app_name} TimeSheet.txt | stnfilter -start=2026-10-16 -end=2026-10-16 | stnreport -columns=0,1
~~~

# SEE ALSO

stnfilter(1), stnreport(1), stn(5)
`

/** FilterHelpText is the stnfilter(1) page, shown by `stnfilter -help`.
 *
 * Example:
 *   fmt.Println(stngo.FmtHelp(cli.FilterHelpText, "stnfilter", stngo.Version, stngo.ReleaseDate, stngo.ReleaseHash))
 */
const FilterHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — filter the output of stnparse by date or text

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads the lines written by stnparse and writes the ones that match.
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
stnparse TimeSheet.txt | {app_name} -start=2026-10-16 -end=2026-10-16
~~~

A week, Sunday through Saturday, mentioning one project:

~~~shell
stnparse TimeSheet.txt | {app_name} -start 2026-10-11 -end 2026-10-17 -match Widgets
~~~

# SEE ALSO

stnparse(1), stnreport(1), stn(5)
`

/** ReportHelpText is the stnreport(1) page, shown by `stnreport -help`.
 *
 * Example:
 *   fmt.Println(stngo.FmtHelp(cli.ReportHelpText, "stnreport", stngo.Version, stngo.ReleaseDate, stngo.ReleaseHash))
 */
const ReportHelpText = `%{app_name}(1) user manual | version {version} {release_hash}
% R. S. Doiel
% {release_date}

# NAME

{app_name} — total the hours in the output of stnparse or stnfilter

# SYNOPSIS

{app_name} [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

{app_name} reads the lines written by stnparse or stnfilter, adds up the hours
for each distinct set of cells and writes a report with the total. The cells are
numbered from zero, and -columns chooses which ones identify a row. Rows are
sorted by their cells.

By default the report is a text table. -format csv and -format json write those
formats, and -itemize writes one row for each entry rather than one for each
set of cells; in the CSV and JSON forms the duration is in fractional hours.

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

-columns LIST
: comma separated, zero indexed cells that identify a row (default 0)

-format NAME
: text (the default), csv or json

-itemize
: report each entry rather than totals

# EXIT STATUS

0
: success

2
: bad command line, or a -columns list that is not numbers

65
: an input line is not output of stnparse

66
: the input file cannot be opened

73
: the output file cannot be created

74
: reading or writing failed part way

# EXAMPLES

Hours by client and activity for one day:

~~~shell
stnparse TimeSheet.txt | stnfilter -start=2026-10-16 -end=2026-10-16 | {app_name} -columns=0,1
~~~

The same as CSV:

~~~shell
stnparse TimeSheet.txt | {app_name} -columns 0,1 -format csv
~~~

# SEE ALSO

stnparse(1), stnfilter(1), stn(5)
`
