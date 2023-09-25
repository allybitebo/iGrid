# regctl

pass 1234

```

regctl is a swiss knife command-line tool for igrid system. It is mostly
used for testing and management as well as maintenance. It contains just
five main commands, which are add, get,list, delete and update that are
used for management of nodes, users and regions.

Usage:
  regctl [command]

Available Commands:
  add         add (users |nodes |regions)
  db          registry database management
  delete      delete (users |nodes |regions) <id>
  get         get (users |nodes |regions) <id>
  help        Help about any command
  list        list (users |nodes |regions )
  update      update (user |node |region)

Flags:
      --address string    the address of regsvc (default "http://localhost")
      --config string     config file (default is $HOME/.regctl.yaml)
  -h, --help              help for regctl
      --password string   user password
      --port string       regsvc port (default ":8080")
      --uuid string       user unique identifier

Use "regctl [command] --help" for more information about a command.


```


### add command
