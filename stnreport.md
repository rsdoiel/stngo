USAGE: stnreport [OPTIONS]

SYNOPSIS

stnreport takes output from stnparse or stnfilter and renders a
report.

	-c	A comma delimited List of zero indexed columns to report
	-columns	A comma delimited List of zero indexed columns to report
	-h	display help
	-help	display help
	-i	input filename
	-input	input filename
	-l	display license
	-license	display license
	-o	output filename
	-output	output filename
	-v	display version
	-version	display version

EXAMPLE

    stnparse -i TimeSheet.txt | stnreport -columns 0,1

This renders columns zero (first column) and one.
%!(EXTRA string=stnreport, string=stnreport, string=stnreport)

stnreport v0.0.5
