FROM ubuntu

ENV VERSION=v1.0

add ./httpsvr /httpsvr

ENTRYPOINT "/httpsvr"
