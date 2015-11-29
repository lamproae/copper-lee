#include <sys/socket.h>
#include <sys/types.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <stdlib.h>

int start_tcp_client(char *name)
{
	int ret = -1;
	int cli_fd = -1;
	char line[1024] = { 0 };
	char srv_ip_addr[20] = { 0 };
	int len;
	struct sockaddr_in srv_addr;

	bzero(&srv_addr, sizeof(srv_addr));
	srv_addr.sin_family = AF_INET;
	srv_addr.sin_port = htons(1000);
	inet_pton(AF_INET, srv_ip_addr, (void *)&srv_addr.sin_addr);

	cli_fd = socket(AF_INET, SOCK_STREAM, 0);
	ret = connect(cli_fd, (struct sockaddr *)&srv_addr, sizeof(srv_addr));
	printf("ret: %d\n", ret);
	write(cli_fd, "Hello world!\n", strlen("Hello world\n"));
	while(read(cli_fd, line, 1024) > 0)
		printf("%s\n", line);

	return 0;
}

int start_udp_client(char *name)
{
	int ret = -1;
	int cli_fd = -1;
	char line[1024] = { 0 };
	char srv_ip_addr[20] = { 0 };
	int len;
	int err;
	struct sockaddr_in srv_addr;

	bzero(&srv_addr, sizeof(srv_addr));
	srv_addr.sin_family = AF_INET;
	srv_addr.sin_port = htons(50333);
	inet_pton(AF_INET, srv_ip_addr, (void *)&srv_addr.sin_addr);

	cli_fd = socket(AF_INET, SOCK_DGRAM, 0);
	ret = connect(cli_fd, (struct sockaddr *)&srv_addr, sizeof(srv_addr));
	printf("ret: %d\n", ret);
	sendto(cli_fd, "Hello world!\n", strlen("Hello world\n"), 0, (struct sockaddr *)&srv_addr, sizeof(srv_addr));
	while(read(0, line, 1024) > 0)
	{
		if (strlen(line) > 0)
		{
			sendto(cli_fd, line, strlen(line), 0, (struct sockaddr *)&srv_addr, sizeof(srv_addr));
			while(err = recvfrom(cli_fd, line, 1024, MSG_DONTWAIT, (struct sockaddr *)&srv_addr, NULL))
			{
				if (err < 0)
					break;

				printf("%s\n", line);
			}
		}
		memset(line, 0, 1024);
	}

	close(cli_fd);
}

int start_raw_client(char *name)
{

}

int start_packet_client(char *name)
{

}

int main(int argc, char **argv)
{
//	start_tcp_client(NULL);
	start_udp_client(NULL);
	return 0;
}
