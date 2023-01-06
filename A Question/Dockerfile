FROM golang:1.19

ENV HTTPS_PROXY="http://fodev.org:8118"

WORKDIR /app

COPY . /tmp
RUN bash /tmp/build.sh
RUN rm -rf /tmp/*

CMD [ "sh", "/app/run.sh" ]