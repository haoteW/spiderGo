# spiderGo

## douban
    douban 
    从电影排行榜首页开始，对其各类电影排行进行二次解析
    并获取排行榜异步接口，访问每个榜单top20电影
    结果是对各榜单top20电影做整理输出

## single spider
    单任务版爬虫 
    一个主执行引擎，通过访问-》解析次级目录-》爬虫任务插入任务队列
    为多级目录做对应解析器，并绑定对应链接，保证各模块独立功能

## concurrent spider
    并发版爬虫
    将单任务爬虫的主要工作（访问->解析）进行多协程并发
    并完成两种调度策略
        1、简单版：提供一个request channel 来对多协程进行调度
        2、队列版：在简单版基础上增加一个worker channel 
                用于监听空闲worker，拆分请求下发与worker监听的高度串行化
    
## docker
    docker run -d --name mysql-container -e MYSQL_ROOT_PASSWORD=123456  -p 3306:3306 mysql
    docker exec -it mysql-container mysql -u root -p

## 补充前端展现样式
![img.png](img.png)

