# API Go Consuming SOAP requests

A Golang API design to consume requests SOAP and return JSON.

## API Stack

**Server:** Golang, Docker, Travis CI

## Run Locally

Clone the project

```bash
  git clone https://github.com/rafaelsanzio/go-consuming-soap
```

Go to the project directory

```bash
  cd go-consuming-soap
```

Install dependencies

- Need to install docker (https://docs.docker.com/desktop/)

Start the server

```bash
  docker-compose up
```

## Running Tests

To run tests, run the following command

```bash
  make test
```

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

## Author

- [@rafaelsanzio](https://www.github.com/rafaelsanzio)
