package thirtynine.mapper;

import org.apache.ibatis.annotations.*;
import thirtynine.pojo.Speech;

import javax.validation.constraints.Size;
import java.util.List;

@Mapper
public interface SpeechDao {

    @Select("select * from speeches where id = #{id}")
    Speech findById(int id);

    @Select("select * from speeches order by id desc")
    @Results({
        @Result(property = "user", column = "user_id", one =@One(select ="thirtynine.mapper.UserDao.getById") )
    })
    List<Speech> findAll();

    @Select("select * from speeches order by id desc limit #{start},#{count}")
    @Results({
            @Result(property = "user", column = "user_id", one =@One(select ="thirtynine.mapper.UserDao.getById") )
    })
    List<Speech> findByPage(int start,int count);

    @Select("select * from speeches where words like #{keyword} order by id desc")
    @Results({
            @Result(property = "user", column = "user_id", one =@One(select ="thirtynine.mapper.UserDao.getById") )
    })
    List<Speech> findByKeyword(String keyword);

    @Select("select count(*) from speeches")
    Integer countSpeeches();

    @Insert("insert into speeches(user_id,words,send_time,img_name,reply_to) values(#{user.id},#{words},#{send_time},#{img_name},#{reply_to})")
    void insertSpeech(Speech speech);

    @Delete("delete from speeches where id =#{id}")
    void deleteSpeech(int id);

    @Update("update speeches set at_top = true where id =#{id}")
    void topSpeech(int id);

    @Update("update speeches set at_top = false")
    void cancelTop();

    @Select("select * from speeches where at_top is true")
    List<Speech> findTop();
}
