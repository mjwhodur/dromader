# This is MakeFile for dromader
PWD := $(pwd)
.PHONY: all
all: server_image node_image app

.PHONY: server_image
server_image:

.PHONY: node_image
node_image:
	cd containers/node && podman build .

.PHONY: app
app:
	cd dromader/cmd/dromader && go build
