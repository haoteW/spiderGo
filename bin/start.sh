docker build --target mysql-builder -t my-mysql-image .
docker build --target nginx-builder -t my-nginx .

docker run -d --name mysql-container -p 3306:3306 my-mysql-image
docker run -d --name my-nginx-container -p 80:80 my-nginx

docker start my-nginx-container
docker start mysql-container

nohup go run main.go > logs/main.log 2>&1 &
nohup go run frontend/starter.go > logs/starter.log 2>&1 &
cd ./frontend/spider_front
cnpm install
nohup  cnpm run dev > ../../logs/frontend.log 2>&1 &
cd -