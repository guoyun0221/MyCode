package thirtynine.utils;

public class EscapeCharacters {
    public static String escape(String src){
        String dst =null;
        //to output multiple spaces and newlines
        //and to avoid xss injection
        dst=src.replaceAll(" ","&nbsp;")
                .replaceAll("<","&lt;")
                .replaceAll(">","&gt;")
                .replaceAll("\r","<br/>");
        return dst;
    }
}
