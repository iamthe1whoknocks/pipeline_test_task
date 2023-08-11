include ./env.env
export $(shell sed 's/=.*//' ./env.env)
# show env variables
env:
	@env | grep PIPELINE

# run via go run
run:
	cd cmd/app && ./pipe

## build binary 
build:
	@cd cmd/app && go build -o pipe