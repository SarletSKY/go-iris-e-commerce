### 安装erlang
wget -O- https://packages.erlang-solutions.com/ubuntu/erlang_solutions.asc | sudo apt-key add -
echo "deb https://packages.erlang-solutions.com/ubuntu bionic contrib" | sudo tee /etc/apt/sources.list.d/erlang-solution.list
sudo apt update
sudo apt install erlang

### 安装rabbitMQ
sudo apt-get install rabbitmq-server -y

### rabbitmq常见的命令
启动：sudo systemctl start rabbitmq-server
停止：sudo  rabbitmqctl stop
插件：sudo rabbitmqctl-plugins 

### 遇到的问题
1. 启动rabbitmq但localhost:15672访问不了的情况
https://blog.csdn.net/qq_52235571/article/details/131561365
解析:实质上就是没有去启动rabbitmq的插件