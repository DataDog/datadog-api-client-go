#!/usr/bin/env bash

if [ "$#" -ne 2 ]; then
	go test -coverpkg=$(go list ./... | grep -v /test | paste -sd "," -) -coverprofile=coverage.txt -covermode=atomic $(go list ./...) -json
else
	if [ "$RERECORD_FAILED_TESTS" == "true" ]; then
		RECORD=true go test -coverpkg=$(go list ./... | grep -v /test | paste -sd "," -) -coverprofile=coverage.txt -covermode=atomic -json $(go list ./...) $1 $2
	else
		go test -coverpkg=$(go list ./... | grep -v /test | paste -sd "," -) -coverprofile=coverage.txt -covermode=atomic -json $(go list ./...) $1 $2
	fi
fi
