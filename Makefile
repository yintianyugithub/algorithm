# 获取当前分支名
CURRENT_BRANCH := $(shell git branch --show-current)
RPC_DIR ?= ./rpc

generate-rpc:
	cd rpc && goctl rpc new $(name) --home ./core/template --style go_zero

generate-api:
	goctl api new $(name) --home ./core/template --style go_zero