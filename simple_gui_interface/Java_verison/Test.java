import java.util.ArrayList;
import java.util.List;

public class Test {
    public static void main(String[] args)  {

        List<String> qs = new ArrayList<>();
        qs.add("input 1: ");
        qs.add("input 2: ");

        SimpleGUI gui = new SimpleGUI("Test Gui Interface", qs);
        gui.createWindow();

        while (true){
            List<String> in = gui.getUserInput();
            for (int i = 0; i < in.size(); i++){
                gui.writeOutput("i: " + i + " value: " + in.get(i) + "\n");
            }
        }
    }
}
