#!/bin/bash
go test ./... -coverprofile=c.out
./tools/x86/gocov convert c.out | ./tools/x86/gocov-html >cov_report.html
rm -rf ./c.out
COV_REPORT_FILE=./cov_report.html
if test -f "$COV_REPORT_FILE"; then
    open $COV_REPORT_FILE
fi
