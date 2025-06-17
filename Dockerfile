FROM ubuntu:22.04
ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update && \
    apt-get install -y curl git make wget postgresql-client

RUN curl -OL https://go.dev/dl/go1.24.3.linux-amd64.tar.gz && \
  tar -C /usr/local -xzf go1.24.3.linux-amd64.tar.gz && \
  rm go1.24.3.linux-amd64.tar.gz

ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app
COPY . .
RUN go mod download
CMD ["go", "test", "-v", "./..."]
