
# USAGE

    stnfilter [OPTIONS]

# SYNOPSIS

stnfilter will filter the output from stnparse based on date or matching text.

```
    -e       end of inclusive date range
    -end     end of inclusive date range
    -h       display help
    -help    display help
    -j       output in JSON format
    -json    output in JSON format
    -l       display license
    -license display license
    -m       match text in annotations
    -match   match text in annotations
    -s       start of inclusive date range
    -start   start of inclusive date range
    -v       display version
    -version display version
```

## EXAMPLES

Filter TimeSheet.tab from July 4, 2015 through July 14, 2015
and render a stream of JSON blobs.

```
    stnfilter -start 2015-07-04 -end 2015-07-14 -json < TimeSheet.tab
```

To render the same in a tab delimited output

```
    stnfilter -start 2015-07-04 -end 2015-07-14 < TimeSheet.tab
```

Typical usage would be in a pipeline with Unix cat and stnparse

```
   cat Time_Sheet.txt | stnparse | stnfilter -start 2015-07-06 -end 2015-07-010
```

Matching a project name "Fred" for the same week would look like

```
    cat Time_Sheet.txt | stnparse | stnfilter -start 2015-07-06 -end 2015-07-010 -match Fred
```

