# Timer
a project to showcase my skills

# Build echo middleware

Example project of test assessment from imaginary company. 

### Build a middleware using echo framework

First of all, you should create a handler which sends how many days left until 13 may (18:40) 2025 and response with HTTP 200 OK status code.

Secondly, build a middleware, which checks HTTP header `User-Role` presents and contains `admin` and prints `admin-user detected` to the console (using default log package or any 3rd party) if so.

## RUN
```shell
go run github.com/Evafag02/Timer/cmd/timer
```
## Test

```shell
curl --location --request GET '127.0.0.1:8080/status' \
--header 'User-Role: admin'
```
