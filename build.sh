#!/bin/bash
GOARM=7 GOARCH=arm go build -o .build/gravemind
orion-packager
