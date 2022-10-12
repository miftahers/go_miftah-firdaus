# Docker

Contents:
- How container work
- How to launch containers using Docker
- How to build and launc your own container images
- How to deploy your application as container


CONTAINER

Application > App dependency (e.g. library, etc.) > Container > Operation System

What is container? A container is not a virtual machine.

Short answer: A container is a process with file system isolation.

Long answer: Everything in linux is a file.
- /dev/sda = hard disk
- /dev/proc = process
- /dev/usb
- /dev/cpu
- /dev/std (input|output)
- /bin/bash... just a binary file

## Docker Basics
- Image
- Container
- Engine
- Registry
- Control plane

## Docker Infrastructure

Visualization of docker infrastructure:

!['Docker'](https://devopedia.org/images/article/101/8323.1565281088.png)

!['Docker Architecture | Learn the Objects and Benefits of Docker'](https://www.educba.com/academy/wp-content/uploads/2019/10/Docker-Architecture-2.png)



## Perintah-perintah dalam docker

FROM        : Getting image from docker registry
RUN         : Execute bash command when building container
ENV         : Set Variable inside container
ADD         : Copy the file with some other process
COPY        : Copy the file
WORKDIR     : Set working file directory
ENTRYPOINT  : Execute command when finish building container
CMD         : Execute command but can be overwrite

## Demo Docker
tutorial docker
1. Create Dockerfile di folder project, satu folder dengan main.go/go.sum/go.mod
2. tulis ini di Dockerfile (edit menggunakan text editor)

```
FROM golang:1.17-alpine // language:version

WORKDIR /app  // define work directory

COPY go.mod ./ // copy file ke workdir
COPY go.sum ./
RUN go mod download // download package yang dipakai app

COPY . . // copy source code yang kita buat (e.g. main.go/controller etc.)

RUN go build -o ./dist // build go file in dist

EXPOSE 3222 // define running port

CMD ["/dist"] // perintah untuk menjalankan app
```

3. Build aplikasi menjadi image

```
// cara melihat image yang sudah dibuat
docker image ls

// cara 1 - tanpa repository dan tag
docker build .

// cara 2 - dengan repository dan default tag
docker build -t backend . //docker build -t <image-name> .

// cara 3 - dengan repository dan custom tag
docker build -t backend:v1.0.0 . // docker build -t <image-name>:<tag name> .

```

4. Create repository di docker hub
5. init di terminal

```

// 1 login
docker login -u miftahers // docker login -u <username>

// 2 lakukan tag image
docker tag backend miftahers/backend // docker tag <image-name> <username>/<repo-name>

// 3 push
docker push miftahers/backend // docker push <username>/<repo-name>
```

6. jalankan docker di terminal

```
// menjalankan docker
docker run backend // docker run <image-name> 

// untuk melihat yang sedang berjalan
1) docker ps
2) docker ps -a

// untuk stop docker yang sedang berjalan /
docker stop <container-id> // container id bisa didapat dengan cek docker yang berjalan

// untuk menghapus docker yang selesai berjalan
docker rm <container-id>

// untuk menjalankan docker dengan -d
docker run -d <image-name>

// untuk menjalankan docker dengan port tertentu supaya bisa di expose dengan browser
docker run -d -p 3200:3222 backend // docker run -d -p <custom-port>:<exposed-port-in-dockerfile> <image-name>


```