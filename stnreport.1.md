%stnreport(1) stnreport user manual
% R. S. Doiel
% August 14, 2022

# NAME

stnreport

# SYNOPSIS

stnreport [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

stnreport takes output from stnparse or stnfilter and renders a
report.

# OPTIONS

-h, -help
: display help

-l, -license
: display license

-v, -version
: display version

-i
: Reading input from a file rather than standard input

-o
: Write output to a file rather than standard output

-format
: Render output as a CSV file or JSON


# EXIT STATUS

0
: success

2
: bad command line, or a bad column list

65
: an input line is not parsed output

66
: input file cannot be opened

73
: output file cannot be created

# EXAMPLES

This renders columns zero (first column) and one.

~~~shell
    stnparse TimeSheet.txt | stnreport -columns 0,1
~~~

This renders columns zero (first column) and one as a CSV file.

~~~shell
    stnparse TimeSheet.txt | stnreport -columns 0,1 -format csv
~~~



