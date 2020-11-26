# swagger
[![Build Status](https://travis-ci.org/joaosoft/swagger.svg?branch=master)](https://travis-ci.org/joaosoft/swagger) | [![codecov](https://codecov.io/gh/joaosoft/swagger/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/swagger) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/swagger)](https://goreportcard.com/report/github.com/joaosoft/swagger) | [![GoDoc](https://godoc.org/github.com/joaosoft/swagger?status.svg)](https://godoc.org/github.com/joaosoft/swagger)

A swagger example using the swagger api [goswagger](https://goswagger.io).
The swagger specification is generated using two possible frameworks
- [go-swagger](https://github.com/go-swagger/go-swagger)
- [go-swag](https://github.com/swaggo/swag)

#How to use ?
```
# to install go-swagger command
make install

# to download swagger client
make cli

# to generate swagger with go-swagger to go-swager/docs folder
make gen-go-swagger

# to generate swagger with go-swag to go-swag/docs folder
make gen-go-swag

# after the generation of the swagger.json file, copy it to the folder ./spec
```

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependecy Management 
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
