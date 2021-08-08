/**
 * @author sdkhadi
 * @email sdkhadiwijaya@gmail.com
 * @create date 2021-08-07 19:27:01
 * @modify date 2021-08-07 19:27:01
 * @desc [description]
 */

docker-clean:
	@docker container rm $$(docker ps -aq) -f
	@echo "Docker successfully removed"

pull-chrome-latest:
	@docker pull selenoid/chrome
	@echo "Image chrome successfully pull"
pull-firefox-latest:
	@docker pull selenoid/firefox
	@echo "Image chrome successfully pull"

selenoid:
	docker run -d -p 4444:4444 --name selenium-hub selenium/hub
	@echo "Docker selenium-hub is running"

selenoid-ui:
	@echo "Successfully build selenoid ui on 127.0.0.1:8080"
