FROM golang:1.14.1-stretch

# This is the release of Vault to pull in.
ARG VAULT_VERSION=1.3.4

# Create a vault user and group first so the IDs get set the same way,
# even as the rest of this may change over time.
RUN addgroup vault && \
    useradd -g vault vault

RUN ln -s /bin/bash /bin/sh.bash && mv /bin/sh.bash /bin/sh

# Set up certificates, our base tools, and Vault.
RUN set -eux; \
    mkdir -p $GOPATH/src $GOPATH/bin $GOPATH/pkg && \
    chmod 777 $GOPATH/src $GOPATH/bin $GOPATH/pkg && \
    apt-get update -y && apt-get install -y ca-certificates libcap2-bin && \
    mkdir -p $GOPATH/src/github.com/hashicorp/vault && \
    git clone https://github.com/wwanglsu/vault.git $GOPATH/src/github.com/hashicorp/vault && \
    cd $GOPATH/src/github.com/hashicorp/vault && \
    make bootstrap && \
    make dev && \
    ln -s /bin/bash /bin/sh.bash && mv /bin/sh.bash /bin/sh && \
    curl -o- https://raw.githubusercontent.com/creationix/nvm/v0.33.11/install.sh | bash && \
    export NVM_DIR="$HOME/.nvm" && \
    [ -s "$NVM_DIR/nvm.sh" ] && \. "$NVM_DIR/nvm.sh" && \
    nvm install 10.15.2 && \
    npm install --global yarn && \
    cd $GOPATH/src/github.com/hashicorp/vault && make static-dist && \
    cd $GOPATH/src/github.com/hashicorp/vault && make dev-ui && \
    cp $GOPATH/src/github.com/hashicorp/vault/bin/vault $GOPATH/bin && \
    cd $HOME && \
    rm -rf $GOPATH/src/github.com $GOPATH/src/golang.org $GOPATH/pkg $HOME/.nvm/.cache && \
    mkdir -p /home/vault && \
    touch /home/vault/.vault-token && \
    chmod 777 /home/vault/.vault-token && \
    curl -o /usr/local/bin/su-exec.c https://raw.githubusercontent.com/ncopa/su-exec/master/su-exec.c; \
    \
    fetch_deps='gcc libc-dev'; \
    apt-get update; \
    apt-get install -y --no-install-recommends $fetch_deps; \
    rm -rf /var/lib/apt/lists/*; \
    gcc -Wall \
        /usr/local/bin/su-exec.c -o/usr/local/bin/su-exec; \
    chown root:root /usr/local/bin/su-exec; \
    chmod 0755 /usr/local/bin/su-exec; \
    rm /usr/local/bin/su-exec.c; \
    \
    apt-get purge -y --auto-remove $fetch_deps

# /vault/logs is made available to use as a location to store audit logs, if desired;
# /vault/file is made available to use as a location with the file storage backend, if desired;
# the server will be started with /vault/config as the config directory so you can add additional config files there.
RUN mkdir -p /vault/logs && \
    mkdir -p /vault/file && \
    mkdir -p /vault/config && \
    chown -R vault:vault /vault

# Expose the logs directory as a volume since there's potentially long-running state in there
VOLUME /vault/logs

# Expose the file directory as a volume since there's potentially long-running state in there
VOLUME /vault/file

# 8200/tcp is the primary interface that applications use to interact with Vault.
EXPOSE 8200

# The entry point script uses dumb-init as the top-level process to reap any
# zombie processes created by Vault sub-processes.
#
# For production derivatives of this container, you shoud add the IPC_LOCK
# capability so that Vault can mlock memory.
#COPY config/local.json /vault/config/local.json
#ENTRYPOINT ["vault"]

COPY docker-entrypoint.sh /usr/local/bin/docker-entrypoint.sh
ENTRYPOINT ["bash", "/usr/local/bin/docker-entrypoint.sh"]

# By default you'll get a single-node development server that stores everything
# in RAM and bootstraps itself. Don't use this configuration for production.
CMD ["server", "-dev"]