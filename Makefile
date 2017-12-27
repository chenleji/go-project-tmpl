default:
	env GOARCH=amd64 GOOS=linux go build

build:
	docker build -t registry.cn-hangzhou.aliyuncs.com/ljchen/go-project-tmpl:v1.0 -f Dockerfile .

push:
	docker push registry.cn-hangzhou.aliyuncs.com/ljchen/go-project-tmpl:v1.0

release: default build push


build_debug:
	docker build -t registry.cn-hangzhou.aliyuncs.com/ljchen/go-project-tmpl:latest -f Dockerfile .

push_debug:
	docker push registry.cn-hangzhou.aliyuncs.com/ljchen/go-project-tmpl:latest

all: default build_debug push_debug
