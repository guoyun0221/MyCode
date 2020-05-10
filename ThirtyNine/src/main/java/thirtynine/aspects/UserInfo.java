package thirtynine.aspects;

import org.aspectj.lang.annotation.Aspect;
import org.aspectj.lang.annotation.Before;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.stereotype.Component;
import org.springframework.web.context.request.RequestAttributes;
import org.springframework.web.context.request.RequestContextHolder;
import thirtynine.pojo.User;
import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpSession;


@Component
@Aspect
public class UserInfo {

    Logger logger = LoggerFactory.getLogger(getClass());

    @Before("execution(* thirtynine.controller.ChatRoom.speak(..))")
    public void record(){
        RequestAttributes requestAttributes = RequestContextHolder.getRequestAttributes();
        HttpServletRequest request =(HttpServletRequest) requestAttributes.resolveReference(RequestAttributes.REFERENCE_REQUEST);
        //now I know how to get session
        HttpSession session = (HttpSession) requestAttributes.resolveReference(RequestAttributes.REFERENCE_SESSION);
        logger.info("-----------speak--------------");
        logger.info("speaker_name: "+((User)session.getAttribute("user")).getName());
        logger.info("speaker_IP: "+request.getRemoteAddr());
    }
}
