FROM golang:1.12.5 as builder

LABEL maintainer="jearzamendia@gmail.com"

# Set the Current Working Directory inside the container
WORKDIR /go/src/github.com/jarzamendia/regionmaker

COPY get.sh .

RUN bash get.sh

# Copy everything from the current directory to the PWD(Present Working Directory) inside the container
COPY . .

# Download dependencies
RUN go get -d -v ./...

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /go/bin/regionmaker .

######## Start a new stage from scratch #######
FROM alpine:3.9  

ENV DOCKER_API_VERSION='1.40'


WORKDIR /root

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /go/bin/regionmaker .

COPY static /root/static
COPY templates /root/templates

ENTRYPOINT ["./regionmaker"]

ARG BUILD_DATE


# Labels.
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.build-date=$BUILD_DATE
LABEL org.label-schema.name="regionmaker"
LABEL org.label-schema.description="Sentinel, a Docker Swarm replicas alert dispatcher."
LABEL org.label-schema.url="https://golang.org/"
LABEL org.label-schema.vendor="Jarza"
LABEL org.label-schema.version="1"
LABEL org.label-schema.docker.cmd="docker run -it -v /var/run/docker.sock:/var/run/docker.sock -v d:/data:/data -p 8080:8080 jarzamendia/regionmaker:1.0.0"