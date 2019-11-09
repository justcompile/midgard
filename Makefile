.PHONY: all
all: build lint


.PHONY: build
build:
	@echo "Building"
	docker build . -t justcompile/migard-web
	docker build . -f Dockerfile.worker -t justcompile/migard-worker

.PHONY: lint
lint:
	@echo "Linting"
	docker run --rm -it justcompile/migard-web bash scripts/lint.sh
	docker run --rm -it justcompile/migard-worker bash scripts/lint.sh

# .PHONY: test
# test:
# 	@echo "Testing"
# 	docker-compose build tests
# 	docker-compose run tests bash scripts/run_tests.sh
# 	docker-compose down

# .PHONY: lambda
# lambda:
# 	@echo "Building and zipping Lambda func"
# 	bash scripts/build_lambda.sh

# .PHONY: publish
# publish:
# 	@echo "--> Publishing API"
# 	docker-compose build proxy
# 	docker tag ${ECR_REPO_BASE}:latest ${ECR_REPO_BASE}:${TAG}
# 	docker push ${ECR_REPO_BASE}:${TAG}