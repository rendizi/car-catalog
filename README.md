To run the API you need to build and configure docker-compose, but before that change the .env MIGRATION_DIR var to the path to the database (which you are going to migrate) and API_URL. 

```
docker-compose build
docker-compose up
```

swagger documentation is located in src, <a href="https://github.com/rendizi/car-catalog/blob/main/BaglanovAlikhan-carCatalog-1.0.0-swagger.yaml">check it</a>
