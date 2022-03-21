FROM nginx:alpine

COPY ./services/op-rcvmapping-nginx/conf/op-rcvmapping-nginx.conf /etc/nginx/conf.d/
COPY ./services/op-rcvmapping-nginx/conf/op-rcvmapping.txt /etc/nginx/conf.d/
COPY ./services/op-rcvmapping-nginx/conf/servername.txt /etc/nginx/conf.d/
COPY ./services/op-rcvmapping-nginx/conf/nginx.conf /etc/nginx/
ENV PORT 80
ENV HOST 0.0.0.0
ARG PORT=80
ARG HOST=0.0.0.0
EXPOSE 80