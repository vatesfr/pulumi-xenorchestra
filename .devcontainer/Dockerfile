ARG UBUNTU_VERSION=latest
FROM --platform=linux/amd64  ubuntu:${UBUNTU_VERSION}

# Update apt-get and install various needed utilities
RUN apt-get update && \
    apt-get install -y curl && \
    apt-get install -y wget && \
    apt-get install -y xz-utils && \
    apt-get install -y make && \
    apt-get install -y gcc && \
    apt-get install -y git

# Install bridged provider prerequisites
# See README.md

# Install go
ARG GO_VERSION=1.23.4
RUN rm -rf /usr/local/go && \
    wget -O ${GO_VERSION}.tar.gz https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf ${GO_VERSION}.tar.gz && \
    rm ${GO_VERSION}.tar.gz

ENV GOPATH=/root/go
ENV PATH=$PATH:/usr/local/go/bin

# Install go linter
RUN mkdir -p $GOPATH/bin && \
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $GOPATH/bin v1.61.0

ENV PATH=$PATH:$GOPATH/bin

# Install pulumictl
ARG PULUMICTL_VERSION=v0.0.46
RUN rm -rf /usr/local/bin/pulumictl && \
    wget -O pulumictl.${PULUMICTL_VERSION}.tar.gz https://github.com/pulumi/pulumictl/releases/download/${PULUMICTL_VERSION}/pulumictl-${PULUMICTL_VERSION}-linux-amd64.tar.gz && \
    tar -C /usr/local/bin -xzf pulumictl.${PULUMICTL_VERSION}.tar.gz

# Install python and related items
RUN apt-get install -y python3 && \
    apt-get install -y python3-setuptools

# Install Pulumi
RUN curl -fsSL https://get.pulumi.com | sh
ENV PATH=$PATH:/root/.pulumi/bin
