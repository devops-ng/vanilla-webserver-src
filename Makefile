
# Go parameters
BINARY_NAME=deployment
all: build package clean
build: 
	cd server && GOOS=linux go build -o $(BINARY_NAME)
package:
	cd server && zip $(BINARY_NAME).zip $(BINARY_NAME)
clean:
	cd server && rm $(BINARY_NAME)