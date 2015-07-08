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

START_WEEK=$(reldate --from="$FOR_DATE" Sunday)
END_WEEK=$(reldate --from="$FOR_DATE" Saturday)
echo "Report for $START_WEEK through $END_WEEK"
cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" |\
    stnparse | stnfilter -start "$START_WEEK" -end "$END_WEEK" | stnreport

226-79:~ rsdoiel$ cat ~/Sandbox/Timekeeping/rpt-time-by-day.sh 
#!/bin/bash
DAY=$(date +"%Y-%m-%d")
if [ "$1" = "" ]; then
	echo "USAGE: rpt-time-by-date.sh YYYY-MM-DD"
else
    DAY="$1"
fi
cat Time_Sheet.txt | shorthand -e "@now := $(date +%H:%M)" | stnparse |\
    stnfilter -start="$DAY" -end="$DAY" | stnreport
