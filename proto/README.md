### protobuf files

В реальном прокете этот пакет должен быть вынесен в отдельный репозиторий, доступный для всех приложений которые кго используют.

Компиляция находясь в целевом каталоге:

    protoc --go_out=plugins=grpc:. *.proto

### examples

    https://github.com/grpc/grpc-go/tree/master/examples
