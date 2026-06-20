dev:
	@templ generate -v -watch -proxy="http://0.0.0.0:1323" -open-browser=false -cmd="go run cmd/server/main.go" & \
	npx @tailwindcss/cli -i ./static/css/input.css -o ./static/css/output.css --watch
