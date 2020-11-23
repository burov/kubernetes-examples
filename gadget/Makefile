IMAGE = "aliakseiburau/gadget"
TAG   = "latest"

docker: docker-build docker-push

docker-build:
	@docker build -f build/docker/Dockerfile -t $(IMAGE):$(TAG) .

docker-push:
	@docker push $(IMAGE):$(TAG)