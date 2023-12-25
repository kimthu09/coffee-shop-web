#!/bin/bash
go test ./... -json -v | ./tools/x86/go-test-report -p -t "Coffee Shop Management - Test Report"

rm ./report/test_report.html
mv ./test_report.html ./report/test_report.html

TEST_REPORT_FILE=./report/test_report.html
if test -f "$TEST_REPORT_FILE"; then
    open $TEST_REPORT_FILE
fi
