% stnparse(1) stngo user manual
% R. S. Doiel
% August, 3, 2022


# NAME

stnparse

# SYNOPSIS

stnparse [OPTIONS] [TIME_DESCRIPTION]

# DESCRIPTION

stnparse parses content in "Standard Timesheet Notation". By default
it parse them into a tabular format but can also optionally
parse them into a stream of JSON blobs.

# OPTIONS

-examples
: display example(s)

-generate-markdown-docs
: generate markdown documentation

-h, -help
: display help

-i, -input
: input filename

-j, -json
: output JSON format

-l, -license
: display license

-o, -output
: output filename

-quiet
: suppress error messages

-v, -version
: display version

# EXAMPLES

This will parse the TimeSheet.txt file into a table.

```shell
	stnparse < TimeSheet.txt
```

This will parse TimeSheet.txt file into a stream of JSON blobs.

```shell
	stnparse -json < TimeSheet.txt
```

# ALSO SEE

- [simple timesheet notation](stn.html)
- [stnfilter](stnfilter.html)
- [stnreport](stnreport.html)
- Website: [https://rsdoiel.github.io/stngo](https://rsdoiel.github.io/stngo)

