FROM arm32v6/alpine
RUN apk --no-cache add ca-certificates
RUN apk update && apk add rsync
WORKDIR /
CMD ["/irsync"]