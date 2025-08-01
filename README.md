//运行单元测试
go test -v ./go-blog-project/database/gorm

//运行服务器并登陆网页端
cd ./go-blog-project
go run .

//开启服务器后，网页登陆注册JWT
http://localhost:5678/login

//单元测试结果
![alt text](image.png)
//登陆注册结果
![alt text](image2.png)