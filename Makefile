include ./env.env
export $(shell sed 's/=.*//' ./env.env)
# show env variables
env:
	@env | grep PIPELINE

# run via go run
run:
	go run main.go