FROM golang:1.14.2

RUN mkdir /out
RUN mkdir -p /go/src/github.com/jenkins-x-labs

WORKDIR /go/src/github.com/jenkins-x-labs

RUN git clone https://github.com/nuxeo/jxlabs-nos-jxl.git jxl && \
  cd jxl && \
  make linux && \
  mv build/linux/jxl /out/jxl

# jenkins-x
ENV JX_VERSION 2.1.3
RUN curl --silent -f -L https://github.com/nuxeo/jxlabs-nos-jx/releases/download/v${JX_VERSION}/jxlabs-nos-jx-linux-amd64.tar.gz | tar xvfCz - /out jx && \
  chmod +x /out/jx

FROM ${DOCKER_REGISTRY}/${DOCKER_REGISTRY_ORG}/jxl-base:0.0.3
EXPOSE 8080

COPY --from=0 /out /usr/local/bin

