package deal

//go:generate go run  ../../../internal/gen/cmd/main.go -name=deal -type=Deal -out ./get.go -short-name=k -template ../../../internal/gen/internal/cmd/get.tmpl
//go:generate go run  ../../../internal/gen/cmd/main.go -name=deal -type=Deal -cmd-names=get -out ./deal.go -template ../../../internal/gen/internal/cmd/root-cmd.tmpl
