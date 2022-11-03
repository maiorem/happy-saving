FROM alpine

# 타임존 설정
RUN apk add --no-cache tzdata
ENV TZ=Asia/Seoul
ARG DISABLE_CACHE
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime \
 && echo $TZ > /etc/timezone


CMD go run main.go