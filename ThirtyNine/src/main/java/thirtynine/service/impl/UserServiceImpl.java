package thirtynine.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import thirtynine.mapper.UserDao;
import thirtynine.pojo.User;
import thirtynine.service.UserService;

import java.util.List;

@Service("UserService")
public class UserServiceImpl implements UserService {

    @Autowired
    private UserDao userDao;

    @Override
    public User getById(int id) {
        return userDao.getById(id);
    }

    @Override
    public User getByName(String name) {
        return userDao.getByName(name);
    }

    @Override
    public List<User> findAll() {
        return userDao.findAll();
    }

    @Override
    public void insertUser(User user) {
        userDao.insertUser(user);
    }

    @Override
    public void updateUser(User user) {
        userDao.updateUser(user);
    }

    @Override
    public void deleteUser(int id) {
        userDao.deleteUser(id);
    }

    @Override
    public boolean login(User user) {
        User user_old =userDao.getByName(user.getName());
        if(user_old!=null){
            //user already in database
            if(user.getPassword().equals(user_old.getPassword())){
                //login successfully
                return true;
            }else{
                //password is wrong
                return false;
            }
        }else {
            //log up
            userDao.insertUser(user);
            return true;
        }
    }
}
