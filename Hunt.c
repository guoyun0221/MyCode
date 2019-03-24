#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <windows.h>

#define xsize 110 
#define ysize 27 

typedef struct{   
	char name[8];
	int x ;
	int y ;
}object;  //角色的参数. x,y为其坐标	

void Repaint(const object a[],const int num);
void Random_Move(int *x ,int *y);
void Reorder(object a[],int num);
void Update_Location(object a[],int num);
void If_Caught(object a[],int num,char huntername[]);

int main()
{
	srand(time(0));
	char huntername[8];
	int hunternum;
	char preyname[8];
	int preynum;
	int i;//遍历用的 
	
	printf("\n请输入捕食者名称: ");	
	scanf("%s",&huntername);
	printf("\n请输入捕食者数量: ");
	scanf("%d",&hunternum) ;
	
	printf("\n请输入猎物名称: ");	
	scanf("%s",&preyname);
	printf("\n请输入猎物数量: ");
	scanf("%d",&preynum);
	
	int num=hunternum+preynum; //数组大小为两种之和 
	object objects[num]; //生成猎人和猎物的数组 
	
	for(i=0;i<hunternum;++i){//初始化hunter
		strcpy(objects[i].name,huntername);//字符串是数组不能直接赋值 
		objects[i].x=rand()%xsize;//随机生成位置坐标 
		objects[i].y=rand()%ysize;
	}
	for(i;i<num;++i){//初始化prey 
		strcpy(objects[i].name,preyname);
		objects[i].x=rand()%xsize;
		objects[i].y=rand()%ysize;
	}
	
	system("color E0");
	while(1){
		Repaint(objects,num);	//重画 
		Update_Location(objects,num);//随机运动，按位置排序 
		If_Caught(objects,num,huntername);//判断是否抓住 
		Sleep(150);
		system("cls");	//清屏 
	}
	return 0;
}

void Repaint(const object a[],const int num)
{
	int t = 0;
	int i = 0;
	int j = 0;
	
	for(i=0;i<a[0].y;++i){//打印a[0] 
		printf("\n");//用回车走到纵坐标位置 
	}
	for(j=0;j<a[0].x;++j){
		printf(" ");//用空格走到横坐标位置 
	}
	printf("%s",a[0].name);
	
	for(t=1;t<num;++t){//根据坐标打印剩下的 //因为要判断i-1所以遍历从1开始 
		if(a[t-1].y<a[t].y){
			for(;i<a[t].y;++i){
				printf("\n");
			}
			for(j=0;j<a[t].x;++j){//不在同一行，从0开始打空格 
				printf(" ");
			}
			printf("%s",a[t].name);
		}else{			
			for(;i<a[t].y;++i){
				printf("\n");
			}
			for(;j<a[t].x;++j){//在同一行不从0打印空格，避免闪现,但还是因为字符串本身的长度而有些闪现 
				printf(" ");
			}
			printf("%s",a[t].name);
		}
	}
}

void Random_Move(int *x ,int *y)
{ 								 //随机向一个方向运动 
	int a =rand()%4;
	switch(a)
	{
		case 0 : ++(*x);break;
		case 1 : --(*x);break;
		case 2 : ++(*y);break;
		case 3 : --(*y);break;
	}
	
	if((*x)<0)(*x)=0;//防止越界 
	if((*y)<0)(*y)=0;
	if((*x)>xsize)(*x)=xsize;
	if((*y)>ysize)(*y)=ysize;
}

void Reorder(object a[],int num)
{									//从小到大排序，y优先 
	int i = 0;
	int j = 0;
	int min;
	object t;	

	for(i=0;i<num;++i){      		//选择排序法 
		min=i;
		for(j=i;j<num;++j){
			if(a[min].y>a[j].y){
				min=j;
			}
			if(a[min].y==a[j].y&&a[min].x>a[j].x){
				min=j;
			}
		}
		if(min!=i){			//交换元素位置 
			t=a[min];
			a[min]=a[i];
			a[i]=t;
		}
	}
}

void Update_Location(object a[],int num)
{
	int i;
	for(i=0;i<num;++i){//全部随机运动一次 
		Random_Move(&a[i].x,&a[i].y);
	}
	
	Reorder(a ,num );//重新排序 
}

void If_Caught(object a[],int num,char huntername[])
{
	int i;
	for(i=0;i<num;++i){
		if((a[i].y-a[i+1].y>-2)&&(a[i].y-a[i-1].y<2)){  //纵坐标满足条件 
			if((a[i].x-a[i+1].x>-4)&&(a[i].x-a[i-1].x<4)){  //横坐标也满足条件(相距够近) 
				if(strcmp(a[i].name,a[i+1].name)){   //如果不是同类//strcmp，相同会返回0 
					if(strcmp(a[i].name,huntername)){  //如果a[i]不是hunter 
						a[i].name[0]='\0';//猎物死掉(不打印出来了) 
					}else{
						a[i+1].name[0]='\0';
					} 
				}
			}
		}	
	}
}

