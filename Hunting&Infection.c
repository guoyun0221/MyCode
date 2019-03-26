#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <windows.h>
#include <conio.h>

#define xsize 110
#define ysize 25
#define MaxNameLength 9   //演化模式 
#define unavailable 391   //拯救模式 
#define hunternamelength 6  //拯救模式 
#define preynamelength 4  //拯救模式 
#define playernamelength 8   //拯救模式

typedef struct{
	int isplayer;//用来标记身份 
	int ishunter;
	int isprey;
	char name[MaxNameLength];
	int namelength;		//后面防止字符串闪现和判断caught用 
	int x ;
	int y ;
}object;  //角色的参数. x,y为其坐标	

void Rescue_Prey(void);
void Watch_Evolution(void);
int Name_Length(char a[]);
void Repaint(const object a[],const int num);
void Random_Move(int *x ,int *y);
void Get_player_Move(int *x ,int *y);
void Update_Location(object a[],int num);
void Reorder(object a[],int num);
int Encounter(object a[],int num,int *savecount);
void Player_Prey(object *prey,int *count);
void Hunter_Prey(object *a,object *b);
int If_None_prey(object a[],int num);

int main()
{
	system("color 74"); 
	srand((time(0)));
	int mode;
					
	printf("\n\t\t\t\tHunting And Infection\n\n");
	printf("猎物碰到猎人会被同化，你可以拯救猎物。小心别被猎人杀死.\n");	//首界面，
	printf("The prey will be infected when encountering hunters, you could save them. Watch out the hunters.\n");
	
	printf("\n选择模式:\n");			//选择模式 
	printf("Choose a way:\n");
	printf("1: 拯救猎物      \t2: 自然演化\n");
	printf("1: Rescueing the prey\t2: Watching the evolution\n");
	scanf("%d",&mode);

	if(mode==1){
		Rescue_Prey();
	}else if(mode==2){
		Watch_Evolution();
	}else{			//处理意外输入 
		printf("\n滚\nFuck off\n"); 
	}
	
	sleep(1);			//游戏结束 
	printf("\nPress enter key to end\n");
	system("pause");
	return 0 ;
}

void Rescue_Prey(void)
{	
	int hunternumber,preynumber;
	int i;
	int savedprey=0;
	int death = 0;
	
	system("cls");		//游戏说明 
	printf("\n\n按'WASD'移动。躲避猎人，拯救猎物。\n");	
	printf("Press 'WASD' to move. Stay away from hunters and save prey.\n\n---------------\n\n");
	
	printf("请输入猎人数量: \n");//参数设置 
	printf("Set the number hunters: \n");
	scanf("%d",&hunternumber);
	printf("\n请输入猎物数量:\n");
	printf("Set the number of prey:\n");
	scanf("%d",&preynumber);
	int number=hunternumber+preynumber+1; 
	object objects[number];
	
	for(i=0;i<hunternumber;++i){//初始化hunter
		objects[i].ishunter=1;
		objects[i].isplayer=0;
		objects[i].isprey=0;
		strcpy(objects[i].name,"hunter");//字符串是数组不能直接赋值 
		objects[i].namelength=hunternamelength;		//就不再去调用namelength的函数了，减少运算 
		objects[i].x=rand()%xsize;//随机生成位置坐标 
		objects[i].y=rand()%ysize;
	}
	for(i;i<number-1;++i){//初始化prey 
		objects[i].ishunter=0;
		objects[i].isplayer=0;
		objects[i].isprey=1;
		strcpy(objects[i].name,"prey");
		objects[i].namelength=preynamelength;
		objects[i].x=rand()%xsize;
		objects[i].y=rand()%ysize;
	}
		objects[i].ishunter=0;		//初始化玩家数据 
		objects[i].isplayer=1;
		objects[i].isprey=0;
		strcpy(objects[i].name,"*player*");
		objects[i].namelength=playernamelength;
		objects[i].x=rand()%xsize;
		objects[i].y=rand()%ysize;
		
	system("cls");
	system("color 84");//光标和背景一个颜色，隐藏光标 
	
	Reorder(objects,number);	//避免第一次乱画 
	while(1){	//无限根据位置清屏重画 
		printf("\t\t-------------------Saved prey: %d-----------------------\n",savedprey);//最上方显示 
		Repaint(objects,number);
		Update_Location(objects,number);	//怪物的随机运动，玩家的运动 ,排序 
		Sleep(50);
		if(Encounter(objects,number,&savedprey)){	//encounter返回death 
			system("cls");
			Repaint(objects,number);//看看怎么死的 
			sleep(1);
			system("cls");
			printf("\n---------\n你救下了%d只猎物,最开始有%d只。你死了。\n",savedprey,preynumber);	
			printf("You saved %d prey out of %d, and you died.\n",savedprey,preynumber);
			break;
		}	
		if(If_None_prey(objects,number)){  	//没有prey了，游戏结束 
			system("cls");
			Reorder(objects,number);//避免闪现 
			Repaint(objects,number);//看一下最后的情况 
			sleep(1);	
			system("cls");
			printf("\n----------\n你救下了%d只猎物,最开始有%d只。\n",savedprey,preynumber);	
			printf("You saved %d prey out of %d.\n",savedprey,preynumber);
			break;
		}
		Reorder(objects,number);//解决玩家a[i]吃掉preya[i-1]之后瞬间跑到a[i-2]后边的问题。 
		system("cls");
	}
}

void Watch_Evolution(void)
{
	char huntername[MaxNameLength];
	int hunternum;
	char preyname[MaxNameLength];
	int preynum;
	int i;//遍历用的 
	
	printf("\n请输入猎人名称: \n");	
	printf("Name hunters: \n") ;
	scanf("%s",&huntername);
	printf("\n请输入捕食者数量: \n");
	printf("Set the number of hunters\n");
	scanf("%d",&hunternum) ;
	
	printf("\n请输入猎物名称: \n");	
	printf("Name prey: \n");
	scanf("%s",&preyname);
	printf("\n请输入猎物数量: \n");
	printf("Set the number of prey: \n");
	scanf("%d",&preynum);
	
	int num=hunternum+preynum;  //数组大小为两种之和 
	object hunter_prey[num];   //生成猎人和猎物的数组 
	
	for(i=0;i<hunternum;++i){   //初始化hunter
		hunter_prey[i].isplayer=0;
		hunter_prey[i].ishunter=1;
		hunter_prey[i].isprey=0;
		strcpy(hunter_prey[i].name,huntername);  //字符串是数组不能直接赋值 
		hunter_prey[i].namelength=Name_Length(hunter_prey[i].name);
		hunter_prey[i].x=rand()%xsize;//随机生成位置坐标 
		hunter_prey[i].y=rand()%ysize;
	}
	for(i;i<num;++i){     //初始化prey 
		hunter_prey[i].isplayer=0;
		hunter_prey[i].isprey=1;
		hunter_prey[i].ishunter=0;
		strcpy(hunter_prey[i].name,preyname);
		hunter_prey[i].namelength=Name_Length(hunter_prey[i].name);
		hunter_prey[i].x=rand()%xsize;
		hunter_prey[i].y=rand()%ysize;
	}
	
	system("cls");
	system("color 84");//隐藏光标 
	
	Reorder(hunter_prey,num); //防止第一下打印出杂乱的一堆
	while(1){
		Repaint(hunter_prey,num);	//重画 
		Update_Location(hunter_prey,num);//随机运动，按位置排序  
		Encounter(hunter_prey,num,0);//判断是否抓住 //不会有savedprey,给个0;
		Sleep(120);
		system("cls");	//清屏 
	}
}

int Name_Length(char a[])
{
	int count=0;
	while(a[count]!='\0'){
		++count;
	}
	return count;
}

void Repaint(const object a[],const int num)
{
	int t = 0;
	int i = 0;
	int j = 0;
	if(a[0].x!=unavailable){
		for(i=0;i<a[0].y;++i){//打印a[0] 
			printf("\n");//用回车走到纵坐标位置 
		}
		for(j=0;j<a[0].x;++j){
			printf(" ");//用空格走到横坐标位置 
		}
		printf("%s",a[0].name);
	}
	for(t=1;t<num;++t){//根据坐标打印剩下的 //因为要判断i-1所以遍历从1开始 
		if(a[t].x!=unavailable){
			if(a[t-1].y<a[t].y){
				for(;i<a[t].y;++i){
					printf("\n");
				}
				for(j=0;j<a[t].x;++j){//不在同一行，从0开始打空格 
					printf(" ");
				}
				printf("%s",a[t].name);
			}else{	
				j=j+a[t-1].namelength;//把前一个字符串本身的长度算进去，避免闪现 
				for(;j<a[t].x;++j){//在同一行不从0打印空格，避免闪现
					printf(" ");
				}
				printf("%s",a[t].name);
			}
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

void Get_player_Move(int *x ,int *y)
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

void Update_Location(object a[],int num)
{
	int i;
	for(i=0;i<num;++i){//全部随机运动一次 
		if(a[i].isplayer==0){
			if(a[i].x!=unavailable){	//剔除已经被拯救的prey 
				Random_Move(&a[i].x,&a[i].y);
			}
		}
		if(a[i].isplayer==1){
			Get_player_Move(&a[i].x,&a[i].y);
		}
	}
	
	Reorder(a , num);//重新排序 
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

int Encounter(object a[],int num,int *savecount)
{
	int i;int death =0;
	for(i=0;i<num-1;++i){
		if(a[i].y-a[i+1].y>-2){  //纵坐标满足条件 
			if((a[i].x-a[i+1].x>=-a[i].namelength)&&(a[i].x-a[i+1].x<=a[i+1].namelength)){ //横坐标也满足条件(相距够近) 
				if(strcmp(a[i].name,a[i+1].name)){   //如果不是同类//strcmp，相同会返回0 
					if((a[i].ishunter==1&&a[i+1].isplayer==1)||(a[i+1].ishunter==1&&a[i].isplayer==1)){ 
						death = 1;
					}
					else if((a[i].isplayer==1&&a[i+1].isprey==1)||(a[i+1].isplayer==1&&a[i].isprey==1)){
						if(a[i].isprey==1){
							Player_Prey(&a[i],savecount);
						}else{
							Player_Prey(&a[i+1],savecount);
						}
					}
					else if((a[i].ishunter==1&&a[i+1].isprey==1)||(a[i+1].ishunter==1&&a[i].isprey==1)){
						if(a[i].ishunter==1){
							Hunter_Prey(&a[i+1],&a[i]);
						}else{
							Hunter_Prey(&a[i],&a[i+1]);
						}	
					}
					//其实不同物种相遇只有这三种情况，我没直接写else只是为了列举出来清晰好看 
				}
			}
		}
	}
	return death;
}

void Player_Prey(object *prey,int *count)//把saved prey移除屏幕 ，并且saved prey+1 
{
	prey->x=unavailable;
	prey->y=unavailable;
	++(*count);
}

void Hunter_Prey(object *a,object *b)//b是hunter ,a是prey
{
	int i;
	for(i=0;i<=b->namelength;++i){	//<=为了把'\0'也复制过去 
		a->name[i]=b->name[i];
	}
	a->namelength=b->namelength;
	a->isprey=0;//改属性 
	a->ishunter=1;
}

int If_None_prey(object a[],int num){
	int i;
	for(i=0;i<num;++i){
		if(a[i].isprey==1&&a[i].x!=unavailable){
			return 0;
		}
	}
	return 1;
}

