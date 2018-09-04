	# compile proto bufers
proto_account:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. services/account/proto/account.proto

proto_adcampaign:
	protoc --proto_path=$(GOPATH)/src:. --micro_out=. --go_out=. services/adCampaign/proto/adCampaign.proto && \
	cd services/adCampaign/proto && ls *.pb.go | xargs -n1 -IX bash -c 'sed s/,omitempty// X > X.tmp && mv X{.tmp,}'

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

git:
	git add . && git commit -m "fix" && git push
