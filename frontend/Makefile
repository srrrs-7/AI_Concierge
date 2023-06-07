init:
	cp compose.override.yml.sample compose.override.yml

up:
	docker compose up -d --build

down:
	docker compose down

dev:
	docker compose run --service-ports --rm ui npm i
	docker compose run --service-ports --rm ui npm run dev

sb:
	docker compose run --service-ports --rm ui npm i
	docker compose run --service-ports --rm ui npm run storybook

storybook:
	npx sb@latest init