#!/bin/bash
go test ./... -coverprofile=c.out
./tools/x86/gocov convert c.out | ./tools/x86/gocov-html >report/cov_report.html
rm -rf ./c.out
COV_REPORT_FILE=./report/cov_report.html
if test -f "$COV_REPORT_FILE"; then
    open $COV_REPORT_FILE
fi
