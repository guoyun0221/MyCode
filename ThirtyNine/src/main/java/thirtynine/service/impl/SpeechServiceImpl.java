package thirtynine.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.springframework.transaction.annotation.Transactional;
import thirtynine.mapper.SpeechDao;
import thirtynine.mapper.UserDao;
import thirtynine.pojo.Speech;
import thirtynine.service.SpeechService;
import thirtynine.utils.EscapeCharacters;
import java.util.List;

@Service("SpeechService")
public class SpeechServiceImpl implements SpeechService {

    private final int size =39;

    @Autowired
    private SpeechDao speechDao;

    @Autowired
    private UserDao userDao;

    @Override
    public void cancelTop() {
        speechDao.cancelTop();
    }

    @Override
    public Speech findTop() {
        List<Speech> speeches = speechDao.findTop();
        if (speeches.isEmpty()){
            return null;
        }
        speeches.get(0).setWords(EscapeCharacters.escape(speeches.get(0).getWords()));
        speeches.get(0).setUser(userDao.getById(speeches.get(0).getUser_id()));
        speeches.get(0).setSpeaker(speeches.get(0).getUser().getName());
        return speeches.get(0);
    }

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
    public Integer getMaxPage() {
        Integer count = speechDao.countSpeeches();
        Integer maxPage = count/size+1;
        return maxPage;
    }

    @Override
    @Transactional
    public void topSpeech(int id) {
        speechDao.cancelTop();
        speechDao.topSpeech(id);
    }

    @Override
    public List<Speech> findByPage(Integer page) {
        int start =(page-1)*size;
        List<Speech> speeches = speechDao.findByPage(start,size);
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
