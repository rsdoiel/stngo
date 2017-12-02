
# USAGE

	stnparse [OPTIONS] [TIME_DESCRIPTION]

## SYNOPSIS



SYNOPSIS

stnparse parses content in "Standard Timesheet Notation". By default
it parse them into a tabular format but can also optionally
parse them into a stream of JSON blobs.



## OPTIONS

```
    -examples                 display example(s)
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -i, -input                input filename
    -j, -json                 output JSON format
    -l, -license              display license
    -o, -output               output filename
    -quiet                    suppress error messages
    -v, -version              display version
```


## EXAMPLES



EXAMPLES

	stnparse < TimeSheet.txt

This will parse the TimeSheet.txt file into a table.

	stnparse -json < TimeSheet.txt

This will parse TimeSheet.txt file into a stream of JSON blobs.



stnparse v0.0.6
