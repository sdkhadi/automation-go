docker-clean:	
	@docker rm -f $$(docker ps -aq)
	@echo "Docker successfully removed"	

pull-chrome:
	@docker pull selenoid/chrome:90.0
	@echo "Image chrome successfully pull"
pull-firefox:
	@docker pull selenoid/firefox:90.0
	@echo "Image chrome successfully pull"

selenoid-hub:
	@docker run -d --name selenoid -p 4444:4444 -v /var/run/docker.sock:/var/run/docker.sock -v ${PWD}:/etc/selenoid/:ro aerokube/selenoid:latest-release
	@echo "Docker selenium-hub is running"

selenoid-ui:
	@docker run -d --name selenoid-ui -p 8080:8080 aerokube/selenoid-ui --selenoid-uri=http://192.168.1.6:4444
	@echo "Docker selenoid-ui is running"
