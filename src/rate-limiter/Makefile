VERSION := 0.0.8
DOCKER_TAG := hxy9243/example_ratelimit:$(VERSION)
NAMESPACE := ratelimit
REDIS_CHART := bitnami/redis
REDIS_VERSION := 23.1.3


docker-build:
	echo "==== Building docker image... ===="
	docker build -t $(DOCKER_TAG) . -f Dockerfile

docker-push: docker-build
	echo "==== Pushing docker image... ===="
	docker push $(DOCKER_TAG)

helm-repo:
	echo "==== Adding helm repo ===="
	helm repo add bitnami https://charts.bitnami.com/bitnami
	helm repo update

helm-redis: helm-repo
	echo "==== Installing redis in k8s cluster ===="
	helm install redis $(REDIS_CHART) --version $(REDIS_VERSION) -n $(NAMESPACE) -f deploy/redis-values.yaml --wait

deploy-service:
	echo "==== Deploying service... ===="
	kubectl apply -f deploy/http_service.yaml -n $(NAMESPACE)
	kubectl get svc -n $(NAMESPACE)

deploy:  docker-push helm-redis deploy-service
	echo "Deploying..."

.PHONY:
all: deploy
