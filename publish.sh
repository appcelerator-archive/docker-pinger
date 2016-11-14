docker push appcelerator/pinger:latest \
	&& docker push appcelerator/pinger:$(cat VERSION | sed 's/[:space:]*$//')

