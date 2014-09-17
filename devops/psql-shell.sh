#!/bin/sh

psql -h 192.168.10.10 -U gongfu -d gongfu

## ~/.pgpass  
## 192.168.10.10:5432:gongfu:gongfu:kungfupanda
## chmod 0600 .pgpass
