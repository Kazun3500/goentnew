# goentnew

run db
```
docker run --name some-postgres -e POSTGRES_PASSWORD=mysecretpassword -p 5432:5432 -d postgres

```
load test python
```
ab -p post_data.json -T application/json -c 20 -n 2000 'http://0.0.0.0:8000/users/' 
```

load test golang

```
ab -p post_data.json -T application/json -c 20 -n 2000 'http://0.0.0.0:4242/market'
```
