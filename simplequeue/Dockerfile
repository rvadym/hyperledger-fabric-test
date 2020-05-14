FROM golang:1.14.2-alpine3.11 as base

CMD ["tail", "-f", "/dev/null"]


#----------------------------------#
# Setup for local development only #
#----------------------------------#
FROM base AS develop

RUN apk add curl vim mc git build-base libc-dev tree htop bash \
    && curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh \
    && go get github.com/onsi/ginkgo/ginkgo \
    && go get github.com/onsi/gomega \
    && go get -u golang.org/x/lint/golint \
    && go get -u github.com/rakyll/gotest \
    && go get -u github.com/golang/mock/mockgen \
    && go get -u github.com/go-delve/delve/cmd/dlv

#----------------------------------#
# Setup for prod only              #
#----------------------------------#
# FROM base AS prod

# RUN CGO_ENABLEG=0 go build main.go

# FROM alpine:3.7
# RUN apk add --no-cache ca-certificates
# COPY --from=0 /go/..../main .
# CMD ["./main"]