FROM alpine:3.19 as build
WORKDIR /build/
COPY . ./
RUN apk add npm go \
    && npm ci && npm run build \
    && go build -o=plateful ./

FROM alpine:3.19
WORKDIR /data/
COPY --from=build /build/plateful /
ENTRYPOINT ["/plateful"]
