
# USAGE

	stnreport [OPTIONS]

## SYNOPSIS



SYNOPSIS

stnreport takes output from stnparse or stnfilter and renders a
report.



## OPTIONS

```
    -c, -columns              a comma delimited List of zero indexed columns to report
    -examples                 display example(s)
    -format                   sets output format, text, csv or JSON
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -i, -input                input filename
    -l, -license              display license
    -o, -output               output filename
    -quiet                    suppress error messages
    -v, -version              display version
```


## EXAMPLES


EXAMPLE

    stnparse -i TimeSheet.txt | stnreport -columns 0,1

This renders columns zero (first column) and one.

    stnparse -i TimeSheet.txt | stnreport -columns 0,1 -format csv

This renders columns zero (first column) and one as a CSV file.


