FROM golang:alpine AS builder
COPY httpenv.go /go
COPY httpenv_test.go /go
RUN go build httpenv.go

FROM builder AS test
RUN go mod init httpenv
CMD ["go","test","-v","./..."]

FROM alpine AS final
RUN apk add --no-cache curl
RUN addgroup -g 1000 httpenv \
    && adduser -u 1000 -G httpenv -D httpenv
COPY --from=builder --chown=httpenv:httpenv /go/httpenv /httpenv
EXPOSE 8888
# we're not changing user in this example, but you could:
# USER httpenv
CMD ["/httpenv"]