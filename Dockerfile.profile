FROM scratch

COPY gopath/bin/profile /profile

ENTRYPOINT ["/profile"]
