FROM buildpack-deps:buster-scm

# docker build -t mydocker --build-arg PROXYADDR=http://172.30.160.1:1080 .
# cat /etc/resolv.conf | grep -oP '(?<=nameserver\ ).*'
ARG PROXYADDR=http://172.30.160.1:1080

ENV http_proxy $PROXYADDR

RUN export http_proxy=$PROXYADDR && \
		echo 'Acquire::http::Proxy "$PROXYADDR";' >> /etc/apt/apt.conf.d/proxy.conf && \
		echo 'Acquire::https::Proxy "$PROXYADDR";' >> /etc/apt/apt.conf.d/proxy.conf

RUN apt-get update

#-----------golang --------- https://learn.go.dev/
RUN wget -P /tmp https://dl.google.com/go/go1.15.6.linux-amd64.tar.gz && \
		tar -C /usr/local -xzf /tmp/go1.15.6.linux-amd64.tar.gz

ENV PATH /usr/local/go/bin:$PATH

#----------- protoc ----------- https://grpc.io/docs/protoc-installation/#install-using-a-package-manager
RUN apt-get install -y protobuf-compiler

#----------- node ----------- https://github.com/nodesource/distributions/blob/master/README.md#installation-instructions
# Using Ubuntu
RUN curl -sL https://deb.nodesource.com/setup_lts.x | bash - && \
		apt-get install -y nodejs



WORKDIR /root/workspace
