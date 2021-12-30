FROM golang:1.17.5 as build

WORKDIR /app
COPY ./ /app
RUN go mod download -x
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .

FROM gcr.io/distroless/base-debian10
COPY --from=build /app/ohagi-api /
ENV PORT=${PORT}
ENTRYPOINT ["/ohagi-api"]
