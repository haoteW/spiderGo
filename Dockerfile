# 使用本地 MySQL 镜像作为基础镜像
FROM mysql:latest

# 设置 MySQL root 用户密码
ENV MYSQL_ROOT_PASSWORD=123456

# 暴露 MySQL 端口
EXPOSE 3306
