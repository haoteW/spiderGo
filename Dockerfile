# 使用本地 MySQL 镜像作为基础镜像
FROM mysql:latest

# 设置 MySQL root 用户密码
ENV MYSQL_ROOT_PASSWORD=123456

# 暴露 MySQL 端口
EXPOSE 3306

# 启动容器后执行的命令
CMD ["mysqld"]

# 授权 root 用户远程访问权限并创建名为 movieInfo 的数据库
RUN mysql -uroot -p123456 -e "GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION; FLUSH PRIVILEGES; CREATE DATABASE IF NOT EXISTS movieInfo;"

FROM nginx

COPY nginx.conf /etc/nginx/conf.d/spiderGo.conf
