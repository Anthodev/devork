ARG GO_VERSION=1.18
ARG CADDY_VERSION=2

FROM golang:${GO_VERSION}
ARG USERNAME=dev
ARG USER_UID=1000
ARG USER_GID=1000

RUN apt update \
    && apt upgrade -y \
    && apt install -y git bash net-tools

# Setup user
RUN addgroup --gid ${USER_GID} ${USERNAME}
RUN adduser $USERNAME --home /home/$USERNAME --disabled-password --uid $USER_UID --gid $USER_GID && \
    mkdir -p /etc/sudoers.d && \
    echo $USERNAME ALL=\(root\) NOPASSWD:ALL > /etc/sudoers.d/$USERNAME && \
    chmod 0440 /etc/sudoers.d/$USERNAME

# Setup shell
USER $USERNAME

WORKDIR /usr/app

COPY . .