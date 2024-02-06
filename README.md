# mscli
Cli for create microservices with golang

## Install

```bash
go install github.com/devalexandre/go-forge
```

# Commands

## Create Architecture
You Can create a new project with the command below
```bash
forge init <name>
```
if you want to create a new project using examples you can use the command below
```bash
forge init <name> -e
```
## Create Service

```bash
forge  create -n <name> -t service
```

## Create Repository

```bash
forge  create -n <name> -t repository
```

## Create Contract

```bash
forge  create -n <name> -t contract
```

