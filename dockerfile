FROM golang:alpine
RUN apk add --update nodejs nodejs-npm bash git make
RUN mkdir -p /go/src/github.com/pschatzmann/tincwebgui
WORKDIR /go/src/github.com/pschatzmann/tincwebgui
RUN git clone  git://github.com/pschatzmann/tincwebgui /go/src/github.com/pschatzmann/tincwebgui
RUN make
