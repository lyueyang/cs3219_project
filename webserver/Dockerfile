# syntax=docker/dockerfile:1

FROM nginx:1.21.1-alpine

WORKDIR /app
RUN rm /etc/nginx/conf.d/default.conf

COPY location.conf /etc/nginx/conf.d/