.PHONY: build clean deploy

build:
	env GOOS=linux go build -ldflags="-s -w" -o bin/checker functions/checker/checker.go
	# env GOOS=linux go build -ldflags="-s -w" -o bin/notifier functions/notifier/notifier.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/notifier-slack functions/notifier-slack/notifier-slack.go

clean:
	rm -rf ./bin

deploy: clean build
	sls deploy --verbose
