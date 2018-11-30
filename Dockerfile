FROM golang:alpine
# Install git and build in the docker image
RUN apk update && apk add git
COPY . $GOPATH/src/github.com/ZJU-DistributedAI/ComputingProvider
# ls show
RUN ls -la $GOPATH/src/github.com/ZJU-DistributedAI/ComputingProvider/public
# setting workdir
WORKDIR $GOPATH/src/github.com/ZJU-DistributedAI/ComputingProvider
RUN go get -d -v
RUN go build -o /go/bin/ComputingProvider
# Define entrypoint
ENTRYPOINT ["/go/bin/ComputingProvider"]