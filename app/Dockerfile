FROM golang:1.14-buster as builder
WORKDIR /go/src/app

ADD main.go /go/src/app
RUN go build -o /go/bin/app

# Response to return
ARG RES
RUN echo $RES > /tmp/data


FROM gcr.io/distroless/base-debian10
EXPOSE 9000
COPY --from=builder /go/bin/app /app
COPY --from=builder /tmp/data /tmp/data
ENTRYPOINT [ "/app", "/tmp/data" ]