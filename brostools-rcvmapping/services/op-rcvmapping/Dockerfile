FROM alpine:3.6

# nginxのインストール
RUN apk update && \
    apk add --no-cache nginx

WORKDIR .

# ドキュメントルート
ADD services/op-rcvmapping/html /html
ADD services/op-rcvmapping/default.conf /etc/nginx/conf.d/default.conf

ENV LANG=en_US.UTF-8
ENV LANGUAGE=en_US.UTF-8
ENV LC_ALL=en_US.UTF-8
ENV TZ=Asia/Tokyo
ENV PORT=8000

# ポート設定
EXPOSE 8000

RUN mkdir -p /run/nginx

# フォアグラウンドでnginx実行
CMD nginx -g "daemon off;"