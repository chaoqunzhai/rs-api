
## APP说明

## Api使用
```shell
1.项目初始化,数据库初始化
 export GOOS=darwin && export GOARCH=amd64 && go build
./go-admin migrate -c config/settings.yml

2.项目启动
./go-admin server -c config/settings.yml

3.生成migrate模板
./go-admin migrate -c config/settings.yml -g false -a true

4.创建app
./go-admin app -n test

```
## 初始化sql问题
1. 有时候MySQL不支持datetime为0的情况,设置db允许时间为空
```shell
[mysqld]
sql_mode=ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION

```
## 构建上传
测试服务器
sudo CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=2.1"  -o dynamic-store  && scp -P 26622 -i ~/.ssh/id_rsa -r dynamic-store root@152.136.36.253:/home/chaoqun

线上服务器
sudo CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-X main.Version=2.3"  -o dynamic-store  && scp -P 26622 -i ~/.ssh/id_rsa -r dynamic-store root@159.75.177.143:/home/chaoqun

### 项目关联
1. 获取本机公钥地址:/Users/zhaichaoqun/.ssh/id_rsa.pub
2. 在github settings中SSH keys / Add new
3. 在本机ssh-add /Users/zhaichaoqun/.ssh/id_rsa 添加私钥尝试访问github
4.  ssh git@github.com 测试是否可访问
### Systemd 方式启动:
```shell
cat > /etc/systemd/system/rs-api.service << "END"
[Unit]
Description=DcStore
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/data/rs-api
## 注:根据可执行文件路径修改
ExecStart=/data/rs-api/rs-api server -c config/settings.yml

# auto restart
StartLimitIntervalSec=0
Restart=always
RestartSec=1

[Install]
WantedBy=multi-user.target

END


systemctl daemon-reload

systemctl start rs-api.service
systemctl status rs-api.service
systemctl enable rs-api.service

```