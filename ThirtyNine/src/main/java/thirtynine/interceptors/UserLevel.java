package thirtynine.interceptors;

import org.springframework.beans.factory.BeanFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.context.support.WebApplicationContextUtils;
import org.springframework.web.servlet.HandlerInterceptor;
import org.springframework.web.servlet.ModelAndView;
import thirtynine.pojo.User;
import thirtynine.service.UserService;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * it's not easy to get session in aop, so I just use interceptor to deal this thing
 * whenever user visits chatroom, he gain one exp
 */
public class UserLevel implements HandlerInterceptor {

    @Autowired
    private UserService userService;

    @Override
    public void postHandle(HttpServletRequest request, HttpServletResponse response, Object handler, ModelAndView modelAndView) throws Exception {
        User user = (User)request.getSession().getAttribute("user");
        int exp = user.getExperience();
        int lv =user.getLevel();
        exp++;
        if(exp>=lv+1){
            //upgrade
            exp=0;
            lv++;
        }
        user.setExperience(exp);
        user.setLevel(lv);
        if (userService == null) {
            /*
             * The loading time of the interceptor is before the spring context,
             * so you need to get service manually
             */
            BeanFactory factory = WebApplicationContextUtils.getRequiredWebApplicationContext(request.getServletContext());
            userService = (UserService) factory.getBean("UserService");
        }
        userService.updateUser(user);
    }
}
