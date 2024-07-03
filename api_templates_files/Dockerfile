# 1st stage, build the app
FROM golang:1.22-bookworm as builder

WORKDIR /builder
RUN mkdir -p bin/output
COPY go.mod go.sum ./

# Ensure go.mod and go.sum are tidy and download dependencies
RUN go mod tidy
RUN go mod download
RUN go fmt ./...

COPY . .

RUN pwd
RUN ls
RUN go build -ldflags="-s -w" -o bin/__MODULE_NAME__ ./cmd/__MODULE_NAME__

# 2nd stage, build the runtime image
FROM golang:1.22-bookworm

LABEL maintainedBy="cdcloud-io"

ARG _APP_NAME=""
ARG _APP_VERSION=""
ARG _APP_COMMIT_SHA=""
ARG _APP_BUILD_ID=""
ARG _APP_BUILD_DATE=""

# Set by pipeline for docker build
### _APP_ENV=<passed in as env var during container startup>
ENV _APP_NAME=${_APP_NAME}
ENV _APP_VERSION=${_APP_VERSION}
ENV _APP_COMMIT_SHA=${_APP_COMMIT_SHA}
ENV _APP_BUILD_ID=${_APP_BUILD_ID}
ENV _APP_BUILD_DATE=${_APP_BUILD_DATE}
ENV _APP_ENTRYPOINT=/opt/app/

RUN echo 'app-sa:x:1001:' >> /etc/group && \
  echo 'app-sa:x:1001:1001::/home/app-sa:/bin/bash' >> /etc/passwd && \
  mkdir -p /home/app-sa && \
  mkdir -p /opt/app/config && \
  chown -R app-sa: /home/app-sa && \
  chown -R app-sa: /opt/app

WORKDIR /opt/app
COPY --from=builder /builder/bin/__MODULE_NAME__ .

ENTRYPOINT [ "/opt/app/__MODULE_NAME__" ]
