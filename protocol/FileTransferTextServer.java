import java.io.*;
import java.net.*;

public class FileTransferTextServer {
    
    public static void main(String argv[]) {
	ServerSocket s = null;
	try {
	    s = new ServerSocket(FileTransferTextConstants.PORT);
	} catch(IOException e) {
	    System.out.println(e);
	    System.exit(1);
	}

	while (true) {
	    Socket incoming = null;
	    try {
		incoming = s.accept();
	    } catch(IOException e) {
		System.out.println(e);
		continue;
	    }

	    new SocketHandler(incoming).start();
	}
    }
}

class SocketHandler extends Thread {

    Socket incoming;
    File clientDir = new File(".");
    BufferedReader reader;
    PrintStream out;

    SocketHandler(Socket incoming) {
	this.incoming = incoming;
    }

    public void run() {
	try {
	    reader = new BufferedReader(new InputStreamReader(
					incoming.getInputStream()));
	    out = new PrintStream(incoming.getOutputStream());
	    
	    while (true) {
		String line = reader.readLine();
		if (line == null) { 
		    break;
		}
		System.out.println("Received request: " + line);

		if (line.startsWith(FileTransferTextConstants.CD)) {
		    changeDirRequest(losePrefix(line, 
						FileTransferTextConstants.CD));
		} else if (line.startsWith(FileTransferTextConstants.DIR)) {
		    directoryRequest();
		} else if (line.startsWith(FileTransferTextConstants.GET)) {
		    // code omitted
		} else if (line.startsWith(FileTransferTextConstants.QUIT)) {
		    break;
		} else {
		    out.print(FileTransferTextConstants.ERROR + 
			      FileTransferTextConstants.CR_LF);
		}
		
	    }
	    incoming.close();
	} catch(IOException e) {
	    e.printStackTrace();
	}
    }


    /**
     * Given that the string starts with the prefix,
     * get rid of the prefix and any following whitespace
     */
    public String losePrefix(String str, String prefix) {
	int index = prefix.length();
	String ret = str.substring(index).trim();
	return ret;

    }

    public void changeDirRequest(String dir) {
	File newDir = new File(clientDir, dir);
	if (newDir.isDirectory()) {
	    clientDir = newDir;
	    try {
		out.print(FileTransferTextConstants.SUCCEEDED + " " +
			  clientDir.getCanonicalPath() +
			  FileTransferTextConstants.CR_LF);
	    } catch(IOException e) {
		e.printStackTrace();
	    }
	} else {
	    out.print(FileTransferTextConstants.ERROR +
		      FileTransferTextConstants.CR_LF);
	}
    }

    public void directoryRequest() {
	String[] fileNames = clientDir.list();
	if (fileNames == null) {
	    out.print(FileTransferTextConstants.ERROR +
		      FileTransferTextConstants.CR_LF);
	}
	for (int n = 0; n < fileNames.length; n++) {
	    out.print(fileNames[n] +
		      FileTransferTextConstants.CR_LF);
	}
	out.print(FileTransferTextConstants.CR_LF);
    }
}
