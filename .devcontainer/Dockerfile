# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.154.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}


ARG HTTP_PROXY $HTTP_PROXY
ARG HTTPS_PROXY $HTTPS_PROXY

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
# RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
#     && apt-get -y install --no-install-recommends protobuf-compiler
RUN apt-get update

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN go get -x   google.golang.org/protobuf/cmd/protoc-gen-go \
                google.golang.org/grpc/cmd/protoc-gen-go-grpc

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g @tarojs/cli" 2>&1


# --------------- custom ------

    # npm install -g cnpm --registry=https://registry.npm.taobao.org && \
RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && \
        npm install -g @tarojs/cli \
        " 2>&1

ENV GO111MODULE on

RUN curl www.google.com && echo $HTTPS_PROXY

RUN go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

RUN git clone --depth 1 https://github.com/googleapis/googleapis.git $GOPATH/src/github.com/googleapis/googleapis

# ENV GOPROXY=https://goproxy.cn,direct

ARG PROTOC_VERSION=3.14.0
ARG PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip

RUN su vscode -c "curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP} && \
    unzip $PROTOC_ZIP -d ~/.local && \
    rm -f $PROTOC_ZIP"

# ENV PATH="$PATH:~/.local/bin"


ENV http_proxy ${HTTP_PROXY}
ENV https_proxy ${HTTPS_PROXY}

RUN git config --global http.proxy ${HTTP_PROXY} && \
    git config --global https.proxy ${HTTPS_PROXY}