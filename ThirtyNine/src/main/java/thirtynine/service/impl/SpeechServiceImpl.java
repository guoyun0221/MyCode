package thirtynine.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import thirtynine.mapper.SpeechDao;
import thirtynine.pojo.Speech;
import thirtynine.service.SpeechService;

import java.util.List;

@Service("SpeechService")
public class SpeechServiceImpl implements SpeechService {

    @Autowired
    private SpeechDao speechDao;

    @Override
    public void insertSpeech(Speech speech) {
        speechDao.insertSpeech(speech);
    }

    @Override
    public void deleteSpeech(int id) {
        speechDao.deleteSpeech(id);
    }

    @Override
    public List<Speech> findAll() {
        return speechDao.findAll();
    }
}
