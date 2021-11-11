# STEP 1 build executable binary
FROM golang as builder

COPY . $GOPATH/src/sample-project/
WORKDIR $GOPATH/src/sample-project/

ENV GO111MODULE=on

#get dependancies
#you can also use dep
RUN go mod download

#build the binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -o /go/bin/app

# STEP 2 run using distroless (usefull because sh not installed)
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/app /
CMD ["/app"]
