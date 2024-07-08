FROM node:alpine as node-builder
COPY client /build
WORKDIR /build
RUN npm run build

FROM golang:alpine as go-builder
COPY server /server
WORKDIR /server
RUN go build

FROM caddy:alpine
WORKDIR serve
COPY Caddyfile .
Copy --from=node-builder /build/dist dist
Copy --from=go-builder /server/po-key-mon .
ENTRYPOINT caddy run --config Caddyfile & ./po-key-mon