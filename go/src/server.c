#include <unistd.h>
#include <stdio.h>
#include <stdlib.h>
#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>

int start_tcp_server(char *name)
{
	int ret = -1;
	int listen_fd = -1;
	int connectted_fd = -1;
	char line[1024] = {0};
	char cli_ip_addr[20];
	int len;
	int rcount;
	struct sockaddr_in srv_addr;
	struct sockaddr_in cli_addr;

	listen_fd = socket(AF_INET, SOCK_STREAM, 0);

	bzero(&srv_addr, sizeof(srv_addr));
	srv_addr.sin_family = AF_INET;
	srv_addr.sin_port = htons(1000);
	srv_addr.sin_addr.s_addr = htonl(INADDR_ANY);

	bind(listen_fd, (struct sockaddr *)&srv_addr, sizeof(srv_addr));
	listen(listen_fd, 1024);
	connectted_fd = accept(listen_fd, (struct sockaddr *)&cli_addr, &len);
	if (connectted_fd > 0)
	{
		inet_ntop(AF_INET, (void *)&cli_addr.sin_addr, cli_ip_addr, 20);
		printf("Received connection from: %s\n", cli_ip_addr);
		while((rcount = read(connectted_fd, line, 1024)) > 0)
		{
			printf("Received: \"%s\" from %s\n", line, cli_ip_addr);
			write(connectted_fd, line, rcount);
		}
	}
		close(listen_fd);
}

int start_udp_server(char *name)
{
	int srv_fd = -1;
	struct sockaddr srv_addr;

	srv_fd = socket(AF_INET, SOCK_DGRAM, 0);

}

int main(int argc, char **argv)
{
	start_tcp_server("liwei");
}
