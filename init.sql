-- 创建名为 movieInfo 的数据库
CREATE DATABASE IF NOT EXISTS movieInfo;

-- 授权 root 用户远程访问权限
GRANT ALL PRIVILEGES ON *.* TO 'root'@'%' IDENTIFIED BY '123456' WITH GRANT OPTION;

-- 刷新权限
FLUSH PRIVILEGES;
