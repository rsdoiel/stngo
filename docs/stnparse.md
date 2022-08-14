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

# EXAMPLES

This will parse the TimeSheet.txt file into a table.

~~~shell
	stnparse < TimeSheet.txt
~~~

This will parse TimeSheet.txt file into a stream of JSON blobs.

~~~shell
	stnparse -json < TimeSheet.txt
~~~


