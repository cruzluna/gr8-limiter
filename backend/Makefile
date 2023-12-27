local:
	redis-server --daemonize yes && go run main.go

dock: 
	docker image build -t gr8-api . && docker run --publish 3000:3000 gr8-api
