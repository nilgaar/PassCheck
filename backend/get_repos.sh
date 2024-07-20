#! /bin/bash

cd ./backend/sources
git clone --depth 1 https://github.com/danielmiessler/SecLists.git
cd ../scripts

go clean
go run build_tables.go

if [ $(ls ../tables/*.txt | wc -l) -lt 10 ]; then
  echo "Error: less than 10 .txt files in tables folder"
  exit 1
fi
