# goNtrace

goNtrace is a web server application written in Go that performs traceroutes to given IP addresses and generates a map image of the geolocations of the hops.

## Project Structure

```
.
  /cmd
    /goNtrace
      main.go
  /pkg
    traceroute.go
    geolocation.go
    map.go
  /api
    handler.go
  go.mod
  go.sum
  README.md
```

## Getting Started

Clone the repository to your local machine:

```
git clone https://github.com/abdulmeLINK/goNtrace.git
```

## Prerequisites

- Go (Download from the [official website](https://golang.org/dl/))

## Installation

Navigate to the project directory and build the project:

```
cd goNtrace/cmd/goNtrace
go build
```

## Usage

Start the server:

```
./goNtrace --serve
```

Generate a map image:

```
./goNtrace --map IP_ADDRESS
```

Replace `IP_ADDRESS` with the IP address you want to trace.

## Built With

- [Go](https://golang.org/)
- [abdulmeLINK/mtr](https://github.com/abdulmeLINK/mtr)
- [gorilla/mux](https://github.com/gorilla/mux)
- [go-staticmaps](https://github.com/flopp/go-staticmaps)

## Author

- [abdulmeLINK](https://github.com/abdulmeLINK)

## License

This project is licensed under the MIT License.

## TODO

- [ ] Use Google Maps for the frontend.
- [ ] Add labels to the markers on the map.
- [ ] Implement a better geolocator to show every IP address location with better accuracy.
