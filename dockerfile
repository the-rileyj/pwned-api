# escape=`

# Compiles *.go program in ./app,
# then transfers the compiled go program and
# all non-*.go files in ./app to /app in the container

# NOTE: golang:alpine sets WORKDIR to /go,
# which is why /go/gogram is used to copy the program
# instead of just /gogram
FROM golang:alpine AS buildenv

# Add ca-certificates to get the proper certs for making requests,
# gcc and musl-dev for any cgo dependencies, and
# git for getting dependencies residing on github
RUN apk add apk update && `
    apk add --no-cache ca-certificates gcc git musl-dev

# Create directory structure properly so that the import paths match up
WORKDIR /go/src/github.com/the-rileyj/pwned-api
RUN mkdir ./functionality

# Copy source files into their correct locations in the directory structure
COPY main.go .
COPY ./functionality/functionality.go ./functionality

# Get dependencies locally, but don't install
RUN go get -d -v ./...

# Compile program with local dependencies
RUN env CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -a -v -o gogram

# Second stage of build, adding in files and running
# newly compiled program
FROM scratch

# Copy the *.go program compiled in the first stage
COPY --from=buildenv /go/gogram /

RUN mkdir secret

COPY ./secret/mailgun.json ./secret/

ENV mailgunFile=./secret/mailgun.json

# Add HTTPS Certificates
COPY --from=buildenv /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Expose ports 80 to host machine
EXPOSE 80

# Run program
ENTRYPOINT ./gogram
