FROM golang:1.22-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /out/minta-api ./cmd/api

FROM scratch
COPY --from=build /out/minta-api /minta-api
EXPOSE 8080
USER 10001:10001
ENTRYPOINT ["/minta-api"]
