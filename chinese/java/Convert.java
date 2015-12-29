
/**
 * convert GB2312 to Unicode
 */

import java.util.Vector;
import java.io.*;

public class Convert {
    public static String convert(byte[] inBytes) {
	try {
	    ByteArrayInputStream in = new ByteArrayInputStream(inBytes);
	    InputStreamReader inReader = new InputStreamReader(in, "GB2312");
	    int ch;
	    StringBuffer outBuff = new StringBuffer();
	    while ((ch = inReader.read()) != -1) {
		outBuff.append((char) ch);
	    }
	    return outBuff.toString();
	} catch(IOException e) {
	    return "";
	}
    }

    /*
    public static String convert(String in) {
	try {
	    StringReader inR = new StringReader(in);
	    InputStreamReader inReader = new InputStreamReader(inR, "GB2312");
	    int ch;
	    StringBuffer outBuff = new StringBuffer();
	    while ((ch = inReader.read()) != -1) {
		outBuff.append((char) ch);
	    }
	    return outBuff.toString();
	} catch(IOException e) {
	    return "";
	}
    }
    */
}

