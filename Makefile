build:
	go build -o card .

.PHONY: run
run:
	nohup ./card > /dev/null 2>&1 &