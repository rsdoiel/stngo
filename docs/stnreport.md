%stnreport(1) stnreport user manual
% R. S. Doiel
% August 14, 2022

# Name

stnreport

# SYNOPSIS

stnreport [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION

stnreport takes output from stnparse or stnfilter and renders a
report.

# OPTIONS

-i
: Reading input from a file rather than standard input

-o
: Write output to a file rather than standard output

-format
: Render output as a CSV file or JSON


# EXAMPLES

This renders columns zero (first column) and one.

~~~shell
    stnparse -i TimeSheet.txt | stnreport -columns 0,1
~~~

This renders columns zero (first column) and one as a CSV file.

~~~shell
    stnparse -i TimeSheet.txt | stnreport -columns 0,1 -format csv
~~~



