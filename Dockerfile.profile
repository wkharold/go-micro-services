FROM scratch

COPY gopath/bin/profile /profile
COPY data/hotels.json /data/hotels.json

ENTRYPOINT ["/profile"]
