## Ally - 基于gin开发的web服务框架

## Features
- excel数据导出
- 文件压缩
- 图像处理（压缩、剪切、合成）开发中...
- 配置文件热加载
- 后台管理功能 开发中...

## Installation
- 下载源码：git clone https://github.com/chenbinhua/ally.git
- 下载依赖：go mod tidy
- 运行项目：go run main.go

## 服务重启
- 进入当前目录：cd ally
- 运行命令: ./server.sh

## Configuration
### 配置文件
- config.yaml(本地)
- config_release.yaml(生产)
### 如何加载配置文件
- 配置文件应与执行文件同一目录
- 基于gin环境变量GIN_MODE匹配，例如GIN_MODE=release，则加载的配置文件为config_release.yaml，默认为config.yaml
### GIN_MODE如何配置生产模式
- vim ~/.bash_profile
- 输入 export GIN_MODE=release
- 使配置生效 source ~/.bash_profile  
### 配置文件内容
```yaml
app:
  name: app
  port: 8080
mysql:
  user: root
  password: *****
  host: 127.0.0.1
  port: 3306
  db: default
redis:
  host: 127.0.0.1
  port: 6379
  db: 1
```


## License
GNU General Public License v3.0
