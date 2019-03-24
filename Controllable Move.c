#include <stdio.h>
#include <conio.h>
#include <windows.h>

#define xsize 110
#define ysize 27 

typedef struct{   
	char name[8];
	int x ;
	int y ;
}Player;  //角色的参数. x,y为其坐标	

void Repaint(char s[],int x,int y);
void Get_Move(int *x ,int *y);

int main()
{
	Player player;  //生成角色 
	printf("使用wasd控制移动\n");
	printf("请输入角色名称: ");
	scanf("%s",&player.name);
	player.x=xsize/2-3;//初始位置 
	player.y=ysize/3*2;
	 
	system("cls"); 
	
	while(1)	//无限清屏重画
	{
		Repaint(player.name,player.x,player.y);
		Get_Move(&player.x,&player.y);
		system("cls");
	}
	return 0;
}

void Repaint(char s[],int x,int y)
{					//根据坐标绘制玩家
	int i=0;
	for(i=0;i<y;++i){
		printf("\n");
	}
	for(i=0;i<x;++i){
		printf(" ");
	}
	printf("%s",s);
}

void Get_Move(int *x ,int *y)
{							 
	char a = getch();//wasd控制坐标变化 从而控制方向 
	
	if(a=='w') --(*y);
	if(a=='s') ++(*y);
	if(a=='a') (*x)-=2; //横向密集小幅度运动不明显 
	if(a=='d') (*x)+=2;
	
	if((*x)<0)(*x)=0;  //防止越界 
	if((*y)<0)(*y)=0;
	if((*x)>xsize)(*x)=xsize;
	if((*y)>ysize)(*y)=ysize;
}

