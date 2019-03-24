#include <stdio.h>
#include <conio.h>
#include <windows.h>

#define xsize 110
#define ysize 24

typedef struct{   
	int x ;
	int y ;
}Player;  //x,y为坐标 

void Paint_Player(const int x,const int y); 
void Get_Move(int *x ,int *y);

int main()
{
	srand(time(0));
	Player player;  
	player.x=xsize/2-3;//初始位置 
	player.y=ysize/3*2;
	
	while(1)	//无限清屏重画
	{
		printf("\t\t\t\t\t\twasd控制角色移动\n");
		Paint_Player(player.x,player.y);
		Get_Move(&player.x,&player.y);
		system("cls");
	}
	return 0;
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

void Paint_Player(const int x,const int y)
{									//在坐标处画角色 
	int i ,j;
	 
	for(i=0;i<y;++i){
		printf("\n");
	}
	
	for(j=0;j<x;++j){
		printf(" ");
	}
	printf("   @\n");
	for(j=0;j<x;++j){
		printf(" ");
	}
	printf("==[!]==\n");
	for(j=0;j<x;++j){
		printf(" ");
	}
	printf("  / \\\n");// 右斜杠是关于逃逸字符什么的，这样才能打出来 
	for(j=0;j<x;++j){
		printf(" ");
	}
	printf(" /   \\\n");
}
/*
		    
		   @
		==[!]==
		  / \
		 /   \
*/
