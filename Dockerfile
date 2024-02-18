# 使用本地 MySQL 镜像作为基础镜像
FROM mysql:5.7 AS mysql-builder

# 设置 MySQL root 用户密码
ENV MYSQL_ROOT_PASSWORD=123456

COPY init.sql /docker-entrypoint-initdb.d/

# 暴露 MySQL 端口
EXPOSE 3306

# 安装 MySQL 客户端工具
# RUN apt-get update && apt-get install -y mysql-client

CMD ["mysqld"]

FROM nginx AS nginx-builder

COPY nginx.conf /etc/nginx/conf.d/default.conf
