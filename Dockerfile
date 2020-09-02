FROM golang:alpine as build

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=goproxy.io \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app
COPY . /app

RUN go mod tidy

RUN go build -o web


# production stage
FROM alpine as production

WORKDIR /app


COPY --from=build /app/web /usr/bin/web
COPY --from=build /app/dist /app/dist

# 声明服务端口
EXPOSE 9933

ENTRYPOINT ["web"]
