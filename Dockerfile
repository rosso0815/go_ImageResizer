# FROM golang:1-alpine as builder
# RUN apk --no-cache --no-progress add git ca-certificates tzdata make \
#     && update-ca-certificates \
#     && rm -rf /var/cache/apk/*
# WORKDIR /go/whoami
# # Download go modules
# COPY go.mod .
# COPY go.sum .
# RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download
# COPY . .
# RUN make build
# # Create a minimal container to run a Golang static binary
# FROM scratch
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /go/whoami/whoami .
# ENV PORT 90
# ENTRYPOINT ["/whoami" ]
# EXPOSE $PORT

FROM  centos:7.7.1908

ENV GOLANG_VERSION 1.13.4
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 702ad90f705365227e902b42d91dd1a40e48ca7f67a2f4b2fd052aaa4295cd95
ENV GOROOT /usr/local/go
ENV PATH $GOROOT/bin:$PATH

#RUN yum update -y
#RUN yum install -y ImageMagick-devel                                          
#RUN yum install -y wget
#RUN yum install -y tar                                              
# RUN yum install -y g++                                              
RUN yum install -y gcc                                              
# RUN yum install -y libc6-dev                                        
#RUN yum install -y make                                             
RUN yum install -y git                                              
#RUN yum install -y python-setuptools                                
#RUN easy_install supervisor                                        
#RUN mkdir -p /var/log/supervisor                                   
#RUN echo_supervisord_conf > /etc/supervisord.conf

RUN yum install -y http://rpms.famillecollet.com/enterprise/remi-release-7.rpm
RUN yum --enablerepo remi install -y ImageMagick7-devel

#RUN yum update -y

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz
#RUN echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz"
# | sha256sum -c - 
RUN tar -C /usr/local -xzf golang.tar.gz                           
RUN rm golang.tar.gz
RUN go version
#
# https://der-linux-admin.de/2016/06/centos-imagemagick-7-installieren/

#RUN  pkg-config --cflags --libs MagickWand

RUN go get -u gopkg.in/gographics/imagick.v3/imagick
COPY go.mod .
COPY go.sum .
# RUN GO111MODULE=on GOPROXY=https://proxy.golang.org go mod download
COPY . .
RUN go env

RUN ls -l

# RUN go test ./mygraphics/
RUN go build -o $GOROOT/bin/app .

#RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH" 
#RUN git clone https://github.com/docker-library/golang.git         
#RUN cd golang                                                      
#RUN sh update.sh                                                   
#RUN cp go-wrapper /usr/local/bin/