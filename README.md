## SHORTY

Shorty is a relatively naive URL shortener built with
minimal dependencies:

* [uuid] for generating unique IDS for URLs.
* [mux] for routing.
* [godotenv] to load environment variables.
* [pflag] for command line flags (I like POSIX flags).

### GETTING THE CODE

The service comes as a web server and uses Redis to persist URL data. A command line
version is also provided.

#### GO GET

```go get github.com/buddhamagnet/shorty ...```

Adapt the ```.env-sample``` file, save as ```.env``` and profit!

#### GO RUN

Just clone the repo and run ```go run web/main.go``` to boot the web service or ```go run cli/main.go``` to run the CLI version.

#### DOCKERIZE ALL THE THINGS

Or clone the repo and fire up the containers via ```docker-compose up```.

###  WEB API

The web service exposes three endpoints, a healthcheck endpoint at root, a URL shortener and a redirection service. 

* To get an ID for a long URL, make a reques to ```http://localhost:8080/shorten?url=<longurl>```. If the URL is valid the service will return an ID.
* To use a short URL, make a request to ```http://localhost:8080/<id>```. The service will issue a 301 redirect to that resource.

### CLI

The command line interface provides the ability to generate IDs from long URLs and
decode IDs back into the original long URLs and does not depend on the web service in any way, just the back end (Redis).

Either build the binary or or run as follows:

* ```go run cli/main.go --url=<url>```
* ```go run cli/main.go --id=<id>```

### TESTS

The codebase is covered for tests that exercise the web service and convenience functions, including table tests. Given more time I would have mocked Redis on
the back end but given the simplicity of this application that seemed like overkill.