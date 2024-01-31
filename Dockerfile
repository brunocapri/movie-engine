FROM --platform=linux/amd64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

ADD movie-engine /usr/bin/movie-engine

CMD ["movie-engine"]
