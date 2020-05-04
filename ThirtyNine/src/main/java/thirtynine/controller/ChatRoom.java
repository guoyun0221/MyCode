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
    public String toChatRoom(Model model,HttpSession session){
        List<Speech> speeches =speechService.findAll();
        for(Speech speech: speeches){
            //to output multiple spaces and newlines
            speech.setWords(speech.getWords().replaceAll(" ","&nbsp;").replaceAll("\r","<br/>"));
            //to get username in thymeleaf
            if(speech.getUser()!=null){
                speech.setSpeaker(speech.getUser().getName());
            }
        }
        model.addAttribute("username",((User)session.getAttribute("user")).getName());
        model.addAttribute("speeches",speeches);
        return "ChatRoom";
    }

    @PostMapping("/speak")
    public String speak(Speech speech, HttpSession session){
        speech.setUser(((User)session.getAttribute("user")));
        speech.setSend_time(DateFormat.SqlDate());
        speechService.insertSpeech(speech);
        return "redirect:ChatRoom";
    }

    @PostMapping("/search")
    public String search(String keyword,Model model,HttpSession session){
        List<Speech> speeches = speechService.findByKeyword(keyword);
        for(Speech speech: speeches){
            speech.setWords(speech.getWords().replaceAll(" ","&nbsp;").replaceAll("\r","<br/>"));
            if(speech.getUser()!=null){
                speech.setSpeaker(speech.getUser().getName());
            }
        }
        model.addAttribute("username",((User)session.getAttribute("user")).getName());
        model.addAttribute("speeches",speeches);
        return "ChatRoom";
    }
}
