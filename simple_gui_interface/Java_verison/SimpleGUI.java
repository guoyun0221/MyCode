import javax.swing.*;
import java.awt.*;
import java.awt.event.ActionEvent;
import java.awt.event.ActionListener;
import java.util.ArrayList;
import java.util.List;

public class SimpleGUI {

    private final String title;
    private final List<String> inputPrompts;
    private String buttonText = "Submit";
    private int basicSize = 24;
    private JFrame jf;
    // text field list for user to input
    private final List<JTextField> inputTexts = new ArrayList<>();
    // mark if user input updated
    private volatile boolean inputUpdated = false;
    private JTextArea outputText;


    /**
     * create an instance of simple gui interface
     * @param title title of this program to be shown
     * @param inputPrompts list of string prompts for user inputs.
     *                     each element in the list corresponds to an input
     */
    public SimpleGUI(String title, List<String> inputPrompts) {
        this.title = title;
        this.inputPrompts = inputPrompts;
    }

    /**
     * create an instance of simple gui interface
     * @param title title of this program to be shown
     * @param inputPrompts list of string prompts for user inputs.
     *                     each element in the list corresponds to an input
     * @param buttonText submit button text. default is "Do it"
     */
    public SimpleGUI(String title, List<String> inputPrompts, String buttonText) {
        this.title = title;
        this.inputPrompts = inputPrompts;
        this.buttonText = buttonText;
    }

    /**
     * create an instance of simple gui interface
     * @param title title of this program to be shown
     * @param inputPrompts list of string prompts for user inputs.
     *                     each element in the list corresponds to an input
     * @param buttonText  text of submit button. default is "Submit"
     * @param basicSize basic size for component. default is 24
     */
    public SimpleGUI(String title, List<String> inputPrompts, String buttonText, int basicSize) {
        this.title = title;
        this.inputPrompts = inputPrompts;
        this.buttonText = buttonText;
        this.basicSize = basicSize;
    }

    public void createWindow(){
        // the window
        this.jf = new JFrame(this.title);
        this.jf.setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
        // To make the window placed in the center of the screen
        this.jf.setLocationRelativeTo(null);
        // set window layout
        this.jf.setLayout(new BorderLayout((int)(basicSize * 0.5), basicSize * 2));
        // create and put components to jf
        createTitleLabel();
        createInputPanel();
        createCenterBlank();
        createOutputTextArea();
        createSubmitButton();

        // fit window size
        this.jf.pack();
        this.jf.setVisible(true);
    }

    /**
     * Wait until user submit input and get it.
     * @return User input list.
     */
    public List<String> getUserInput(){
        List<String> ret = new ArrayList<>();
        while (!SimpleGUI.this.inputUpdated) {
            // wait input updated
        }
        // get user input list
        for (int i = 0; i < SimpleGUI.this.inputPrompts.size(); i++){
            ret.add(SimpleGUI.this.inputTexts.get(i).getText());
        }
        // reset inputUpdated to false
        SimpleGUI.this.inputUpdated = false;
        // return user input
        return ret;
    }

    /**
     * Write output to output text area
     * @param s content to be written
     */
    public void writeOutput(String s){
        this.outputText.append(s);
//        this.outputText.paintImmediately(this.outputText.getBounds());
        jf.repaint();
    }

    private void createTitleLabel(){
        // put title to top of the window
        JLabel titleLabel = new JLabel(this.title, JLabel.CENTER);
        titleLabel.setFont(new Font(Font.MONOSPACED, Font.BOLD,this.basicSize * 2));
        this.jf.add(titleLabel, BorderLayout.NORTH);
    }

    private void createInputPanel(){
        // panel for input prompt labels and input texts

        // to make it looks better, the rows are set at least 3
        int rows = inputPrompts.size();
        if (inputPrompts.size() == 1){
            rows = 3;
        } else if (inputPrompts.size() == 2){
            rows = 4;
        }
        JPanel inputPanel = new JPanel(new GridLayout(rows, 2, 0, basicSize * 2));

        // fill grid with blank at top
        if(inputPrompts.size() < 3){
            JLabel blank11 = new JLabel();
            JLabel blank12 = new JLabel();
            inputPanel.add(blank11);
            inputPanel.add(blank12);
        }

        // add inputPromptLabel-inputText pair to jp
        for (String labelText : inputPrompts) {
            // if the text is too short, pad left spaces
            int targetLen = basicSize / 2;
            if (labelText.length() < targetLen) {
                labelText = padLeftSpaces(labelText, targetLen);
            }
            JLabel inputPromptLabel = new JLabel(labelText);
            inputPromptLabel.setFont(new Font(Font.MONOSPACED, Font.BOLD, this.basicSize));
            JTextField inputTextField = new JTextField();
            inputTextField.setFont(new Font(Font.MONOSPACED, Font.BOLD, this.basicSize));
            inputPanel.add(inputPromptLabel);
            inputPanel.add(inputTextField);
            // add input text field to list
            this.inputTexts.add(inputTextField);
        }

        // fill grid with blank at bottom
        if(inputPrompts.size() < 3){
            // fill black to fit grid
            JLabel blank21 = new JLabel();
            JLabel blank22 = new JLabel();
            inputPanel.add(blank21);
            inputPanel.add(blank22);
        }
        this.jf.add(inputPanel, BorderLayout.WEST);
    }

    private void createCenterBlank(){
        // Nothing in center layout
        // just leave some blank space between east component and west component
        JLabel titleLabel = new JLabel("");
        this.jf.add(titleLabel, BorderLayout.CENTER);
    }

    private void createOutputTextArea(){
        // add text area for output
        this.outputText = new JTextArea(0, (int)(this.basicSize * 1.2));
        this.outputText.setLineWrap(true);
        this.outputText.setEditable(false);
        this.outputText.setFont(new Font(Font.MONOSPACED, Font.BOLD,this.basicSize));
        JScrollPane scroll = new JScrollPane(this.outputText);
        scroll.setVerticalScrollBarPolicy(ScrollPaneConstants.VERTICAL_SCROLLBAR_ALWAYS);
        this.jf.add(scroll, BorderLayout.EAST);
    }

    private void createSubmitButton(){
        JButton button = new JButton(this.buttonText);
        button.setFont(new Font(Font.MONOSPACED, Font.BOLD, (int)(this.basicSize * 1.5)));

        button.addActionListener(new ActionListener() {
            @Override
            public void actionPerformed(ActionEvent e) {
                // input updated
                SimpleGUI.this.inputUpdated = true;
            }
        });
        this.jf.add(button, BorderLayout.SOUTH);
    }


    private String padLeftSpaces(String originalStr, int length){
        // pad original string to length, with left spaces.
        if (originalStr.length() >= length) {
            return originalStr;
        }
        StringBuilder sb = new StringBuilder();
        while (sb.length() < length - originalStr.length()) {
            sb.append(' ');
        }
        sb.append(originalStr);

        return sb.toString();
    }
}
