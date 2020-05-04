package thirtynine.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.WebMvcConfigurer;
import thirtynine.interceptors.LoginInterceptor;
import thirtynine.interceptors.UserLevel;

@Configuration
public class MyMVCConfig implements WebMvcConfigurer {

    @Override
    public void addInterceptors(InterceptorRegistry registry) {
        registry.addInterceptor(new LoginInterceptor()).addPathPatterns("/**")
                .excludePathPatterns("/login");
        registry.addInterceptor(new UserLevel()).addPathPatterns("/ChatRoom");
    }
}
