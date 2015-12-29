
/**
 * FileTransferTextClient.java
 */

/**
 * WARNING: the following code is okay as procedural code
 * but it sucks as O/O code
 */

import java.io.*;
import java.net.*;

public class FileTransferTextClient {

    private final static String  UI_DIR = "dir";
    private final static String  UI_CD = "cd";
    private final static String  UI_GET = "get";
    private final static String  UI_QUIT = "quit";

    protected Socket sock;
    protected BufferedReader reader;
    protected BufferedReader console;
    protected PrintStream writer;

    public static void main(String[] args){

	if (args.length != 1) {
	    System.err.println("Usage: Client address");
	    System.exit(1);
	}
	new FileTransferTextClient(args[0]);
    }

    public FileTransferTextClient(String server) {

	InetAddress address = null;
	try {
	    address = InetAddress.getByName(server);
	} catch(UnknownHostException e) {
	    e.printStackTrace();
	    System.exit(2);
	}

	sock = null;
	InputStream in = null;
	OutputStream out = null;
	try {
	    sock = new Socket(address, FileTransferTextConstants.PORT);
	    in = sock.getInputStream();
	    out = sock.getOutputStream();
	} catch(IOException e) {
	    e.printStackTrace();
	    System.exit(3);
	}

	reader = new BufferedReader(new InputStreamReader(in));
	writer = new PrintStream(out);
	console = new BufferedReader(new InputStreamReader(System.in));

	while (true) {
	    String line = null;
	    try {
		System.out.print("Enter request: ");
		line = console.readLine();
		System.out.println("Request was " + line);
	    } catch(IOException e) {
		e.printStackTrace();
		exit();
	    }

	    if (line.equals(UI_DIR)) {
		directoryRequest();
	    } else if (line.startsWith(UI_CD)) {
		changeDirRequest(losePrefix(line, 
					    UI_CD));
	    } else if (line.startsWith(UI_GET)) {
		getFileRequest(losePrefix(line, 
					  UI_GET));
	    } else if (line.equals(UI_QUIT)) {
		exit();
	    } else {
		System.out.println("Unrecognised command");
	    }
	}
    }

    /**
     * Given that the string starts with the prefix,
     * get rid of the prefix and any whitespace
     */
    public String losePrefix(String str, String prefix) {
	int index = prefix.length();
	String ret = str.substring(index).trim();
	return ret;
    }

    public void exit() {
	try {
	    writer.print(FileTransferTextConstants.QUIT +
			 FileTransferTextConstants.CR_LF);
	    reader.close();
	    writer.close();
	    sock.close();
	} catch(Exception e) {
	    e.printStackTrace();
	}
	System.exit(0);
    }

    public void directoryRequest() {
	writer.print(FileTransferTextConstants.DIR + 
		     FileTransferTextConstants.CR_LF);

	System.out.println("Dir listing is:");
	String line = null;
	while (true) {
	    try {
		line = reader.readLine();
	    } catch(IOException e) {
		break;
	    }
	    if (line.equals("")) {
		break;
	    }
	    System.out.println(line);
	}
    }

    public void changeDirRequest(String dir) {
	writer.print(FileTransferTextConstants.CD + " " + dir + 
		     FileTransferTextConstants.CR_LF);

	String response = null;
	try {
	    response = reader.readLine();
	} catch (IOException e) {
	    e.printStackTrace();
	    return;
	}

	if (response.equals(FileTransferTextConstants.ERROR)) {
	    System.out.println("Error in DIR request");
	} else if (response.startsWith(FileTransferTextConstants.SUCCEEDED)) {
	    String newdir = losePrefix(response, 
				       FileTransferTextConstants.SUCCEEDED);
	    System.out.println("Changed dir to " + newdir);
	} else {
	    System.out.println("Illegal response from server" +
			       response);
	}
    }

    public void getFileRequest(String filename) {
	// code omitted
    }
} // FileTransferTextClient








