
Simple Timesheet Notation
-------------------------

Misc Notes
==========

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

Todo
====

+ Make sure local timezone is handled consistently in all tools when converting from YYYY-MM-DD to RFC3339.
    + double check reldate needs to be adjusted to local timezone.
+ Write middle (e.g. extract middle lines of a file, all but the last N lines of a file, file starting at line N)
+ middle - a utility to pull out the middle of a file using line index numbers (negative numbers count form end of file)
    + support negative indexes (relative to end of file) via a masking buffer
    + create a buffer the size of the last index to end of file, stream file through buffer
    + decide line range uses absolute line numbers or start plus next number of lines
+ Integrate reldate pkg into stnfilter 
+ Migrate out ok to its own test module

