mocks:
	# Run mockgen to generate mock interfaces
	mockgen -source=./internal/controller/controller.go -destination=mocks/mock_controller.go -package=mocks
	mockgen -source=./pkg/utils/utils.go -destination=mocks/mock_utils.go -package=mocks
	# Print a message indicating the process is complete
	echo "Mock interfaces generated successfully."