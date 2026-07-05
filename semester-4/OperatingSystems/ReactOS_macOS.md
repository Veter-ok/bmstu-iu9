# Set up ReactOS for Mac OS Apple Silicon

> Самый простой способо собрать iso на MacOS это создать образ linux через docker-контейнер

Создание docker-контейнера
```bash
docker run -it --platform linux/amd64 --name rosbe ubuntu:20.04 bash
```

Теперь мы находимся в нашем  docker-контейнере, куда мы можем скачать React Build Environment.

```bash
apt update
apt install -y software-properties-common
apt install -y bison flex
add-apt-repository -y ppa:reactos/rosbe-unix
apt update

# Устанавливаем RosBE и необходимые инструменты
apt install -y rosbe-unix git nano cmake ninja-build

source /usr/RosBE/RosBE.sh   # Активируем окружение RosBE
rosbe version  # Должна показаться версия
```

Качаем исходники
```bash
git clone https://github.com/reactos/reactos.git
cd reactos
```

Билдим ISO

```bash
./configure.sh
cd output-MinGW-i386
ninja bootcd
```

Перебрасываем iso на свой рабочий стол
```bash
docker ps # чекаем ID контейнера
docker cp <ID контейнера>:/reactos/output-MinGW-i386/bootcd.iso ~/Desktop/ReactOS.iso
```

```bash
qemu-img create -f qcow2 ReactOS.qcow2 10G

qemu-system-i386 -m 4G \
  -drive if=ide,index=0,media=disk,file=ReactOS.qcow2 \
  -cdrom ReactOS.iso \      
  -boot order=d \
  -rtc base=localtime -serial file:ReactOS.log

qemu-system-i386 -m 4G -drive if=ide,index=0,media=disk,file=ReactOS.qcow2 -boot order=c -rtc base=localtime -serial file:ReactOS.log
```


Если нужно заново открыть docker-контейнер, то возпользуйтесь следующими командами
```bash
docker start <имя_или_ID_контейнера>
docker exec -it <имя_или_ID_контейнера> bash
```