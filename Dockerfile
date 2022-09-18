FROM golang:1.17.5-alpine3.15 as stage

WORKDIR /usr/app

COPY . .

RUN go mod init meli 
RUN go mod tidy
RUN go build -o meli
RUN go build -o menu --tags=menu

FROM alpine
RUN apk update && apk add bash 
RUN apk add --no-cache libc6-compat && \ 
    apk add build-base 

WORKDIR /usr/app 
COPY --from=stage /usr/app/meli .
COPY --from=stage /usr/app/menu .

RUN chmod +x menu
RUN chmod +x meli
RUN crontab -l | { cat; echo "0 */1 * * *   /usr/app/meli"; } | crontab -
CMD ["crond","-f"] 