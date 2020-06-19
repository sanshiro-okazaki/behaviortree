.ONESHELL:

define MSG

   	Usage: make [OPTION]

	OPTIONS:
		run s=0.5: スキルA実行確率を指定して実行

endef

export MSG
.PHONY: help
help:
	@echo "$$MSG"

.PHONY: run
run:
	go mod tidy && go build -o bt && SKILL_A=$(s) ./bt
