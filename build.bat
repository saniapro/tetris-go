set CGO_ENABLED=0
go build -trimpath -ldflags="-s -w" -o tetris.exe cmd/tetris/main.go