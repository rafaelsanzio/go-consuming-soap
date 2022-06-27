# API Go Consuming SOAP requests

A Golang API design to consume requests SOAP and return JSON.

## API Stack

**Server:** Golang, Docker, Travis CI

## API Reference

#### Get all continents

```http
  GET /continents
```

#### Example Response

```json
[
  {
    "Code": "AF",
    "Name": "Africa"
  },
  {
    "Code": "AN",
    "Name": "Antarctica"
  },
  {
    "Code": "AS",
    "Name": "Asia"
  },
  {
    "Code": "EU",
    "Name": "Europe"
  },
  {
    "Code": "OC",
    "Name": "Ocenania"
  },
  {
    "Code": "AM",
    "Name": "The Americas"
  }
]
```

#### Get countries

```http
  GET /countries
```

| Query Parameters | Type  | Description                                      |
| :--------------- | :---- | :----------------------------------------------- |
| `limit`          | `int` | **Not Required**. Limit of item to fetch         |
| `offset`         | `int` | **Not Required**. Offset to start to fetch items |

#### Example Response

```json
[
  {
    "Code": "AX",
    "Name": "Åland Islands"
  },
  {
    "Code": "AF",
    "Name": "Afghanistan"
  },
  {
    "Code": "AL",
    "Name": "Albania"
  },
  {
    "Code": "DZ",
    "Name": "Algeria"
  },
  {
    "Code": "AS",
    "Name": "American Samoa"
  }
]
```

#### Get countries info by country code

```http
  GET /countries/{code}
```

| Parameter | Type     | Description                             |
| :-------- | :------- | :-------------------------------------- |
| `code`    | `string` | **Required**. The country code to fetch |

#### Example Response

```json
{
  "Capital": "Brasilia",
  "Currency": {
    "Code": "BRL",
    "Name": "Brazil Real"
  },
  "Flag": "http://www.oorsprong.org/WebSamples.CountryInfo/Flags/Brazil.jpg",
  "PhoneCode": 55
}
```
