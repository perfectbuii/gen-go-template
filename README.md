<div align="center">

<img height="120" alt="Thanks for visiting my repository" width="100%" src="https://raw.githubusercontent.com/BrunnerLivio/brunnerlivio/master/images/marquee.svg" />

# Gen Go Template

<p align="center">
<a href="https://github.com/harish-sethuraman/readme-components">
 <img  src="https://readme-components.vercel.app/api?component=logo&fill=black&logo=go&animation=spin&svgfill=15d8fe">  
 </a>
 <a href="https://github.com/harish-sethuraman/readme-components">
 <img  src="https://readme-components.vercel.app/api?component=logo&fill=black&logo=github&animation=spin">  
 </a>
</p>

[![CI](https://github.com/idodod/protoc-gen-fieldmask/actions/workflows/ci.yml/badge.svg)](https://github.com/idodod/protoc-gen-fieldmask/actions/workflows/ci.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/idodod/protoc-gen-fieldmask)](https://goreportcard.com/report/github.com/idodod/protoc-gen-fieldmask)
![GitHub release (latest SemVer)](https://img.shields.io/github/v/release/idodod/protoc-gen-fieldmask)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/idodod/protoc-gen-fieldmask)
![GitHub](https://img.shields.io/github/license/idodod/protoc-gen-fieldmask)

<i>A curated list of Go Gen Tools READMEs</i>

<a href="https://github.com/perfectbuii/gen-go-template/stargazers"><img src="https://img.shields.io/github/stars/perfectbuii/gen-go-template" alt="Stars Badge"/></a>
<a href="https://github.com/perfectbuii/gen-go-template/network/members"><img src="https://img.shields.io/github/forks/perfectbuii/gen-go-template" alt="Forks Badge"/></a>
<a href="https://github.com/perfectbuii/gen-go-template/pulls"><img src="https://img.shields.io/github/issues-pr/perfectbuii/gen-go-template" alt="Pull Requests Badge"/></a>
<a href="https://github.com/perfectbuii/gen-go-template/issues"><img src="https://img.shields.io/github/issues/perfectbuii/gen-go-template" alt="Issues Badge"/></a>
<a href="https://github.com/perfectbuii/gen-go-template/graphs/contributors"><img alt="GitHub contributors" src="https://img.shields.io/github/contributors/perfectbuii/gen-go-template?color=2b9348"></a>
<a href="https://github.com/perfectbuii/gen-go-template/blob/master/LICENSE"><img src="https://img.shields.io/github/license/perfectbuii/gen-go-template?color=2b9348" alt="License Badge"/></a>

</div>

## Features:
🔭 Main features of this repo to generate everything to make our coding process super fast while keeping the architecture and code style very good.<br />

🔭 Auto generate protobuf files.<br />
🔭 Auto generate mock interface for DDD. <br />
🔭 Auto generate all layer of DDD. <br />
🔭 Auto generate sql query with struct mapping and entities. <br />
🔭 Auto migrate with Postgres. <br />
🔭 Add pgroonga extension for search. <br />
🔭 Run concurrency for searching team and hub. <br />

## Project Structure:

```sh
.
├── cmd
│   ├── gen-layer
│   │   ├── internal
│   │   │   ├── generator.go
│   │   │   └── step.go
│   │   ├── main.go
│   │   ├── models
│   │   │   ├── cli_step.go
│   │   │   ├── feature.go
│   │   │   └── template.go
│   │   ├── templates
│   │   │   ├── cucumber
│   │   │   │   ├── create.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── delivery
│   │   │   │   ├── create.tpl
│   │   │   │   ├── default.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── godog
│   │   │   │   ├── create.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   ├── repository
│   │   │   │   ├── create.tpl
│   │   │   │   ├── default.tpl
│   │   │   │   ├── delete.tpl
│   │   │   │   ├── list.tpl
│   │   │   │   ├── retrieve.tpl
│   │   │   │   └── update.tpl
│   │   │   └── service
│   │   │       ├── create.tpl
│   │   │       ├── default.tpl
│   │   │       ├── delete.tpl
│   │   │       ├── list.tpl
│   │   │       ├── retrieve.tpl
│   │   │       └── update.tpl
│   │   └── utils
│   │       ├── parser
│   │       │   └── parser.go
│   │       └── steps.go
│   ├── protoc-gen-custom
│   │   ├── internal
│   │   │   └── generator.go
│   │   └── main.go
│   └── server
│       └── main.go
├── config
│   └── config.go
├── database
│   ├── migrations
│   │   ├── 0001_migrate.up.sql
│   │   └── 0002_migrate.up.sql
│   └── queries
│       ├── hub.sql
│       ├── team.sql
│       └── user.sql
├── developments
│   ├── bdd_test.Dockerfile
│   ├── docker-compose.yml
│   ├── gen-go.sh
│   ├── groonga-build.sh
│   ├── postgres.Dockerfile
│   ├── proto.Dockerfile
│   └── sqlc.yaml
├── docs
│   ├── html
│   │   └── index.html
│   └── swagger
│       ├── auth.swagger.json
│       ├── call.swagger.json
│       ├── customer.swagger.json
│       ├── data.swagger.json
│       ├── enum.swagger.json
│       ├── file.swagger.json
│       ├── hub.swagger.json
│       ├── project.swagger.json
│       ├── search.swagger.json
│       ├── task.swagger.json
│       ├── team.swagger.json
│       └── user.swagger.json
├── go.mod
├── go.sum
├── integration_test
│   ├── bdd_setup.go
│   ├── bdd_test.go
│   ├── create_hub.go
│   ├── create_team.go
│   ├── features
│   │   ├── create_hub.feature
│   │   ├── create_team.feature
│   │   └── search_hub.feature
│   └── search_hub.go
├── internal
│   ├── deliveries
│   │   ├── grpc
│   │   │   ├── hub.go
│   │   │   ├── search.go
│   │   │   ├── team.go
│   │   │   └── user.go
│   │   └── http
│   │       ├── auth.go
│   │       ├── http.go
│   │       └── user.go
│   ├── middleware
│   │   └── authentication.go
│   ├── models
│   │   ├── db.go
│   │   ├── hub.sql.go
│   │   ├── models.go
│   │   ├── querier.go
│   │   ├── team.sql.go
│   │   └── user.sql.go
│   ├── repositories
│   │   ├── hub.go
│   │   ├── search.go
│   │   ├── team.go
│   │   └── user.go
│   └── services
│       ├── hub.go
│       ├── hub_test.go
│       ├── search.go
│       ├── team.go
│       └── user.go
├── Makefile
├── mocks
│   ├── dbtx.go
│   ├── hub_repository.go
│   ├── hub_service.go
│   ├── querier.go
│   ├── search_repository.go
│   ├── search_service.go
│   ├── team_repository.go
│   ├── team_service.go
│   ├── user_repository.go
│   └── user_service.go
├── pb
│   ├── auth_enums.pb.go
│   ├── auth_grpc.pb.go
│   ├── auth_methods.pb.go
│   ├── auth.pb.go
│   ├── auth.pb.gw.go
│   ├── auth.pb.validate.go
│   ├── enum_enums.pb.go
│   ├── enum_methods.pb.go
│   ├── enum.pb.go
│   ├── enum.pb.validate.go
│   ├── hub_enums.pb.go
│   ├── hub_grpc.pb.go
│   ├── hub_methods.pb.go
│   ├── hub.pb.go
│   ├── hub.pb.gw.go
│   ├── hub.pb.validate.go
│   ├── search_enums.pb.go
│   ├── search_grpc.pb.go
│   ├── search_methods.pb.go
│   ├── search.pb.go
│   ├── search.pb.gw.go
│   ├── search.pb.validate.go
│   ├── team_enums.pb.go
│   ├── team_grpc.pb.go
│   ├── team_methods.pb.go
│   ├── team.pb.go
│   ├── team.pb.gw.go
│   ├── team.pb.validate.go
│   ├── user_enums.pb.go
│   ├── user_grpc.pb.go
│   ├── user_methods.pb.go
│   ├── user.pb.go
│   ├── user.pb.gw.go
│   └── user.pb.validate.go
├── proto
│   ├── enum.proto
│   ├── hub.proto
│   ├── options
│   │   ├── annotations.pb.go
│   │   ├── annotations.proto
│   │   └── doc.go
│   ├── search.proto
│   ├── team.proto
│   └── user.proto
├── README.md
├── target
├── transform
│   ├── hub_transformer.go
│   ├── options.go
│   ├── search_transformer.go
│   ├── team_transformer.go
│   └── user_transformer.go
└── utils
    ├── authenticate
    │   ├── authenticator.go
    │   ├── jwt.go
    │   ├── jwt_test.go
    │   ├── paseto.go
    │   ├── paseto_test.go
    │   ├── payload.go
    │   └── token.go
    ├── crypto
    │   ├── sha256.go
    │   └── sha256_test.go
    ├── helper
    │   ├── validation.go
    │   └── validation_test.go
    ├── pathutils
    │   └── path.go
    ├── string.go
    └── transformhelpers
        └── helpers.go
```

# ARCHITECTURE

![clean architecture](https://raw.githubusercontent.com/phungvandat/clean-architecture/dev/images/clean-arch.png)

<div align="center">

## Installation:

Make sure you have Go installed ([download](https://golang.org/dl/)).

Install make for start the server.

For Linux:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">sudo apt install make</a></i></pre>
</h2>

For Macos:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">brew install make</a></i></pre>
</h2>

## How to review my online test:

You run: 

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make start-project</a></i></pre>
</h2>

After that check my integration test to know flow of prj:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make bdd-test</a></i></pre>
</h2>

## How to start server:

You only need run one command:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make start-project</a></i></pre>
</h2>

Some detail step to debug in local ^^

First of all, you must start postgres:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make start-postgres</a></i></pre>
</h2>

After that should migrate:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make migrate</a></i></pre>
</h2>

Start server with cmd/terminal:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make run</a></i></pre>
</h2>

## Tools:

Generate sql:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-sql</a></i></pre>
</h2>

Generate proto:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-proto</a></i></pre>
</h2>

Generate layer by DDD (delivery, service, repository):

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-layer</a></i></pre>
</h2>

## Unit test:

Generate mock:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make gen-mock</a></i></pre>
</h2>

Run all test:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make test</a></i></pre>
</h2>

## Integration test:

<h2 align="center">
<pre><i><a href="https://rednafi.github.io/reflections" target="_blank">make bdd-test</a></i></pre>
</h2>

</div>

## License:

MIT

**Free Software, Hell Yeah!**
