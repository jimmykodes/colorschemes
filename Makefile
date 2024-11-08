all: colors/*.vim

colors/%.vim: templates/%.yaml base.yaml
	yq eval-all 'select(fileIndex == 0) *=  select(fileIndex == 1)' base.yaml $< | go run main.go
