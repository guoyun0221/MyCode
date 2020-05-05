package thirtynine.pojo;

import org.springframework.stereotype.Component;
import java.sql.Date;

@Component
public class Speech {
    private int id;
    private int user_id;

    public int getUser_id() {
        return user_id;
    }

    public void setUser_id(int user_id) {
        this.user_id = user_id;
    }

    private User user;//identify speaker
    private String speaker;//it is just user.name, but ${speech.user.name} lead a exception
    //I don't know why, so I have to use this.
    private String words;
    private String send_time;

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

    @Override
    public String toString() {
        return "Speech{" +
                "id=" + id +
                ", user=" + user +
                ", speaker='" + speaker + '\'' +
                ", words='" + words + '\'' +
                ", send_time='" + send_time + '\'' +
                '}';
    }
}
