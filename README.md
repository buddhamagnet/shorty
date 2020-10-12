## SHORTY

Shorty is a relatively naive URL shortener built with
minimal dependencies:

* [uuid](https://github.com/google/uuid) for generating unique IDS for URLs.
* [mux](https://github.com/gorilla/mux) for routing.
* [godotenv](https://github.com/joho/godotenv) to load environment variables.
* [pflag](https://github.com/spf13/pflag) for command line flags (I like POSIX flags).
* [redigo](https://github.com/gomodule/redigo) - Redis client.

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

* To get an ID for a long URL, make a POST request to ```http://localhost:8080/shorten``` with a JSON payload like ```{"url":"http://www.google.com"}```. If the URL is valid the service will return an ID.
* To use a short URL, make a request to ```http://localhost:8080/<id>```. The service will issue a 301 redirect to that resource.

* I added a GraphQL layer for shortening URLs, to spin this up, navigate to the ```graphql``` folder, run ```yarn``` and ```yarn run start```, navigate to port 5000 and issue a request, example:

```javascript
mutation shorten {
  shorten(url: "http://www.google.com") {
    id
  }
}
```

#### EXAMPLES

POST request to generate a short URL, using [httpie](https://httpie.org/):

```txt
http -f POST http://localhost:8080/shorten <<<'{"url":"http://google.com"}'
```

```json
{
    "id": "c9bd4a"
}
```

### CLI

The command line interface provides the ability to generate IDs from long URLs and
decode IDs back into the original long URLs and does not depend on the web service in any way, just the back end (Redis).

Either build the binary or or run as follows:

* ```go run cli/main.go --url=<url>```
* ```go run cli/main.go --service=decode --id=<id>```

### TESTS

The codebase is covered for tests that exercise the web service and convenience functions, including table tests. The CLI version is also comprehensively tested for all flag permutations. Given more time I would have mocked Redis on the back end but given the simplicity of this application that seemed like overkill.
