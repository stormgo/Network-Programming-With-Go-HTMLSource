import java.io.*;
import java.util.Vector;
import java.util.StringTokenizer;

public class FileEntryDataBase implements EntryDataBase {

    private BufferedReader in;

    public FileEntryDataBase(String fileName) throws java.io.IOException {
	FileInputStream fin = null;
	fin = new FileInputStream(fileName);
	in = new BufferedReader(new InputStreamReader(fin, "GB2312"));
    }

    public Entry get() {
	String line;

	// the entries are pre-sorted in pinyin 
	try {
	    if ((line = in.readLine()) != null) {
		Entry entry = makeEntry(line);
		return entry;
		// System.out.println("Added entry" + entry.toString());
	    }
	} catch(IOException e) {
	    e.printStackTrace();
	    return null;
	}
	return null;
    }

    /**
     * Make an Entry from a line in the CEDICT file.
     * The syntax of entries in this file is
     * <chinese character(s)> [ <pinyin> ] / <english> / <english> ...
     */
    private Entry makeEntry(String line) {
	String chinese = "";
	String pinyin = "";
	Vector english = new Vector();

	int len = line.length();
	int n = 0;
	while ((n < len) && line.charAt(n) != ' ') {
	    chinese += line.charAt(n);
	    n++;
	}

	n++; // skip space
	n++; //skip [
	while ((n < len) && (line.charAt(n) != ']')) {
	    pinyin += line.charAt(n);
	    n++;
	}
	n++; // skip ]
	n++; // skip ' '
	n++; //skip /
	
	while (n < len) {
	    String englishWord = "";
	    while (line.charAt(n) != '/') {
		englishWord += line.charAt(n); 
		n++;
	    }
	    english.add(englishWord);
	    n++;
	}
	return new Entry(chinese,pinyin, english);
    }
}
