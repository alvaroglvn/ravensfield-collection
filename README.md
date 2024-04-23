# The Ravensfield Collection

## Acknowledgements

This project uses the following open-source software. Thanks to the maintainers and contributors of the following libraries:

- [gosec](https://github.com/securego/gosec)
- [staticcheck](https://github.com/dominikh/go-tools)
- [chi](https://github.com/go-chi/chi)
- [godotenv](https://github.com/joho/godotenv)
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [blackfriday](https://github.com/russross/blackfriday)
  - [go-webp](https://github.com/kolesa-team/go-webp)

## Development commands

```text
docker run -d \
    --name ravensfield \
    -e NODE_ENV=development \
    -e url=http://localhost:8081 -p 8081:2368 \
    --mount type=bind,source="$(pwd)"/theme,target=/var/lib/ghost/content/themes/ravensfield \
    ghost 
```  

```gow
-c run main.go
```
