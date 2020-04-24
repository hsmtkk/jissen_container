FROM ubuntu:20.04

RUN apt upddate \
 && apt install -y cron \
 && apt clean \
 && rm -rf /var/lib/apt/lists/*

COPY ./task.sh /usr/local/bin
COPY ./cron-example /etc/cron.d
RUN chmod 0644 /etc/cron.d/cron-example

ENTRYPOINT ["cron", "-f"]
