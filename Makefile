all: colors/*.vim

colors/*.vim: templates/*.yaml
	@go run main.go
