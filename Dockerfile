FROM golang:alpine
RUN go version

ADD . /go/src/max.work
WORKDIR /go/src/max.work

# Uncomment if you want a hardcoded port at image build time
# ENV PORT 8080
# EXPOSE 8080

# Install make
RUN apk add --update make

# Compile app
RUN make
# Run app
CMD ["/go/src/max.work/bin/max.work"]