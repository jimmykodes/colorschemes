NVIM_DIR=../colorschemes.nvim
WEZTERM_DIR=../colorschemes.wezterm
K9S_DIR=../colorschemes.k9s
GHOSTTY_DIR=../colorschemes.ghostty

all: wezterm k9s nvim ghostty

# Create a list of source YAML files
TEMPLATES := $(wildcard templates/*.yaml)

WEZTERM_TARGETS := $(patsubst templates/%.yaml,$(WEZTERM_DIR)/%.toml,$(TEMPLATES))
GHOSTTY_TARGETS := $(patsubst templates/%.yaml,$(GHOSTTY_DIR)/%,$(TEMPLATES))
K9S_TARGETS := $(patsubst templates/%.yaml,$(K9S_DIR)/%.yaml,$(TEMPLATES))
NVIM_TARGETS := $(patsubst templates/%.yaml,$(NVIM_DIR)/colors/%.vim,$(TEMPLATES))

.PHONY: wezterm
wezterm: $(WEZTERM_TARGETS)

# Rule to create TOML files from YAML templates
$(WEZTERM_DIR)/%.toml: templates/%.yaml
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml $< | \
		go run main.go wezterm --wezterm-dir $(WEZTERM_DIR)

.PHONY: ghostty
ghostty: $(GHOSTTY_TARGETS)

$(GHOSTTY_DIR)/%: templates/%.yaml
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml $< | \
		go run main.go ghostty --dir $(GHOSTTY_DIR)
	
.PHONY: k9s
k9s: $(K9S_TARGETS)

$(K9S_DIR)/%.yaml: templates/%.yaml
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml $< | \
		go run main.go k9s --k9s-dir $(K9S_DIR)

.PHONY: nvim
nvim: $(NVIM_TARGETS)

$(NVIM_DIR)/colors/%.vim: templates/%.yaml
	yq eval-all 'select(fileIndex == 0) * select(fileIndex == 1)' base.yaml $< | \
		go run main.go nvim --nvim-dir $(NVIM_DIR)
