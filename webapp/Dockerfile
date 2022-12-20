FROM golang:1.19 AS build

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /webapp


FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /webapp /webapp
COPY . /

EXPOSE 3000

USER nonroot:nonroot

ENTRYPOINT [ "/webapp" ]