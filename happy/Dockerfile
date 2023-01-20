FROM golang:1.18-alpine AS build

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# 모듈 카피
COPY go.mod .

# 의존성 다운로드 (모듈 변화가 없다면 패스)
RUN go mod download

# 소스 전체 카피
COPY . .

RUN go build -o main main.go

FROM alpine:latest

# 타임존 설정
RUN sed 's/https/http/g' -i /etc/apk/repositories
RUN apk update
RUN apk add --no-cache tzdata
ENV TZ=Asia/Seoul
ARG DISABLE_CACHE
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
 && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 8000

CMD ["/app/main"]
