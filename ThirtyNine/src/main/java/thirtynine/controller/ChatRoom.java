package thirtynine.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.multipart.MultipartFile;
import thirtynine.pojo.Speech;
import thirtynine.pojo.User;
import thirtynine.service.SpeechService;
import thirtynine.utils.DateFormat;
import javax.imageio.ImageIO;
import javax.servlet.http.HttpSession;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.util.List;
import java.util.Map;
import java.util.UUID;

@Controller
public class ChatRoom {

    @Autowired
    private SpeechService speechService;

    @RequestMapping({"/","/ChatRoom"})
    public String toChatRoom(Model model, HttpSession session, Integer page){
        if(page==null){
            page=1;
        }
        Integer maxPage = speechService.getMaxPage();
        Speech topSpeech = speechService.findTop();
        List<Speech> speeches = speechService.findByPage(page);
        model.addAttribute("username",((User)session.getAttribute("user")).getName());
        model.addAttribute("top",topSpeech);
        model.addAttribute("speeches",speeches);
        model.addAttribute("currentPage",page);
        model.addAttribute("maxPage",maxPage);
        return "ChatRoom";
    }

    @PostMapping("/speak")
    public String speak(Speech speech,HttpSession session,@RequestParam("file") MultipartFile file){
        if (!file.isEmpty()){//upload file
            try {
                //if it's a image
                BufferedImage bi = ImageIO.read(file.getInputStream());
                if(bi==null){
                    session.setAttribute("wrong","NotImage");
                    return "redirect:wrong";
                }
                if(file.getSize()>5*1024*1024){
                    session.setAttribute("wrong","ImageTooLarge");
                    return "redirect:wrong";
                }
            } catch (IOException e) {
                e.printStackTrace();
            }
            //set path to store files
            String path =System.getProperty("user.dir")+"/src/main/resources/static/uploadedFiles/";
            // path: On my linux server
//            String path ="/root/ThirtyNine/static/uploadedFiles/";
            File f = new File(path);
            if(!f.exists()){
                f.mkdir();
            }
            //get file name
            String id =UUID.randomUUID().toString();
            String fileName =id+ file.getOriginalFilename();
            //file upload
            try{
                file.transferTo(new File(path,fileName));
            }catch (IOException e){
                e.printStackTrace();
            }
            //save file name to database
            speech.setImg_name(fileName);
        }
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

}
