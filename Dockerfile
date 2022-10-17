
FROM golang:1.19-alpine3.16 AS buildstg
#can be git commit sha in the build pipeline [CI/CD] pass it as build arg `docker build ... --build-arg REVISION_NUM=xxxx` 
ARG REVISION_NUM=""

WORKDIR /app
# Download necessary Go modules
# we can use vendoring mode to avoid downloading modules in case of offline builds are needed
COPY go.mod ./
COPY go.sum ./
RUN go mod download 

COPY . /app
# we can also use -mod=vendor instead if we don't want to download modules everytime we build 
RUN go build -o dups -ldflags "-X 'github.com/msolimans/dups/pkg/build.Revision=$REVISION_NUM#$(date -u +%Y-%m-%d)' -X 'github.com/msolimans/dups/pkg/build.Time=$(date)'" cmd/main.go
# we can use even smaller image like `scratch` which saves cost in registeries and more secured than alpine however it won't give access to shell, busybox can be used in cases where we need shell
FROM alpine
COPY --from=buildstg /app/dups /dups
ENTRYPOINT ["/dups"]