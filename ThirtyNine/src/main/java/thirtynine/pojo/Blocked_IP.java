package thirtynine.pojo;

import org.springframework.stereotype.Component;

@Component
public class Blocked_IP {
    private String IP;

    public String getIP() {
        return IP;
    }

    public void setIP(String IP) {
        this.IP = IP;
    }

    @Override
    public String toString() {
        return "Blocked_IP{" +
                "IP='" + IP + '\'' +
                '}';
    }
}
