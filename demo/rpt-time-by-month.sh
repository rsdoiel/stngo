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
