
# USAGE

	stnfilter [OPTIONS]

## SYNOPSIS



SYNOPSIS

stnfilter will filter the output from stnparse based on date or matching text.



## OPTIONS

```
    -e, -end                  end of inclusive date range
    -examples                 display examples(s)
    -generate-markdown-docs   generate markdown documentation
    -h, -help                 display help
    -i, -input                input file name
    -j, -json                 output JSON format
    -l, -license              display license
    -m, -match                Match text annotations
    -o, -output               output file name
    -quiet                    suppress error message
    -s, -start                start of inclusive date range
    -v, -version              display version
```


## EXAMPLES



EXAMPLES

Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

    stnfilter -start 2015-07-04 -end 2015-07-14 -json < TimeSheet.tab

To render the same in a tab delimited output

    stnfilter -start 2015-07-04 -end 2015-07-14 < TimeSheet.tab

Typical usage would be in a pipeline with Unix cat and stnparse

   cat Time_Sheet.txt | stnparse | stnfilter -start 2015-07-06 -end 2015-07-010

Matching a project name "Fred" for the same week would look like

    cat Time_Sheet.txt | stnparse | stnfilter -start 2015-07-06 -end 2015-07-010 -match Fred



