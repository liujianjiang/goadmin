### jenkeins
docker pull jenkins/jenkins
mkdir -p /data/jenkins_home/
chown -R 1000:1000 /data/jenkins_home/
docker run -d -p 10240:8080 -p 10241:50000 -v /www/jenkins_data:/var/jenkins_home -v /etc/localtime:/etc/localtime --name myjenkins jenkins/jenkins

### 启动
docker-compose up
docker-compose up --build

