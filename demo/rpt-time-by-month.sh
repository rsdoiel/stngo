#!/bin/bash
FOR_DATE=$(date +"%Y-%m")
if [ "$1" = "--help" ] || [ "$1" = "-h" ]; then
	echo "USAGE: rpt-time-by-month.sh YYYY-MM"
	echo "    Without a date it reports the current week."
	exit 1
elif [ "$1" != "" ]; then
	FOR_DATE=$1
fi

function startMonth {
    echo "$FOR_DATE-01"
}

function endMonth {
	reldate --from="$FOR_DATE-01" --end-of-month
}

echo "Report for $(startMonth $FOR_DATE) through $(endMonth $FOR_DATE)"
cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" | stnparse |\
    stnfilter -start "$(startMonth $FOR_DATE)" -end "$(endMonth $FOR_DATE)" | stnreport
