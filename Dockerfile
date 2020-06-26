FROM node:12-alpine as yarn-build

WORKDIR /app
ADD ./js-hello-world/ ./

RUN yarn install
RUN yarn build

FROM golang:1.14-alpine

RUN mkdir -p ./dist/

RUN apk add --no-cache git
RUN go get github.com/NYTimes/gziphandler

ADD ./main.go ./
RUN go build -o serve main.go
COPY --from=yarn-build /app/build/ ./dist/

CMD ["./serve"]
