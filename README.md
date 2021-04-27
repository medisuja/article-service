# ARTICLE SERVICE

## Install dependencies
Install dependency using go mod 

`go mod download`


## Create config

```cp config.example.yaml config.yaml```

And adjust your settings accordingly.


## Migrate Database

```goose up```


## Run application

`gin -p 3000 main.go`

And you are ready to Go!