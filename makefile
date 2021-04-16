install:
	go get -t -d github.com/tebeka/selenium && \
	go get -t -d gopkg.in/yaml.v2 && \
	go get -t -d github.com/namsral/flag 
	
install-openjdk:
	sudo apt-get install xvfb openjdk-11-jre

run:
	#go run main.go --alsologtostderr  --download_browsers --download_latest
	go run main.go

build-ci: build
	sh ci/build.sh


.PHONY: build
build: main.go
	CGO_ENABLED=0 GOOS=linux \
    go build -installsuffix cgo -o goinsta $<

release: build
	- upx -9 goinsta