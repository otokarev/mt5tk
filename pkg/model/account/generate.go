package account

//go:generate go run  ../../../internal/gen/cmd/main.go -name=account -type=Account -get-path=/api/user/account/get -id-name=login -out ./default.go -template ../../../internal/gen/pkg/model/default.tmpl
