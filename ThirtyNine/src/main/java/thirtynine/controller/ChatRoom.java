package thirtynine.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import thirtynine.pojo.Speech;
import thirtynine.pojo.User;
import thirtynine.service.SpeechService;

import javax.servlet.http.HttpSession;
import java.util.Collection;

@Controller
public class ChatRoom {

    @Autowired
    private SpeechService speechService;

    @RequestMapping({"/","/ChatRoom"})
    public String toChatRoom(Model model){
        Collection<Speech> speeches =speechService.findAll();
        model.addAttribute("speeches",speeches);
        return "ChatRoom";
    }

    @PostMapping("/speak")
    public String speak(Speech speech, HttpSession session){
        speechService.insertSpeech(speech);
        return "redirect:ChatRoom";
    }
}
