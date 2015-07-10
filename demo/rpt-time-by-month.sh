#!/bin/bash

# Get month and year in YYYY-MM format.
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

START_OF_MONTH="${FOR_DATE:0:7}-01"
END_OF_MONTH=$(reldate --from $START_OF_MONTH --end-of-month)

# Now that we have date in the format needed, create a pipeline for the report.
echo "Report for $START_OF_MONTH through $END_OF_MONTH"
cat Time_Sheet.txt | shorthand -e "@now := $NOW" | stnparse |\
    stnfilter -start "$START_OF_MONTH" -end "$END_OF_MONTH" | stnreport
