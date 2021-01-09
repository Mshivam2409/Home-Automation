FROM ubuntu

RUN apt update

RUN wget https://golang.org/dl/go1.15.6.linux-amd64.tar.gz && \
    tar -C /usr/local -xzf go1.15.6.linux-amd64.tar.gz && \
    export PATH=$PATH:/usr/local/go/bin

RUN apt install -y software-properties-common && \
    add-apt-repository ppa:deadsnakes/ppa && \
    apt update && \
    apt install -y python3.8

RUN apt-get install -y gnupg && \
    wget -qO - https://www.mongodb.org/static/pgp/server-4.4.asc | apt-key add - && \
    echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu focal/mongodb-org/4.4 multiverse" | tee /etc/apt/sources.list.d/mongodb-org-4.4.list && \
    apt-get install -y mongodb-org

RUN wget -O- https://packages.erlang-solutions.com/ubuntu/erlang_solutions.asc | apt-key add - && \
    echo "deb https://packages.erlang-solutions.com/ubuntu focal contrib" | tee /etc/apt/sources.list.d/rabbitmq.list && \
    apt update && \
    apt install -y erlang apt-transport-https && \
    wget -O- https://dl.bintray.com/rabbitmq/Keys/rabbitmq-release-signing-key.asc | apt-key add - && \
    wget -O- https://www.rabbitmq.com/rabbitmq-release-signing-key.asc | apt-key add - && \
    echo "deb https://dl.bintray.com/rabbitmq-erlang/debian focal erlang-22.x" | tee /etc/apt/sources.list.d/rabbitmq.list && \
    apt update && \
    apt install -y rabbitmq-server

RUN apt install -y nginx