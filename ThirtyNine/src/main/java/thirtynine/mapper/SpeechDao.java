package thirtynine.mapper;

import org.apache.ibatis.annotations.Delete;
import org.apache.ibatis.annotations.Insert;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import thirtynine.pojo.Speech;

import java.util.List;

@Mapper
public interface SpeechDao {

    @Select("select * from speeches")
    List<Speech> findAll();

    @Insert("insert into speeches(user,words) values(#{user},#{words})")
    void insertSpeech(Speech speech);

    @Delete("delete from speeches where id =#{id}")
    void deleteSpeech(int id);
}
