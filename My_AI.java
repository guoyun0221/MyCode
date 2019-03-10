import java.util.Scanner;
public class My_AI
{
	public static void replacechar(String s)
	{
		String s1 = s.replaceAll("我","我也");
		String s2 = s1.replace('你','我');
		String s3 = s2.replace('吗','啊');
		String s4 = s3.replace('？','。');
		String s5 = s4.replace('几','3');
		String s6 = s5.replaceAll("什么","那个");
		String s7 = s6.replaceAll("名字","李华");
		String s8 = s7.replace('哪','这');
		String s9 = s8.replaceAll("傻逼","不傻逼");
		String s10 =s9.replaceAll("有没","");
		String s11=s10.replace('么','样');
		String s12=s11.replace('哈','嗯');
		String s13=s12.replaceAll("可以","不可以");
		String s14=s13.replaceAll("能","不能");
		String s15=s14.replaceAll("谁","我");
		String s16=s15.replaceAll("会","不会");
		String s17=s16.replace('怎','这');
		String s18=s17.replaceAll("他","你也");
		String s19=s18.replaceAll("今天","昨天");
		String s20=s19.replaceAll("真","确实");
		
		if(!s20.equals(s)){
			System.out.println("-"+s20+"\n");
		}else{
			System.out.println("-听不懂，滚!\n");
		}
	}
	public static void main(String[] args)
	{
		Scanner in = new Scanner (System.in);
		String s=in.nextLine();
		while(!s.equals("0")){
			replacechar(s);
			s=in.nextLine();
		}
	
	}
}
