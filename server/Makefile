generate_grpc_code:
	protoc \
    --go_out=internal/game \
    --go_opt=paths=source_relative \
    --go-grpc_out=internal/game \
    --go-grpc_opt=paths=source_relative \
    internal/game/game.proto