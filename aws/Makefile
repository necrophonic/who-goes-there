.PHONY: build clean deploy remove

build:
	GOOS=linux go build -ldflags="-s -w" -o bin/checker functions/checker/checker.go
	GOOS=linux go build -ldflags="-s -w" -o bin/notifier-slack functions/notifier-slack/notifier-slack.go

clean:
	rm -rf ./bin

remove:
	sls remove --verbose

deploy: clean build
	sls deploy --verbose
