# lokidb CLI
LokiDB CLI (connect to server)

---

### Table of Contents
- [lokidb CLI](#lokidb-cli)
    - [Table of Contents](#table-of-contents)
    - [Featurs](#featurs)
    - [Installation](#installation)
      - [Docker](#docker)
    - [Usage](#usage)


### Featurs
- interactive shell
- autocompleate
- gRPC connection
- slim docker image


### Installation
#### Docker
```shell
docker run -it yoyocode/lokidb-cli -addr <grpc_host>:<grpc_port>
```

### Usage
```shell
$ lokidb-cil --help
Usage of /tmp/go-build1997907887/b001/exe/main:
  -addr string
        the address to connect to (default "localhost:50051")
```

```shell
$ lokidb-cli
>>> help
commands:
    get 
    set 
    del 
    keys 
    flush 
    bye 
    help 

>>> keys
1) a
2) b
3) mom

>>> get asdlg


>>> get a
vavava

>>> exit
```