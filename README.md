# RIO
An http request tool for debugging and testing. Declare a request as a json file.

## Installation
```bash
$ go install
```

Check for the comands
```
$ rio -help
```

## How to use
Create a json file with the following structure. e.g. example.json

```json
{
  "name": "Example Request 1",
  "method": "POST",
  "body": {
    "key1": "value1",
    "key2": "value2"
  },
  "headers": {
    "Content-Type": "application/json"
  },
  "url": "http://localhost:8080"
}
```

Then run the following command
```bash
$ rio send <path>/example.json
```
