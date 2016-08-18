[![Travis](https://travis-ci.org/ilijamt/postfixbeat.svg?branch=master)](https://travis-ci.org/ilijamt/postfixbeat)
[![GoReportCard](https://goreportcard.com/badge/github.com/ilijamt/postfixbeat)](https://goreportcard.com/badge/github.com/ilijamt/postfixbeat)

# Postfixbeat

Welcome to Postfixbeat.

Ensure that this folder is at the following location:
`${GOPATH}/github.com/ilijamt`

## Getting Started with Postfixbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.6
* [Glide](https://github.com/Masterminds/glide) >= 0.10.0

### Build

To build the binary for Postfixbeat run the command below. This will generate a binary
in the same directory with the name postfixbeat.

```
make
```

### Run

To run Postfixbeat with debugging output enabled, run:

```
./postfixbeat -c postfixbeat.yml -e -d "*"
```


### Test

To test Postfixbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`


### Package

To be able to package Postfixbeat the requirements are as follows:

 * [Docker Environment](https://docs.docker.com/engine/installation/) >= 1.10
 * $GOPATH/bin must be part of $PATH: `export PATH=${PATH}:${GOPATH}/bin`

To cross-compile and package Postfixbeat for all supported platforms, run the following commands:

```
cd dev-tools/packer
make deps
make images
make
```

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `etc/fields.yml`.
To generate etc/postfixbeat.template.json and etc/postfixbeat.asciidoc

```
make update
```


### Cleanup

To clean  Postfixbeat source code, run the following commands:

```
make fmt
make simplify
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Postfixbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/github.com/ilijamt
cd ${GOPATH}/github.com/ilijamt
git clone https://github.com/ilijamt/postfixbeat
```
