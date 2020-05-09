package thirtynine.config;

import org.springframework.context.annotation.Configuration;
import org.springframework.web.servlet.config.annotation.InterceptorRegistry;
import org.springframework.web.servlet.config.annotation.ResourceHandlerRegistry;
import org.springframework.web.servlet.config.annotation.ViewControllerRegistry;
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

    @Override
    public void addResourceHandlers(ResourceHandlerRegistry registry) {
        registry.addResourceHandler("/uploadedFiles/**").addResourceLocations("file:"+System.getProperty("user.dir")+"/src/main/resources/static/uploadedFiles/");
        //On my linux server
        //registry.addResourceHandler("/uploadedFiles/**").addResourceLocations("file:/root/uploadedFiles/");
    }

    @Override
    public void addViewControllers(ViewControllerRegistry registry) {
        registry.addViewController("/wrong").setViewName("wrong");
    }
}
