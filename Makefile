.DEFAULT_GOAL := help

REPO := devopsgig

IMAGE_TEST_NAME := restapi_test
IMAGE_PROD_NAME := restapi

CONTAINER_TEST_NAME := "${IMAGE_TEST_NAME}"
CONTAINER_PROD_NAME := "${IMAGE_PROD_NAME}"

.PHONY: help build/test/image test build/prod/image

help:
	@echo "------------------------------------------------------------------------"
	@echo "devopsgig REST API"
	@echo "------------------------------------------------------------------------"
	@grep -E '^[a-zA-Z_/%\-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build/test/image: ## build test image
	@docker build -t  "${REPO}"/"${IMAGE_TEST_NAME}" -f ./resources/test/Dockerfile .

build/prod/image: ## build prod image
	@docker build -t  "${REPO}"/"${IMAGE_PROD_NAME}" -f ./resources/prod/Dockerfile .

test: build/test/image ## run unit tests
	@docker run -it --rm --name "${CONTAINER_TEST_NAME}" "${REPO}"/"${IMAGE_TEST_NAME}"
	@printf "Removing "${REPO}"/"${IMAGE_TEST_NAME}" image\n\n"
	@docker rmi "${REPO}"/"${IMAGE_TEST_NAME}"

run: clean build/prod/image ## start the server
	@docker run -d -p 8080:8080 --name "${CONTAINER_PROD_NAME}" "${REPO}"/"${IMAGE_PROD_NAME}"

clean: ## remove running container
	@./scripts/rm-container.sh "${REPO}"/"${IMAGE_PROD_NAME}" "${CONTAINER_PROD_NAME}" &> /dev/null
