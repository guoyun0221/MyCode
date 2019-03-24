#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <windows.h>

void Endless_Stripe(void);
void Bilibili_TV(void);

int main()
{
	srand(time(0));
	int a = 0;
	printf("1无限循环条纹, 2小电视\n");
	scanf("%d",&a);
	
	switch(a)
	{
		case 1:
			Endless_Stripe();break;
		case 2:
			Bilibili_TV();break;
	}
	return 0 ;
}

void Endless_Stripe(void)
{
	while(1){
		printf("\t\t|\t        *       \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t       ***      \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t      *****     \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t     *******    \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t    *********   \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t   ***********  \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t  ************* \t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t|\t ***************\t|\t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
						//the middle
		printf("\t\t \t  ************* \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t   ***********  \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t    *********   \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t     *******    \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t      *****     \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
		printf("\t\t \t       ***      \t \t\t%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d%d\n",rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10,rand()%10);
			Sleep(50);
	}
}

void Bilibili_TV(void)
{
	printf("\n");
		Sleep(70);
	printf("              *                   *\n");
		Sleep(70);
	printf("               *                 *\n");
		Sleep(70);
	printf("                *               *\n");
		Sleep(70);
	printf("          * * * * * * * * * * * * * * *\n");
		Sleep(70);
	printf("          *                           *\n");
		Sleep(70);
	printf("          *      *            *       *\n");
		Sleep(70);
	printf("          *    *               *      *\n");
		Sleep(70);
	printf("          *  *                  *     *\n");
		Sleep(70);
	printf("          *                           *\n");
		Sleep(70);
	printf("          *      *    **    *         *\n");
		Sleep(70);
	printf("          *       *  *  *  *          *\n");
		Sleep(70);
	printf("          *        **    **           *\n");
		Sleep(70);
	printf("          *                           *\n");
		Sleep(70);
	printf("          * * * * * * * * * * * * * * *\n");
		Sleep(70);
	printf("      \n  ");
		Sleep(70);	
}

