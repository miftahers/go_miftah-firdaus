1.  connection refused from by localhost:3306
note: WORKDIR should written "/app" 
```
https://stackoverflow.com/questions/57566060/panic-dial-tcp-127-0-0-13306-connect-connection-refused
```

2. Networking in compose
note: sangat penting bahwa tuliskan db host nya itu nama compose db -> @(mysql-container:<port>) jangan @(localhost:<port>)
```
https://docs.docker.com/compose/networking/
```
3. tutorial docker compose sederhana dari freecodecamp

```
https://www.freecodecamp.org/news/run-multiple-containers-with-docker-compose/
```
4. `wait-for` penting banget buat nungguin suatu kondisi , contoh nunggu mysql nya nyala baru jalanin golang

```
https://stackoverflow.com/questions/63198731/how-to-use-wait-for-it-in-docker-compose-file
```