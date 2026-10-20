%stnreport(1) user manual | version 0.0.15 82f8419
% R. S. Doiel
% 2026-10-20

# NAME

stnreport — total the hours in the output of stnparse or stnfilter

# SYNOPSIS

stnreport [OPTIONS] [INPUT_FILENAME [OUTPUT_FILENAME]]

# DESCRIPTION

stnreport reads the lines written by stnparse or stnfilter, adds up the hours
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
stnparse TimeSheet.txt | stnfilter -start=2026-10-16 -end=2026-10-16 | stnreport -columns=0,1
~~~

The same as CSV:

~~~shell
stnparse TimeSheet.txt | stnreport -columns 0,1 -format csv
~~~

# SEE ALSO

stnparse(1), stnfilter(1), stn(5)

