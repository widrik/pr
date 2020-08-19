FROM golang:latest as builder
WORKDIR /pr/
ENV APP_ROOT /pr
COPY . .
RUN make build

FROM alpine:3.11
COPY --from=builder /pr/bin/pr /pr/bin/pr
RUN chmod -R +x /pr/bin
CMD ["/pr/bin/pr"]