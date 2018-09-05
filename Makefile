# compile proto bufers
proto_account:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. services/account/proto/account.proto

proto_adcampaign:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. services/adCampaign/proto/adCampaign.proto && \
	cd services/adCampaign/proto && ls *.pb.go | xargs -n1 -IX bash -c "sed -e '/int64/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"

proto: proto_account proto_adcampaign

# build docker images
build_web:
	docker build -t nanosapp/web:v0.0.1 services/web

build_account:
	docker build -t nanosapp/account:v0.0.1 services/account

build_adcampaign:
	docker build -t nanosapp/adcampaign:v0.0.1 services/adCampaign

# ensure dependencies
dep_web:
	cd services/web && dep ensure -update
dep_account:
	cd services/account && dep ensure -update
dep_adcampaign:
	cd services/adCampaign && dep ensure -update

dep: dep_web dep_account dep_adcampaign

git_push:
	git add . && git commit -m "fix" && git push

# deploy to Kubernetes
reg: 
	minikube docker-env
	eval $(minikube docker-env)
	
deploy_web:
	kubectl create -f deployments/web/deployment.yaml
	kubectl create -f deployments/web/service.yaml

deploy_account:
	kubectl create -f deployments/account/deployment.yaml
	kubectl create -f deployments/account/service.yaml

deploy_adcampaign:
	kubectl create -f deployments/adCampaign/deployment.yaml
	kubectl create -f deployments/adCampaign/service.yaml

deploy_dbs:
	kubectl create -f deployments/db/account/volume.yaml
	kubectl create -f deployments/db/account/deployment.yaml
	kubectl create -f deployments/db/account/service.yaml
	kubectl create -f deployments/db/adCampaign/deployment.yaml
	kubectl create -f deployments/db/adCampaign/service.yaml
	
deploy: deploy_dbs deploy_web deploy_account deploy_adcampaign

# clean unused docker images and containers
clean:
	@echo "Remove all non running containers"
	-docker rm `docker ps -q -f status=exited`
	@echo "Delete all untagged/dangling (<none>) images"
	-docker rmi `docker images -q -f dangling=true`
	
# all steps for 
web: git_push dep_web build_web deploy_web
	
account: proto_account git_push dep_account build_account deploy_account
	
adcampaign: proto_adcampaign git_push dep_adcampaign build_adcampaign deploy_adcampaign
	
all: deploy_dbs build_web deploy_web build_account deploy_account build_adcampaign deploy_adcampaign clean
	
