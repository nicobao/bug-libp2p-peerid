#!/bin/bash

echo "=== JS VERSION ==="
cd js
npm start

printf "\n"
echo "=== GO VERSION ==="
cd ../go
go run main.go
