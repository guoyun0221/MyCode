import java.util.Scanner;
public class My_AI
{
	public static void replacechar(String s)
	{
		String s1="0";
		String s2="0";
		String s3="0";
		String s4="0";
		String s5="0";
		String s6="0";
		String s7="0";
		String s8="0";
		String s9="0";
		String s10="0";
		String s11="0";
		String s12="0";
		String s13="0";
		String s14="0";
		String s15="0";
		String s16="0";
		String s17="0";
		String s18="0";
		String s19="0";
		String s20="0";

		s1 = s.replaceAll("我","我也");
		s2 = s1.replace('你','我');
		s3 = s2.replace('吗','啊');
		s4 = s3.replace('？','。');
		s5 = s4.replace('几','3');
		s6 = s5.replaceAll("什么","那个");
		s7 = s6.replaceAll("名字","李华");
		s8 = s7.replace('哪','这');
		s9 = s8.replaceAll("傻逼","不傻逼");
		s10 =s9.replaceAll("有没","");
		s11=s10.replace('么','样');
		s12=s11.replace('哈','嗯');
		s13=s12.replaceAll("可以","不可以");
		s14=s13.replaceAll("能","不能");
		s15=s14.replaceAll("谁","我");
		s16=s15.replaceAll("会","不会");
		s17=s16.replace('怎','这');
		s18=s17.replaceAll("他","你也");
		s19=s18.replaceAll("今天","昨天");
		s20=s19.replaceAll("真","确实");
		
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
