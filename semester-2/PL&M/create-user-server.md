#################
###   Linux   ###
#################

0.
mkdir /home/testjava
useradd testjava
passwd testjava
chown testjava:testjava /home/testjava
nano /etc/passwd
testj:x:1000:1000::/home/testj:/bin/bash

1.
https://www.oracle.com/java/technologies/downloads/
wget https://download.oracle.com/java/21/latest/jdk-21_linux-x64_bin.tar.gz
wget https://download.oracle.com/java/23/latest/jdk-23_linux-x64_bin.tar.gz

2.
tar -xvf jdk-21_linux-x64_bin.tar.gz
tar -xvf jdk-23_linux-x64_bin.tar.gz

3.
ls -lsa
total 191068
     4 drwxr-xr-x 5 testj testj      4096 Feb  6 17:11 .
     4 drwxr-xr-x 3 root  root       4096 Feb  6 17:07 ..
     4 drwx------ 2 testj testj      4096 Feb  6 17:09 .cache
     4 drwx------ 3 testj testj      4096 Feb  6 17:09 .gnupg
     4 drwxrwxr-x 9 testj testj      4096 Feb  6 17:11 jdk-21.0.2
191048 -rw-rw-r-- 1 testj testj 195631048 Jan  5 21:51 jdk-21_linux-aarch64_bin.tar.gz

4.
nano .bashrc
export JAVA_HOME=$HOME/jdk-23.0.2
export PATH=$JAVA_HOME/bin:$PATH

5.
nano .profile
export JAVA_HOME=$HOME/jdk-23.0.2
export PATH=$JAVA_HOME/bin:$PATH

6.
javac -version
javac 23.0.2

7.
nano Dog.java

8.
public class Dog{
 private String name;
 private int age;
 private String color;
 public Dog(){
  age =0;
  System.out.println("Собака ▒^`одила▒^a▒^l");
 }
 public void setName(String varName){
 name = varName;
 }
 public void setAge(int varAge){
 age = varAge;
 }
  public int getAge(){
  System.out.println("Age method="+age);
  return age;
 }
 public static void main(String []args){
  Dog Buldog = new Dog();
  Dog Dvorniazhka = new Dog();
  System.out.println("Age="+Buldog.age);
  Buldog.age=10;
  System.out.println("Age="+Buldog.age);
  Dvorniazhka.setAge(105);
  System.out.println("Age="+Dvorniazhka.age);
  Dvorniazhka.getAge();
 }
}

9.
javac Dog.java

10.
java Dog

###################
###   Windows   ###
###################

1.
скачиваем https://download.oracle.com/java/21/latest/jdk-21_windows-x64_bin.zip

2.
распаковываем например в C:/javalec

3. 
заходим в настройки: Пуск -> Параметры, далее в открывшемся окне в окне поиска: "Изменение системных переменных среды" -> клик по кнопке "Переменные среды..." -> Выбираем "Path" -> нажимаем кнопку "Изменить" -> в открывшемся окне нажимаем кнопку "Добавить" -> добавляем: C:\javalec\bin -> сохраняем изменения и выходим

4. 
открываем консоль Windows: cmd
C:\Users\danila>javac -version
javac 21.0.2

5.
в любой паке windows создаем текстовый файл Factorian.java

import java.util.stream.IntStream;
public class Factorial 
{ 
 public static void main(String [] args) 
 { 
  if (args.length == 0) 
     { 
      System.out.println("Usage: java Factorial x"); 
     } 
  else 
     { 
      int n = Integer.parseInt(args [0]);
      var numbers = IntStream.range(1, n+1); 
      int f = numbers.reduce(1, (r,x)-> r*x); 
      System.out.println(f); 
     } 
 } 
}

6. 

C:\Users\danila\Desktop\iu_java>javac Factorial.java

7.

C:\Users\danila\Desktop\iu_java>java Factorial
Error: A JNI error has occurred, please check your installation and try again
Exception in thread "main" java.lang.UnsupportedClassVersionError: Factorial has been compiled by a more recent version of the Java Runtime (class file version 65.0), this version of the Java Runtime only recognizes class file versions up to 52.0
        at java.lang.ClassLoader.defineClass1(Native Method)
        at java.lang.ClassLoader.defineClass(Unknown Source)
        at java.security.SecureClassLoader.defineClass(Unknown Source)
        at java.net.URLClassLoader.defineClass(Unknown Source)
        at java.net.URLClassLoader.access$100(Unknown Source)
        at java.net.URLClassLoader$1.run(Unknown Source)
        at java.net.URLClassLoader$1.run(Unknown Source)
        at java.security.AccessController.doPrivileged(Native Method)
        at java.net.URLClassLoader.findClass(Unknown Source)
        at java.lang.ClassLoader.loadClass(Unknown Source)
        at sun.misc.Launcher$AppClassLoader.loadClass(Unknown Source)
        at java.lang.ClassLoader.loadClass(Unknown Source)
        at sun.launcher.LauncherHelper.checkAndLoadMain(Unknown Source)

Проблема приведенная выше возникает из-за того, что на поем ПК установлено две версии Java и по умолчанию запускается не та версия Java-машины, которая стоит в C:\javalec\bin\
А вот эта C:\ProgramData\Oracle\Java\javapath\java.exe, см. далее:
C:\Users\danila\Desktop\iu_java>where java
C:\ProgramData\Oracle\Java\javapath\java.exe
C:\javalec\bin\java.exe


8.
Если указать абсолютный путь к последней версии Java-машины, то все заработает:
C:\Users\danila\Desktop\iu_java>C:\javalec\bin\java.exe Factorial
Usage: java Factorial x

C:\Users\danila\Desktop\iu_java>

---------------------

scp -r Test.java someuser@1kurs.yss.su:/home/someuser

---------------------

net1.yss.su
IP-адрес сервера: 185.104.251.226
Пользователь: root
Пароль: fMs0m69gIGQ3

net2.yss.su
IP-адрес сервера: 185.102.139.161
Пользователь: root
Пароль: Up5b0A1wiLMQ

net3.yss.su
IP-адрес сервера: 185.102.139.168
Пользователь: root
Пароль: gOsQ5p7FUJ9w

---------------------

HTTP-сервер на Java

1. Пример 1 (источник информации: https://habr.com/ru/articles/441150/)

import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.io.PrintWriter;
import java.net.ServerSocket;
import java.net.Socket;
import java.nio.charset.StandardCharsets;

public class HttpServer {

    public static void main(String[] args) {
        try (ServerSocket serverSocket = new ServerSocket(8080)) {
            System.out.println("Server started!");
            
            while (true) {
                // ожидаем подключения
                Socket socket = serverSocket.accept();
                System.out.println("Client connected!");

                // для подключившегося клиента открываем потоки 
                // чтения и записи
                try (BufferedReader input = new BufferedReader(new InputStreamReader(socket.getInputStream(), StandardCharsets.UTF_8));
                     PrintWriter output = new PrintWriter(socket.getOutputStream())) {

                    // ждем первой строки запроса
                    while (!input.ready()) ;

                    // считываем и печатаем все что было отправлено клиентом
                    System.out.println();
                    while (input.ready()) {
                        System.out.println(input.readLine());
                    }

                    // отправляем ответ
                    output.println("HTTP/1.1 200 OK");
                    output.println("Content-Type: text/html; charset=utf-8");
                    output.println();
                    output.println("<p>Привет всем!</p>");
                    output.flush();
                    
                    // по окончанию выполнения блока try-with-resources потоки, 
                    // а вместе с ними и соединение будут закрыты
                    System.out.println("Client disconnected!");
                }
            }
        } catch (IOException ex) {
            ex.printStackTrace();
        }
    }
}

2. Пример 2 (источник информации: https://stackoverflow.com/questions/3732109/simple-http-server-in-java-using-only-java-se-api)

import java.io.IOException;
import java.io.OutputStream;
import java.net.InetSocketAddress;

import com.sun.net.httpserver.HttpExchange;
import com.sun.net.httpserver.HttpHandler;
import com.sun.net.httpserver.HttpServer;

public class Test {

    public static void main(String[] args) throws Exception {
        HttpServer server = HttpServer.create(new InetSocketAddress(8000), 0);
        server.createContext("/test", new MyHandler());
        server.setExecutor(null); // creates a default executor
        server.start();
    }

    static class MyHandler implements HttpHandler {
        @Override
        public void handle(HttpExchange t) throws IOException {
            String response = "This is the response";
            t.sendResponseHeaders(200, response.length());
            OutputStream os = t.getResponseBody();
            os.write(response.getBytes());
            os.close();
        }
    }

}
