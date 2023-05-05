FROM php:5.6-fpm

RUN sed -i '/deb.debian.org\/debian stretch/d' /etc/apt/sources.list
RUN sed -i '/security.debian.org\/debian-security stretch\/updates/d' /etc/apt/sources.list
RUN echo "deb http://archive.debian.org/debian/ stretch main" >> /etc/apt/sources.list
RUN echo "deb http://archive.debian.org/debian-security/ stretch/updates main" >> /etc/apt/sources.list
RUN echo "Acquire::Check-Valid-Until false;" > /etc/apt/apt.conf.d/99disablechecks

COPY ./docker/php/php.ini /usr/local/etc/php/

RUN set -xe; \
    apt-get update -yqq && \
    apt-get install -yqq \
    curl

WORKDIR /var/www
