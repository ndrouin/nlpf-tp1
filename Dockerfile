FROM golang
MAINTAINER ndrouin

RUN mkdir /go/src/project
RUN apt-get update -y
RUN apt-get install vim -y
WORKDIR /go/src/project

RUN go get -u github.com/kataras/iris/iris
RUN go get github.com/tools/godep

WORKDIR /go/src/project

RUN git clone https://github.com/ndrouin/nlpf-tp1.git 

RUN godep save

EXPOSE 80
