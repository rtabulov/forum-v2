# yet another forum

### local build

build

`go build -o ./bin/main ./cmd`

run

`./bin/main`

### run with Docker

build

`docker build -t forum .`

run

`docker run -dp 8080:8080 forum`
