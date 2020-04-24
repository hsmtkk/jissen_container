FROM ubuntu:20.04

RUN apt update \
 && apt install -y cron \
 && apt clean \
 && rm -rf /var/lib/apt/lists/*

COPY ./build/cronjob/task.sh /usr/local/bin
COPY ./build/cronjob/cron-example /etc/cron.d
RUN chmod 0644 /etc/cron.d/cron-example

ENTRYPOINT ["cron", "-f"]
