package nyb

//Zones contains time zone info in json format
var Zones = []byte(`[
  {
    "countries": [
      {
        "name": "United States of America (Minor Outlying Islands)",
        "cities": [
          "Baker Island",
          "Howland Island"
        ]
      }
    ],
    "offset": -12
  },
  {
    "countries": [
      {
        "name": "American Samoa",
        "cities": [
          "Pago Pago"
        ]
      },
      {
        "name": "Niue",
        "cities": [
          "Alofi"
        ]
      },
      {
        "name": "United States of America",
        "cities": [
          "Midway Atoll"
        ]
      }
    ],
    "offset": -11
  },
  {
    "countries": [
      {
        "name": "Cook Islands",
        "cities": [
          "Avarua",
          "Rarotonga"
        ]
      },
      {
        "name": "French Polynesia",
        "cities": [
          "Papeete",
          "Vaitape (Bora Bora)"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Hawaii",
          "Hilo",
          "Honolulu",
          "Wailuku"
        ]
      },
      {
        "name": "United States of America",
        "cities": [
          "Johnston Atoll"
        ]
      }
    ],
    "offset": -10
  },
  {
    "countries": [
      {
        "name": "French Polynesia",
        "cities": [
          "Atuona",
          "Eiao",
          "Fatu Huku",
          "Taiohae"
        ]
      }
    ],
    "offset": -9.5
  },
  {
    "countries": [
      {
        "name": "French Polynesia",
        "cities": [
          "Gambier Islands",
          "Rikitea"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Alaska",
          "Anchorage"
        ]
      }
    ],
    "offset": -9
  },
  {
    "countries": [
      {
        "name": "Canada",
        "cities": [
          "British Columbia",
          "Surrey",
          "Vancouver",
          "Yukon"
        ]
      },
      {
        "name": "Mexico",
        "cities": [
          "Mexicali",
          "Tijuana"
        ]
      },
      {
        "name": "Pitcairn Islands",
        "cities": [
          "Adamstown"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "California",
          "Los Angeles",
          "Nevada",
          "Oregon",
          "San Diego",
          "San Francisco",
          "San Jose",
          "Seattle"
        ]
      }
    ],
    "offset": -8
  },
  {
    "countries": [
      {
        "name": "Canada",
        "cities": [
          "Alberta",
          "Calgary",
          "Edmonton",
          "Northwest Territories"
        ]
      },
      {
        "name": "Mexico",
        "cities": [
          "Chihuahua",
          "Ciudad Juárez"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Albuquerque",
          "Arizona",
          "Colorado",
          "Denver",
          "El Paso",
          "Idaho",
          "Montana",
          "New Mexico",
          "Phoenix",
          "Utah",
          "Wyoming"
        ]
      }
    ],
    "offset": -7
  },
  {
    "countries": [
      {
        "name": "Belize",
        "cities": []
      },
      {
        "name": "Canada",
        "cities": [
          "Winnipeg"
        ]
      },
      {
        "name": "Costa Rica",
        "cities": []
      },
      {
        "name": "Ecuador",
        "cities": [
          "Puerto Ayora"
        ]
      },
      {
        "name": "El Salvador",
        "cities": []
      },
      {
        "name": "Guatemala",
        "cities": []
      },
      {
        "name": "Honduras",
        "cities": []
      },
      {
        "name": "Mexico",
        "cities": [
          "Mexico City"
        ]
      },
      {
        "name": "Nicaragua",
        "cities": []
      },
      {
        "name": "United States",
        "cities": [
          "Alabama",
          "Arkansas",
          "Austin",
          "Chicago",
          "Dallas",
          "Houston",
          "Illinois",
          "Iowa",
          "Kansas",
          "Louisiana",
          "Memphis",
          "Milwaukee",
          "Minneapolis",
          "Minnesota",
          "Mississippi",
          "Missouri",
          "Nebraska",
          "North Dakota",
          "Oklahoma",
          "Omaha",
          "San Antonio",
          "South Dakota",
          "Tennessee",
          "Texas",
          "Wisconsin"
        ]
      }
    ],
    "offset": -6
  },
  {
    "countries": [
      {
        "name": "Bahamas",
        "cities": []
      },
      {
        "name": "Brazil",
        "cities": [
          "Rio Branco"
        ]
      },
      {
        "name": "Canada",
        "cities": [
          "Montreal",
          "Quebec",
          "Toronto"
        ]
      },
      {
        "name": "Cayman Islands",
        "cities": []
      },
      {
        "name": "Colombia",
        "cities": []
      },
      {
        "name": "Cuba",
        "cities": []
      },
      {
        "name": "Ecuador",
        "cities": [
          "Guayaquil"
        ]
      },
      {
        "name": "Haiti",
        "cities": []
      },
      {
        "name": "Jamaica",
        "cities": []
      },
      {
        "name": "Mexico",
        "cities": [
          "Cancún"
        ]
      },
      {
        "name": "Panama",
        "cities": []
      },
      {
        "name": "Peru",
        "cities": []
      },
      {
        "name": "United States",
        "cities": [
          "Atlanta",
          "Baltimore",
          "Boston",
          "Detroit",
          "Florida",
          "Indianapolis",
          "Kentucky",
          "Massachusetts",
          "Miami",
          "Michigan",
          "New York",
          "Ohio",
          "South Carolina",
          "Tampa",
          "Vermont",
          "Virginia",
          "Virginia Beach",
          "Washington",
          "West Virginia"
        ]
      }
    ],
    "offset": -5
  },
  {
    "countries": [
      {
        "name": "Anguilla",
        "cities": []
      },
      {
        "name": "Antigua and Barbuda",
        "cities": []
      },
      {
        "name": "Aruba",
        "cities": []
      },
      {
        "name": "Barbados",
        "cities": []
      },
      {
        "name": "Bermuda",
        "cities": []
      },
      {
        "name": "Bolivia",
        "cities": []
      },
      {
        "name": "Bonaire, Sint Eustatius and Saba",
        "cities": []
      },
      {
        "name": "Brazil",
        "cities": [
          "Amazonas"
        ]
      },
      {
        "name": "British Virgin Islands",
        "cities": []
      },
      {
        "name": "Canada",
        "cities": [
          "Halifax",
          "Nova Scotia"
        ]
      },
      {
        "name": "Curaçao",
        "cities": []
      },
      {
        "name": "Dominica",
        "cities": []
      },
      {
        "name": "Dominican Republic",
        "cities": []
      },
      {
        "name": "Grenada",
        "cities": []
      },
      {
        "name": "Guadeloupe",
        "cities": []
      },
      {
        "name": "Guyana",
        "cities": []
      },
      {
        "name": "Martinique",
        "cities": []
      },
      {
        "name": "Montserrat",
        "cities": []
      },
      {
        "name": "Puerto Rico",
        "cities": []
      },
      {
        "name": "Saint Lucia",
        "cities": []
      },
      {
        "name": "Trinidad and Tobago",
        "cities": []
      },
      {
        "name": "U.S. Virgin Islands",
        "cities": []
      },
      {
        "name": "Venezuela",
        "cities": []
      }
    ],
    "offset": -4
  },
  {
    "countries": [
      {
        "name": "Canada",
        "cities": [
          "Grand Falls-Windsor",
          "Paradise",
          "St Johns",
          "Stephenville"
        ]
      }
    ],
    "offset": -3.5
  },
  {
    "countries": [
      {
        "name": "Argentina",
        "cities": [
          "Buenos Aires",
          "Córdoba",
          "Rosario"
        ]
      },
      {
        "name": "Brazil",
        "cities": [
          "Alagoas",
          "Amapá",
          "Bahia",
          "Belém",
          "Campo Grande",
          "Ceará",
          "Fortaleza",
          "Maranhão",
          "Mato Grosso",
          "Mato Grosso do Sul",
          "Paraíba",
          "Pará",
          "Pernambuco",
          "Piauí",
          "Recife",
          "Rio Grande do Norte",
          "Salvador",
          "Sergipe",
          "Tocantins"
        ]
      },
      {
        "name": "Chile",
        "cities": [
          "Puente Alto",
          "Santiago"
        ]
      },
      {
        "name": "Falkland Islands",
        "cities": []
      },
      {
        "name": "French Guiana",
        "cities": []
      },
      {
        "name": "Greenland",
        "cities": [
          "Nuuk"
        ]
      },
      {
        "name": "Paraguay",
        "cities": []
      },
      {
        "name": "Saint Pierre and Miquelon",
        "cities": []
      },
      {
        "name": "Suriname",
        "cities": []
      },
      {
        "name": "Uruguay",
        "cities": []
      }
    ],
    "offset": -3
  },
  {
    "countries": [
      {
        "name": "Brazil",
        "cities": [
          "Belo Horizonte",
          "Brasília",
          "Curitiba",
          "Espírito Santo",
          "Federal District",
          "Goiás",
          "Minas Gerais",
          "Paraná",
          "Porto Alegre",
          "Rio Grande do Sul",
          "Rio de Janeiro",
          "Santa Catarina",
          "São Paulo"
        ]
      },
      {
        "name": "South Georgia and South Sandwich Islands",
        "cities": []
      }
    ],
    "offset": -2
  },
  {
    "countries": [
      {
        "name": "Azores",
        "cities": [
          "Ponta Delgada"
        ]
      },
      {
        "name": "Cape Verde",
        "cities": [
          "Praia"
        ]
      },
      {
        "name": "Greenland",
        "cities": [
          "Ittoqqortoormiit"
        ]
      }
    ],
    "offset": -1
  },
  {
    "countries": [
      {
        "name": "Burkina Faso",
        "cities": [
          "Bobo-Dioulasso",
          "Ouagadougou"
        ]
      },
      {
        "name": "Canary Islands",
        "cities": []
      },
      {
        "name": "Faroe Islands",
        "cities": []
      },
      {
        "name": "Gambia",
        "cities": []
      },
      {
        "name": "Ghana",
        "cities": []
      },
      {
        "name": "Guinea",
        "cities": []
      },
      {
        "name": "Guinea-Bissau",
        "cities": []
      },
      {
        "name": "Iceland",
        "cities": []
      },
      {
        "name": "Ireland",
        "cities": [
          "Dublin"
        ]
      },
      {
        "name": "Isle of Man",
        "cities": []
      },
      {
        "name": "Jersey",
        "cities": []
      },
      {
        "name": "Liberia",
        "cities": []
      },
      {
        "name": "Mali",
        "cities": []
      },
      {
        "name": "Mauritania",
        "cities": []
      },
      {
        "name": "Morocco",
        "cities": []
      },
      {
        "name": "Portugal",
        "cities": []
      },
      {
        "name": "Senegal",
        "cities": []
      },
      {
        "name": "Sierra Leone",
        "cities": []
      },
      {
        "name": "Togo",
        "cities": []
      },
      {
        "name": "United Kingdom",
        "cities": [
          "London"
        ]
      },
      {
        "name": "Western Sahara",
        "cities": []
      }
    ],
    "offset": 0
  },
  {
    "countries": [
      {
        "name": "Albania",
        "cities": []
      },
      {
        "name": "Andorra",
        "cities": []
      },
      {
        "name": "Angola",
        "cities": []
      },
      {
        "name": "Austria",
        "cities": []
      },
      {
        "name": "Belgium",
        "cities": []
      },
      {
        "name": "Benin",
        "cities": []
      },
      {
        "name": "Bosnia and Herzegovina",
        "cities": []
      },
      {
        "name": "Cameroon",
        "cities": []
      },
      {
        "name": "Chad",
        "cities": []
      },
      {
        "name": "Croatia",
        "cities": []
      },
      {
        "name": "Czechia",
        "cities": []
      },
      {
        "name": "Denmark",
        "cities": []
      },
      {
        "name": "Equatorial Guinea",
        "cities": []
      },
      {
        "name": "France",
        "cities": []
      },
      {
        "name": "Gabon",
        "cities": []
      },
      {
        "name": "Germany",
        "cities": []
      },
      {
        "name": "Gibraltar",
        "cities": []
      },
      {
        "name": "Hungary",
        "cities": []
      },
      {
        "name": "Italy",
        "cities": []
      },
      {
        "name": "Liechtenstein",
        "cities": []
      },
      {
        "name": "Luxembourg",
        "cities": []
      },
      {
        "name": "Macedonia",
        "cities": []
      },
      {
        "name": "Malta",
        "cities": []
      },
      {
        "name": "Monaco",
        "cities": []
      },
      {
        "name": "Montenegro",
        "cities": []
      },
      {
        "name": "Netherlands",
        "cities": []
      },
      {
        "name": "Niger",
        "cities": []
      },
      {
        "name": "Nigeria",
        "cities": []
      },
      {
        "name": "Norway",
        "cities": []
      },
      {
        "name": "Poland",
        "cities": []
      },
      {
        "name": "San Marino",
        "cities": []
      },
      {
        "name": "Serbia",
        "cities": []
      },
      {
        "name": "Slovakia",
        "cities": []
      },
      {
        "name": "Slovenia",
        "cities": []
      },
      {
        "name": "Spain",
        "cities": []
      },
      {
        "name": "Sweden",
        "cities": []
      },
      {
        "name": "Switzerland",
        "cities": []
      },
      {
        "name": "Tunisia",
        "cities": []
      },
      {
        "name": "Vatican City",
        "cities": []
      }
    ],
    "offset": 1
  },
  {
    "countries": [
      {
        "name": "Botswana",
        "cities": [
          "Gaborone"
        ]
      },
      {
        "name": "Bulgaria",
        "cities": [
          "Sofia"
        ]
      },
      {
        "name": "Burundi",
        "cities": [
          "Bujumbura"
        ]
      },
      {
        "name": "Cyprus",
        "cities": []
      },
      {
        "name": "Egypt",
        "cities": []
      },
      {
        "name": "Estonia",
        "cities": []
      },
      {
        "name": "Finland",
        "cities": []
      },
      {
        "name": "Greece",
        "cities": []
      },
      {
        "name": "Israel",
        "cities": []
      },
      {
        "name": "Jordan",
        "cities": []
      },
      {
        "name": "Latvia",
        "cities": []
      },
      {
        "name": "Lebanon",
        "cities": [
          "Beirut"
        ]
      },
      {
        "name": "Lesotho",
        "cities": [
          "Maseru"
        ]
      },
      {
        "name": "Libya",
        "cities": [
          "Tripoli"
        ]
      },
      {
        "name": "Lithuania",
        "cities": []
      },
      {
        "name": "Malawi",
        "cities": [
          "Lilongwe"
        ]
      },
      {
        "name": "Moldova",
        "cities": []
      },
      {
        "name": "Mozambique",
        "cities": []
      },
      {
        "name": "Namibia",
        "cities": []
      },
      {
        "name": "Palestinian Territories",
        "cities": []
      },
      {
        "name": "Romania",
        "cities": []
      },
      {
        "name": "Russia",
        "cities": [
          "Kaliningrad"
        ]
      },
      {
        "name": "Rwanda",
        "cities": []
      },
      {
        "name": "South Africa",
        "cities": [
          "Cape Town"
        ]
      },
      {
        "name": "Sudan",
        "cities": []
      },
      {
        "name": "Swaziland",
        "cities": []
      },
      {
        "name": "Syria",
        "cities": []
      },
      {
        "name": "Ukraine",
        "cities": [
          "Kyiv"
        ]
      },
      {
        "name": "Zambia",
        "cities": []
      },
      {
        "name": "Zimbabwe",
        "cities": []
      }
    ],
    "offset": 2
  },
  {
    "countries": [
      {
        "name": "Bahrain",
        "cities": [
          "Manama"
        ]
      },
      {
        "name": "Belarus",
        "cities": [
          "Minsk"
        ]
      },
      {
        "name": "Comoros",
        "cities": [
          "Moroni"
        ]
      },
      {
        "name": "Djibouti",
        "cities": [
          "Djibouti"
        ]
      },
      {
        "name": "Eritrea",
        "cities": []
      },
      {
        "name": "Ethiopia",
        "cities": [
          "Addis Ababa"
        ]
      },
      {
        "name": "Iraq",
        "cities": [
          "Baghdad"
        ]
      },
      {
        "name": "Kenya",
        "cities": [
          "Nairobi"
        ]
      },
      {
        "name": "Kuwait",
        "cities": [
          "Al Ahmadi"
        ]
      },
      {
        "name": "Madagascar",
        "cities": [
          "Antananarivo"
        ]
      },
      {
        "name": "Qatar",
        "cities": [
          "Doha"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Moscow",
          "Saint Petersburg"
        ]
      },
      {
        "name": "Saudi Arabia",
        "cities": [
          "Jeddah",
          "Mecca"
        ]
      },
      {
        "name": "Somalia",
        "cities": [
          "Mogadishu"
        ]
      },
      {
        "name": "South Sudan",
        "cities": [
          "Juba",
          "Malakal"
        ]
      },
      {
        "name": "Tanzania",
        "cities": []
      },
      {
        "name": "Turkey",
        "cities": []
      },
      {
        "name": "Uganda",
        "cities": []
      },
      {
        "name": "Yemen",
        "cities": []
      }
    ],
    "offset": 3
  },
  {
    "countries": [
      {
        "name": "Iran",
        "cities": [
          "Isfahan",
          "Karaj",
          "Mashhad",
          "Qom",
          "Shiraz",
          "Tabriz",
          "Tehran"
        ]
      }
    ],
    "offset": 3.5
  },
  {
    "countries": [
      {
        "name": "Armenia",
        "cities": [
          "Yerevan"
        ]
      },
      {
        "name": "Azerbaijan",
        "cities": [
          "Baku",
          "Ganja"
        ]
      },
      {
        "name": "Georgia",
        "cities": [
          "Kutaisi",
          "Tbilisi"
        ]
      },
      {
        "name": "Mauritius",
        "cities": [
          "Port Louis",
          "Vacoas"
        ]
      },
      {
        "name": "Oman",
        "cities": [
          "Muscat"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Samara",
          "Tolyatti"
        ]
      },
      {
        "name": "Réunion",
        "cities": [
          "Saint-Denis"
        ]
      },
      {
        "name": "Seychelles",
        "cities": [
          "Victoria"
        ]
      },
      {
        "name": "United Arab Emirates",
        "cities": [
          "Abu Dhabi",
          "Dubai"
        ]
      }
    ],
    "offset": 4
  },
  {
    "countries": [
      {
        "name": "Afghanistan",
        "cities": [
          "Kabul",
          "Kandahar",
          "Mazari Sharif"
        ]
      }
    ],
    "offset": 4.5
  },
  {
    "countries": [
      {
        "name": "French Southern Territories",
        "cities": [
          "Port-aux-Français"
        ]
      },
      {
        "name": "Kazakhstan",
        "cities": [
          "Aqtöbe"
        ]
      },
      {
        "name": "Maldives",
        "cities": [
          "Malé"
        ]
      },
      {
        "name": "Pakistan",
        "cities": [
          "Faisalabad",
          "Gujranwala",
          "Hyderabad",
          "Islamabad",
          "Karachi",
          "Lahore",
          "Multan",
          "Peshawar",
          "Quetta",
          "Rawalpindi"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Chelyabinsk",
          "Yekaterinburg"
        ]
      },
      {
        "name": "Tajikistan",
        "cities": [
          "Dushanbe",
          "Khujand"
        ]
      },
      {
        "name": "Turkmenistan",
        "cities": [
          "Ashkabad",
          "Turkmenabat"
        ]
      },
      {
        "name": "Uzbekistan",
        "cities": [
          "Namangan",
          "Tashkent"
        ]
      }
    ],
    "offset": 5
  },
  {
    "countries": [
      {
        "name": "India",
        "cities": [
          "Ahmedabad",
          "Bangalore",
          "Chennai",
          "Hyderabad",
          "Kanpur",
          "Kolkata",
          "Mumbai",
          "New Delhi",
          "Pune",
          "Surat"
        ]
      },
      {
        "name": "Sri Lanka",
        "cities": [
          "Colombo"
        ]
      }
    ],
    "offset": 5.5
  },
  {
    "countries": [
      {
        "name": "Nepal",
        "cities": [
          "Biratnagar",
          "Kathmandu",
          "Pokhara"
        ]
      }
    ],
    "offset": 5.75
  },
  {
    "countries": [
      {
        "name": "Bangladesh",
        "cities": [
          "Chittagong",
          "Comilla",
          "Cox’s Bāzār",
          "Dhaka",
          "Jessore",
          "Khulna",
          "Narsingdi",
          "Rajshahi",
          "Rangpur",
          "Tongi"
        ]
      },
      {
        "name": "Bhutan",
        "cities": [
          "Thimphu"
        ]
      },
      {
        "name": "British Indian Ocean Territory",
        "cities": []
      },
      {
        "name": "Kazakhstan",
        "cities": [
          "Almaty",
          "Astana"
        ]
      },
      {
        "name": "Kyrgyzstan",
        "cities": [
          "Bishkek",
          "Osh"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Omsk"
        ]
      }
    ],
    "offset": 6
  },
  {
    "countries": [
      {
        "name": "Burma",
        "cities": [
          "Naypyidaw",
          "Yangon"
        ]
      },
      {
        "name": "Cocos [Keeling,] Islands",
        "cities": []
      }
    ],
    "offset": 6.5
  },
  {
    "countries": [
      {
        "name": "Cambodia",
        "cities": [
          "Phnom Penh",
          "Takeo"
        ]
      },
      {
        "name": "Christmas Island",
        "cities": []
      },
      {
        "name": "Indonesia",
        "cities": [
          "Bandung",
          "Bekasi",
          "Depok",
          "Jakarta",
          "Medan",
          "Palembang",
          "Semarang",
          "South Tangerang",
          "Surabaya",
          "Tangerang"
        ]
      },
      {
        "name": "Laos",
        "cities": [
          "Pakxe",
          "Vientiane"
        ]
      },
      {
        "name": "Mongolia",
        "cities": [
          "Khovd"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Krasnoyarsk",
          "Novokuznetsk",
          "Novosibirsk"
        ]
      },
      {
        "name": "Thailand",
        "cities": [
          "Bangkok",
          "Chon Buri",
          "Mueang Nonthaburi",
          "Udon Thani"
        ]
      },
      {
        "name": "Vietnam",
        "cities": []
      }
    ],
    "offset": 7
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Mandurah",
          "Perth",
          "Western Australia"
        ]
      },
      {
        "name": "Brunei",
        "cities": [
          "Bandar Seri Begawan"
        ]
      },
      {
        "name": "China",
        "cities": [
          "Beijing",
          "Chengdu",
          "Chongqing",
          "Dongguan",
          "Guangzhou",
          "Nanjing",
          "Shanghai",
          "Shenzhen",
          "Tianjin",
          "Wuhan"
        ]
      },
      {
        "name": "Hong Kong",
        "cities": []
      },
      {
        "name": "Indonesia",
        "cities": [
          "Balikpapan",
          "Banjarmasin",
          "Makassar"
        ]
      },
      {
        "name": "Macau",
        "cities": []
      },
      {
        "name": "Malaysia",
        "cities": [
          "Klang",
          "Kota Bharu",
          "Kuala Lumpur"
        ]
      },
      {
        "name": "Mongolia",
        "cities": [
          "Erdenet",
          "Ulan Bator"
        ]
      },
      {
        "name": "Philippines",
        "cities": [
          "Manila"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Irkutsk"
        ]
      },
      {
        "name": "Singapore",
        "cities": []
      },
      {
        "name": "Taiwan",
        "cities": []
      }
    ],
    "offset": 8
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Eucla",
          "Madura",
          "Mundrabilla"
        ]
      }
    ],
    "offset": 8.75
  },
  {
    "countries": [
      {
        "name": "East Timor",
        "cities": [
          "Dili"
        ]
      },
      {
        "name": "Indonesia",
        "cities": [
          "Ambon City",
          "Jayapura"
        ]
      },
      {
        "name": "Japan",
        "cities": [
          "Osaka Prefecture",
          "Tokyo",
          "Yokohama"
        ]
      },
      {
        "name": "Palau",
        "cities": []
      },
      {
        "name": "Russia",
        "cities": [
          "Chita",
          "Yakutsk"
        ]
      },
      {
        "name": "South Korea",
        "cities": [
          "Seoul"
        ]
      },
      {
        "name": "North Korea",
        "cities": [
          "Pyongyang"
        ]
      }
    ],
    "offset": 9
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Darwin",
          "Northern Territory"
        ]
      }
    ],
    "offset": 9.5
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Brisbane",
          "Gold Coast",
          "Queensland"
        ]
      },
      {
        "name": "Guam",
        "cities": [
          "Hagåtña"
        ]
      },
      {
        "name": "Micronesia",
        "cities": [
          "Moen"
        ]
      },
      {
        "name": "Northern Mariana Islands",
        "cities": []
      },
      {
        "name": "Papua New Guinea",
        "cities": [
          "Lae",
          "Port Moresby"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Khabarovsk",
          "Vladivostok"
        ]
      }
    ],
    "offset": 10
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Adelaide",
          "South Australia"
        ]
      }
    ],
    "offset": 10.5
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Australian Capital Territory",
          "Canberra",
          "New South Wales",
          "Sydney",
          "Tasmania",
          "Victoria"
        ]
      },
      {
        "name": "Micronesia",
        "cities": [
          "Palikir"
        ]
      },
      {
        "name": "New Caledonia",
        "cities": [
          "Noumea"
        ]
      },
      {
        "name": "Norfolk Island",
        "cities": [
          "Kingston"
        ]
      },
      {
        "name": "Solomon Islands",
        "cities": [
          "Honiara"
        ]
      },
      {
        "name": "Vanuatu",
        "cities": [
          "Port Vila"
        ]
      }
    ],
    "offset": 11
  },
  {
    "countries": [
      {
        "name": "Kiribati",
        "cities": [
          "Gilbert Islands"
        ]
      },
      {
        "name": "Marshall Islands",
        "cities": [
          "Majuro"
        ]
      },
      {
        "name": "Nauru",
        "cities": []
      },
      {
        "name": "Russia",
        "cities": [
          "Petropavlovsk-Kamchatsky"
        ]
      },
      {
        "name": "Tuvalu",
        "cities": []
      },
      {
        "name": "United States of America (Minor Outlying Islands)",
        "cities": [
          "Wake Island"
        ]
      },
      {
        "name": "Wallis and Futuna",
        "cities": [
          "Mata-Utu"
        ]
      }
    ],
    "offset": 12
  },
  {
    "countries": [
      {
        "name": "Fiji",
        "cities": [
          "Suva"
        ]
      },
      {
        "name": "Kiribati",
        "cities": [
          "Kanton Island"
        ]
      },
      {
        "name": "New Zealand",
        "cities": [
          "Auckland",
          "Wellington"
        ]
      },
      {
        "name": "Tokelau",
        "cities": []
      },
      {
        "name": "Tonga",
        "cities": [
          "Nuku'alofa"
        ]
      }
    ],
    "offset": 13
  },
  {
    "countries": [
      {
        "name": "New Zealand",
        "cities": [
          "Chatham Islands"
        ]
      }
    ],
    "offset": 13.75
  },
  {
    "countries": [
      {
        "name": "Kiribati",
        "cities": [
          "Tabwakea Village"
        ]
      },
      {
        "name": "Samoa",
        "cities": [
          "Apia"
        ]
      }
    ],
    "offset": 14
  }
]
`)
