gen:
	protoc -I=proto/ \
	--go_out=. --go_opt=module=github.com/MobileStore-Grpc/product \
	--go-grpc_out=. --go-grpc_opt=module=github.com/MobileStore-Grpc/product \
	--grpc-gateway_out=. --grpc-gateway_opt=module=github.com/MobileStore-Grpc/product \
	--openapiv2_out=swagger \
	proto/*.proto proto/google/api/*.proto

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go --port 8080

rest:
	go run cmd/server/main.go --port 8082 --type rest --endpoint 0.0.0.0:8080

client:
	go run cmd/client/main.go --address 0.0.0.0:8080

build-image:
	docker build -t mobilestore-product:v1.0.0 .

run:
	docker run -d --name product -p 9080:8080 mobilestore-product:v1.0.0