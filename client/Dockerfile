FROM node:alpine as builder
COPY . /build

WORKDIR /build
RUN npm run build

FROM caddy:alpine
WORKDIR serve
COPY Caddyfile .
Copy --from=builder /build/dist dist
ENTRYPOINT caddy run --config Caddyfile