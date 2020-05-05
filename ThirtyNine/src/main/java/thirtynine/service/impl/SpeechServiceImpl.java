package thirtynine.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import thirtynine.mapper.SpeechDao;
import thirtynine.mapper.UserDao;
import thirtynine.pojo.Speech;
import thirtynine.service.SpeechService;
import thirtynine.utils.EscapeCharacters;
import java.util.List;

@Service("SpeechService")
public class SpeechServiceImpl implements SpeechService {

    @Autowired
    private SpeechDao speechDao;

    @Autowired
    private UserDao userDao;

    @Override
    public Speech findById(int id) {
        Speech speech =speechDao.findById(id);
        speech.setUser(userDao.getById(speech.getUser_id()));
        speech.setSpeaker(speech.getUser().getName());
        return speech;
    }

    @Override
    public List<Speech> findAll() {
        List<Speech> speeches =speechDao.findAll();
        for(Speech speech: speeches){
            if(speech.getWords()!=null){
                //escape some characters
                speech.setWords(EscapeCharacters.escape(speech.getWords()));
            }
            if(speech.getUser()!=null){
                //to get username in thymeleaf
                speech.setSpeaker(speech.getUser().getName());
            }
        }
        return speeches;
    }

    @Override
    public List<Speech> findByKeyword(String keyword) {
        keyword='%'+keyword+'%';
        List<Speech> speeches =speechDao.findByKeyword(keyword);
        for(Speech speech: speeches){
            speech.setWords(EscapeCharacters.escape(speech.getWords()));
            if(speech.getUser()!=null){
                speech.setSpeaker(speech.getUser().getName());
            }
        }
        return speeches;
    }

    @Override
    public void insertSpeech(Speech speech) {
        speechDao.insertSpeech(speech);
    }

    @Override
    public void deleteSpeech(int id) {
        speechDao.deleteSpeech(id);
    }

}
