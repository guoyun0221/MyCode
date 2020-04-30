package thirtynine.pojo;

import org.springframework.stereotype.Component;

@Component
public class Speech {
    private int id;
    private int user;//identified by user's id
    private String words;

    public int getId() {
        return id;
    }

    public void setId(int id) {
        this.id = id;
    }

    public int getUser() {
        return user;
    }

    public void setUser(int user) {
        this.user = user;
    }

    public String getWords() {
        return words;
    }

    public void setWords(String words) {
        this.words = words;
    }

    @Override
    public String toString() {
        return "Speech{" +
                "id=" + id +
                ", user=" + user +
                ", words='" + words + '\'' +
                '}';
    }
}
