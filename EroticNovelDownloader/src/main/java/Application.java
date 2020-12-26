import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;

import java.io.*;
import java.nio.charset.Charset;
import java.util.Scanner;
import java.util.regex.Matcher;
import java.util.regex.Pattern;

public class Application {

    FileOutputStream fileOutputStream;
    OutputStreamWriter outputStreamWriter;
    BufferedWriter writer = null;

    public Application(){
        File f = new File("EroticNovel.txt");
        fileOutputStream = null;
        outputStreamWriter = null;
        try{
            fileOutputStream = new FileOutputStream(f,true);
            outputStreamWriter = new OutputStreamWriter(fileOutputStream, "UTF-8");
            writer = new BufferedWriter(outputStreamWriter);
        }catch (Exception e){
            e.printStackTrace();
        }
    }

    public static void main(String[] args) {
        Application app = new Application();
        app.crawl();
    }

    public void crawl(){
        Scanner scanner = new Scanner(System.in);
        System.out.println("base url: ");
        String baseUrl = scanner.nextLine();
        System.out.println("pages: ");
        int maxPage = scanner.nextInt();

        for(int page = 1; page <= maxPage; page ++){
            String web = get_web(baseUrl + page + ".htm");
            String str = analyze(web);
            save_info(str);
        }
        closeShit();
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
                web_content = EntityUtils.toString(entity, Charset.forName("GBK"));
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

    public String analyze(String web){
//        System.out.println(web);
        String expr = "<!--HTMLBUILERPART0-->([\\s\\S]+?)<!--/HTMLBUILERPART0-->";
        Matcher m = Pattern.compile(expr).matcher(web);

        if(m.find()){
            String str = m.group(1);
            str = str.replaceAll("<BR><BR>", "\n");
            System.out.println(str);
            return str;
        }
        return null;
    }

    public void save_info(String str){

        try{
            writer.write(str);
            writer.newLine();
            writer.flush();
        } catch (Exception e) {
            e.printStackTrace();
        }
    }
    public void closeShit(){
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
