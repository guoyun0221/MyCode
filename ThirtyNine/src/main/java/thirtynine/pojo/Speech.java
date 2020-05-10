package thirtynine.pojo;

import org.springframework.stereotype.Component;
import java.sql.Date;

@Component
public class Speech {
    private int id;
    private int user_id;
    private User user;//identify speaker
    private String speaker;//it is just user.name, but ${speech.user.name} leads a exception
    //I don't know why, so I have to use this.
    private String words;
    private String send_time;
    private String img_name;//uploaded img name
    private Integer reply_to;
    private boolean at_top;

    @Override
    public String toString() {
        return "Speech{" +
                "id=" + id +
                ", user_id=" + user_id +
                ", user=" + user +
                ", speaker='" + speaker + '\'' +
                ", words='" + words + '\'' +
                ", send_time='" + send_time + '\'' +
                ", img_name='" + img_name + '\'' +
                ", reply_to=" + reply_to +
                ", at_top=" + at_top +
                '}';
    }

    public boolean isAt_top() {
        return at_top;
    }

    public void setAt_top(boolean at_top) {
        this.at_top = at_top;
    }

    public Integer getReply_to() {
        return reply_to;
    }

    public void setReply_to(Integer reply_to) {
        this.reply_to = reply_to;
    }

    public String getImg_name() {
        return img_name;
    }

    public void setImg_name(String img_name) {
        this.img_name = img_name;
    }

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public User getUser() {
        return user;
    }

    public void setUser(User user) {
        this.user = user;
    }

    public String getSpeaker() {
        return speaker;
    }

    public void setSpeaker(String speak) {
        this.speaker = speak;
    }

    public String getWords() {
        return words;
    }

    public void setWords(String words) {
        this.words = words;
    }

    public String getSend_time() {
        return send_time;
    }

    public void setSend_time(String send_time) {
        this.send_time = send_time;
    }

    public int getUser_id() {
        return user_id;
    }

    public void setUser_id(int user_id) {
        this.user_id = user_id;
    }

}
