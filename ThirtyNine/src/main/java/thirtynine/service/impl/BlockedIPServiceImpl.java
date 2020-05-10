package thirtynine.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import thirtynine.mapper.Blocked_IP_Dao;
import thirtynine.pojo.Blocked_IP;
import thirtynine.service.BlockedIPService;

import java.util.List;

@Service("BlockedIPService")
public class BlockedIPServiceImpl implements BlockedIPService {

    @Autowired
    private Blocked_IP_Dao blocked_ip_dao;

    @Override
    public List<Blocked_IP> getAll() {
        return blocked_ip_dao.getAllBlockedIp();
    }
}
