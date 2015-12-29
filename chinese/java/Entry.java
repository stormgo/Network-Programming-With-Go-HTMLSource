

/**
 * Entry in the dictionary
 */

import java.util.Vector;
import java.util.StringTokenizer;

public class Entry {

    public String chinese = "";
    public String pinyin = "";
    public Vector english;

    public Entry(String ch, String pi, Vector en) {
	chinese = ch;
	pinyin = pi;
	english = en;
    }

    public String toString() {
	StringBuffer ret = new StringBuffer("/" + chinese + "/" + 
					    pinyin + "[" + 
					    (String) english.get(0));
	for (int n = 1; n < english.size(); n++) {
	    ret.append("|" + (String) english.get(n));
	}
	ret.append("]");
	return ret.toString();
    }

    public String getEnglish() {
	String eng = (String) english.get(0);
	for (int n = 1; n < english.size(); n++) {
	    eng += " | " + (String) english.get(n);
	}
	return eng;
    }

    public boolean englishMatch(String pattern) {
	for (int n = 0; n < english.size(); n++) {
	    String englishWord = (String) english.get(n);
	    if (englishWord.equals(pattern)) {
		return true;
	    }
	}
	return false;
    }

    /**
     * Place an accent above the first vowel in the string
     * This is not quite correct e.g. liang has the accent
     * above the 'a', not the 'i'.
     * But I don't know the proper rule for this
     */
    public String accentPinyin() {
	StringTokenizer st = new StringTokenizer(pinyin);
	Vector words = new Vector();
	while (st.hasMoreTokens()) {
	    String word = st.nextToken();
	    String newWord = addAccent(word);
	    words.add(newWord);
	}
	StringBuffer buff = new StringBuffer();
	int size = words.size();
	for (int n = 0; n < size - 1; n++) {
	    buff.append(words.get(n) + " ");
	}
	buff.append(words.get(size-1));
	return buff.toString();
    }

    public String addAccent(String original) {
	int length = original.length();
	char accentValue = original.charAt(length - 1);
	if ((accentValue < '1') || (accentValue > '4')) {
	    return original;
	}
	char buff[] = new char[length - 1];
	int n;
	for (n = 0; n < length - 1; n++) {
	    char ch = original.charAt(n);
	    if (isVowel(ch)) {
		ch = addAccent(ch, accentValue);
		buff[n] = ch;
		break;
	    }
	    buff[n] = ch;
	}
	for (n++; n < length - 1; n++) {
	    buff[n] = original.charAt(n);
	}
	return new String(buff);
    }

    private char addAccent(char ch, char accentValue) {
	switch (ch) {
	case 'a': 
	    switch (accentValue) {
	    case '1':
		ch = '\u0101';
		break;
	    case '2':
		ch = '\u00e1';
		break;
	    case '3':
		ch = '\u0103';
		break;
	    case '4':
		ch = '\u00e0';
		break;
	    }
	    break;
	case 'e':
	    switch (accentValue) {
	    case '1':
		ch = '\u0113';
		break;
	    case '2':
		ch = '\u00e9';
		break;
	    case '3':
		ch = '\u0115';
		break;
	    case '4':
		ch = '\u00e8';
		break;
	    }
	    break;
	case 'i':
	    switch (accentValue) {
	    case '1':
		ch = '\u012b';
		break;
	    case '2':
		ch = '\u00e9';
		break;
	    case '3':
		ch = '\u012d';
		break;
	    case '4':
		ch = '\u00ec';
		break;
	    }
	    break;
	case 'o':
	    switch (accentValue) {
	    case '1':
		ch = '\u014d';
		break;
	    case '2':
		ch = '\u00f3';
		break;
	    case '3':
		ch = '\u014f';
		break;
	    case '4':
		ch = '\u00f2';
		break;
	    }
	    break;
	case 'u':
	    switch (accentValue) {
	    case '1':
		ch = '\u016b';
		break;
	    case '2':
		ch = '\u00fa';
		break;
	    case '3':
		ch = '\u016d';
		break;
	    case '4':
		ch = '\u00f9';
		break;
	    }
	    break;
	}
	return ch;
    }

    private boolean isVowel(char ch) {
	if (ch == 'a' ||
	    ch == 'e' ||
	    ch == 'i' ||
	    ch == 'o' ||
	    ch == 'u') {
	    return true;
	}
	return false;
    }


}
