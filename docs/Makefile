SDIR:=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))/../../api)
ODIR:=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))
DIR:=$(strip $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST)))))


all: checkout build

checkout:

build:
	echo $(DIR)
	docker run --rm -v $(SDIR)/:/api -v $(ODIR):/build create-adoc convert -i /api/swagger.yaml -d /build
	docker run -v $(DIR):/docs -v $(ODIR):/build create-html --doctype book /docs/index.adoc  -D /build
	docker run -v $(DIR):/docs -v $(ODIR):/build create-pdf  --doctype book /docs/index.adoc -D /build

clean:
	rm -fr *.html *.pdf paths.adoc definitions.adoc overview.adoc security.adoc
