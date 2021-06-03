FROM golang:1.16.0 as build

WORKDIR /app
COPY ./ /app
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM gcr.io/distroless/base-debian10
COPY --from=build /app/ohagi /
ENV PORT=${PORT}
ENTRYPOINT ["/ohagi"]
