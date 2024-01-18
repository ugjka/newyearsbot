# NewYearsBot

## Preparing for 2025, updates landing in December as usual
## I'm doing it again for the greater good

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
- `!time <location>` get the current time in a location
- `!time` UTC
- `!help` show help

The command prefix `!` can be changed using the -prefix flag

## Pro-tip

- make sure your system's time is synchronized with NTP

## Installation

### 2024 versions

Arch linux PKGBUILD in archlinux folder

Prebuilt Linux binaries are available on the releases [page](https://github.com/ugjka/newyearsbot/releases)

### Using make

You need to have make, go, go-tools

Build with `make`

Install with `make install`

Uninstall with `make uninstall`

Clean up with `make clean`

### Specifying channel key

-channels "#channelname:channelkey, #channelname2:channelkey2"

### Yaml config file

See example config: [settings.yaml](settings.yaml)

Useful if you wanna run multiple bot instances across different IRC hosts
