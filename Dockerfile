FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates
COPY ./static /app/static

ADD hyrox /usr/bin/hyrox

CMD ["hyrox"]