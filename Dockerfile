FROM golang:1.9-alpine

COPY . /go/src/fast-text-server

WORKDIR /go/src/fast-text-server

RUN \
    apk update && \
    apk add bash make g++ ca-certificates openssl && \
    update-ca-certificates  && \
    wget https://github.com/facebookresearch/fastText/archive/master.zip && \
    unzip master.zip && \
    cd fastText-master && make && cp fasttext /usr/local/bin && cd -  && \
    go build

EXPOSE 8080

CMD ./fast-text-server --model=$MODEL



