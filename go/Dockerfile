FROM ubuntu

RUN apt-get update
RUN apt-get -y install python3.6
RUN apt-get install -y curl
RUN rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.13

RUN curl -sSL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz \
		| tar -v -C /usr/local -xz

ENV PATH /usr/local/go/bin:$PATH

RUN mkdir -p /go/src /go/bin && chmod -R 777 /go
ENV GOROOT /usr/local/go
ENV PATH /go/bin:$PATH
WORKDIR /go/src/github.com/codemirror

COPY . .

RUN env GOOS=linux go build main.go

ENTRYPOINT ./main