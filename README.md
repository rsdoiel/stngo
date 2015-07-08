
# stngo

Golang implementation of Simple Timesheet Notation utilities and packages.

+ *shorthand* - a simple label/substition processor for text strings and file inclusion. Works with standard input and output.
+ *reldate* - a utility to generate relative dates in the YYYY-MM-DD format expected by Simple Timesheet Notation
+ *stnparse* - translates a standard input and output turning Simple Timesheet Notation into a tab delimited table with RFC3339 dates or JSON blobs.
+ *stnfilter* - filters the output of *stnparse* by dates or text string
+ *stnreport* - summarizes the tab delimited output of *stnfilter* or *stnparse* yielding a simple table showing hours and first annotations

This makes it easy to process and generate reports with Bash and common Unix utilities (e.g. *date*)

## Examples of using these utilities in a Unix pipeline

In this example we are filtering entries for a specific date.

### Report durations of activities by day

```shell
    #!/bin/bash
    DAY=$(date +"%Y-%m-%d")
    if [ "$1" = "" ]; then
        echo "USAGE: rpt-time-by-date.sh YYYY-MM-DD"
    else
        DAY="$1"
    fi
    # Now that we have date in the format needed, create a pipeline for the report.
    cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" | stnparse |\
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

    # Now that we have date in the format needed, create a pipeline for the report.
    START_WEEK=$(reldate --from="$FOR_DATE" Sunday)
    END_WEEK=$(reldate --from="$FOR_DATE" Saturday)
    echo "Report for $START_WEEK through $END_WEEK"
    cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" |\
        stnparse | stnfilter -start "$START_WEEK" -end "$END_WEEK" | stnreport
```

### Report durations of activities by month

```shell
    #!/bin/bash
    FOR_DATE=$(date +"%Y-%m")
    if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
    	echo "USAGE: rpt-time-by-month.sh YYYY-MM"
    	echo "    Without a date it reports the current week."
    	exit 1
    elif [ "$1" != "" ]; then
    	FOR_DATE="$1-01"
    fi

    END_OF_MONTH=$(reldate --from $FOR_DATE --end-of-month)

    # Now that we have date in the format needed, create a pipeline for the report.
    echo "Report for $FOR_DATE through $(reldate --from="$FOR_DATE" --end-of-month)"
    cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" | stnparse |\
        stnfilter -start "$FOR_DATE" -end "$END_OF_MONTH" | stnreport
```

### Report durations of activities by year

```shell
    #!/bin/bash
    FOR_DATE=$(date +"%Y")
    if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
        echo "USAGE: rpt-time-by-week.sh YYYY"
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
    cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" | stnparse |\
        stnfilter -start "$(startYear $FOR_DATE)" -end "$(endYear $FOR_DATE)" | stnreport
```
