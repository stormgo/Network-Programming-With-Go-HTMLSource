import java.awt.*;
import javax.swing.*;

public class HelloWorld extends JFrame {

	private JButton JButPress;

	public HelloWorld(String title) {
		super(title);
		getContentPane().setLayout(new FlowLayout());
		JLabel myLabel = new JLabel("Hello World!!!!");
		getContentPane().add(myLabel);
		JButPress = new JButton("Press Me");
		getContentPane().add(JButPress);
	      setDefaultCloseOperation(JFrame.EXIT_ON_CLOSE);
		setSize(250, 250);
		show();	
	}


	public static void main(String[] args) {
	     new HelloWorld("");
	}
}
