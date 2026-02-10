.PHONY: build test test-docker clean

build:
	go build ./...

test:
	go test ./...

test-docker:
	docker compose -f docker-compose.test.yml build
	@echo "=== Testing Alpine ==="
	docker compose -f docker-compose.test.yml run --rm test-alpine
	@echo ""
	@echo "=== Testing Debian ==="
	docker compose -f docker-compose.test.yml run --rm test-debian
	@echo ""
	@echo "=== Testing Ubuntu ==="
	docker compose -f docker-compose.test.yml run --rm test-ubuntu
	@echo ""
	@echo "=== Testing Fedora ==="
	docker compose -f docker-compose.test.yml run --rm test-fedora
	@echo ""
	@echo "=== Testing Arch Linux ==="
	docker compose -f docker-compose.test.yml run --rm test-archlinux
	@echo ""
	@echo "=== Testing openSUSE ==="
	docker compose -f docker-compose.test.yml run --rm test-opensuse

test-docker-clean:
	docker compose -f docker-compose.test.yml down --rmi all

clean:
	rm -rf bin/
	go clean
