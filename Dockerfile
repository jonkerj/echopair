FROM golang:1.17 as builder
RUN mkdir /workdir
WORKDIR /workdir
COPY . .
RUN go mod download
RUN go build -o /echopair main.go

FROM scratch
ENV GOTRACEBACK=single
COPY --from=builder /echopair /echopair
ENTRYPOINT ["/echopair"]
