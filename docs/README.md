# Compiling suika-game-go

On a Unix/Linux shell:

```
env GOOS=js GOARCH=wasm go build -o suika-game-go.wasm github.com/demouth/suika-game-go

cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
```
