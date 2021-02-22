# yet another forum

### run with Docker

#### build

`docker build -t forum .`

#### run

`docker run -dp 8080:8080 forum`

### local build

#### build

`go build -o ./tmp/main ./cmd`

#### run

`./tmp/main`
