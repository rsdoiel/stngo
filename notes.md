
# Simple Timesheet Notation

## Misc Notes

Random notes on what this Golang package should look like.

Tokens

+ *shorthandDefinishtionMArker* - space colon equals ("#mything := this is my things", "#mything" is the symbol, "this is my things" is the value) 
+ *entryMarker* - entry marker is to new line characters
+ *cellMarker* - a semi-colon
+ *rangeMarker* - space dash space, separates two times (do I need to suppoert date ranges?)
+ *activeDate* - A single date on a line ending in an *entryMarker*. Date is in form of YYYY-MM-DD
+ *timeElement* - in either 12 or 25 hour notations. 24:00 is always midnight, 12:00 is always noon, 00:00 is start of day, cannot cross date boundry
+ *textBlock* - everything else, formatting is preserved and assuemd to be Markdown friend

Commands

+ stnexpand - Takes an Simple timesheet notation file and expands all the short hand found.
+ stnparse - generates a steam of syntax elements encountered and their values as JSON blobs
+ stnfilter - filters syntax stream by date rand returning a JSON blobs
+ stnreport - report generator that renders the JSON blobs into a userful format (e.e.g plain text summary, comma separated value file)
  + if there are more cells of data than the report requires the right hand cells are concatendated by semi-colons





