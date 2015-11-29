#include <stdio.h>
#include <string.h>
#include <unistd.h>
#include <fcntl.h>
#include <errno.h>
#include <sys/select.h>
#include <sys/types.h>
#include <sys/ioctl.h>

int fd_map[FD_SETSIZE] = { 0 };

void call_popen(void)
{
	char cmd[64] = { 0 };
	char cmd1[64] = { 0 };
	char cmd2[64] = { 0 };
	char line[1024] = { '\0' };
	FILE *fp = NULL;
	FILE *fp1 = NULL;
	FILE *fp2 = NULL;
	int fd = -1;
	int fd1 = -1;
	int fd2 = -1;
	int rcount = -1;
	int rrcount = -1;
	int tcount = 0;
	int nfds = 1;
	int ret = 0;
	int i;
	int read_finish = 1;
	fd_set rfds;

	sprintf(cmd, "ps axjf");
	sprintf(cmd1, "ls -R /sbin");
	sprintf(cmd2, "find /bin");

	if ((fp = popen(cmd, "r")) != NULL && (fp1 = popen(cmd1, "r")) != NULL &&
			(fp2 = popen(cmd2, "r")) != NULL)
	{
		fd = fileno(fp);
		fd1 = fileno(fp1);
		fd2 = fileno(fp2);
		FD_ZERO(&rfds);
		FD_SET(fd, &rfds);
		FD_SET(fd1, &rfds);
		FD_SET(fd2, &rfds);
		fd_map[fd] = 1;
		fd_map[fd1] = 1;
		fd_map[fd2] = 1;

again:
		while((ret = (select(FD_SETSIZE, &rfds, NULL, NULL, NULL))) > 0)
		if (ret > 0)
		{
			read_finish = 1;
			if (FD_ISSET(fd, &rfds))
			{
				ioctl(fd, FIONREAD, &rcount);
				printf("Received %d bytes to be read!\n", rcount);
				while((rrcount = read(fd, line, 1023)) > 0)
				{
					ioctl(fd, FIONREAD, &rcount);
					printf("Keep %d bytes to be read!\n", rcount);

					tcount += rrcount;
					line[1024] = '\0';
					printf("[%s][%d]: %s\n", __func__, rrcount, line);
					if (rcount == 0)
					{
						FD_SET(fd, &rfds);
						goto again;
					}
				}

				printf("Totally read count: %d\n", tcount);
				FD_CLR(fd, &rfds);
				fd_map[fd] = 0;
				pclose(fp);
			}

			if (FD_ISSET(fd1, &rfds))
			{
				ioctl(fd1, FIONREAD, &rcount);
				while((rrcount = read(fd1, line, 1023)) > 0)
				{
					line[1024] = '\0';
					printf("[%s][%d]: %s\n", __func__, rrcount, line);
				}
				FD_CLR(fd1, &rfds);
				fd_map[fd1] = 0;
				pclose(fp1);
			}

			if (FD_ISSET(fd2, &rfds))
			{
				ioctl(fd2, FIONREAD, &rcount);
				while((rrcount = read(fd2, line, 1023)) > 0)
				{
					line[1024] = '\0';
					printf("[%s][%d]: %s\n", __func__, rrcount, line);
				}
				FD_CLR(fd2, &rfds);
				fd_map[fd2] = 0;
				pclose(fp2);
			}

			for (i = 0; i <= FD_SETSIZE; i++)
			{
				if (fd_map[i] == 1)
				{
					read_finish = 0;
					FD_SET(i, &rfds);
				}
			}

			if (read_finish)
				break;
		}
		else if (ret == 0)
			printf("Call select expired!\n");
		else
			printf("Call select Error with message: %s\n", strerror(errno));
	}

	return;
}

int main(int argc, char **argv)
{
	call_popen();
}
