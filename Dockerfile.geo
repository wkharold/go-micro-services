FROM scratch

COPY gopath/bin/geo /geo

ENTRYPOINT ["/geo"]
