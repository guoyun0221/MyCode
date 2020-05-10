package thirtynine.interceptors;


import org.springframework.web.servlet.HandlerInterceptor;
import thirtynine.pojo.User;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

public class ManageInterceptor implements HandlerInterceptor {
    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        User user = (User)request.getSession().getAttribute("user");
        if(!user.getName().equals("root")){
            request.getSession().setAttribute("wrong","NoPermission");
            request.getRequestDispatcher("/wrong").forward(request,response);
            return false;
        }
        return true;
    }
}
