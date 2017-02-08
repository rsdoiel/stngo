
# USAGE

    stnparse [OPTIONS] [TIME_DESCRIPTION]

## SYNOPSIS

stnparse parses content in "Standard Timesheet Notation". By default
it parse them into a tabular format but can also optionally
parse them into a stream of JSON blobs.

```
    -h       display help
    -help    display help
    -i       input filename
    -input   input filename
    -json    output as JSON format
    -l       display license
    -license display license
    -o       output filename
    -output  output filename
    -v       display version
    -version display version
```

## EXAMPLES

```
    stnparse < TimeSheet.txt
```

This will parse the TimeSheet.txt file into a table.

```
    stnparse -json < TimeSheet.txt
```

This will parse TimeSheet.txt file into a stream of JSON blobs.

