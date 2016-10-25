FROM golang
MAINTAINER ndrouin

RUN mkdir /go/src/project

#install iris framework
WORKDIR /go/src/project

RUN go get -u github.com/kataras/iris/iris
RUN go get github.com/tools/godep
RUN godep save

#get project files
WORKDIR /go/src/project

RUN git clone https://github.com/ndrouin/nlpf-tp1.git 


EXPOSE 80
