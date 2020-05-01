package thirtynine.utils;

import java.sql.Date;
import java.text.SimpleDateFormat;

public class DateFormat {
    public static String SqlDate(){
        long l = System.currentTimeMillis();
        Date time=new Date(l);
        SimpleDateFormat sdf=new SimpleDateFormat("yyyy-MM-dd HH:mm:ss");
        return sdf.format(time);
    }
}
