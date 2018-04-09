FROM arm32v6/alpine:3.7
RUN apk --no-cache add ca-certificates
RUN apk update && apk add rsync
COPY irsync /
WORKDIR /
ENTRYPOINT ["/irsync"]