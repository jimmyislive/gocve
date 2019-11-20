FROM jimmyislive/goubuntu:1.0.2

RUN apt-get update && \
    apt-get install -y build-essential sqlite3 postgresql-client
RUN useradd -ms /bin/bash gouser
USER gouser
WORKDIR /home/gouser
