package thirtynine.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestParam;
import thirtynine.pojo.User;
import thirtynine.service.UserService;
import thirtynine.utils.EscapeCharacters;

import javax.servlet.http.HttpSession;

@Controller
@RequestMapping("/user")
public class UserHome {

    @Autowired
    private UserService userService;

    @RequestMapping("/info")
    public String show(@RequestParam("username") String username, Model model){
        User user= userService.getByName(username);
        if(user.getDescription()!=null){
            user.setDescription(EscapeCharacters.escape(user.getDescription()));
        }
        model.addAttribute("user",user);
        return "user/info";
    }

    @RequestMapping("/edit")
    public String toEdit(Model model, HttpSession session){
        model.addAttribute("user",(User)session.getAttribute("user"));
        return "user/edit";
    }

    @RequestMapping("/editDone")
    public String editDone(HttpSession session,String description){
        User user=(User)session.getAttribute("user");
        user.setDescription(description);
        userService.updateUser(user);
        return "redirect:/ChatRoom";
    }
}
