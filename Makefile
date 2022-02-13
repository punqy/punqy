#!make
include .env

start: ## Builds, (re)creates, starts, and attaches to containers for a service in the background
	$(info Make: Starting "$(ENV)" environment containers.)
	@TAG=$(TAG) docker-compose $(COMPOSE_FILE_PATH) up -d

stop: ## Stop running containers without removing them
	$(info Make: Stopping "$(ENV)" environment containers.)
	@docker-compose stop

down: ## Stops containers and removes containers, networks, volumes, and images created by `up`
	$(info Make: Stopping and removing "$(ENV)" environment containers, networks, and volumes.)
	@docker-compose down --remove-orphans

clear: ## Stops containers and removes containers, networks, volumes with static informations
	$(info Make: Stopping and removing "$(ENV)" environment containers, networks, and volumes with data.)
	@docker-compose down -v --remove-orphans

restart: ## Stop and start containers
	$(info Make: Restarting "$(ENV)" environment containers.)
	@make -s stop
	@make -s start

push: ## Pushing image to hub
	$(info Make: Pushing "$(TAG)" tagged image.)
	@docker push $(IMAGE):$(TAG)

pull: ## Pulling images from hub
	$(info Make: Pulling images from hub)
	@docker-compose pull

clean: ## Remove unused data without prompt for confirmation
	@docker system prune --volumes --force

up: ## Up the project
	$(info Make: Starting the project)
	@make -s start
