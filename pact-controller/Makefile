IMAGE = "aliakseiburau/pact-controller"
TAG   = "latest"

docker: docker-build docker-push

docker-build:
	@docker build -f docker/Dockerfile -t $(IMAGE):$(TAG) .

docker-push:
	@docker push $(IMAGE):$(TAG)