%stnparse(1) user manual | version 0.0.15 9c12c25
% R. S. Doiel
% 2026-10-20

# NAME

stnparse — parse Simple Timesheet Notation

# SYNOPSIS

stnparse [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

stnparse reads a timesheet written in Simple Timesheet Notation and writes one
line for each entry: the start and end time as RFC 3339 timestamps, then the
cells of the entry, separated by tabs. With -json it writes a JSON array
instead. Pipe the result to stnfilter and stnreport.

Times are local time. The hours in a timesheet are the hours on the clock where
it was written, so 8:00 is 8:00 whatever the date, and each entry gets the UTC
offset in force on its own day, summer or winter.

Times are written without am or pm, so stnparse works them out from the order
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
stnparse TimeSheet.txt
~~~

Parse it into JSON:

~~~shell
stnparse -json TimeSheet.txt
~~~

Report the hours for one day:

~~~shell
stnparse TimeSheet.txt | stnfilter -start=2026-10-16 -end=2026-10-16 | stnreport -columns=0,1
~~~

# SEE ALSO

stnfilter(1), stnreport(1), stn(5)

