# Go Language

## Commands

```shell
# Initialize module
go mod init directory/module_name

# Install external modules
go get

# Update modules
go get -u

# Get required modules and remove unused modules
go mod tidy -v

# Run
go run .

# Build
go build

# Build Optimization (Windows Commands)
go build -ldflags="-s -w" main.go

# Build Optimization (Linux Commands)
GOOS=linux go build -ldflags="-s -w" main.go
```
