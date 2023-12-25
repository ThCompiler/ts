VERSION=local

.PHONY: release
release:
	git tag $(VERSION); \
	git push origin $(VERSION)

.PHONY: clean
changelog:
	sh ./workflow/changes.sh $(VERSION) > CURRENT-CHANGELOG.md
