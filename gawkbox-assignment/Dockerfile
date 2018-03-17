FROM golang

COPY . usr/src/gawkbox-assignment

WORKDIR /usr/src/gawkbox-assignment

RUN go build .

ENTRYPOINT [ "./gawkbox-assignment" ]