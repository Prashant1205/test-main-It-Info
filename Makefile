startProject:
	go run main.go

DockerRun:
	docker run -d -p 8080:8080 ltinfo

DockerBuild:
	docker build -t ltinfo .

DockerRemoveImage:
	docker rmi ltinfo

DockerRemoveContainer:
	docker kill ltinfo
	docker rm ltinfo	

.PHONY: startProject DockerRun DockerBuild DockerRemoveImage DockerRemoveContainer