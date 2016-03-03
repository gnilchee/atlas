#
# Atlas - A simple Go based http server
#
FROM debian:jessie-backports

MAINTAINER https://github.com/gnilchee/atlas

# Install Go
RUN \
  mkdir -p /etc/atlas/ssl && \
  mkdir -p /goroot && mkdir -p /gopath/static && \
  apt-get -y update && apt-get -y upgrade && apt-get -y install curl && \
  curl https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz | tar xvzf - -C /goroot --strip-components=1

# Set environment variables.
ENV GOROOT /goroot
ENV GOPATH /gopath
ENV PATH $GOROOT/bin:$GOPATH/bin:$PATH

ADD atlas.go /gopath/
ADD certfile.crt /etc/atlas/ssl/
ADD keyfile.key /etc/atlas/ssl/

VOLUME /gopath/static

WORKDIR /gopath

EXPOSE 4443

CMD go run atlas.go
