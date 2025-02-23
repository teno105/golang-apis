# golang-apis

`golang-apis` 는 Golang으로 작성된 json 파일들을 gin으로 받아오는 프로젝트입니다.
게임이 실행될때 접속할 서버정보와 공지사항, 개인정보 정책 등을 얻어옵니다.

## 프로젝트 폴더 구조
```plaintext
golang-apis/
│
├── cmd/
│   └── golang-apis/
│        └── main.go
│
├── infra/
│   ├── db.json
│   └── file.proto
│
├── internal/
│   └── models/
│        ├── in_game_board.go
│        ├── latest_policy.go
│        ├── maintenance.go
│        ├── notice.go
│        ├── store_link.go
│        └── version_infos.go
│
├── data/
│   └── 11/
│        ├── in_game_board.json
│        ├── latest_policy.json
│        ├── maintenance.json
│        ├── notice.json
│        ├── store_link.json
│        └── version_infos.go
│
├── go.mod
├── Makefile
└── README.md
```