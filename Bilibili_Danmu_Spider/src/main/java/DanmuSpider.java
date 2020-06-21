import com.alibaba.fastjson.JSON;
import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import org.apache.http.util.EntityUtils;
import org.w3c.dom.Document;
import org.w3c.dom.Element;
import org.w3c.dom.Node;
import org.w3c.dom.NodeList;
import org.xml.sax.InputSource;

import javax.xml.parsers.DocumentBuilder;
import javax.xml.parsers.DocumentBuilderFactory;
import java.io.*;
import java.nio.charset.StandardCharsets;
import java.util.Scanner;
import java.util.Set;

public class DanmuSpider {
    public static void main(String[] args) {
        DanmuSpider spider = new DanmuSpider();
        Scanner scanner = new Scanner(System.in);
        System.out.println("Input BV number(eg. BV1sz411e78d)");
        String BV = scanner.nextLine();
        scanner.close();
        String json_web = spider.get_web_content("https://api.bilibili.com/x/player/pagelist?bvid="+BV+"&jsonp=jsonp");
        String cid = spider.get_cid(json_web);
        String danmuku =  spider.get_web_content("https://comment.bilibili.com/"+cid+".xml");
        spider.get_danmu(danmuku,BV);
    }


    public String get_web_content(String url){
        CloseableHttpClient client = HttpClientBuilder.create().build();
        HttpGet get = new HttpGet(url);
        get.addHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.36");
        CloseableHttpResponse response = null;
        String web_content = null;

        try{
            response = client.execute(get);
            HttpEntity entity = response.getEntity();
            if (entity!=null){
                web_content = EntityUtils.toString(entity, StandardCharsets.UTF_8);
            }
        } catch (Exception e) {
            e.printStackTrace();
        }finally{
            try {
                if (client != null) {
                    client.close();
                }
                if (response != null) {
                    response.close();
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
        return web_content;
    }

    public String get_cid(String json){
        String cid = null;
        StringBuffer sb =null;
        JSONObject jb = JSONObject.parseObject(json);
        Set<String> jsonSet = jb.keySet();
        String innerJsonStr =  jb.getString("data");
        sb = new StringBuffer();
        JSONArray objects = JSON.parseArray(innerJsonStr);
        JSONObject obj = (JSONObject)objects.get(0);
        cid = obj.getString("cid");
        return cid;
    }

    public void get_danmu(String danmuku, String BV){
        DocumentBuilderFactory factory = DocumentBuilderFactory.newInstance();
        DocumentBuilder builder = null;
        Document doc = null;

        FileOutputStream fileOutputStream = null;
        OutputStreamWriter outputStreamWriter = null;
        BufferedWriter writer = null;

        try {
            fileOutputStream = new FileOutputStream(BV+".txt",true);
            outputStreamWriter = new OutputStreamWriter(fileOutputStream,"UTF-8");
            writer = new BufferedWriter(outputStreamWriter);

            builder = factory.newDocumentBuilder();
            doc = builder.parse(new InputSource(new StringReader(danmuku)));
            Element message = doc.getDocumentElement();
            NodeList list = message.getChildNodes();
            if (list != null){
                for (int i = 0; i < list.getLength(); i++){
                    Node node = list.item(i);
                    if(node.getNodeName() == "d"){
                        //appear time by sec
                        String t_sec = node.getAttributes().getNamedItem("p").getNodeValue();
                        int time_sec = Integer.valueOf(t_sec.substring(0,t_sec.indexOf(".")));
                        int min = time_sec/60;
                        int sec = time_sec%60;
                        writer.write(min+":"+sec+" ");
                        writer.write(node.getFirstChild().getNodeValue());
                        writer.newLine();
                    }
                }
            }
        } catch (Exception e) {
            e.printStackTrace();
        }finally {
            if (writer != null){
                try {
                    writer.flush();
                    writer.close();
                } catch (Exception e) {
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
            if(fileOutputStream != null){
                try {
                    fileOutputStream.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }
}
