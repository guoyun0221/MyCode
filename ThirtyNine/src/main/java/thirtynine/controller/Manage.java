package thirtynine.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import thirtynine.pojo.Speech;
import thirtynine.service.SpeechService;

import java.util.List;

@Controller
public class Manage {

    @Autowired
    private SpeechService speechService;

    @RequestMapping("/manage")
    public String listSpeech(Model model,Integer page){
        if(page==null){
            page=1;
        }
        Integer maxPage = speechService.getMaxPage();
        Speech topSpeech = speechService.findTop();
        List<Speech> speeches = speechService.findByPage(page);
        model.addAttribute("top",topSpeech);
        model.addAttribute("speeches",speeches);
        model.addAttribute("currentPage",page);
        model.addAttribute("maxPage",maxPage);
        return "manage";
    }

    @RequestMapping("/manage/delete")
    public String deleteSpeech(int id){
        speechService.deleteSpeech(id);
        return "redirect:/manage";
    }

    @RequestMapping("/manage/top")
    public String topSpeech(int id){
        speechService.topSpeech(id);
        return "redirect:/manage";
    }

    @RequestMapping("/manage/cancelTop")
    public String cancelTop(){
        speechService.cancelTop();
        return "redirect:/manage";
    }
}
