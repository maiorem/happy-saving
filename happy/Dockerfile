FROM golang:1.19-alpine AS build

WORKDIR /app

ENV GO111MODULE=on
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY . .

RUN go build -o main main.go

FROM alpine:latest

# 타임존 설정
RUN apk add --no-cache tzdata
ENV TZ=Asia/Seoul
ARG DISABLE_CACHE
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
 && echo $TZ > /etc/timezone

WORKDIR /app

COPY --from=build /app/main .

EXPOSE 80

CMD ["/app/main"]