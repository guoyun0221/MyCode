package thirtynine.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import thirtynine.pojo.Speech;
import thirtynine.pojo.User;
import thirtynine.service.SpeechService;
import thirtynine.utils.DateFormat;
import javax.servlet.http.HttpSession;
import java.util.List;

@Controller
public class ChatRoom {

    @Autowired
    private SpeechService speechService;

    @RequestMapping({"/","/ChatRoom"})
    public String toChatRoom(Model model,HttpSession session,Integer page){
        if(page==null){
            page=1;
        }
        Integer maxPage = speechService.getMaxPage();
        List<Speech> speeches = speechService.findByPage(page);
        model.addAttribute("username",((User)session.getAttribute("user")).getName());
        model.addAttribute("speeches",speeches);
        model.addAttribute("currentPage",page);
        model.addAttribute("maxPage",maxPage);
        return "ChatRoom";
    }

    @PostMapping("/speak")
    public String speak(Speech speech, HttpSession session){
        speech.setUser(((User)session.getAttribute("user")));
        speech.setSend_time(DateFormat.SqlDate());
        speechService.insertSpeech(speech);
        return "redirect:ChatRoom";
    }

    @RequestMapping("/search")
    public String search(String keyword,Model model,HttpSession session){
        List<Speech> speeches = speechService.findByKeyword(keyword);
        model.addAttribute("username",((User)session.getAttribute("user")).getName());
        model.addAttribute("speeches",speeches);
        return "ChatRoom";
    }

    @RequestMapping("/toReply")
    public String toReply(int reply_id,Model model){
        Speech speech=speechService.findById(reply_id);
        model.addAttribute(speech);
        return "reply";
    }

    public String toReply(Speech speech,HttpSession session){
        speech.setUser(((User)session.getAttribute("user")));
        speech.setSend_time(DateFormat.SqlDate());
        speechService.insertSpeech(speech);
        return "redirect:ChatRoom";
    }
}
