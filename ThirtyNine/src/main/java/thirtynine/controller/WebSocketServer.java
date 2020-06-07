package thirtynine.controller;

import org.springframework.stereotype.Component;
import javax.websocket.*;
import javax.websocket.server.ServerEndpoint;
import java.io.IOException;
import java.util.concurrent.ConcurrentHashMap;

@ServerEndpoint("/ws")
@Component
public class WebSocketServer {

    private static ConcurrentHashMap<Session,WebSocketServer> webSocketMap = new ConcurrentHashMap<>();
    private Session session;

    @OnOpen
    public void onOpen(Session session){
        this.session=session;
        webSocketMap.put(session,this);
    }

    @OnClose
    public void onClose(Session session){
        webSocketMap.remove(session);
    }

    @OnMessage
    public void onMessage(String message){
    }

    public static void sendMessage(){
        for(WebSocketServer server: webSocketMap.values()){
            try {
                server.session.getBasicRemote().sendText("message from server");
            } catch (IOException e) {
                e.printStackTrace();
            }
        }
    }
}
