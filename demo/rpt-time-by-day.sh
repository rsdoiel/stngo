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
