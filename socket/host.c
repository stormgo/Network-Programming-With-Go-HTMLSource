#include <stdio.h>
#include <sys/param.h>

#include <sys/types.h>
#include <sys/socket.h>
#include <netdb.h>
#include <netinet/in.h>
#include <arpa/inet.h>


#define SIZE (MAXHOSTNAMELEN+1)

int main(void)
{
    char name[SIZE];
    struct hostent *entry;

    if (gethostname(name, SIZE)
                != 0) {
        fprintf(stderr,
             "unknown name\n");
        exit(1);
    }
    printf("host name is %s\n",
            name);

    if ((entry =
        gethostbyname(name)) ==
              NULL) {
        fprintf(stderr,
        "no host name info\n");
        exit(2);
    }
    printf("offic. name: %s\n",
                entry->h_name);
    printf("address: %s\n",
           inet_ntoa(
       *(struct in_addr *) (entry->h_addr_list[0])));

    exit(0);
}
