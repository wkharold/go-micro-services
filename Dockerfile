FROM scratch

COPY gopath/bin/api /api

ENTRYPOINT ["/api"]
