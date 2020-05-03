package thirtynine.mapper;

import org.apache.ibatis.annotations.*;
import org.springframework.stereotype.Repository;
import thirtynine.pojo.User;

import java.util.List;

@Mapper
public interface UserDao {

    @Select("select * from users where id = #{id}")
    User getById(int id);

    @Select("select * from users where name=#{name}")
    User getByName(String name);

    @Select("select * from users")
    List<User> findAll();

    @Insert("insert into users(name,password) values(#{name},#{password})")
    void insertUser(User user);

    @Update("update users set name=#{name},password=#{password},level=#{level},experience=#{experience},description=#{description} where id=#{id}")
    void updateUser(User user);

    @Delete("delete from users where id=#{id}")
    void deleteUser(int id);
}
