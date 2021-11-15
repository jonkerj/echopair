FROM golang:1.17 as builder
RUN mkdir /workdir
WORKDIR /workdir
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 go build -o /echopair main.go

FROM scratch
ENV GOTRACEBACK=single
COPY --from=builder /echopair /echopair
ENTRYPOINT ["/echopair"]
