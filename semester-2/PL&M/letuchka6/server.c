#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>

int main() {
    int server_fd = socket(AF_INET, SOCK_STREAM, 0);
    if (server_fd < 0) {
        fprintf(stderr, "Socket error: %s\n", strerror(errno));
        return 1;
    }

    struct sockaddr_in addr = {0};
    addr.sin_family = AF_INET;
    addr.sin_port = htons(6060);
    addr.sin_addr.s_addr = htonl(INADDR_ANY);

    if (bind(server_fd, (struct sockaddr*)&addr, sizeof(addr)) < 0) {
        fprintf(stderr, "Bind error: %s\n", strerror(errno));
        close(server_fd);
        return 1;
    }

    if (listen(server_fd, 1) < 0) {
        fprintf(stderr, "Listen error: %s\n", strerror(errno));
        close(server_fd);
        return 1;
    }

    printf("Server waiting for connections on 127.0.0.1:6060...\n");

    struct sockaddr_in client_addr;
    socklen_t client_len = sizeof(client_addr);
    int client_fd = accept(server_fd, (struct sockaddr*)&client_addr, &client_len);
    if (client_fd < 0) {
        fprintf(stderr, "Accept error: %s\n", strerror(errno));
        close(server_fd);
        return 1;
    }

    int received_num;
    if (recv(client_fd, &received_num, sizeof(received_num), 0) <= 0) {
        fprintf(stderr, "Recv error: %s\n", strerror(errno));
    } else {
        received_num = ntohl(received_num);
        printf("Square number: %d\n", received_num*received_num);
    }

    close(client_fd);
    close(server_fd);
    return 0;
}