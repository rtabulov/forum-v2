# THE Forum

## Available on heroku

[https://rtforum.herokuapp.com/](https://rtforum.herokuapp.com/)

## Docker

### build

`docker build -t forum .`

### run

`docker run -dp 8080:8080 forum`

## Manual

### build

`go build -o ./tmp/main ./cmd`

### run

`./tmp/main`
