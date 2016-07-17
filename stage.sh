#!/bin/bash
echo "Setting up..."
mkdir -p .build/etc/systemd/system;
mkdir -p .build/etc/gravemind;
mkdir -p .build/usr/local/bin;
mkdir -p .build/usr/local/share/gravemind;

echo "Copying config and template files..."
cp configs/orion.json .build/usr/local/share/gravemind/;
cp configs/gravemind.service .build/etc/systemd/system/;
cp configs/gravemind.json .build/etc/gravemind/;
cp configs/dashboard.html .build/usr/local/share/gravemind/;

echo "Building gravemind..."
go build -o .build/usr/local/bin/gravemind;

echo "Packaging orion archive..."
orion-packager;

echo "Done"
