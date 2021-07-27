FROM golang:1.16.6-buster as builder

ADD . /src
WORKDIR /src

RUN make build

FROM gcr.io/distroless/base
COPY --from=builder /src/bin/cuneum-api /bin
CMD ["/bin/cuneum-api"]
