
glue:
	cp "${GOROOT}/misc/wasm/wasm_exec.js" ./www/

build:
	GOOS=js GOARCH=wasm go build -o ./www/color.wasm ./main.go