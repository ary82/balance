FROM golang:alpine3.20 as base

WORKDIR /app

COPY . .


RUN apk add npm
RUN npm install -g pnpm
RUN pnpm install

RUN apk add make

RUN make clean
RUN go mod download
RUN make build

FROM alpine:3.20
COPY --from=base /app/main /app/
COPY --from=base /app/static /app/static

WORKDIR /app

EXPOSE 8080

CMD [ "./main" ]
