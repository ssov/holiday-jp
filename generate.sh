#!/bin/bash

set -euo pipefail

HOLIDAY_CSV_URL=https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv

echo 'package holiday' > var.go
echo 'import "time"' >> var.go
echo 'var holidays = map[time.Time]struct{}{' >> var.go

holidays=$(curl -Lfs ${HOLIDAY_CSV_URL} | nkf -w | tail -n+2)
for line in ${holidays}
do
    line=$(echo ${line} | tr -d "\r")
    date=$(echo ${line} | cut -d ',' -f1)
    time=$(LC_ALL=en_US.UTF-8 date +"time.Date(%Y, time.%B, %e, 0, 0, 0, 0, jst)" --date="${date}")
    echo -e "\t${time}: struct{}{}," >> var.go
done

echo "}" >> var.go

gofmt -w var.go
