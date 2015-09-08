
# Simple Timesheet Notation

## Misc Notes

Random notes on what this Golang package should look like.


### Commands

+ stnparse - generates a steam of syntax elements encountered and their values as JSON blobs
+ stnfilter - filters syntax stream by date rand returning a JSON blobs
+ stnreport - report generator that renders the JSON blobs into a useful format (e.e.g plain text summary, comma separated value file)
  + if there are more cells of data than the report requires the right hand cells are concatenated by semi-colons


### date lines

Older STN files used a MM/DD/YYYY format. What is the best way to handle those files?

+ a dedicated conversion tool?
+ add support for multiple date formats?


## shorthand ideas

+ text substitutions defined with LABEL := STRING
+ file inclusion defined with LABEL :< PATH TO FILE TO INCLUDE
+ support middle of file extraction negative index refers to lines from end of file
	+ middle 6,-10 would mean the buffer size would be ten lines and when you hit eof the buf will be discarded.
	+ LABEL :< #,# PATH TO FILE FRAGMENT TO INCLUDE
+ support secondary output :> #,# PATH OF FILE TO WRITE


## Todo

+ Make sure local timezone is handled consistently in all tools when converting from YYYY-MM-DD to RFC3339.
    + double check reldate needs to be adjusted to local timezone.
+ Add support for including files via shorthand (e.g. LABEL :< #,# PATH_TO_FILE_TO_BE_INCLUDED)
+ Write middle
+ middle - a utility to pull out the middle of a file using line index numbers (negative numbers count form end of file)
    + support negative indexes (relative to end of file) via a masking buffer
    + create a buffer the size of the last index to end of file, stream file through buffer
    + decide line range uses absolute line numbers or start plus next number of lines

