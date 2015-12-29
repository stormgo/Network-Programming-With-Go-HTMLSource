
import java.awt.*;
import java.awt.event.*;
import java.io.*;
import java.lang.reflect.*;

public class GUIShell extends Frame implements KeyListener {

    protected TextArea text;
    protected PipedInputStream inPipe = new PipedInputStream();
    protected PipedOutputStream outInPipe;
    protected PrintStream outInStream;

    protected PipedOutputStream outPipe = new PipedOutputStream();
    protected PipedInputStream inOutPipe;
    
    protected PipedOutputStream errPipe = new PipedOutputStream();
    protected PipedInputStream inErrPipe;

    /**
     * run Main class in its own threadgroup
     * That way we can check how many threads are still
     * running so we know when to exit readline()
     * Looks like a bug? Every time a thread exits it throws
     * a pipeline read exception, so we have to catch and ignore
     * these and check how many threads are still alive
     *
     * I suspect there are still bugs in this
     */
    protected ThreadGroup threadGroup;

    static public void main(String[] argv) {
	GUIShell shell = new GUIShell(argv);
	shell.addWindowListener(new WindowAdapter() {
            public void windowClosing(WindowEvent e) {
                System.exit(0);
            }
          }
        );
	shell.setVisible(true);
    }

    public GUIShell(String[] argv) {
	try {
	    outInPipe = new PipedOutputStream(inPipe);
	    outInStream = new PrintStream(outInPipe);
	    inOutPipe = new PipedInputStream(outPipe);
	    inErrPipe = new PipedInputStream(errPipe);
	} catch(IOException e) {
	    System.err.println(e);
	    System.exit(1);
	}

	text = new TextArea(30, 50);
	text.setBackground(new Color(255, 255, 255));
	text.setFont(new Font("Monospaced", Font.BOLD, 16));

	add(text, BorderLayout.CENTER);
	pack();

	center();

	System.setIn(inPipe);
	System.setOut(new PrintStream(outPipe));
	System.setErr(new PrintStream(errPipe));

	text.addKeyListener(this);

	new ProcessReader(new BufferedReader(
		      new InputStreamReader(inOutPipe))).start();
	new ProcessReader(new BufferedReader(
		      new InputStreamReader(inErrPipe))).start();

	threadGroup = new ThreadGroup("Main"); 
	new RunProcess(argv).start();
    }

    public void keyTyped(KeyEvent evt) {
	char ch = evt.getKeyChar();
	outInStream.print(ch);
    }

    public void keyReleased(KeyEvent evt) {

    }

    public void keyPressed(KeyEvent evt) {

    }

    /** 
     * Center in the screen
     */
    protected void center() {
	Rectangle bounds = getBounds();
	Dimension screenSize = Toolkit.getDefaultToolkit().getScreenSize();
	
	int x = (screenSize.width - bounds.width)/2;
	int y = (screenSize.height - bounds.height)/2;
	setLocation(x, y);
    }

    class RunProcess extends Thread {
	String [] argv;

	RunProcess(String[] argv) {
	    super(threadGroup, "Main");
	    this.argv = argv;
	}

	/**
	 * Run the class by invoking method 
	 *    static void main(String[] argv)
	 * The class name is give in the property "guiShellFor"
	 */
	public void run() {
            String guiShellFor = System.getProperty("guiShellFor");
            try {
                // Class cls = Class.forName("SimpleThread");
                Class cls = Class.forName(guiShellFor);
		setTitle(cls.getName());
                Method method = cls.getMethod("main",
                                              new Class[] {argv.getClass()});
                method.invoke(null, new Object[] {argv});
            } catch(Exception e) {
                System.err.println(e.toString());
            }

	    /*
            System.out.close();
            System.err.close();
	    */
	}
    }

    class ProcessReader extends Thread {
	protected BufferedReader in;

	ProcessReader(BufferedReader in) {
	    this.in = in;
	}

	public void run() {
	    String line = null;
	    while (true) {
		try {
		    line = in.readLine();
		} catch(IOException e) {
		    // e.printStackTrace();
		    // text.append("IO error " + e);
		    // System.out.println("Threads in Main " +
		    //		   threadGroup.activeCount());
		    if (threadGroup.activeCount() == 0) {
			break;
		    }
		    continue;
		}
		if ((line == null) || line.equals("")) {
		    break;
		}
		text.append(line + "\n");
	    }
	    text.append("Program finished\n");
        }
    }
}
