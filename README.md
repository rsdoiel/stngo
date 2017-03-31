
# stngo

Golang implementation of Simple Timesheet Notation plus dome additional utilities and go packages.

Main Simple Timesheet Notation utilities:

+ *stnparse* - translates a standard input and output turning Simple Timesheet Notation into a tab delimited table with RFC3339 dates or JSON blobs.
+ *stnfilter* - filters the output of *stnparse* by dates or text string
+ *stnreport* - summarizes the tab delimited output of *stnfilter* or *stnparse* yielding a simple table showing hours and first annotations

Helpful extra utilities:

+ [shorthand](https://github.com/rsdoiel/shorthand) - a simple label expander for text strings, file inclusion and simple Bash statements. Works with standard input and output.
+ [reldate](https://github.com/rsdoiel/reldate) - a utility to generate relative dates in the YYYY-MM-DD format expected by Simple Timesheet Notation. It easy to process and generate reports with Bash and common Unix utilities (e.g. *date*, *sed*, *tr*)
+ [timefmt](https://github.com/rsdoiel/timefmt) - Golang style date/time formatting

For details of Simple Timesheet Notation markup see [stn.md](stn.html).

For details on using *shorthand* with *stnparse* or generate HTML in
reports see [shorthand](http://rsdoiel.github.io/shorthand).


## Examples of using these utilities in a Unix pipeline

In this example we are filtering entries for a specific date.

### Report durations of activities by day

```shell
    #!/bin/bash

    # Get today in YYYY-MM-DD format
    DAY=$(date +"%Y-%m-%d")
    # If you normally use 12hr notation then use %I:%M otherwise for 23hr format use %H:M
    NOW=$(date +%I:%M)

    if [ "$1" = "" ]; then
        echo "USAGE: rpt-time-by-date.sh YYYY-MM-DD"
    else
        DAY="$1"
    fi

    # Now that we have date in the format needed, create a pipeline for the report.
    cat Time_Sheet.txt | shorthand -e "@now := $NOW" | stnparse |\
        stnfilter -start="$DAY" -end="$DAY" | stnreport
```

### Report durations of activities by week

In this example we use the *reldate* utility from this package to capture the start and end of the work week.

```shell
    #!/bin/bash
    #
    # Report time for current week of the requested week starting with $1.
    #
    RELDATE=$(which reldate)
    FOR_DATE=$(date +"%Y-%m-%d")
    CUR_WEEK_DAY=$(date +%u)

    # If you normally use 12hr notation then use %I:%M otherwise for 23hr format use %H:M
    NOW=$(date +%I:%M)

    # Make sure we have reldate command available.
    if [ "$RELDATE" = "" ]; then
        echo "Missing reldate command. See https://github.com/rsdoiel/reldate"
        exit 1
    fi

    # See if we are asking for help
    if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
        echo "USAGE: rpt-time-by-week.sh YYYY-MM-DD"
        echo "    Without a date it reports the current week."
        exit 1
    elif [ "$1" != "" ]; then
        FOR_DATE=$(reldate --from=$1 0 days)
    fi

    START_WEEK=$(reldate --from="$FOR_DATE" Sunday)
    END_WEEK=$(reldate --from="$FOR_DATE" Saturday)

    # Now that we have date in the format needed, create a pipeline for the report.
    echo "Report for $START_WEEK through $END_WEEK"
    cat Time_Sheet.txt | shorthand -e "@now := $NOW" |\
        stnparse | stnfilter -start "$START_WEEK" -end "$END_WEEK" | stnreport
```

### Report durations of activities by month

```shell
    #!/bin/bash

    # Get the month/year in YYYY-MM format.
    FOR_DATE=$(date +"%Y-%m")
    # If you normally use 12hr notation then use %I:%M otherwise for 23hr format use %H:M
    NOW=$(date +%I:%M)

    if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
        echo "USAGE: rpt-time-by-month.sh YYYY-MM"
        echo "    Without a date it reports the current week."
        exit 1
    elif [ "$1" != "" ]; then
        FOR_DATE="$1"
    fi

    START_OF_MONTH="$FOR_DATE-01"
    END_OF_MONTH=$(reldate --from $START_OF_MONTH --end-of-month)

    # Now that we have date in the format needed, create a pipeline for the report.
    echo "Report for $START_OF_MONTH through $END_OF_MONTH"
    cat Time_Sheet.txt | shorthand -e "@now := $NOW" | stnparse |\
        stnfilter -start "$START_OF_MONTH" -end "$END_OF_MONTH" | stnreport
```

### Report durations of activities by year

```shell
    #!/bin/bash

    # Get the year in YYYY format.
    FOR_DATE=$(date +"%Y")
    # If you normally use 12hr notation then use %I:%M otherwise for 23hr format use %H:M
    NOW=$(date +%I:%M)

    if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
        echo "USAGE: rpt-time-by-week.sh YYYY-MM-DD"
        echo "    Without a date it reports the current week."
        exit 1
    elif [ "$1" != "" ]; then
        FOR_DATE=$1
    fi  

    function startYear {
        echo "$FOR_DATE-01-01"
    }

    function endYear {
        echo "$FOR_DATE-12-31"
    }

    # Now that we have date in the format needed, create a pipeline for the report.
    echo "Report for $(startYear $FOR_DATE) through $(endYear $FOR_DATE)"
    cat Time_Sheet.txt | shorthand -e "@now := $NOW" | stnparse |\
        stnfilter -start "$(startYear $FOR_DATE)" -end "$(endYear $FOR_DATE)" | stnreport
```


## Installation

_stngo_ and its commands can be installed with the *go get* command.

```
    go get github.com/rsdoiel/stngo/...
```



