### Golang을 활용한 gRPC 인증 구현 - README

# Golang 기반 gRPC 인증

이 저장소는 Golang에서 gRPC 서비스를 인증 방식으로 보호하는 기본적인 구현을 포함하고 있습니다. 토큰 기반 인증과 인터셉터(interceptor)를 활용한 요청 검증을 다룹니다.

## 주요 기능
- 인증 미들웨어를 포함한 gRPC 서버
- 토큰 기반 인증을 사용하는 gRPC 클라이언트


## 시작하기

### 사전 준비
- Go 1.18 이상
- Protocol Buffers (`protoc`)
- gRPC 및 관련 Go 패키지

