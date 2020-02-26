### go-service

A simple service that returns the square of an input number.

#### Endpoints
##### GET /hb

Acts as a health endpoint, when the service is up and running returns `200 OK` and `hello` as a body

##### POST /square

Takes `input` and returns the result of squaring it as a response. 

#### Development
No external libs are used at the moment.

##### Running the tests
`go test`

##### Building
`go build`

#### Deployment

Every new release (`git tag`) is deployed to am ECS cluster through terraform.
The respective container is published on `docker.io/0xdiba/go-service` 
