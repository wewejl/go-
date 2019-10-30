#!/bin/bash
rm *.db
go build -o cli *.go 
./cli println
