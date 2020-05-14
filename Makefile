.PHONY : help
help : Makefile
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

##
##- On host computer ----------------------------------------------------------------------------------------

##
##-- Docker

.PHONY : install
build:                                             ## Installs all dependencies.
	@bash ./scripts/install.sh

.PHONY : build
build:                                             ## Builds docker images.
	@bash ./scripts/build.sh

.PHONY : start
start:                                             ## Starts entire setup.
	@bash ./scripts/start.sh

.PHONY : stop
stop:                                              ## Stops entire setup.
	@bash ./scripts/stop.sh

.PHONY : ssh
ssh:                                               ## SSH into dev container.
	@bash ./scripts/ssh.sh

.PHONY : ssh-chaincode
ssh-chaincode:                                     ## SSH into chaincode container.
	docker exec -it chaincode sh

.PHONY : ssh-cli
ssh-cli:                                           ## SSH into cli container.
	docker exec -it cli bash


##
##- Inside containers ---------------------------------------------------------------------------------------

##
##-- simplequeuedev


.PHONY : format
format:                                            ## Fix go code formatting.
	go fmt ./...

.PHONY : test
test:                                              ## Stops entire setup.
	@bash ./scripts/test.sh

.PHONY : generate-mocks
generate-mocks:                                    ## Generates mocks for contracts.
	@bash ./scripts/generate-mocks.sh

##
##-- chaincode

## * nothing available *

##
##-- cli

## * nothing available *