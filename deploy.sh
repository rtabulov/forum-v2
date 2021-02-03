go build -o ./tmp/main ./cmd
git add tmp
git commit -m "build"
git push heroku main
