FROM golang:latest as builder

WORKDIR /pr
ENV APP_ROOT /pr
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN make build

FROM alpine:3.11
COPY --from=builder /pr /pr
RUN chmod -R +x /pr/bin
CMD ["/pr/bin/pr"]