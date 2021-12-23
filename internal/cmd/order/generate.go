package order

//go:generate go run  ../../../internal/gen/cmd/main.go -name=order -type=Order -out ./get.go -short-name=k -template ../../../internal/gen/internal/cmd/get.tmpl
//go:generate go run  ../../../internal/gen/cmd/main.go -name=order -type=Order -cmd-names=get -out ./order.go -template ../../../internal/gen/internal/cmd/root-cmd.tmpl
