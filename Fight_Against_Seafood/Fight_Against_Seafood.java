import javax.swing.*;
import java.awt.*;
import java.awt.event.*;
import java.util.Timer;
import java.util.TimerTask;

public class Fight_Against_Seafood{
	public static void main(String[] args){
		panorama pano=new panorama();
		pano.draw_panorama();
		
	}
}

class character{

	int x;  	//location
	int y;
	boolean left;  // mark its direction;
	int pic_index;	//to draw player:0 for right stand, 1 for left ; 2 and 3 for right walk,
	Image pic;		//4 and 5 for left ; 6 for right attack ,7 for left; 8 for right jump, 9 for left ;
	boolean being_attacked;
	boolean attacking;

	int attack;    // attribute of character 
	int MAX_HP;
	int HP;
	
	
}

class Monster extends character{
	Monster(){
		pic_index=1;
		x=700;y=440;
		left=true;
		attack =5;
		MAX_HP=30;
		HP=MAX_HP;
		
	}
	
	public void death(){
		pic_index=0;	// change to coin
		x=x+30;
		y=y+70;
	}

	public void move(JPanel jp,Player player){
		Runnable thread_job = new monster_move(jp,player);
		Thread mon = new Thread(thread_job);
		mon.start();
	}
	public void newone(JPanel jp,Player player){
		
		pic_index++;
		MAX_HP=(int)(MAX_HP*1.2);
		HP=MAX_HP;
		x=700;y=440;
		left=true;
		attack = (int)(attack*1.2);
		move(jp,player);
	}

	class monster_move implements Runnable{
		JPanel jp;
		Player player;
		monster_move(JPanel jp , Player player){
			this.jp=jp;
			this.player=player;
		}
		public void run(){
			while(pic_index!=0){
				if(left){
					x=x-2;
					jp.repaint();
					if(x<=0){
						x=0;
						left =false;
					}
					try{
						Thread.sleep(10);
					}catch(Exception ex){
						ex.printStackTrace();
					}
				}
				if(!left){
					x=x+2;
					jp.repaint();
					if(x>1000){
						x=1000;
						left=true;
					}
					try{
						Thread.sleep(10);
					}catch(Exception ex){
						ex.printStackTrace();
					}
					
				}
				if(((int)(Math.random()*200))==1){		//attack
					attacking = true;
					if((player.y-y)>=-30){
						if((left&&(x-player.x)>=70&&(x-player.x)<=330)||(!left&&(x-player.x)>=-150&&(x-player.x)<=120)){
							player.HP=player.HP-attack;
							if(player.HP<=0){
								player.HP=0;		//death here;
								player.death=true;
							}
						}
					}
					for(int i = 0;i<10;++i){
						if(left)x=x-2;
						else x=x+2;
						jp.repaint();
						try{
							Thread.sleep(30);
						}catch(Exception ex){
							ex.printStackTrace();
						}
					}
					attacking = false;
				}

			
			}

		}
	}
}

class Player extends character{
	
	boolean on_platform;
	boolean A_pressed;
	boolean D_pressed;
	boolean death;

	int MAX_MP;
	int MP;
	int MAX_XP;
	int XP;
	int money ;
	double HP_recovery;
	double MP_recovery;
	int killed;

	Player(){		//set initial value
		x=30;
		y=430;
		pic_index=0;	
		MAX_HP=100;
		HP=MAX_HP;
		MAX_MP=100;
		MP=0;
		attack=10;
		money = 20;
		HP_recovery=0.01;
		MP_recovery=0.01;
	}	
	
	public void recovery(JPanel jp){
		Runnable thread_job = new player_recovery(jp);
		Thread recover = new Thread(thread_job);
		recover.start();
	}


	
	class player_recovery implements Runnable{
		JPanel jp;
		player_recovery(JPanel jp){
			this.jp=jp;
		}
		public void run(){
			while(HP!=0){
				HP=(int)(HP+HP_recovery*MAX_HP);
				MP=(int)(MP+MP_recovery*MAX_MP);
				if(HP>=MAX_HP)HP=MAX_HP;
				if(MP>=MAX_MP)MP=MAX_MP;
				jp.repaint();
				try{
					Thread.sleep(2000);
				}catch(Exception ex){
					ex.printStackTrace();
				}
			}
		}
	}
	
	public void stand(){
		if(left) pic_index=1;	
		else pic_index=0;
	}
	public void walk_right(JPanel jp){
		x=x+10;
		if(x>900)x=900;	//in case of over the boundary
		left = false;
		pic_index++;
		if(pic_index>3||pic_index<2)pic_index=2;
		
		if(x>760&&on_platform){				//fall down from the platform
			on_platform = false;
			fall_platform(this,jp);
		}
	}
	public void walk_left(JPanel jp){
		x=x-10;
		if(x<-50)x=-50;	//in case of over the boundary
		left = true;
		pic_index++;
		if(pic_index>5||pic_index<4)pic_index=4;
		
		if(x<340&&on_platform){			
			on_platform=false;
			fall_platform(this,jp);
		}
	}
	public void to_attacking(){
		attacking = true;
		if(left)pic_index=7;
		else pic_index=6;
	}
	
	public void jump(JPanel jp,Player player){
		class jump extends TimerTask{
			boolean down;
			public void run(){
				if(player.on_platform){			//on the platform
					if(A_pressed)player.x=x-5;
					if(x<-50)x=-50;
					if(D_pressed)player.x=x+5;
					if(x>900)x=900;
					if(down==false){
						player.y=player.y-10;
						jp.repaint();	
						if(player.y<=-55)down=true;
					}else {		
						player.y=player.y+10;
						jp.repaint();
						if(player.y>=115){
							player.y=115;
							if(player.left){		
								player.pic_index=1;
							}else{
								player.pic_index=0;
							}
							down = false;
							if(x<340||x>760){
								on_platform=false;
								fall_platform(player,jp);
							}
							jp.repaint();
							this.cancel();
						}
					}
				}else{
					if(player.x<600||player.x>720){		//in the ground
						if(A_pressed)player.x=x-5;
						if(x<-50)x=-50;
						if(D_pressed)player.x=x+5;
						if(x>900)x=900;
						if(down==false){
							player.y=player.y-10;
							jp.repaint();	
							if(player.y<=230)down=true;
						}else {		
							player.y=player.y+10;
							jp.repaint();
							if(player.y>=430){
								player.y=430;
								if(player.left){		
									player.pic_index=1;
								}else{
									player.pic_index=0;
								}
								down = false;
								jp.repaint();
								this.cancel();
							}
						}
	
					}else{				//in the spring
						if(player.D_pressed==false&&player.A_pressed==false){
							if(down==false){
								player.y=player.y-10;
								jp.repaint();	
								if(player.y<=90)down=true;
							}else{		
								player.y=player.y+10;
								jp.repaint();
								if(player.y>=115){
									player.y=115;
									if(player.left){		
										player.pic_index=1;
									}else{
										player.pic_index=0;
									}
									down = false;
									jp.repaint();
									player.on_platform =true;
									this.cancel();
								}
							}
						}else{

							if(A_pressed)player.x=x-5;
							if(x<-50)x=-50;
							if(D_pressed)player.x=x+5;
							if(x>900)x=900;
							if(down==false){
								player.y=player.y-10;
								jp.repaint();	
								if(player.y<=230)down=true;
							}else {		
								player.y=player.y+10;
								jp.repaint();
								if(player.y>=430){
									player.y=430;
									if(player.left){		
										player.pic_index=1;
									}else{
										player.pic_index=0;
									}
									down = false;
									jp.repaint();
									this.cancel();
								}
							}
						
						}
					}
				}	
			}
		}	
		if(player.left)player.pic_index=9;
		else player.pic_index=8;
		Timer jump = new Timer();
		jump.schedule(new jump(),0,15);
	}

	public void fall_platform(Player player , JPanel jp){
		class fall extends TimerTask{
			public void run(){
				if(A_pressed)player.x=x-5;
				if(x<-50)x=-50;
				if(D_pressed)player.x=x+5;
				if(x>900)x=900;
				player.y=player.y+10;
				jp.repaint();
				if(player.y>=430){
					player.y=430;
					if(player.left){		
						player.pic_index=1;
					}else{
						player.pic_index=0;
					}
					jp.repaint();
					this.cancel();
				}
			}
		}

		if(player.left)player.pic_index=9;
		else player.pic_index=8;
		Timer fall = new Timer();
		fall.schedule(new fall(),0,15);
	}
}


class shop{
	Player player;
	JLabel MAX_HP_label;
	JLabel attack_label;
	JPanel jp;
	
	shop(Player player,JPanel jp){
		this.player=player;
		this.jp=jp;
	}
	public void open(){
		JFrame frame = new JFrame();
		frame.setVisible(true);
		frame.setSize(500,300);
		JPanel sp = new JPanel();
		Color LightYellow  = new Color(255,248,220);
		sp.setBackground(LightYellow);
		frame.getContentPane().add(sp);
				
		MAX_HP_label = new JLabel("            MAX_HP: "+player.MAX_HP+"           ");
		MAX_HP_label.setFont(new java.awt.Font("Dialog", 1, 24));
		if(player.MAX_HP<200)MAX_HP_label.setForeground(Color.black);
		if(player.MAX_HP>600&&player.MAX_HP<1000)MAX_HP_label.setForeground(Color.blue);
		Color purple = new Color(128,0,128);
		if(player.MAX_HP>3000&&player.MAX_HP<5000)MAX_HP_label.setForeground(purple);
		if(player.MAX_HP>10000)MAX_HP_label.setForeground(Color.red);
		sp.add(MAX_HP_label);
		
		JButton hp_button = new JButton("remold");
		sp.add(hp_button);
		hp_button.addActionListener(new hp_change());

		attack_label = new JLabel("            attack: "+player.attack+"             ");
		attack_label.setFont(new java.awt.Font("Dialog", 1, 24));
		if(player.attack<30)attack_label.setForeground(Color.black);
		if(player.attack>100&&player.attack<200)attack_label.setForeground(Color.blue);
		if(player.attack>500&&player.attack<800)attack_label.setForeground(purple);
		if(player.attack>2000)attack_label.setForeground(Color.red);
		sp.add(attack_label);

		JButton attack_button = new JButton("remold");
		sp.add(attack_button);
		attack_button.addActionListener(new attack_change());
	}	
	class hp_change implements ActionListener{
		public void actionPerformed(ActionEvent event){
			if(player.money>=10){
				MAX_HP_label.setText("            MAX_HP: "+player.MAX_HP+"          ");
				player.money=player.money-10;
				int quality =(int)(Math.random()*100);
				if(quality<60)player.MAX_HP=(int)(Math.random()*100+100);
				if(quality>=60&&quality<85)player.MAX_HP=(int)(Math.random()*400+600);
				if(quality>=85&&quality<95)player.MAX_HP=(int)(Math.random()*2000+3000);
				if(quality>=95&&quality<=100)player.MAX_HP=(int)(Math.random()*10000+10000);
		
				if(player.MAX_HP<=200)MAX_HP_label.setForeground(Color.black);
				if(player.MAX_HP>=600&&player.MAX_HP<=1000)MAX_HP_label.setForeground(Color.blue);
				Color purple = new Color(128,0,128);
				if(player.MAX_HP>=3000&&player.MAX_HP<=5000)MAX_HP_label.setForeground(purple);
				if(player.MAX_HP>=10000)MAX_HP_label.setForeground(Color.red);
				MAX_HP_label.setText("            MAX_HP: "+player.MAX_HP+"          ");
			}else{
				MAX_HP_label.setText("      Insufficient money         ");
			}
			jp.repaint();
		}
		
	}
	class attack_change implements ActionListener{
		public void actionPerformed(ActionEvent event){
			if(player.money>=10){
				attack_label.setText("            attack: "+player.attack+"             ");
				player.money=player.money-10;
				int quality =(int)(Math.random()*100);
				if(quality<60)player.attack=(int)(Math.random()*30);
				if(quality>=60&&quality<85)player.attack=(int)(Math.random()*100+100);
				if(quality>=85&&quality<95)player.attack=(int)(Math.random()*300+500);
				if(quality>=95&&quality<=100)player.attack=(int)(Math.random()*2000+2000);

				if(player.attack<30)attack_label.setForeground(Color.black);
				if(player.attack>100&&player.attack<200)attack_label.setForeground(Color.blue);
				Color purple = new Color(128,0,128);
				if(player.attack>500&&player.attack<800)attack_label.setForeground(purple);
				if(player.attack>2000)attack_label.setForeground(Color.red);
				attack_label.setText("            attack: "+player.attack+"             ");
			}else{
				attack_label.setText("         Insufficient money         ");
			}
			jp.repaint();
		}	
	}
}

class panorama{

	Player player ;		//component
	Monster monster;
	scenery sce;

	panorama(){			
		player = new Player();
		monster = new Monster();
	}
	
	
	public void draw_panorama(){		//create the window and get it the panel 
		JFrame frame = new JFrame();
		frame.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		frame.setVisible(true);
		frame.setSize(1200,700);

		sce = new scenery();  
		frame.getContentPane().add(sce);
		frame.addKeyListener(sce);
		frame.addMouseListener(sce);
		
		monster.move(sce,player);
		player.recovery(sce);
	}

	public void fight(){
		if((player.y-monster.y)>=-30){
			if(player.left&&(player.x-monster.x)<=130&&(player.x-monster.x)>=-60){
				if(player.attacking){
					monster.being_attacked = true;
					sce.repaint();
					monster.x=monster.x-10;
					monster.HP=monster.HP-player.attack;
					if(monster.HP<=0){
						monster.death();
						monster.HP=0;
						player.killed++;
					}
				}
			}
			if(!player.left&&(monster.x-player.x)>=120&&(monster.x-player.x)<=320){
				if(player.attacking){
					monster.being_attacked = true;
					sce.repaint();
					monster.x=monster.x+10;
					monster.HP=monster.HP-player.attack;
					if(monster.HP<=0){
						monster.HP=0;
						monster.death();
						player.killed++;
					}
				}
			}

		}
	}
	
	class scenery extends JPanel implements KeyListener,MouseListener{
	
		public void paintComponent(Graphics g){

			Graphics2D g2d = (Graphics2D)g;  			//background 
			Color skyblue = new Color(135,206,235);
			Color wheat = new Color(245,222,179);
			GradientPaint gradient = new GradientPaint(600,0,skyblue,600,700,wheat);
			g2d.setPaint(gradient);
			g2d.fillRect(0,0,this.getWidth(),this.getHeight());

			Image menu = new ImageIcon("menu.png").getImage();		//draw menu;
			g.drawImage(menu,1100,20,this);

			Color grown = new Color(244,164,96); 		//paint ground
			g.setColor(grown);

			g.fillRect(0,575,this.getWidth(),this.getHeight());
			Image platform = new ImageIcon("platform.png").getImage();	//draw platform
			g.drawImage(platform,500,250,this);

			Image spring = new ImageIcon("spring.png").getImage();	//draw spring
			g.drawImage(spring,750,570,this);

			Image house = new ImageIcon("house.png").getImage();	//draw the house
			g.drawImage(house,520,80,this);
			if((player.x>=420&&player.x<=580)&&player.y==115){
				g.setFont(new Font("Tahoma", Font.BOLD,22));
				Color gray = new Color(105,105,105);
				g.setColor(gray);
				g.drawString("press S to come in",550,80);
			}
			
			Color grey = new Color(211,211,211);		//draw the player's attribute
			g.setFont(new Font("Tahoma", Font.BOLD,24));
			Color gray = new Color(105,105,105);
			g.setColor(gray);
			g.drawString("HP",10,50);				
			g.drawString("MP",10,80);
			g.setColor(grey);					
			g.fillRect(50,30,300,20);
			g.fillRect(50,60,300,20);
			Color red = new Color(255,0,0);
			g.setColor(red);
			g.fillRect(50,30,(int)(300*(player.HP/(float)player.MAX_HP)),20);
			Color blue = new Color(0,0,255);
			g.setColor(blue);
			g.fillRect(50,60,(int)(300*(player.MP/(float)player.MAX_MP)),20);
			g.setFont(new Font("Tahoma", Font.BOLD,15));
			g.setColor(gray);
			g.drawString(player.HP+"/"+player.MAX_HP,360,45);
			g.drawString(player.MP+"/"+player.MAX_MP,360,75);
			Image coin = new ImageIcon("coin.png").getImage();
			g.drawImage(coin,10,100,this);
			Color gold = new Color(255,165,0);
			g.setColor(gold);
			g.setFont(new Font("Tahoma", Font.BOLD,30));
			g.drawString(""+player.money,60,130);
			g.setColor(red);
			g.drawString("killed: "+player.killed,150,130);

			player.pic = new ImageIcon("player_"+player.pic_index+".png").getImage();	//draw player
			g.drawImage(player.pic,player.x,player.y,this);
			
			monster.pic = new ImageIcon("monster_"+monster.pic_index+".png").getImage();	//
			g.drawImage(monster.pic,monster.x,monster.y,this);						//draw monster
			if(monster.pic_index!=0){
				g.setColor(gray);
				g.fillRect(monster.x,monster.y-10,150,10);
				g.setColor(red);
				g.fillRect(monster.x,monster.y-10,(int)(150*(monster.HP/(float)monster.MAX_HP)),10);
			}

			if(monster.being_attacked){
				if(player.left){
					Image hit= new ImageIcon("hit.png").getImage();	 //draw hit effect
					g.drawImage(hit,monster.x+80,monster.y+50,this);
				}else{
					Image hit= new ImageIcon("hit.png").getImage();	
					g.drawImage(hit,monster.x,monster.y+50,this);
				}
			}
			if(monster.pic_index==0){						//eat coin
				if(monster.x-player.x>=80&&monster.x-player.x<=180&&monster.y-player.y<=100){
					player.money=player.money+player.killed;
					monster.newone(this,player);
				}
			}
			
			if(monster.attacking){
				if(monster.left){
					Image attack= new ImageIcon("attack_1.png").getImage();	 //draw hit effect
					g.drawImage(attack,monster.x-100,monster.y+20,this);
				}else{
					Image attack= new ImageIcon("attack_0.png").getImage();	
					g.drawImage(attack,monster.x+80,monster.y+20,this);
				}
			}
			if(player.death){
				g.setColor(Color.red);
				g.setFont(new Font("Tahoma", Font.BOLD,150));
				g.drawString("You died",200,400);
			}	
		
		}
		
		public void keyTyped(KeyEvent e){ 	//deal with player's input
			
		}

		public void keyPressed(KeyEvent e){
			if(!player.death){
				int key = e.getKeyCode();
				if(key==KeyEvent.VK_D){
					player.walk_right(this);
					player.D_pressed=true;
					
				}
				if(key==KeyEvent.VK_A){
					player.walk_left(this);
					player.A_pressed=true;
				}
				if(key==KeyEvent.VK_W||key==KeyEvent.VK_K){		//ready to jump
					player.jump(this,player);
				}
				if(key==KeyEvent.VK_J){
					player.to_attacking();
					if(monster.pic_index!=0)
						fight();  //everytime press attack , judge fight;
				}
				if(key==KeyEvent.VK_S&&player.x>=420&&player.x<=580&&player.y==115){
					shop sp = new shop(player,this);
					sp.open();
				}
				repaint();
			}
		}

		public void keyReleased(KeyEvent e){
			int key = e.getKeyCode();
			if(key==KeyEvent.VK_K||key==KeyEvent.VK_W){		
							//jump ,no rush to stand standard
			}else if(key==KeyEvent.VK_A||key==KeyEvent.VK_D){
				player.A_pressed=false;		player.D_pressed=false;		
					// stand there keep its position	
			}else if (key==KeyEvent.VK_J){
				player.stand();
				player.attacking =false;
				monster.being_attacked=false;
			}else{
				player.stand();
			}

			repaint();
		}
		
		public void mouseClicked(MouseEvent e){
			int x = e.getX();
			int y = e.getY();
			if(x>=1115&&x<=1170&&y>=55&&y<=115){
				Menu menu = new Menu();
				menu.open();
			}
		}

		public void mousePressed(MouseEvent e){
		}
		public void mouseReleased(MouseEvent e){
		}
		public void mouseEntered(MouseEvent e){
		}
		public void mouseExited(MouseEvent e){
		}

	}
}

class Menu {
	JTextArea help = new JTextArea(14,50);
	public void open(){
		JFrame frame = new JFrame();
		frame.setVisible(true);
		frame.setSize(700,400);
		JPanel menu = new JPanel();
		Color LightYellow  = new Color(255,248,220);
		menu.setBackground(LightYellow);
		frame.getContentPane().add(menu);
		
		JScrollPane scroller = new JScrollPane(help);
		help.setLineWrap(true);
		
		scroller.setVerticalScrollBarPolicy(ScrollPaneConstants.VERTICAL_SCROLLBAR_ALWAYS);
		scroller.setHorizontalScrollBarPolicy(ScrollPaneConstants.HORIZONTAL_SCROLLBAR_NEVER);
		
		menu.add(scroller);
		help.setBackground(LightYellow);
		help.setForeground(Color.blue);
		Font font = new Font("黑体",Font.PLAIN,20);
		help.setText("   A向左走;D向右走;j攻击;W或K跳跃\n   怪物每次出现会更强(攻击力和血量是之前的1.2倍)\n");
		help.append("   生命和法术回复都是每两秒回复最大值的百分之一\n");
		help.append("   打死怪物掉落金币,怪物金币掉落数量等于杀死的怪物数量\n");
		help.append("   唯一提升战斗力的方法在小屋里是花10金币改造一下属性(remold),刷属性可能变好也可能变差\n");
		help.append("   各种品质概率如下:普通(黑色)60%;稀有(蓝色)25%;史诗(紫色)10%;传说(红色)5%\n");
		help.append("   目前没有技能体系\n");
		help.setFont(font);
	}
}
