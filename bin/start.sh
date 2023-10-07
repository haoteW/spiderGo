docker build -t my-mysql-image .
#docker start mysql-container
docker run -d --name mysql-container -p 3306:3306 my-mysql-image

nohup go run main.go > logs/main.log 2>&1 &
nohup go run frontend/starter.go > logs/starter.log 2>&1 &
nohup go run frontend/frontend.go > logs/frontend.log 2>&1 &