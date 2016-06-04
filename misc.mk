count-go-lines:
	@echo -n 'Lines of go code: '
	@find . -type f -name '*.go' | grep -vF /vendor/ | xargs cat | grep -v '^//' | grep -v '^$$' | wc -l
.PHONY: count-go-lines
