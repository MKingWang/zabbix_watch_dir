GOCMD=go
GOBUILD=$(GOCMD) build 
BIN_SERVER=notifyServer
BIN_CLIENT=notifyCli

all:ser cli 

ser:
	$(GOBUILD) -o $(BIN_SERVER)

cli:
	$(GOBUILD)  -o $(BIN_CLIENT) client/*.go

clean:
	/bin/rm -rf $(BIN_CLIENT) $(BIN_SERVER)