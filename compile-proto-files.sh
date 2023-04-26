for file in "$@"
do
    echo "Compiling $file"
    protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative "$file"
done