VERSION=local

.PHONY: release
release:
	git tag $(VERSION); \
	git push origin $(VERSION)

.PHONY: changelog
changelog:
	sh ./workflow/changes.sh $(VERSION) > CURRENT-CHANGELOG.md
