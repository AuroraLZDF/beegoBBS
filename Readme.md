Nginx配置：
```bash
server {
    listen       80;
    server_name  local.beego.com;

    charset utf-8;
    access_log  /var/log/nginx/local.beego.com.access.log;
    error_log  /var/log/nginx/local.beego.com.errors.log;

    location /(css|js|fonts|img)/ {
        access_log off;
        expires 1d;

        root /www/go/src/github.com/xxxxxxx/beegoBBS/static;
        try_files $uri @backend;
    }

    location / {
        try_files /_not_exists_ @backend;
    }

    location @backend {
        proxy_set_header X-Forwarded-For $remote_addr;
        proxy_set_header Host            $http_host;

        proxy_pass http://127.0.0.1:8080;
    }
}
```