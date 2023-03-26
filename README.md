# tutorial-microservices-go

Todo:
- Refactor: store the var env in a file
- Add test
- Container test
- Run services independently
- Graceful shutdown


How to test it:
1) 
```
make up_build
make postgres-migrate
```

2) 
```
  cd front-end
  go run ./cmd/web
```

3) Open http://localhost/ in your browser

To connect to Compass:   
```
mongodb://admin:password@localhost:27018/logs?authSource=admin&readPreference=primary&appname=MongoDB%20Compass&ssl=false
```
