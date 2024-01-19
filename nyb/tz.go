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
          "Hilo",
          "Honolulu",
          "Johnston Atoll",
          "Wailuku"
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
          "Alaska",
          "Anchorage",
          "Fairbanks",
          "Juneau",
          "Unalaska"
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
          "Albuquerque",
          "Arizona",
          "Colorado",
          "Denver",
          "El Paso",
          "Idaho",
          "Montana",
          "New Mexico",
          "Phoenix",
          "Salt Lake City",
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
        "name": "Chile",
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
        "name": "Greenland",
        "cities": [
          "Qaanaaq",
          "Thule Air Base"
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
          "Alagoas",
          "Amapá",
          "Bahia",
          "Belém",
          "Brasilia",
          "Ceará",
          "Fortaleza",
          "Maranhão",
          "Paraíba",
          "Rio de Janeiro",
          "Salvador",
          "Santa Catarina",
          "Sergipe",
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
        "name": "Greenland",
        "cities": [
          "Ittoqqortoormiit"
        ]
      },
      {
        "name": "Portugal",
        "cities": [
          "Ponta Delgada"
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
        "cities": [
          "Reykjavik"
        ]
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
          "Cardiff",
          "Edinburgh",
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
        "name": "Congo",
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
        "cities": [
          "Athens"
        ]
      },
      {
        "name": "Israel",
        "cities": []
      },
      {
        "name": "Latvia",
        "cities": [
          "Riga"
        ]
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
        "name": "South Sudan",
        "cities": [
          "Juba"
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
          "Belushya Guba",
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
          "Dar es Salaam",
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
          "Port-aux-Francais"
        ]
      },
      {
        "name": "French Southern Territories",
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
          "Male",
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
          "Quetta"
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
          "Ashgabat",
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
          "Bengaluru",
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
          "Colombo",
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
        "cities": [
          "Diego Garcia"
        ]
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
          "Pontianak",
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
          "Hovd",
          "Khovd"
        ]
      },
      {
        "name": "Russia",
        "cities": [
          "Khatanga",
          "Krasnoyarsk",
          "Norilsk",
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
          "Eucla",
          "Madura"
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
          "Osaka Prefecture",
          "Tokyo",
          "Yokohama"
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
          "Chita",
          "Tiksi",
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
          "Cairns",
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
          "Komsomolsk-on-Amur",
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
          "Hobart",
          "Melbourne",
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
          "Kanton"
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
