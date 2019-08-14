all: dailymotion_updater

dailymotion_updater: main.go dm/ChannelList.go dm/dmapi.go vidispine/MetadataFieldGroup.go vidispine/PortalExtraFieldData.go vidispine/VSCommunicator.go vidispine/VSFieldGroup.go
	go build

dailymotion_updater.linux64: dailymotion_updater
	GOOS=linux GOARCH=amd64 go build -o dailymotion_updater.linux64

docker: dailymotion_updater.linux64 Dockerfile
	./build-docker-image.sh
    
test:
	go test ./...

clean:
	rm -f dailymotion_updater dailymotion_updater.linux64
