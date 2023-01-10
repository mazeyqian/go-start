# go-gin-gee

Gee provides several services for everyday life. The project is based on Gin [1], and the structure refers to ProjectLayout [3]. Some daily scripts in the folder Scripts depend on Script [4], which can run by the command Run.

**Table of Contents**

- [go-gin-gee](#go-gin-gee)
  - [Install](#install)
  - [Script Examples](#script-examples)
  - [API Examples](#api-examples)
    - [Save Data](#save-data)
  - [Build](#build)
    - [Linux](#linux)
    - [Mac](#mac)
    - [Windows](#windows)
  - [Deploy](#deploy)
  - [Contributing](#contributing)
  - [References](#references)

## Install

```
git clone git@github.com:mazeyqian/go-gin-gee.git
```

## Script Examples

Change Git name and email for different projects.

```
go run scripts/change-git-user/main.go -path="/Users/X/Web" -username="Your Name" -useremail="your@email.com"
```

Git pull all projects in a folder.

```
go run scripts/batch-git-pull/main.go
```

Transfer apple note table to markdown table. 

```
go run scripts/transfer-notes-to-md-table/main.go
```

More in folder `scripts`.

## API Examples

Domain is `https://feperf.com`.

### Save Data

**Description:**

Save the data for searching.

**Path:** **/api/gee/create-alias2data**
 
**Method:** **POST**

**Params:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| alias | string | Alias | Yes |
| data | string | Data | Yes |
| public | bool | Public | Yes |

**Example:**

```shell
curl --location --request POST 'https://feperf.com/api/gee/create-alias2data' \
--header 'Content-Type: application/json' \
--data-raw '{
    "alias": "alias example",
    "data": "data example",
    "public": true
}'
```

**Returns:**

| Params | Type | Description | Required |
| :-------- | :--------| :------ | :------ |
| id | int | ID | Yes |
| alias | string | Alias | Yes |
| data | string | Data | Yes |

**Example:**

Success:

```json
{
    "id": 2,
    "created_at": "2023-01-07T11:14:24.572495702+08:00",
    "updated_at": "2023-01-07T11:14:24.57882362+08:00",
    "alias": "alias example",
    "data": "data example"
}
```

Failure:

```json
{
    "code": 400,
    "message": "data exist"
}
```

## Build

Default:

```
go build cmd/api/main.go
```

### Linux

It's usually useful to run the command `chmod u+x script-name-linux-amd64` if the permission error happens.

```
GOOS=linux GOARCH=amd64 go build -o dist/api-linux-amd64 cmd/api/main.go
```

### Mac

```
GOOS=darwin GOARCH=amd64 go build -o dist/api-mac-darwin-amd64 cmd/api/main.go
```

### Windows

```
GOOS=windows GOARCH=amd64 go build -o dist/api-windows-amd64 cmd/api/main.go
```

## Deploy

Supervisor Config:

```
[program:api]
directory=/web/go-gin-gee
command=/web/go-gin-gee/dist/api-linux-amd64 -configpath="data/config.json"
autostart=true
autorestart=true
stderr_logfile=/web/go-gin-gee/log/api.err
stdout_logfile=/web/go-gin-gee/log/api.log
```

## Contributing

```
# All Dependences
go mod download

# Add
go get github.com/example/name
```

If `i/o timeout`, run the command to replace the proxy: 

```
go env -w GOPROXY=https://goproxy.cn
```

It's necessary to run the command `go run scripts/init/main.go` when serving the project first.

```
# Serve
go run cmd/api/main.go

# Restart
# cd /web/go-gin-gee
go run scripts/restart/main.go
```

Visit: `http://127.0.0.1:3000/api/ping`.

```
pong/v1.0.0/2022-09-29 04:52:43
```

## References

1. [Gin Web Framework](https://github.com/gin-gonic/gin)
2. [lo](https://github.com/samber/lo)
3. [project-layout](https://github.com/golang-standards/project-layout)
4. [script](https://github.com/bitfield/script)
5. [go-shadowsocks2](https://github.com/shadowsocks/go-shadowsocks2)
