
# Simple Timesheet Notation


## Misc Notes

Random notes on what this Golang package should look like.

Tokens

+ *ShorthandDefinitionMarker* - space colon equals ("#mything := this is my things", "#mything" is the symbol, "this is my things" is the value)
+ *EntryMarker* - entry marker is to new line characters
+ *CellMarker* - a semi-colon
+ *RangeMarker* - space dash space, separates two times (do I need to suppoert date ranges?)
+ *ActiveDate* - A single date on a line ending in an *entryMarker*. Date is in form of YYYY-MM-DD
+ *TimeElement* - in either 12 or 25 hour notations. 24:00 is always midnight, 12:00 is always noon, 00:00 is start of day, cannot cross date boundry
+ *TextBlock* - everything else, formatting is preserved and assuemd to be Markdown friend

Commands

+ shorthand - Takes an Simple timesheet notation file and expands all the shorthand found.
+ stnparse - generates a steam of syntax elements encountered and their values as JSON blobs
+ stnfilter - filters syntax stream by date rand returning a JSON blobs
+ stnreport - report generator that renders the JSON blobs into a userful format (e.e.g plain text summary, comma separated value file)
  + if there are more cells of data than the report requires the right hand cells are concatendated by semi-colons


# Todo

+ Figure out how to handle timezone offset (e.g. assume local, allow of explicit timezone, etc)

```go
    // Example getting local timezone
    now := time.Now();
    location := now.Location()
    timezone := now.Zone()
    fmt.Printf("Now: %v, Location: %v, Timezone: %v\n", now, location, timezone)
```


## shorthand ideas

+ text substitutions defined with LABEL := STRING
+ file inclusion defined with LABEL :< PATH TO FILE TO INCLUDE
+ support middle of file exstraction negative index refers to lines from end of file
	+ creat a buffer the size of the last index to end of file, stream file through buffer
	+ middle 6,-10 would mean the buffer size would be ten lines and when you hit eof the buf wil be discarded.
	+ LABEL :< #,# PATH TO FILE FRAGMENT TO INCLUDE
+ support secondary output :> #,# PATH TO WRITE


## date lines

Older STN files used a MM/DD/YYYY format. What is the best way to handle those files?

+ a dedicated covertion tool?
+ add support for multiple date formats?


