# SOLID Go
This repository is a basic client-server code to operate CRUD on Elasticsearch with SOLID principle.

## Prerequisites
- [Docker](https://docs.docker.com/engine/install/)
- [Go](https://go.dev/doc/install)

## How to Run 
Let's look at Makefile to run a command that we need.

### Run Elasticsearch
Execute `make upES` to turn on our Elasticsearch locally.

### Running CRUD Server
Execute `make run` to run our CRUD Server.

### Stop Elasticsearch
Execute `make downES` to teardown running Elasticsearch.