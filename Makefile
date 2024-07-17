.PHONY: env-example css css-watch templ-watch watch

env-example:
	@echo "Creating .env.example file"
	@sed 's/=.*/=/' .env > .env.example
	@echo ".env.example file created."

css:
	tailwindcss -i pkg/web/static/css/main.css -o pkg/web/static/css/lib_build_minify.css --minify

css-watch:
	tailwindcss -i pkg/web/static/css/main.css -o pkg/web/static/css/lib_build.css --watch

templ-watch:
	@templ generate --watch --proxy="https://localhost:3333" --open-browser=false

watch:
	@echo "Starting combined watch..."
	@trap 'kill 0' INT; \
	make css-watch & \
	make templ-watch & \
	air & \
	wait