package thirtynine;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.jdbc.DataSourceAutoConfiguration;

@SpringBootApplication
public class ThirtyNine {
    public static void main(String[] args) {
        SpringApplication.run(ThirtyNine.class,args);
    }
}
