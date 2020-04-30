package thirtynine.service;

import org.springframework.stereotype.Service;
import thirtynine.pojo.User;

import java.util.List;

public interface UserService {

    User getById(int id);

    User getByName(String name);

    List<User> findAll();

    void insertUser(User user);

    void updateUser(User user);

    void deleteUser(int id);

    boolean login(User user);
}
