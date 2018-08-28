# Go bindings for HPE OneSphere APIs

## Prerequisites

go1.9.2 and above.
You can install the latest version from:

[https://golang.org/dl](https://golang.org/dl)

## Usage

#### Install the OneSphere package

```sh
go get github.com/HewlettPackard/hpe-onesphere-go
```

#### Import the package

```go
import (
  "fmt"
  "github.com/HewlettPackard/hpe-onesphere-go"
)
```

#### Connect to the OneSphere server

```go
osbinding.Connect("https://onesphere-host-url", "username", "password")
```

#### Make calls to the OneSphere API

```go
fmt.Println("Status:", onesphere.GetStatus())
```

example output

```json
{"service":"OK","database":""}
```

#### Disconnect from the OneSphere server

```go
osbinding.Disconnect()
```

## Full Example

see [sample/main.go](./sample/main.go)

## Run the sample project

`go get` the OneSphere project

```sh
go get github.com/HewlettPackard/hpe-onesphere-go
```

You must set the OneSphere `host` url, `user`, and `password` flags.

Replace these values:

- `ONESPHERE_HOST_URL`
- `YOUR_ONESPHERE_USERNAME`
- `YOUR_ONESPHERE_PASSWORD`
```sh
go run $GOPATH/src/github.com/HewlettPackard/hpe-onesphere-go/sample/main.go \
  -host=https://ONESPHERE_HOST_URL \
  -user=YOUR_ONESPHERE_USERNAME \
  -password=YOUR_ONESPHERE_PASSWORD
```

_alternatively_ use environment variables
```sh
host=https://ONESPHERE_HOST_URL \
  user=YOUR_ONESPHERE_USERNAME \
  password=YOUR_ONESPHERE_PASSWORD \
  go run $GOPATH/src/github.com/HewlettPackard/hpe-onesphere-go/sample/main.go
```

## APIs

All the APIs return data in JSON format the same as those returned from HPE OneSphere composable APIs.

### Not Implemented Yet

Some APIs are not yet implemented.
