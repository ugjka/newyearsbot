# NewYearsBot

## Preparing for 2021

[![Build Status](https://travis-ci.org/ugjka/newyearsbot.svg?branch=master)](https://travis-ci.org/ugjka/newyearsbot)
[![codecov](https://codecov.io/gh/ugjka/newyearsbot/branch/master/graph/badge.svg)](https://codecov.io/gh/ugjka/newyearsbot)
[![Go Report Card](https://goreportcard.com/badge/github.com/ugjka/newyearsbot/nyb)](https://goreportcard.com/report/github.com/ugjka/newyearsbot/nyb)
[![Donate](https://share.ugjka.net/paypal.svg)](https://www.paypal.me/ugjka)
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fugjka%2Fnewyearsbot.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Fugjka%2Fnewyearsbot?ref=badge_shield)

Irc Bot for celebrating New Year Eve

Posts New Years for each timezone when they happen

```text
<hny2020> Happy New Year in American Samoa (Pago Pago), Niue (Alofi), United States of America (Midway Atoll)
<ugjka> hny !next
<hny2020> Next New Year in 25 minutes 9 seconds in Cambodia (Phnom Penh, Takeo), Christmas Island, Indonesia (Bandung, Bekasi, Depok, Jakarta, Medan, Palembang, Semarang, South Tangerang, Surabaya, Tangerang), Laos (Pakxe, Vientiane), Mongolia (Khovd), Russia (Krasnoyarsk, Novokuznetsk, Novosibirsk), Thailand (Bangkok, Chon Buri, Mueang Nonthaburi, Udon Thani), Vietnam
<ugjka> hny !last
<hny2020> Last NewYear 34 minutes 55 seconds ago in Australia (Mandurah, Perth, Western Australia), Brunei (Bandar Seri Begawan), China (Beijing, Chengdu, Chongqing, Dongguan, Guangzhou, Nanjing, Shanghai, Shenzhen, Tianjin, Wuhan), Hong Kong, Indonesia (Balikpapan, Banjarmasin, Makassar), Macau, Malaysia (Klang, Kota Bharu, Kuala Lumpur), Mongolia (Erdenet, Ulan Bator), Philippines (Manila), Russia (Irkutsk), Singapore, Taiwan
<ugjka> hny Riga
<hny2020> ugjka: New Year in Riga, Rīga, Vidzeme, LV-1050, Latvia will happen in 5 hours 25 minutes
<hny2020> Happy New Year in Nepal (Biratnagar, Kathmandu, Pokhara)
<hny2020> Next New Year in 14 minutes 57 seconds in India (Ahmedabad, Bangalore, Chennai, Hyderabad, Kanpur, Kolkata, Mumbai, New Delhi, Pune, Surat), Sri Lanka (Colombo)
```

## Bot's commands

- `hny !next` upcoming new year
- `hny !last` previous new year
- `hny !remaining` number of remaining timezones
- `hny <location>` get new year status for location
- `hny !time <location>` get current time in location
- `hny !time` UTC time
- `hny !help` show help

The `hny` part can be changed by defining a different trigger

## Pro-tip

- make sure your system's time is synchronized with NTP

## Installation

### 2021 versions

Arch linux PKGBUILD in archlinux folder

RPM on [releases page](https://github.com/ugjka/newyearsbot/releases)

DEB on [releases page](https://github.com/ugjka/newyearsbot/releases)

### Using make

You need make, go, go-tools

Build with `make all`

Install with `make install`

## License

[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Fugjka%2Fnewyearsbot.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Fugjka%2Fnewyearsbot?ref=badge_large)
