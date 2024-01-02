FROM golang:1.18 AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

# ENV GOARCH=amd64

# RUN go build -o /go/bin/app
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-go

# FROM gcr.io/distroless/base-debian11

# COPY --from=build /go/bin/app /app

EXPOSE 8081

# USER nonroot:nonroot

CMD ["/docker-go"]