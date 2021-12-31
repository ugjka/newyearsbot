# NewYearsBot

## Preparing for 2023
[![Go Report Card](https://goreportcard.com/badge/github.com/rhinosf1/newyearsbot/nyb)](https://goreportcard.com/report/github.com/rhinosf1/newyearsbot/nyb)

New Year's Eve IRC party bot

Posts New Years for each timezone when they happen

```text
<hny2021> Happy New Year in American Samoa (Pago Pago), Niue (Alofi), United States of America (Midway Atoll)
<ugjka> !next
<hny2021> Next New Year in 25 minutes 9 seconds in Cambodia (Phnom Penh, Takeo), Christmas Island, Indonesia (Bandung, Bekasi, Depok, Jakarta, Medan, Palembang, Semarang, South Tangerang, Surabaya, Tangerang), Laos (Pakxe, Vientiane), Mongolia (Khovd), Russia (Krasnoyarsk, Novokuznetsk, Novosibirsk), Thailand (Bangkok, Chon Buri, Mueang Nonthaburi, Udon Thani), Vietnam
<ugjka> !last
<hny2021> Last NewYear 34 minutes 55 seconds ago in Australia (Mandurah, Perth, Western Australia), Brunei (Bandar Seri Begawan), China (Beijing, Chengdu, Chongqing, Dongguan, Guangzhou, Nanjing, Shanghai, Shenzhen, Tianjin, Wuhan), Hong Kong, Indonesia (Balikpapan, Banjarmasin, Makassar), Macau, Malaysia (Klang, Kota Bharu, Kuala Lumpur), Mongolia (Erdenet, Ulan Bator), Philippines (Manila), Russia (Irkutsk), Singapore, Taiwan
<ugjka> !hny Riga
<hny2021> New Year in Riga, Rīga, Vidzeme, LV-1050, Latvia will happen in 5 hours 25 minutes
<hny2021> Happy New Year in Nepal (Biratnagar, Kathmandu, Pokhara)
<hny2021> Next New Year in 14 minutes 57 seconds in India (Ahmedabad, Bangalore, Chennai, Hyderabad, Kanpur, Kolkata, Mumbai, New Delhi, Pune, Surat), Sri Lanka (Colombo)
```

## Bot's commands

- `!next` upcoming new year
- `!previous` previous new year
- `!remaining` number of remaining timezones
- `!hny <location>` get new year status for location
- `!time <location>` get current time in location
- `!time` UTC time
- `!help` show help

The command prefix `!` can be changed using the -prefix flag

## Pro-tip

- make sure your system's time is synchronized with NTP

## Installation

### 2021 versions

Arch linux PKGBUILD in archlinux folder

### Using make

You need make, go, go-tools

Build with `make`

Install with `make install`

Uninstall with `make uninstall`

Clean up with `make clean`
