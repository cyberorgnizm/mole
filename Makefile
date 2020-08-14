# Run build for SPA and server code
build:
	@yarn
	@yarn run build
	@go build -o mole.out main.go
run:
	@./mole.out