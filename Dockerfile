# syntax = docker/dockerfile:1.2

FROM cirrusci/flutter AS web-build
WORKDIR /app
COPY ./app .
RUN --mount=type=cache,target=/root/.pub-cache \
    flutter pub get && flutter build web

# get modules, if they don't change the cache can be used for faster builds
FROM golang AS base
ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
WORKDIR /src
COPY go.* .
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

# build th application
FROM base AS build
COPY . .
# copy built web files from flutter
COPY --from=web-build /app/build/web ./app/build/web
# temp mount all files instead of loading into image with COPY
# temp mount go build cache
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -ldflags="-w -s" -o /app/main ./cmd/dashboard/*.go

# Import the binary from build stage
FROM gcr.io/distroless/static:nonroot as prd
COPY --from=build /app/main /
# this is the numeric version of user nonroot:nonroot to check runAsNonRoot in kubernetes
USER 65532:65532
ENTRYPOINT ["/main"]
