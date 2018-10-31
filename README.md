# onOff_project

A project with golang, to transmit electric current and using pinout of the raspberry

## Getting Started

For help getting started with Golang, view the online
[documentation](https://golang.org/).

### Installing

```
go get github.com/gorilla/mux
go get github.com/stianeikeland/go-rpio
```

## Running the project

```
go run main.go
```

## Example Request

To close the circuit call:

```
http://<IpAddress>:8080/on/?ver=<versionRaspberry>&pin=<pinout>
```

To open the circuit call:

```
http://<IpAddress>:8080/off/?ver=<versionRaspberry>&pin=<pinout>
```

## Built With

- [Golang](https://golang.org/) - the programming language used
- [Raspberry Pi](https://www.raspberrypi.org/) - the device used, you can use all the version available

## Authors

- **Giovanni Lasagna** - _Initial work_ - [lassjr](https://github.com/lassjr)
