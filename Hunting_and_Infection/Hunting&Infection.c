#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <windows.h>
#include <conio.h>

#define xsize 110
#define ysize 26
#define MaxNameLength 9   //演化模式 
#define unavailable 391   //拯救模式 
#define playernamelength 8  //拯救模式 + 逃亡模式 

typedef struct{
	int isplayer;//用来标记身份 
	int ishunter;
	int isprey;
	char name[MaxNameLength];
	int namelength;		//后面防止字符串闪现和判断caught用 
	int x ;
	int y ;
}object;  //角色的参数. x,y为其坐标	

void Flee(void);
void Rescue_Prey(void);
void Watch_Evolution(void);
int Name_Length(char a[]);
void initialize(object objects[],int hunternumber,int preynumber,int player,char huntername[],char preyname[]);
void Repaint(const object a[],const int num);
void Chase(object a[],int num);
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
					
	printf("\n\t\t\t\t\tHunting And Infection\n\n");
	printf("猎物碰到猎人会被同化。你可以拯救猎物。小心别被猎人杀死.\n");	//首界面，
	printf("The prey will be infected when encountering hunters, you could save them. Watch out the hunters.\n\n");
	printf("---------------------------\n");
	
	printf("选择模式:\n");			//选择模式 
	printf("Choose a way:\n\n");
	printf("1: 逃亡   \t2: 拯救猎物          \t3: 自然演化\n");
	printf("1: fleeing\t2: Rescuing the prey\t3: Watching the evolution\n");
	scanf("%d",&mode);

	if(mode==1){
		Flee();
	}else if(mode==2){
		Rescue_Prey();
	}else if(mode==3){			
		Watch_Evolution();
	}else{		//处理异常输入 
		printf("滚!\nFuck off!\n");
	}
	
	sleep(1);			//游戏结束 
	printf("\nPress enter key to end\n");
	system("pause");
	return 0 ;
}

void Flee(void)
{
	int hunternumber;
	int i;
	int death = 0;
	int score = 0;
	
	system("cls");			//游戏说明 
	printf("\n\n按'WASD'移动，别被猎人抓住。\n");
	printf("Press 'WASD' to move. Don't be caught.\n\n---------------------\n\n");
	
	printf("请输入猎人数量: \n");	//创建
	printf("Set the number of hunters\n");
	scanf("%d",&hunternumber);
	int num = hunternumber + 1; //玩家数据用一个元素存储 
	object hunter_player[num];
	
	initialize(hunter_player,hunternumber,0,1,"hunter","");//初始化 

	system("cls");
	system("color 84");//光标和背景一个颜色，隐藏光标 
	
	Reorder(hunter_player,num);	//避免第一次乱画 
	while(death!=1){		//不死就循环 
		printf("\n\t\t-----------------------Score: %d----------------------------\n",score);//最上方显示 
		Repaint(hunter_player,num);	 
		Chase(hunter_player,num); 		//怪物动作 
		for(i=0;i<num;++i){	// 	找到玩家元素 
			if(hunter_player[i].isplayer==1){
			 	Get_player_Move(&hunter_player[i].x,&hunter_player[i].y);//获取玩家操作 
		 		break;//减少运算
			}
		}
		death=Encounter(hunter_player,num,0);//没有拯救，给个0  
		++score;//每轮循环，即玩家每动一下，加一分 
		Reorder(hunter_player,num);//按坐标对数组元素排序 
		Sleep(50);
		system("cls");//清屏，准备下一次重画 
	}
			//死了跳出循环 
	printf("\n\t\t-----------------------Score: %d----------------------------\n",score);
	Repaint(hunter_player,num);//看看怎么死的
	sleep(1);
	system("cls");
	printf("\n------------------------\n你的得分是: %d\n",score);
	printf("Your scroe is: %d\n",score);
	
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
	
	printf("请输入猎人数量: \n");		//创建 
	printf("Set the number of hunters: \n");
	scanf("%d",&hunternumber);
	printf("\n请输入猎物数量:\n");
	printf("Set the number of prey:\n");
	scanf("%d",&preynumber);
	int number=hunternumber+preynumber+1; 
	object objects[number];
	
	initialize(objects,hunternumber,preynumber,1,"hunter","prey");//初始化 

	system("cls");
	system("color 84");//隐藏光标 
	
	Reorder(objects,number);	//避免第一次乱画 
	while(1){	//无限根据位置清屏重画 
		printf("\n\t\t----------------------Saved prey: %d--------------------------\n",savedprey);//最上方显示 
		Repaint(objects,number);
		Update_Location(objects,number);	//怪物的随机运动，玩家的运动 ,排序 
		if(Encounter(objects,number,&savedprey)){	//encounter返回death 
			system("cls");
			printf("\n\t\t----------------------Saved prey: %d--------------------------\n",savedprey);
			Repaint(objects,number);//看看怎么死的 
			sleep(1);
			system("cls");
			printf("\n\n--------------\n你救下了%d只猎物,最开始有%d只。你死了。\n",savedprey,preynumber);	
			printf("You saved %d prey out of %d, and you died.\n",savedprey,preynumber);
			break;
		}	
		if(If_None_prey(objects,number)){  	//没有prey了，游戏结束 
			system("cls");
			Reorder(objects,number);//避免闪现 
			printf("\n\t\t----------------------Saved prey: %d--------------------------\n",savedprey);
			Repaint(objects,number);//看一下最后的情况 
			sleep(1);	
			system("cls");
			printf("\n\n---------------\n你救下了%d只猎物,最开始有%d只。\n",savedprey,preynumber);	
			printf("You saved %d prey out of %d.\n",savedprey,preynumber);
			break;
		}
		Reorder(objects,number);//解决玩家a[i]吃掉preya[i-1]之后瞬间跑到a[i-2]后边的问题。 
		Sleep(50);
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
	
	initialize(hunter_prey,hunternum,preynum,0,huntername,preyname);//初始化各种类参数 
	
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

void initialize(object objects[],int hunternumber,int preynumber,int player,char huntername[],char preyname[])
{
	int num=hunternumber+preynumber+player;
	int i;
	
	for(i=0;i<hunternumber;++i){//初始化hunter
		objects[i].ishunter=1;
		objects[i].isplayer=0;
		objects[i].isprey=0;
		strcpy(objects[i].name,huntername);
		objects[i].namelength=Name_Length(huntername);;		
		objects[i].x=rand()%xsize;
		objects[i].y=rand()%ysize;
	}
	for(i=0;i<preynumber;++i){//初始化prey 
		objects[hunternumber+i].ishunter=0;
		objects[hunternumber+i].isplayer=0;
		objects[hunternumber+i].isprey=1;
		strcpy(objects[hunternumber+i].name,preyname);
		objects[hunternumber+i].namelength=Name_Length(preyname);;
		objects[hunternumber+i].x=rand()%xsize;
		objects[hunternumber+i].y=rand()%ysize;
	}
	if(player){		//有玩家参与的话 
		objects[num-1].ishunter=0;		//初始化玩家数据 
		objects[num-1].isplayer=1;		//用最后一个元素存储玩家数据 
		objects[num-1].isprey=0;
		strcpy(objects[num-1].name,"*player*");
		objects[num-1].namelength=playernamelength;
		objects[num-1].x=rand()%xsize;
		objects[num-1].y=rand()%ysize;
	}
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
		if(a[t].x!=unavailable){	//去掉被移除屏幕的 
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

void Chase(object a[],int num)
{
	int x,y;	
	int i ;
	
	for(i=0;i<num;++i){		//记录玩家位置
		if(a[i].isplayer==1){
			x=a[i].x;
			y=a[i].y;
		}
	}
	
	for(i=0;i<num;++i){
		if(rand()%2){	//随机从x或y方向趋近玩家 
			if(a[i].x<x) a[i].x+=2;//没有设置等号，如果是玩家自己，坐标不变 
			if(a[i].x>x) a[i].x-=2;
		}else{
			if(a[i].y<y) ++a[i].y;
			if(a[i].y>y) --a[i].y;
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
	
	if(a=='w'||a=='W') --(*y);
	if(a=='s'||a=='S') ++(*y);
	if(a=='a'||a=='A') (*x)-=2; //横向密集小幅度运动不明显 
	if(a=='d'||a=='D') (*x)+=2;
	
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
		if(a[i].y-a[i+1].y>-2){  //纵坐标满足条件 ,上下相差一行以内 
			if((a[i].x-a[i+1].x>=-a[i].namelength)&&(a[i].x-a[i+1].x<=a[i+1].namelength)){ //横坐标也满足条件(相距够近) 
				if(strcmp(a[i].name,a[i+1].name)){   //如果不是同类//strcmp，相同会返回0 
					if((a[i].ishunter==1&&a[i+1].isplayer==1)||(a[i+1].ishunter==1&&a[i].isplayer==1)){ 
						death = 1; //玩家遇到hunter就死 
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
			return 0; 		//只要遇到prey就直接返回0; 
		}
	}
	return 1;//没有prey返回1; 
}
