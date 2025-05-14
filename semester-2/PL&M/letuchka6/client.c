#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <unistd.h>
#include <stdio.h>
#include <errno.h>
#include <string.h>

int main() {
    int sock = socket(AF_INET, SOCK_STREAM, 0);
    if (sock < 0) {
        fprintf(stderr, "Socket error: %s\n", strerror(errno));
        return 1;
    }

    struct sockaddr_in server_addr = {0};
    server_addr.sin_family = AF_INET;
    server_addr.sin_port = htons(6060);
    inet_pton(AF_INET, "127.0.0.1", &server_addr.sin_addr);

    if (connect(sock, (struct sockaddr*)&server_addr, sizeof(server_addr)) < 0) {
        fprintf(stderr, "Connect error: %s\n", strerror(errno));
        close(sock);
        return 1;
    }

    int network_num = htonl(11);
    if (send(sock, &network_num, sizeof(network_num), 0) <= 0) {
        fprintf(stderr, "Send error: %s\n", strerror(errno));
    } else {
        printf("Sent number: %d\n", 11);
    }

    close(sock);
    return 0;
}