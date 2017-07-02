# mock-client

Build `mock-client-M.m.P-I.x86_64.rpm`
and   `mock-client_M.m.P-I_amd64.deb`
where "M.m.P-I" is Major.minor.Patch-Iteration.

## Usage

A program that simulates servers:

1. Unix Domain Socket server

### Invocation

```console
mock-client socket --socket-file /var/run/xyz.sock
```

## Development

### Dependencies

#### Set environment variables

```console
export GOPATH="${HOME}/go"
export PATH="${PATH}:${GOPATH}/bin:/usr/local/go/bin"
export PROJECT_DIR=${GOPATH}/src/github.com/docktermj
```

#### Download project

```console
mkdir -p ${PROJECT_DIR}
cd ${PROJECT_DIR}
git clone git@github.com:docktermj/mock-client.git
```

#### Download dependencies

```console
cd ${PROJECT_DIR}/mock-client
make dependencies
```

### Build

#### Local build

```console
cd ${PROJECT_DIR}/mock-client
make build-local
```

The results will be in the `${GOPATH}/bin` directory.

#### Docker build

```console
cd ${PROJECT_DIR}/mock-client
make build
```

The results will be in the `.../target` directory.

### Test

```console
cd ${PROJECT_DIR}/mock-client
make test-local
```

### Install

#### RPM-based

Example distributions: openSUSE, Fedora, CentOS, Mandrake

##### RPM Install

Example:

```console
sudo rpm -ivh mock-client-M.m.P-I.x86_64.rpm
```

##### RPM Update

Example: 

```console
sudo rpm -Uvh mock-client-M.m.P-I.x86_64.rpm
```

#### Debian

Example distributions: Ubuntu

##### Debian Install / Update

Example:

```console
sudo dpkg -i mock-client_M.m.P-I_amd64.deb
```

### Cleanup

```console
make clean
```
