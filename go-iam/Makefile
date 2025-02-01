# name app build
APP_NAME = gobe

# dir and main
RESOURCES_DIR = src/resources
MAIN_DIR = src/main
DI_DIR = $(MAIN_DIR)/di
MAIN_FILE = $(MAIN_DIR)/main.go

# default build if not args in make
all: build

# Install and generate swag
swag_install:
	@echo "Installing swag..."
	@go install github.com/swaggo/swag/cmd/swag@latest

swag_gen:
	@echo "Generating swag.."
	@swag init --generalInfo $(MAIN_FILE) --output $(RESOURCES_DIR)/swagger

swag_build: swag_install
	@echo "Remove Dir Swag"
	@rm -rf $(RESOUCES_DIR)/swagger
	@echo "Generating swag.."
	@swag init --generalInfo $(MAIN_FILE) --output $(RESOURCES_DIR)/swagger

# Install DI
di_install:
	@echo "Installing wire..."
	@go install github.com/google/wire/cmd/wire@latest

di_gen:
	@echo "Generating wire..."
	@cd $(DI_DIR) && wire gen

# Build the application
build: 
	@echo "Building the application..."
	@go build -v -o $(APP_NAME) $(MAIN_FILE)

# run the application
run: 
	@echo "Running the application..."
	@go run $(MAIN_FILE)