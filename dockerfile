FROM golang:1.9

# setup go
ENV GOBIN $GOPATH/bin
ENV PATH $GOBIN:/usr/local/go/bin:$PATH

COPY build $GOBIN
RUN crawler --version

ENV COMMAND "--help"
CMD crawler $COMMAND
