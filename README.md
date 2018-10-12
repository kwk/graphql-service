# GraphQL Service

## Development
The following guide is mainly targeted towards a Linux or macOS development
machine.

#### Pre-requisites

Have the following installed on your machine:

- [GoLang from 1.8+](https://golang.org/dl/)
- set the environment variable `GOPATH`.
- get code:

```sh
git clone https://github.com/corinnekrych/graphql-service $GOPATH/src/github.com/corinnekrych/graphql-service
```

#### Build

```
cd $GOPATH/src/github.com/corinnekrych/graphql-service
make build
```

#### Run

```
make server
```

* to test with `curl`:

```
curl -XPOST -d '{"query": "{ workItems {name} }"}' localhost:8000/graphql
curl -XPOST -d '{"query": "{ filter(spaceId: \"e8864cfe-f65a-4351-85a4-3a585d801b45\") {name} }"}' localhost:8000/graphql
```
* or use GraphiQL, go to http://localhost:8000/ and enter in editor:
```
{
  	filter(spaceId: "e8864cfe-f65a-4351-85a4-3a585d801b45") {name}
}
```
