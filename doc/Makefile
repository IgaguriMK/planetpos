.PHONY: html
html:
	pandoc -s \
		-c style.css \
		-H extra-header.html \
		--number-sections \
		--section-divs \
		--toc \
		--mathjax \
		doc.md -o doc.html
