FROM golang
MAINTAINER ndrouin

RUN mkdir /go/src/project

#install vim
RUN apt-get update -y
RUN apt-get install vim -y
ADD vimrc.local /etc/vim
ADD colors /etc/vim/colors

#install iris framework
WORKDIR /go/src/project

RUN go get -u github.com/kataras/iris/iris
RUN go get github.com/tools/godep
RUN godep save

#get project files
WORKDIR /go/src/project

RUN git clone https://github.com/ndrouin/nlpf-tp1.git 


EXPOSE 80
