# Canary Testing
An Application POC to display how with two different containers, you can route traffic via Headers with Istio.

## Running app locally
To run locally simply run:
`go run main.go /path/to/file`

Once running to see the response:
`curl http://localhost:9000`

The application has two endpoints:

`/` -  returns `<File Contents>`

`/headers` - returns headers sent to app

### Docker
With the docker image, setting the build arg RES, will create a file with the contents of the value
```
docker build --build-arg RES=test-a -t canary-testing:a .
docker build --build-arg RES=test-b -t canary-testing:b .
```

### Routing with Virtual Services
With the pod being labeled, with `release: canary`, you can utilize virtual service and destination rules to route traffic to the new canary.

Define routing subsets:
```yaml
# Destination Rule
...
    subsets:
    - labels:
        release: canary
      name: canary
    - labels:
        release: production
      name: production
```

```yaml
# Virtual Service
...
- name: canary-route
  match:
    - headers:
        release:
          exact: canary
  route:
    - destination:
        host: app-canary-testing
        subset: canary
...
```

## Deploy with Helm
1. Update [values.yaml](values.yaml) with host, ingress, or any other environment specific information.
2. `helm upgrade -n $ISTIO_ENABLED_NAMESPACE app $PWD`