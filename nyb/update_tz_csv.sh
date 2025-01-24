#!/bin/bash
curl https://timezonedb.com/files/TimeZoneDB.csv.zip > time_zone.csv.zip
unzip -o time_zone.csv.zip time_zone.csv
rm time_zone.csv.zip