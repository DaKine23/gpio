#!/usr/bin/env bash
mkdir -p ./target/arm
mkdir -p ./target/linux/x64
mkdir -p ./target/windows/x64


GOOS=linux go build main.go
mv main target/linux/x64/gpio
GOOS=windows go build main.go
mv main.exe target/windows/x64/gpio.exe
GOOS=linux GOARCH=arm go build main.go
mv main target/arm/gpio


