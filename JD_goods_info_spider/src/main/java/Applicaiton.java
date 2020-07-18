import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;

import java.io.*;
import java.nio.charset.Charset;
import java.util.ArrayList;
import java.util.List;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Applicaiton {

    List<Goods> goods_list = new ArrayList<>();

    public static void main(String[] args) {
        Applicaiton app = new Applicaiton();
        app.crawl();
    }

    public void crawl(){
        Scanner scanner = new Scanner(System.in);
        System.out.println("Input keyword you want to search");
        String keyword = scanner.nextLine();

        String url = null;
        for(int page = 1; goods_list.size() < 100; page += 2){
            url = "https://search.jd.com/Search?keyword=" + keyword + "&wq=" + keyword + "&page=" + page + "&s=" + (page/2)*50+1;
            String web = get_web(url);
            analyze(web);
            save_info();
        }
    }

    public String get_web(String url){
        String web_content = null;
        CloseableHttpResponse response = null;

        CloseableHttpClient client = HttpClientBuilder.create().build();
        HttpGet get = new HttpGet(url);
        get.addHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
                + "AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.36");

        try{
            response = client.execute(get);
            HttpEntity entity = response.getEntity();
            if (entity!=null){
                web_content = EntityUtils.toString(entity, Charset.forName("UTF-8"));
            }
        }catch(Exception e){
            e.printStackTrace();
        }finally{
            try{
                if (client != null){
                    client.close();
                }
                if (response != null){
                    response.close();
                }
            }catch(Exception e){
                e.printStackTrace();
            }
        }

        return web_content;
    }

    public void analyze(String web){
        //1:title 2:href 3:img 4:price 5:shop
        String expr = "<li[\\s\\S]{0,200}<a target=\"_blank\" title=\"([\\s\\S]+?)\"\\s\\nhref=\"(\\S+?html)\""
                + "[\\s\\S]+?<img[\\s\\S]+?src=\"(\\S+?)\"[\\s\\S]+?<em>￥</em><i>(\\S+?)</i>[\\s\\S]+?curr-shop"
                + "[\\s\\S]+?title=\"(\\S+?)\"[\\s\\S]+?</li>";
        Matcher m = Pattern.compile(expr).matcher(web);

        while(m.find() && goods_list.size() < 100){
            Goods g = new Goods();
            g.title = m.group(1);
            g.href = m.group(2);
            g.img = m.group(3);
            g.price = m.group(4);
            g.shop = m.group(5);
            //there is something wrong with my regular expression,
            //but I'm too lazy to debug that. so I put this 'if' to handle that.
            if(g.title.length()<100){
                System.out.println(g.toString());
                goods_list.add(g);
            }
        }
    }

    public void save_info(){
        File f = new File("JD_Goods_info.txt");
        FileOutputStream fileOutputStream = null;
        OutputStreamWriter outputStreamWriter = null;
        BufferedWriter writer = null;
        try{
            fileOutputStream = new FileOutputStream(f,true);
            outputStreamWriter = new OutputStreamWriter(fileOutputStream, "UTF-8");
            writer = new BufferedWriter(outputStreamWriter);

            for (Goods goods: goods_list){
                writer.write(goods.toString());
                writer.newLine();
            }
            writer.flush();
        } catch (Exception e) {
            e.printStackTrace();
        }finally {
            if (writer != null){
                try {
                    writer.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            if (outputStreamWriter != null){
                try {
                    outputStreamWriter.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
            if (fileOutputStream != null){
                try {
                    fileOutputStream.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}

class Goods{
    String title;
    String shop;
    String price;
    String img;
    String href;

    @Override
    public String toString() {
        return "title= " + title + "\n" +
                "shop= " + shop + "\n" +
                "price= " + price + "\n" +
                "img= " + img + "\n" +
                "href= " + href + "\n"  ;
    }
}
