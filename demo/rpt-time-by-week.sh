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
