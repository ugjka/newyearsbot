package nyb

// Zones contains time zone information in JSON format
var Zones = []byte(`[
  {
    "countries": [
      {
        "name": "United States",
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
        "name": "United States",
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
          "Papeete"
        ]
      },
      {
        "name": "Leeward Islands",
        "cities": [
          "Bora Bora",
          "Huahine",
          "Raiatea",
          "Tahaa"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Adak",
          "Hawaii",
          "Johnston Atoll"
        ]
      },
      {
        "name": "Windward Islands",
        "cities": [
          "Maiao",
          "Mehetia",
          "Moorea",
          "Tahiti",
          "Tetiaroa"
        ]
      }
    ],
    "offset": -10
  },
  {
    "countries": [
      {
        "name": "Marquesas Islands",
        "cities": [
          "Fatu Hiva",
          "Hiva Oa",
          "Nuku Hiva",
          "Tahuata",
          "Ua Huka",
          "Ua Pou"
        ]
      }
    ],
    "offset": -9.5
  },
  {
    "countries": [
      {
        "name": "Gambier Islands",
        "cities": [
          "Akamaru",
          "Aukena",
          "Mangareva",
          "Taravai"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Alaska"
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
          "Vancouver"
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
          "Las Vegas",
          "Portland",
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
          "Inuvik",
          "Northwest Territories",
          "Whitehorse",
          "Yukon"
        ]
      },
      {
        "name": "Mexico",
        "cities": [
          "Ciudad Juárez",
          "Hermosillo"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Arizona",
          "Boise",
          "Colorado",
          "Montana",
          "New Mexico",
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
          "Arkansas",
          "Fargo",
          "Houston",
          "Huntsville",
          "Illinois",
          "Iowa",
          "Louisiana",
          "Minnesota",
          "Mississippi",
          "Missouri",
          "Nashville",
          "Oklahoma",
          "Omaha",
          "Sioux Falls",
          "Wichita",
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
        "name": "Chile",
        "cities": [
          "Easter Island"
        ]
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
          "Detroit",
          "Indianapolis",
          "Louisville",
          "Massachusetts",
          "Miami",
          "New York",
          "Ohio",
          "South Carolina",
          "Virginia",
          "Washington"
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
        "name": "Bonaire",
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
          "Halifax"
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
        "name": "Greenland",
        "cities": [
          "Qaanaaq"
        ]
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
        "name": "Saba",
        "cities": []
      },
      {
        "name": "Saint Lucia",
        "cities": []
      },
      {
        "name": "Sint Eustatius",
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
          "Mary's Harbour",
          "Paradise (Newfoundland and Labrador)",
          "St. John's",
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
          "Rio de Janeiro",
          "Salvador",
          "São Paulo"
        ]
      },
      {
        "name": "Chile",
        "cities": [
          "Puente Alto",
          "Punta Arenas",
          "Santiago"
        ]
      },
      {
        "name": "Falkland Islands",
        "cities": [
          "Stanley"
        ]
      },
      {
        "name": "French Guiana",
        "cities": [
          "Cayenne"
        ]
      },
      {
        "name": "Paraguay",
        "cities": [
          "Asuncion"
        ]
      },
      {
        "name": "Saint Pierre and Miquelon",
        "cities": []
      },
      {
        "name": "Suriname",
        "cities": [
          "Paramaribo"
        ]
      },
      {
        "name": "Uruguay",
        "cities": [
          "Montevideo"
        ]
      }
    ],
    "offset": -3
  },
  {
    "countries": [
      {
        "name": "Brazil",
        "cities": [
          "Fernando de Noronha, Pernambuco"
        ]
      },
      {
        "name": "Greenland",
        "cities": [
          "Ittoqqortoormiit",
          "Kangerlussuaq",
          "Nuuk"
        ]
      },
      {
        "name": "South Georgia and the South Sandwich Islands",
        "cities": [
          "King Edward Point"
        ]
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
        "name": "Cabo Verde",
        "cities": [
          "Praia"
        ]
      },
      {
        "name": "Cape Verde",
        "cities": [
          "Praia"
        ]
      },
      {
        "name": "Portugal",
        "cities": [
          "Azores"
        ]
      }
    ],
    "offset": -1
  },
  {
    "countries": [
      {
        "name": "Burkina Faso",
        "cities": []
      },
      {
        "name": "Canary Islands",
        "cities": []
      },
      {
        "name": "Cote d'Ivoire",
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
        "name": "Greenland",
        "cities": [
          "Danmarkshavn"
        ]
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
        "cities": []
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
        "name": "Portugal",
        "cities": [
          "Lisbon"
        ]
      },
      {
        "name": "Saint Helena",
        "cities": []
      },
      {
        "name": "Sao Tome and Principe",
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
        "name": "Algeria",
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
        "name": "Bosnia-Herzegovina",
        "cities": []
      },
      {
        "name": "CAR",
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
        "name": "Congo-Brazzaville",
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
        "name": "DRC",
        "cities": [
          "Kinshasa"
        ]
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
        "name": "Morocco",
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
        "name": "North Macedonia",
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
        "name": "Vatican",
        "cities": []
      }
    ],
    "offset": 1
  },
  {
    "countries": [
      {
        "name": "Botswana",
        "cities": []
      },
      {
        "name": "Bulgaria",
        "cities": []
      },
      {
        "name": "Burundi",
        "cities": []
      },
      {
        "name": "Cyprus",
        "cities": []
      },
      {
        "name": "DRC",
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
        "name": "Eswatini",
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
        "name": "Latvia",
        "cities": []
      },
      {
        "name": "Lebanon",
        "cities": []
      },
      {
        "name": "Lesotho",
        "cities": []
      },
      {
        "name": "Libya",
        "cities": []
      },
      {
        "name": "Lithuania",
        "cities": []
      },
      {
        "name": "Malawi",
        "cities": []
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
        "name": "South Sudan",
        "cities": []
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
        "name": "Ukraine",
        "cities": []
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
        "cities": []
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
        "cities": []
      },
      {
        "name": "Eritrea",
        "cities": []
      },
      {
        "name": "Ethiopia",
        "cities": []
      },
      {
        "name": "Iraq",
        "cities": [
          "Baghdad"
        ]
      },
      {
        "name": "Jordan",
        "cities": []
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
        "cities": []
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
          "Riyadh"
        ]
      },
      {
        "name": "Somalia",
        "cities": [
          "Mogadishu"
        ]
      },
      {
        "name": "Syria",
        "cities": [
          "Damascus"
        ]
      },
      {
        "name": "Tanzania",
        "cities": [
          "Dodoma"
        ]
      },
      {
        "name": "Turkey",
        "cities": [
          "Istanbul"
        ]
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
          "Izhevsk",
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
        "name": "France",
        "cities": [
          "Amsterdam Island",
          "Port-aux-Français"
        ]
      },
      {
        "name": "Kazakhstan",
        "cities": [
          "Aqtöbe",
          "Oral"
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
          "Islamabad",
          "Karachi"
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
          "Dushanbe"
        ]
      },
      {
        "name": "Turkmenistan",
        "cities": [
          "Ashgabat"
        ]
      },
      {
        "name": "Uzbekistan",
        "cities": [
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
          "Bangalore",
          "Delhi",
          "Lucknow",
          "Mumbai"
        ]
      },
      {
        "name": "Sri Lanka",
        "cities": [
          "Sri Jayawardenepura Kotte"
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
          "Dhaka"
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
        "cities": [
          "Diego Garcia"
        ]
      },
      {
        "name": "Kazakhstan",
        "cities": [
          "Astana"
        ]
      },
      {
        "name": "Kyrgyzstan",
        "cities": [
          "Bishkek"
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
        "name": "Cocos (Keeling) Islands",
        "cities": [
          "Home Island",
          "West Island"
        ]
      },
      {
        "name": "Myanmar",
        "cities": [
          "Naypyidaw",
          "Yangon"
        ]
      }
    ],
    "offset": 6.5
  },
  {
    "countries": [
      {
        "name": "Cambodia",
        "cities": [
          "Phnom Penh"
        ]
      },
      {
        "name": "Christmas Island",
        "cities": []
      },
      {
        "name": "Indonesia",
        "cities": [
          "Jakarta"
        ]
      },
      {
        "name": "Laos",
        "cities": [
          "Vientiane"
        ]
      },
      {
        "name": "Mongolia",
        "cities": [
          "Hovd"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Krasnoyarsk",
          "Norilsk",
          "Novosibirsk"
        ]
      },
      {
        "name": "Thailand",
        "cities": [
          "Bangkok"
        ]
      },
      {
        "name": "Vietnam",
        "cities": [
          "Hanoi"
        ]
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
          "Shanghai",
          "Shenzhen",
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
          "Kuala Lumpur"
        ]
      },
      {
        "name": "Mongolia",
        "cities": [
          "Ulaanbaatar"
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
        "cities": [
          "Taipei"
        ]
      }
    ],
    "offset": 8
  },
  {
    "countries": [
      {
        "name": "Australia",
        "cities": [
          "Cocklebiddy",
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
          "Jayapura",
          "Manokwari"
        ]
      },
      {
        "name": "Japan",
        "cities": [
          "Tokyo"
        ]
      },
      {
        "name": "North Korea",
        "cities": [
          "Pyongyang"
        ]
      },
      {
        "name": "Palau",
        "cities": [
          "Ngerulmud"
        ]
      },
      {
        "name": "Russia",
        "cities": [
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
        "name": "Timor-Leste",
        "cities": [
          "Dili"
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
          "Alice Springs",
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
          "Port Moresby"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Khabarovsk",
          "Verkhoyansk",
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
          "Aṉangu Pitjantjatjara Yankunytjatjara",
          "Broken Hill",
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
          "Canberra",
          "Lord Howe Island",
          "Macquarie Island",
          "Melbourne",
          "Sydney",
          "Tasmania"
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
        "name": "Russia",
        "cities": [
          "Magadan",
          "Srednekolymsk",
          "Yuzhno-Sakhalinsk"
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
        "name": "Fiji",
        "cities": [
          "Suva"
        ]
      },
      {
        "name": "France",
        "cities": [
          "Wallis and Futuna"
        ]
      },
      {
        "name": "Kiribati",
        "cities": [
          "Gilbert Islands",
          "Tarawa"
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
        "cities": [
          "Yaren"
        ]
      },
      {
        "name": "Norfolk Island",
        "cities": [
          "Kingston"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Anadyr",
          "Petropavlovsk-Kamchatsky",
          "Pevek"
        ]
      },
      {
        "name": "Tuvalu",
        "cities": [
          "Funafuti"
        ]
      },
      {
        "name": "United States",
        "cities": [
          "Wake Island"
        ]
      }
    ],
    "offset": 12
  },
  {
    "countries": [
      {
        "name": "Kiribati",
        "cities": [
          "Kanton",
          "Phoenix Islands"
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
        "name": "Samoa",
        "cities": [
          "Apia"
        ]
      },
      {
        "name": "Tokelau",
        "cities": [
          "Atafu",
          "Fakaofo"
        ]
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
          "Kiritimati",
          "Line Islands"
        ]
      }
    ],
    "offset": 14
  }
]`)
