
local:
	go run app/services/sales-api/main.go

# ==============================================================================
# Modules support

tidy:
	go mod tidy
	go mod vendor