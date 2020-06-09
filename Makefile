build-hero:
	./builds/hero/build.sh
run-hero:
	docker run  -p 9000:9000 gcr.io/money-money-279309/hero
deploy-hero:
	./deployments/cloud_run/hero_server.sh
