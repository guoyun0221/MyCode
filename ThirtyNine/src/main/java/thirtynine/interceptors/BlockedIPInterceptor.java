package thirtynine.interceptors;

import org.springframework.beans.factory.BeanFactory;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.context.support.WebApplicationContextUtils;
import org.springframework.web.servlet.HandlerInterceptor;
import thirtynine.pojo.Blocked_IP;
import thirtynine.service.BlockedIPService;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import java.util.List;

public class BlockedIPInterceptor implements HandlerInterceptor {

    @Autowired
    private BlockedIPService blockedIPService;

    @Override
    public boolean preHandle(HttpServletRequest request, HttpServletResponse response, Object handler) throws Exception {
        if(blockedIPService==null){
            BeanFactory factory= WebApplicationContextUtils.getRequiredWebApplicationContext(request.getServletContext());
            blockedIPService = (BlockedIPService) factory.getBean("BlockedIPService");
        }
        List<Blocked_IP> IPs = blockedIPService.getAll();
        String userIP= request.getRemoteAddr();
        for(Blocked_IP ip:IPs){
            if(userIP.equals(ip.getIP())){
                request.getSession().setAttribute("message","BlockedIP");
                request.getRequestDispatcher("/error").forward(request,response);
                return false;
            }
        }
        return true;
    }
}
