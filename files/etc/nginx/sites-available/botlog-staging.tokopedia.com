server {
    server_name  botlog-staging.tokopedia.com;
    listen 80;

    access_log /var/log/nginx/botlog.access.log main;
    error_log  /var/log/nginx/botlog.error.log;

    root /var/www/botlog/public;

    location / {
        proxy_pass http://127.0.0.1:8910;
    }
}