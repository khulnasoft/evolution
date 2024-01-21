MAKE ?= make
SHELL ?= /bin/bash

README_PATH=./README.md
MAINTAINERS_PATH=./MAINTAINERS.md
MAINTAINERS_YAML_PATH=./maintainers.yaml
REPOSITORIES_YAML_PATH=./repositories.yaml

.PHONY: all
all: gen-readme gen-maintainers

.PHONY: clean
clean:
	+@$(MAKE) -C utils clean

.PHONY: gen-readme
gen-readme: utils
	+./utils/bin/utils readme -o $(README_PATH) -r $(REPOSITORIES_YAML_PATH)

.PHONY: gen-maintainers
gen-maintainers: utils
	+./utils/bin/utils maintainers -o $(MAINTAINERS_PATH) -r $(REPOSITORIES_YAML_PATH) -m $(MAINTAINERS_YAML_PATH)

.PHONY: utils
utils:
	+@$(MAKE) -C utils
