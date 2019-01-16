
PREFIX=localhost:5000/
TAG=latest

all:

OUTPUTDIR=_output

output:
	mkdir -p $(OUTPUTDIR)

build: storage-build itineraries-server-build ui-build

images: storage-image-build itineraries-server-image-build ui-image-build

storage-validate-swagger:
	cd storage/swagger && swagger validate ./swagger.yaml

storage-build: output
	CGO_ENABLED=0  GOOS=linux go build -i -installsuffix cgo -ldflags '-w' -o $(OUTPUTDIR)/storage storage/cmd/storage/main.go

storage-image-build: storage-build
	cp -f $(OUTPUTDIR)/storage storage/image
	cd  storage/image && docker build .  -t $(PREFIX)storage:$(TAG)
	rm -rf storage/storage

storage-generate-server:
	cd storage/swagger && swagger generate server --target ../pkg/gen  --flag-strategy pflag --exclude-main --name storage --spec ./swagger.yaml

storage-generate-client:
	cd storage/swagger && swagger generate client --target ../pkg/gen --name storage --spec ./swagger.yaml

itineraries-server-validate-swagger:
	cd itineraries-server/swagger && swagger validate ./swagger.yaml

itineraries-server-build: output
	CGO_ENABLED=0  GOOS=linux go build -i -installsuffix cgo -ldflags '-w' -o $(OUTPUTDIR)/itineraries-server itineraries-server/cmd/itineraries-server/main.go

itineraries-server-image-build: itineraries-server-build
	cp -f $(OUTPUTDIR)/itineraries-server itineraries-server/image
	cd  itineraries-server/image && docker build . -t $(PREFIX)itineraries-server:$(TAG)
	rm -f itineraries-server/image/itineraries-server

itineraries-server-generate-server:
	cd itineraries-server/swagger && swagger generate server --target ../pkg/gen --flag-strategy pflag  --exclude-main --name itineraries --spec ./swagger.yaml

itineraries-server-generate-client:
	cd itineraries-server/swagger && swagger generate client --target ../pkg/gen --name itineraries --spec ./swagger.yaml

ui-build: $(OUTPUTDIR)
	cd ui && go-bindata -o=assets/bindata.go --nocompress --nometadata --pkg=assets templates/... static/...
	CGO_ENABLED=0  GOOS=linux go build -i -installsuffix cgo -ldflags '-w' -o $(OUTPUTDIR)/ui ui/cmd/main.go

ui-image-build: ui-build
	cp -f $(OUTPUTDIR)/ui ui/image
	cd  ui/image && docker build . -t $(PREFIX)ui:$(TAG)
	rm -rf ui/image/ui

clean: $(OUTPUTDIR)
	rm -rf $(OUTPUTDIR)

#TMP target for local tests
start-ui: ${OUTPUTDIR}
	cd ui && ./ui

start-itineraries-server: $(OUTPUTDIR)
	output/itineraries-server --port=41807
