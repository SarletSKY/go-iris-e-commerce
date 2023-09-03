### 安装erlang
wget -O- https://packages.erlang-solutions.com/ubuntu/erlang_solutions.asc | sudo apt-key add - <br />
echo "deb https://packages.erlang-solutions.com/ubuntu bionic contrib" | sudo tee /etc/apt/sources.list.d/erlang-solution.list <br />
sudo apt update <br />
sudo apt install erlang <br />

### 安装rabbitMQ
sudo apt-get install rabbitmq-server -y

### rabbitmq常见的命令
启动：sudo systemctl start rabbitmq-server <br />
停止：sudo  rabbitmqctl stop <br />
插件：sudo rabbitmqctl-plugins 

### rabbitmq 五种生产模式
1.简单模式 1对1  <br />
2.工作模式，1对多个消费者，消费者负责顺序输出 <br />
3.订阅模式，也是1对多个消费者，但提供了交换机，对每个消费者进行输入，其中队列要绑定交换机(交换机类型广播) <br />
4.路由模式，在订阅模式的基础上，可以指定发送到哪个消费者(交换机用direct)指定数值用到key <br />
5.话题模式，和订阅模式不同的是，可以通过规则来接收哪些数据，(交换机用topic) [#与*] <br />


### 遇到的问题
1. 启动rabbitmq但localhost:15672访问不了的情况 
https://blog.csdn.net/qq_52235571/article/details/131561365
解析:实质上就是没有去启动rabbitmq的插件