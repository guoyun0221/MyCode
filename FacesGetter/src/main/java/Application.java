import org.apache.http.HttpEntity;
import org.apache.http.client.methods.CloseableHttpResponse;
import org.apache.http.client.methods.HttpGet;
import org.apache.http.impl.client.CloseableHttpClient;
import org.apache.http.impl.client.HttpClientBuilder;
import java.io.*;
import java.util.Random;
import java.util.Scanner;

public class Application {

    public static void main(String[] args) {
        Application app = new Application();

        Scanner scanner = new Scanner(System.in);
        System.out.println("input the number of pics you want to get");
        int count = scanner.nextInt();

        for (int i = 0; i < count; i++) {
            app.get_image();
        }
    }

    public void get_image(){
        String url = "https://thispersondoesnotexist.com/image";

        CloseableHttpClient client = HttpClientBuilder.create().build();
        HttpGet get = new HttpGet(url);
        get.addHeader("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) "
                + "AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.81 Safari/537.36");

        CloseableHttpResponse response = null;
        OutputStream os = null;

        try{
            response = client.execute(get);
            HttpEntity entity = response.getEntity();

            File imgs = new File("imgs");
            if (!imgs.exists()){
                imgs.mkdir();
            }
            String imageName = new Random().nextInt() + ".jpeg";
            File imageFile = new File(imageName);
            os = new FileOutputStream(imgs + "/" + imageFile);

            entity.writeTo(os);

            System.out.println(imageName);

        }catch(Exception e){
            e.printStackTrace();
        }finally {
            if(response != null){
                try {
                    response.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }

            if(os != null){
                try {
                    os.close();
                } catch (IOException e) {
                    e.printStackTrace();
                }
            }
        }
    }

}
