## Lenslocked

Web app developed during the Jon Calhoun's web development course 

https://www.usegolang.com/

Deploy on fly.io:

```
docker build . --platform linux/amd64 --tag=lenslocked:latest \
&& fly deploy --local-only --image lenslocked:latest
```

Deployment:

https://lenslocked.fly.dev/galleries/1/edit
