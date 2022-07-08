### jenkeins
docker pull jenkins/jenkins
mkdir -p /data/jenkins_home/
chown -R 1000:1000 /data/jenkins_home/
docker run -d -p 10240:8080 -p 10241:50000 -v /www/jenkins_data:/var/jenkins_home -v /etc/localtime:/etc/localtime --name myjenkins jenkins/jenkins

### 前端编译
cd web
sudo yarn && yarn build

### 前端发布
git pull

### 后端发布
docker-compose up
docker-compose up --build admin-server chromedp-headless-shell

### mysql无法连接
```
方法一：关闭防火墙
centos关闭防火墙的操作为
systemctl stop firewalld
方法二： 在防火墙上开发指定端口
firewall-cmd --zone=public --add-port=2181/tcp --permanent
firewall-cmd --reload
```

