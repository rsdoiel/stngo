#!/bin/bash

# Get year in YYYY format.
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
