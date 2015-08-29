[![build status](https://secure.travis-ci.org/rsdoiel/stn.png)](http://travis-ci.org/rsdoiel/stn)
stn - Simple Time Notation
===============================


# Overview

I've often found it necessary to keep track of time spent on projects or
activities.  Every "tracking system" I've used has worked for me at some level
accept one. I always forget to use them. I forget because they break my
workflow. I work with allot of text files so the system that works for me
is a text file log. Over time I have simplified that format which has made it
easy to enter, read and parse by a program. It was inspired
by other simple formats like Textile and markdown but rather than focus
on rendering to HTML typically Simple Timesheet Notation is rendered to
a flat delimited text stream or JSON encoded data structure.

Here's the summary view of the notation. Date are entered on single line by
themselves. Dates are in the YYYY-MM-DD format (e.g. November 10, 2012 would be
typed 2012-11-10) and are applied to all following time entries until another
date is encountered.  If no date is indicated then the assumed date is today.

Time entries take up a single line and start with a time range. Time ranges are
in the form of HH:MM and do not assume 24hr or 12hr representation. Time entries
are assumed to be within a single calendar day (e.g. 00:00 to 24:00). The start
time and end time are separated by a hyphen. The range itself is terminated by
a semi-colon. The rest the line is made up of a semi-colon delimited list of
"tags" or phrases (e.g. descriptions). These are typically are transformed into
columns (e.g. in tab delimited output) or cells of an array (in the case of JSON).

Here's an example of a recording of travel and a meeting on November 6, 2012 -

```shell
	2012-11-06

	7:45 - 8:30; travel; train to meeting

	8:30 - 12:00; meeting; Standing committee for secret world domination by miniature sentient petunias.
```

This example suggests to column (or array cells) associated with the each time
range. Here is an example of that markup rendered as a tab delimit table -

```shell
	2012-11-06T06:45:00-08:00	2012-11-06T07:30:00-08:00	travel	train to meeting
	2012-11-06T07:30:00-08:00	2012-11-06T11:00:00-08:00	meeting	Standing committee for secret world domination by miniature sentient petunias.
```

Note that "travel" and "train" as well as "meeting" and "Standing" are separated
by tabs in this example.

A JSON presentation might look like -

```json
	[{
		"Start":"2012-11-06T06:45:00-08:00",
		"End":"2012-11-06T07:30:00-08:00",
		"Annotations":["travel","train to meeting"]
	},
	{
		"Start":"2012-11-06T07:30:00-08:00",
		"End":"2012-11-06T11:00:00-08:00",
		"Annotations":[
			"meeting",
			"Standing committee for secret world domination by miniature sentient petunias."
		]
	}]
```

Integration then is just a matter of pushing the data into the appropriate database
or service.


# Notation details

## Dates of entries

A line which contains only a numerically formatted a date (e.g. YYYY-MM-DD
format) indicates the start of a log entries for a particular day.  It is
typed only once. All time entries after that are affiliated with that day.


E..g Say I'm entering time for November 3, 2011. I would note it one a single
line as-

```
	2011-11-03
```

Any following entries would then refer to that date until a new date was encountered.


## Time entries

An entry is a line which indicates an activity to be tracked. A time entry
consists of a time range and one or more semi-colon delimited keywords or phrases.
A time range is made up for two time elements separated by space, dash and space.
The time element should be in the HH:MM format and does not assume AM/PM or
24 hour notation. The range is in the form of "HH:MM - HH:MM".

I practice I usually use the first semi-colon delimited element to name a project
and follow that by a short activity description. The notation doesn't assume the
meaning of the semi-colon elements and just treats them as columns of text.

Here is an example entry about debugging code on a project call "timesheet"
from 8:30 AM until 1:00 PM (i.e. 13:00) I would note it as

```
	8:30 - 1:00; timesheet; debugging parsing code.
```

If could also look like this

```
	8:30 - 13:00; timesheet; debugging parsing code.
```

### 12 hour versus 24 hour time notation

If you not using a twelve hour clock it is assume the first time before the
dash is the start time and the second entry is the end time.  Calculating hours
then evolves looking at the relationship of those two times to each other.  If
the start time is smaller then the end time then simple subtraction of the
start from the end calculates hours spent.  If that is not the case (i.e. you
have crossed the noon boundary) then you will need to normalize the values
before subtracting the start from end time. A time range of "8:00 - 1:00"
would normalize to "8:00 - 13:00".

## Embedding extra text

If the line is not a date or time entry it is ignored.  This allows pre-processors
like *shorthand* to integrate easily with a simple timesheet notation file. It
also lets you add notes to self or extended descriptions without worry they
will get pushed to whatever is down the pipeline after the stn parse is done
its job.

## Example timesheet entries for a day

In the following example is time entries for November 19, 2011 working on
simple timesheet parsing project.

```text
	2012-11-19

	// A meta entry which records pounds and meters associated with date
	// Nov. 19, 2011
	{lbs: 175.0, meters: 2.9}

	8:30 - 11:00; timesheet notation; Writing a README.md describing my simple timesheet notation.

	11:00 - 12:00; timesheet notation; Drafting a NodeJS example parser for the notation.

	1:00 - 3:00; timesheet notation; debugging parser, realizing I can simplify my notation further and drop the first semi-colon.

	Realized I need to keep some backward compatibility for my parse so I don't have to rewrite/edit my ascii timesheet file.
```

The lines starting with "//" and "{" are ignored since they are not recognized as
a date or time entry. Likewise the last line starting with "Realized" is skipped.
