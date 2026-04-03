NVIM_DIR=./
K9S_DIR=./extras/k9s
GHOSTTY_DIR=./extras/ghostty

all: nvim

# Create a list of source YAML files
TEMPLATES := $(wildcard templates/*.yaml)

TARGETS := $(patsubst templates/%.yaml,$(NVIM_DIR)/colors/%.lua,$(TEMPLATES))

.PHONY: nvim
nvim: $(TARGETS)

$(NVIM_DIR)/colors/%.lua: templates/%.yaml
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml $< | \
		go run ./cmd/generate --nvim-dir $(NVIM_DIR) --k9s-dir $(K9S_DIR) --ghostty-dir $(GHOSTTY_DIR)
