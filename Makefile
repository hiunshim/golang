# Check if all Go files are formatted
check-format:
	@OUTPUT=$$(gofmt -l .); \
	if [ -n "$$OUTPUT" ]; then \
		echo "The following files are not formatted. Run 'make format' to fix them:"; \
		echo "$$OUTPUT"; \
		exit 1; \
	fi

# Format all Go files
format:
	gofmt -w .

