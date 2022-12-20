FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /api


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /api /api
COPY .env /

EXPOSE 5000

USER nonroot:nonroot

ENTRYPOINT [ "/api" ]