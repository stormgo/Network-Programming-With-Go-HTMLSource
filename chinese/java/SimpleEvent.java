import java.awt.*;
import java.awt.event.*;
import javax.swing.*;
public class SimpleEvent extends JFrame implements ActionListener {

	public static void main(String[] args) {
		new SimpleEvent("");
	}
	public SimpleEvent(String title) {
		super(title);
		JButton jbutPress = new JButton("Click me");
		jbutPress.addActionListener(this);
		getContentPane().add(jbutPress);	
		pack();
		show();
	}
	public void actionPerformed(ActionEvent evt) {
		System.out.println("Event generated" + evt);

		try {
			Thread.sleep(100000);
		}catch(Exception e) {}
	}
}
