.PHONY: all build zip clean

# variables
OUTPUT=build/bootstrap
ZIP=build/serverless-user-service.zip
SRC=src/cmd/user/main.go
GOOS=linux
GOARCH=amd64

# main command
all: clean build zip

# build file
build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(OUTPUT) $(SRC)
	chmod +x $(OUTPUT)

# zip files
zip: build
	zip -j $(ZIP) $(OUTPUT)

# remove created files
clean:
	rm -f $(OUTPUT) $(ZIP)
