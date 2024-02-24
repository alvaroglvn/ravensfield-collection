# The Ravensfield Collection

## Acknowledgements

This project uses the following open-source software. Thanks to the maintainers and contributors of the following libraries:

- [gosec](https://github.com/securego/gosec)
- [staticcheck](https://github.com/dominikh/go-tools)
- [chi](https://github.com/go-chi/chi)
- [godotenv](https://github.com/joho/godotenv)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [blackfriday](https://github.com/russross/blackfriday)

## Development commands

```bash
docker run -d --name raven-ghost -e NODE_ENV=development -e url=http://localhost:8081 -p 8081:2368 ghost
```

```gow
-c run main.go
```
