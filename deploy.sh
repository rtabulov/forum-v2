go build -o ./main ./cmd
git add main
git commit -m "build"
git push heroku main
rm main
