# starAPI
API to calculate planets, houses and star positions

## Stars

`http://localhost:8000/run-star?birthdate=1.12.1986&utctime=10:15&stars=Antares,Aldebaran,Algol,Regulus`

```
[
   {
      "altName":"alSco",
      "distance":"5238842987525771.000000000",
      "house":"11.3672040",
      "latitude":"-4.5682411",
      "longitude":"249.5720103",
      "magnitude":"0.910",
      "speed":"0.0000786",
      "starName":"Antares"
   },
   {
      "altName":"alTau",
      "distance":"630479610141846.500000000",
      "house":"5.1034017",
      "latitude":"-5.4682563",
      "longitude":"69.6102578",
      "magnitude":"0.860",
      "speed":"0.0000858",
      "starName":"Aldebaran"
   },
   {
      "altName":"bePer",
      "distance":"850750236492794.500000000",
      "house":"7.0517324",
      "latitude":"22.4275615",
      "longitude":"55.9894234",
      "magnitude":"2.120",
      "speed":"0.0000647",
      "starName":"Algol"
   },
   {
      "altName":"alLeo",
      "distance":"750223063514369.375000000",
      "house":"7.9588747",
      "latitude":"0.4643776",
      "longitude":"149.6464118",
      "magnitude":"1.400",
      "speed":"0.0001850",
      "starName":"Regulus"
   }
]
```

## Houses

`http://localhost:8000/run-houses?birthdate=1.12.1986&utctime=10:15&latitude=-25.42777778&longitude=-49.27305556&altitude=935&housesystem=P`
```
[
   {
      "longitude":"275.7652920",
      "name":"house  1"
   },
   {
      "longitude":"299.0812398",
      "name":"house  2"
   },
   {
      "longitude":"324.3213977",
      "name":"house  3"
   },
   {
      "longitude":"353.9595168",
      "name":"house  4"
   },
   {
      "longitude":"28.2012878",
      "name":"house  5"
   },
   {
      "longitude":"63.4325388",
      "name":"house  6"
   },
   {
      "longitude":"95.7652920",
      "name":"house  7"
   },
   {
      "longitude":"119.0812398",
      "name":"house  8"
   },
   {
      "longitude":"144.3213977",
      "name":"house  9"
   },
   {
      "longitude":"173.9595168",
      "name":"house 10"
   },
   {
      "longitude":"208.2012878",
      "name":"house 11"
   },
   {
      "longitude":"243.4325388",
      "name":"house 12"
   },
   {
      "longitude":"275.7652920",
      "name":"Ascendant"
   },
   {
      "longitude":"173.9595168",
      "name":"MC"
   },
   {
      "longitude":"174.4548852",
      "name":"ARMC"
   },
   {
      "longitude":"47.0830946",
      "name":"Vertex"
   },
   {
      "longitude":"264.9100976",
      "name":"equat. Asc."
   },
   {
      "longitude":"254.4053543",
      "name":"co-Asc. W.Koch"
   },
   {
      "longitude":"306.9309899",
      "name":"co-Asc Munkasey"
   },
   {
      "longitude":"74.4053543",
      "name":"Polar Asc."
   }
]
```

## Planets

`http://localhost:8000/run-planets?birthdate=1.12.1986&utctime=10:15&latitude=-25.42777778&longitude=-49.27305556&altitude=935&housesystem=P`
```
[
   {
      "dailySpeed":"1.0140575",
      "latitude":"-0.0000170",
      "longitude":"248.9206269",
      "name":"Sun"
   },
   {
      "dailySpeed":"15.1472622",
      "latitude":"-3.6372108",
      "longitude":"245.1097231",
      "name":"Moon"
   },
   {
      "dailySpeed":"1.1001474",
      "latitude":"2.2811197",
      "longitude":"228.9030032",
      "name":"Mercury"
   },
   {
      "dailySpeed":"0.2057819",
      "latitude":"0.9652533",
      "longitude":"215.4511866",
      "name":"Venus"
   },
   {
      "dailySpeed":"0.6809279",
      "latitude":"-1.2442962",
      "longitude":"333.6053004",
      "name":"Mars"
   },
   {
      "dailySpeed":"0.0763301",
      "latitude":"-1.3085515",
      "longitude":"343.8609421",
      "name":"Jupiter"
   },
   {
      "dailySpeed":"0.1185810",
      "latitude":"1.4984671",
      "longitude":"251.8207019",
      "name":"Saturn"
   },
   {
      "dailySpeed":"0.0597571",
      "latitude":"-0.1079613",
      "longitude":"261.7538311",
      "name":"Uranus"
   },
   {
      "dailySpeed":"0.0348851",
      "latitude":"1.0182110",
      "longitude":"274.5555780",
      "name":"Neptune"
   },
   {
      "dailySpeed":"0.0362193",
      "latitude":"15.9514438",
      "longitude":"218.5203073",
      "name":"Pluto"
   },
   {
      "dailySpeed":"-0.0529067",
      "latitude":"0.0000000",
      "longitude":"18.1126749",
      "name":"mean Node"
   }
]
```

## Help (required parameters)
`http://localhost:8000/options`
```
{
   "endpoints":{
      "/run-houses":{
         "description":"Endpoint to run the houses binary.",
         "parameters":{
            "altitude":"string (required) - Altitude in meters.",
            "birthdate":"string (required) - Birthdate in the format of 'dd.mm.yyyy', 'dd-mm-yyyy', or 'dd/mm/yyyy'.",
            "debug":"bool - Enable debug mode.",
            "housesystem":"string (required) - House system (P for Placidus, R for Regiomontanus).",
            "latitude":"string (required) - Latitude in decimal format.",
            "longitude":"string (required) - Longitude in decimal format.",
            "utctime":"string (required) - UTC time in the format of 'hh:mm'."
         }
      },
      "/run-planets":{
         "description":"Endpoint to run the planets binary.",
         "parameters":{
            "altitude":"string (required) - Altitude in meters.",
            "birthdate":"string (required) - Birthdate in the format of 'dd.mm.yyyy', 'dd-mm-yyyy', or 'dd/mm/yyyy'.",
            "debug":"bool - Enable debug mode.",
            "housesystem":"string (required) - House system (P for Placidus, R for Regiomontanus).",
            "latitude":"string (required) - Latitude in decimal format.",
            "longitude":"string (required) - Longitude in decimal format.",
            "utctime":"string (required) - UTC time in the format of 'hh:mm'."
         }
      },
      "/run-star":{
         "description":"Endpoint to run the star binary.",
         "parameters":{
            "birthdate":"string (required) - Birthdate in the format of 'dd.mm.yyyy', 'dd-mm-yyyy', or 'dd/mm/yyyy'.",
            "debug":"bool - Enable debug mode.",
            "stars":"string (required) - Comma-separated list of stars.",
            "utctime":"string (required) - UTC time in the format of 'hh:mm'."
         }
      }
   }
}
```
