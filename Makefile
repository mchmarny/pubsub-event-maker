all: test
build:
	go build -o ./bin/eventmaker -v
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./bin/eventmaker-linux

clean:
	go clean
	rm -f ./bin/eventmaker

run:
	go run *.go --project ${GCP_PROJECT} --topic test-topic --sources 3 --freq 2s

