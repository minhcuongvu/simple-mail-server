# Use an official Ubuntu image as a base
FROM ubuntu:22.04

# Set non-interactive frontend for package installations
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update && apt-get install -y \
    postfix \
    mailutils \
    libsasl2-modules \
    sasl2-bin \
    curl \
    telnet \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Install Go manually (or the latest stable version)
RUN curl -fsSL https://go.dev/dl/go1.23.2.linux-amd64.tar.gz -o go1.23.2.linux-amd64.tar.gz \
    && tar -C /usr/local -xzf go1.23.2.linux-amd64.tar.gz \
    && rm go1.23.2.linux-amd64.tar.gz
# Update PATH to include the new Go binary
ENV PATH="/usr/local/go/bin:${PATH}"

RUN mkdir /app
COPY . /app/
WORKDIR /app

RUN go mod tidy

# # Ensure the necessary directories for mail storage exist
# RUN mkdir -p /var/mail

# Expose port 25 for Postfix (SMTP)
EXPOSE 25

CMD ["./scripts/setup.sh"]
