# AB Testing
An Application POC to display how with two different containers, you can route traffic via Headers with Istio.

### Running locally
To run locally simply run:
`go run main.go /path/to/file`

Once running to see the response:
`curl http://localhost:9000` should return `<File Contents>`

### Docker
With the docker image, setting the build arg RES, will create a file with the contents of the value
```
docker build --build-arg RES=test-a -t danquack/ab-testing:a .
docker build --build-arg RES=test-b -t danquack/ab-testing:b .

docker push danquack/ab-testing:a
docker push danquack/ab-testing:b
```

### Routing with Virtual Services
TODO