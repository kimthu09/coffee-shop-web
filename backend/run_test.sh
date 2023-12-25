#!/bin/bash
go test ./... -json -v | ./tools/x86/go-test-report -p -t "Coffee Shop Management - Test Report"

TEST_REPORT_FILE=./test_report.html
if test -f "$TEST_REPORT_FILE"; then
    open $TEST_REPORT_FILE
fi
