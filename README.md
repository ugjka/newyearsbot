# NewYearsBot

## Preparing for 2019

[![Build Status](https://travis-ci.org/ugjka/newyearsbot.svg?branch=master)](https://travis-ci.org/ugjka/newyearsbot)
[![Coverage Status](https://coveralls.io/repos/lawrencewoodman/roveralls/badge.svg?branch=master)](https://coveralls.io/r/ugjka/newyearsbot/nyb?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ugjka/newyearsbot/nyb)](https://goreportcard.com/report/github.com/ugjka/newyearsbot/nyb)

Irc Bot for celebrating New Year Eve

Posts New Years for each timezone when they happen

```text
<HNYbot18> Happy New Year in American Samoa (Pago Pago), Niue (Alofi), United States of America (Midway Atoll)
<ugjka> hny !next
<hny2019> Next New Year in 25 minutes 9 seconds in Cambodia (Phnom Penh, Takeo), Christmas Island, Indonesia (Bandung, Bekasi, Depok, Jakarta, Medan, Palembang, Semarang, South Tangerang, Surabaya, Tangerang), Laos (Pakxe, Vientiane), Mongolia (Khovd), Russia (Krasnoyarsk, Novokuznetsk, Novosibirsk), Thailand (Bangkok, Chon Buri, Mueang Nonthaburi, Udon Thani), Vietnam
<ugjka> hny !last
<hny2019> Last NewYear 34 minutes 55 seconds ago in Australia (Mandurah, Perth, Western Australia), Brunei (Bandar Seri Begawan), China (Beijing, Chengdu, Chongqing, Dongguan, Guangzhou, Nanjing, Shanghai, Shenzhen, Tianjin, Wuhan), Hong Kong, Indonesia (Balikpapan, Banjarmasin, Makassar), Macau, Malaysia (Klang, Kota Bharu, Kuala Lumpur), Mongolia (Erdenet, Ulan Bator), Philippines (Manila), Russia (Irkutsk), Singapore, Taiwan
<ugjka> hny Riga
<hny2019> ugjka: New Year in Riga, Rīga, Vidzeme, LV-1050, Latvia will happen in 5 hours 25 minutes
<hny2019> Happy New Year in Nepal (Biratnagar, Kathmandu, Pokhara)
<hny2019> Next New Year in 14 minutes 57 seconds in India (Ahmedabad, Bangalore, Chennai, Hyderabad, Kanpur, Kolkata, Mumbai, New Delhi, Pune, Surat), Sri Lanka (Colombo)
```

## Has GTK3 gui (standalone cli tool also)

![alt text](https://i.imgur.com/hMjsn34.png "Main window")

![alt text](https://i.imgur.com/ze0V82J.png "Bot status")

## Bot's commands

* `hny !next` upcoming new year
* `hny !last` previous new year
* `hny <location>` query location
* `hny !help` show help

The `hny` part can be changed by defining a different trigger

## What's new and great in 2018/2019

* target date wraps around after last zone
* added `hny !last` command to print where previous new year happened
* uses Nominatim for geocoding (no more google api)
* you can specify different Nominatim server if you want
* timezone lookup is done from tz shapefile (no more google api)
* code much more readable

## What's not so great

* timezone shapefile is loaded in memory which increases ram usage by 60 to 80MB
* timezone lookup on slow hardware might be slow

## Installation

### 2019 versions are ready

Arch linux PKGBUILD in archlinux folder

RPM on [releases page](https://github.com/ugjka/newyearsbot/releases)

DEB on [releases page](https://github.com/ugjka/newyearsbot/releases)

### Using make

You need make, go, go-tools, gtk3, glib2

Build with `make all` (`make cli` if you just want the commandline utility)

Install with `make install`
