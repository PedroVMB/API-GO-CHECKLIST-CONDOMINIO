FROM golang:1.22.1

WORKDIR /app

COPY go.mod ./
COPY main.go ./

RUN go build -o /server

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /server /server

EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]
