package thirtynine.mapper;

import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Select;
import thirtynine.pojo.Blocked_IP;

import java.util.List;

@Mapper
public interface Blocked_IP_Dao {

    @Select("select * from blocked_IP")
    List<Blocked_IP> getAllBlockedIp();
}
