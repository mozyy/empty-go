# See here for image contents: https://github.com/microsoft/vscode-dev-containers/tree/v0.154.2/containers/go/.devcontainer/base.Dockerfile

# [Choice] Go version: 1, 1.15, 1.14
ARG VARIANT="1"
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# [Option] Install Node.js
ARG INSTALL_NODE="true"
ARG NODE_VERSION="lts/*"
RUN if [ "${INSTALL_NODE}" = "true" ]; then su vscode -c "umask 0002 && . /usr/local/share/nvm/nvm.sh && nvm install ${NODE_VERSION} 2>&1"; fi

# [Optional] Uncomment this section to install additional OS packages.
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends protobuf-compiler

# [Optional] Uncomment the next line to use go get to install anything else you need
RUN go get -x   google.golang.org/protobuf/cmd/protoc-gen-go \
                google.golang.org/grpc/cmd/protoc-gen-go-grpc

# [Optional] Uncomment this line to install global node packages.
# RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && npm install -g @tarojs/cli" 2>&1


# --------------- custom ------


RUN su vscode -c "source /usr/local/share/nvm/nvm.sh && \
        npm install -g cnpm --registry=https://registry.npm.taobao.org && \
        cnpm install -g @tarojs/cli \
        " 2>&1

ENV http_proxy=$HTTP_PROXY
ENV https_proxy=$HTTP_PROXY
ENV GO111MODULE=on

RUN echo $http_proxy && echo $GO111MODULE && go get github.com/asim/go-micro/cmd/protoc-gen-micro/v3

RUN git clone --depth 1 https://github.com/googleapis/googleapis.git $GOPATH/src/github.com/googleapis/googleapis

RUN go env -w GOPROXY=https://goproxy.cn,direct
