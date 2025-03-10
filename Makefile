compile_grpc:
	protoc -I=api/grpc_consumer -I=internal/InputStreamShard -I=internal/base_rpc -I=api/client_stream_creator \
	  --go_out=. --go-grpc_out=. \
	  --go_opt=module=go_video_streamer \
	  --go-grpc_opt=module=go_video_streamer \
	  api/grpc_consumer/VideoRCV.proto \
	  api/client_stream_creator/ClientStreamCreator.proto \
	  internal/InputStreamShard/InputStreamShard.proto \
	  internal/base_rpc/RPCStructs.proto

build_gocv_base_image:
	go mod download
	GOCV_HOME=$$(go list -m -f '{{.Dir}}' gocv.io/x/gocv) && \
	cd "$$GOCV_HOME" && \
	docker build -f Dockerfile -t flowweaver:gocv-base .

build_flowweaver_grpc:
	docker build -f Dockerfile.grpc_local_mem -t flowweaver:grpc_local_mem .

swagger_http_rest_api:
	swag init -g cmd/http_rest_api/main.go -o ./docs

.PHONY: build_gocv_base_image compile_grpc