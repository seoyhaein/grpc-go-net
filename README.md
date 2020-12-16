## grpc-go-net

#### Grpc 서버 api 정리

* gprc
> https://github.com/grpc/grpc-go   --- go
>
> 
* protobuf
``` 
https://github.com/protocolbuffers/protobuf-go - currently go protobuf official site. 
https://pkg.go.dev/google.golang.org/protobuf  - documentation 
https://developers.google.com/protocol-buffers/docs/overview - protobuf developing guide
```

* grpcurl 
> https://github.com/fullstorydev/grpcurl 

* grpcurl 명령어
```
proto file 읽기 
grpcurl --import-path proto 위치 -proto  proto 파일 이름 list 

plaintext insecure server 접속의 경우
grpcurl --plaintext 서버주소(localhost:50051) list > grpc 서비스 호출  
grpcurl --plaintext 서버주소(localhost:50051) list 서비스 이름 > 해당 서비스의 rpc api 호출 

grpc 메서드 살펴보기
grpcurl --plaintext 서버주소(localhost:50051) describe 서비스.메서드 이름

grpc 메세지 살펴보기
grpcurl --plaintext 서버주소(localhost:50051) describe .메세지이름

grpc 메세지 json 형태로 바꾸기
grpcurl --plaintext -ㅡmgs-template 서버주소(localhost:50051) describe .메세지이름

grpc 메서드 호출
grpcurl --plaintext -d '{
"name": "seoyhaein"
}' 서버주소(localhost:50051) 서비스 이름/메서드 이름
```
* for test
> https://github.com/nicholasjackson/all-things-microservices/tree/master/grpc-curl 