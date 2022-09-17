gen :
	protoc demo.proto --go_out=plugins=grpc:.

clear:
	rm proto/*.go