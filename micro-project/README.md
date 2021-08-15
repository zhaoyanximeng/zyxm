user生产protobuf文件: protoc --proto_path=./proto/user --micro_out=proto/user --go_out=proto/user user.proto

设置注册中心：
    set MICRO_REGISTRY=etcd
    set MICRO_REGISTRY_ADDRESS=localhost:22379（对应etcd端口）