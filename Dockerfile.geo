FROM golang

COPY gopath/bin/geo /geo
COPY data/geo.json /data/geo.json

ENTRYPOINT ["/geo"]
