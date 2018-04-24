PID      = /tmp/awesome-project-watcher.pid
GO_FILES = $(wildcard *.go)
APP      = ./
serve: restart
	@fswatch -o . | xargs -n1 -I{}  make restart || make kill

kill:
	@if [ -s $(PID) ]; then kill `cat $(PID)`; echo "" > $(PID); fi

before:
	@echo "\n\n Building go code ... \n\n"
restart: kill before
	@go build server.go && ./server & echo $$! > $(PID)


.PHONY: serve restart kill before # let's go to reserve rules names
