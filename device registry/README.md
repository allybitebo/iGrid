# registry
igrid registry

contains two programs
- regctl - a command line tool
- regsvc - a registry server

### build the binaries
```bash
cd cmd/regctl
go build -o ../bin/regctl
cd ../regsvc
go build -o ../bin/regsvc
cd ../bin
```

### run the registry server
go to bin directory after building them and execute

```bash
./regsvc
```

### use regctl
```bash
./regctl
```
after typing the command it will return help message on how to use the tool
before starting using it PostgreSQL should be installed.
refctl offer a utility to test and set up the database after installation.
run
```bash
./regctl db
```
a help message will pop up on how to use the utility
