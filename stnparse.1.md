%stnparse(1) %stnparse user manual
% R. S. Doiel
% Auguest 14, 2022

# NAME

stnparse

# SYNOPSIS

stnparse [OPTIONS] [INPUT_FILENAME] [OUTPUT_FILENAME]

# DESCRIPTION
	
stnparse parses content in "Standard Timesheet Notation". By default
it parse them into a tabular format but can also optionally
parse them into a stream of JSON blobs.

# OPTIONS

-h, -help
: display help

-l, -license
: display license

-v, -version
: display version

# EXIT STATUS

0
: success

2
: bad command line

65
: an entry line could not be parsed (it is reported and skipped, the rest are still parsed)

66
: input file cannot be opened

73
: output file cannot be created

# EXAMPLES

This will parse the TimeSheet.txt file into a table.

~~~shell
	stnparse TimeSheet.txt
~~~

This will parse TimeSheet.txt file into a stream of JSON blobs.

~~~shell
	stnparse -json TimeSheet.txt
~~~


