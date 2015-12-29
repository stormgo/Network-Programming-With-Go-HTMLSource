/**
 * main dictionary class
 */

import javax.swing.*;
import java.awt.*;
import java.awt.event.*;
import java.io.*;

public class ChineseDictionary extends JFrame
    implements ActionListener {

    public EntryTable entries;
    private Font chineseFont = new Font("ZYSong18030",
					Font.BOLD, 12);
    private Font largeChineseFont = new Font(// "Bitstream Cyberbit", 
					     "ZYSong18030",
					     Font.BOLD, 18);
    private static String ZHONG_WEN = "\u4e2d\u6587";
    private static int SEARCH_CHINESE = 1;
    private static int SEARCH_ENGLISH = 2;

    private int searchType;

    // gui objects with scope of this object
    private JRadioButton searchEnglish;
    private JRadioButton searchChinese;
    private JTextField textEnterBox;
    private JTextArea results;

    public static void main(String [] argv) {

	ChineseDictionary dict = new ChineseDictionary();
	dict.pack();
	dict.setVisible(true);
    }

    public ChineseDictionary() {
	EntryDataBase dbase = null;
	try {
	    dbase = new FileEntryDataBase("/usr/local/src/chinese/cedict.gb");
	    entries = new EntryTable(dbase);
	} catch(IOException e) {
	    e.printStackTrace();
	    System.exit(1);
	}

	/*
	Font[] allfonts = GraphicsEnvironment.getLocalGraphicsEnvironment().getAllFonts();
	Font okFont = null;
	System.out.println("Font count: " + allfonts.length);
	char sample = '\u4e00';
	for (int n = 0; n < allfonts.length; n++) {
	    String name = allfonts[n].getFontName();
	    System.out.println("Font: " + name);
	    if (allfonts[n].canDisplay(sample)) {
		System.out.println("found font " + allfonts[n].getFontName());
		okFont = new Font(allfonts[n].getFontName(),
				       Font.PLAIN, 16);
		// break;
	    }
	}
	if (okFont == null) {
	    System.out.println("no font found");
	    System.exit(1);
	}
	*/

	setLayout();
	addListeners();
    }

    public void addListeners() {
	searchChinese.addActionListener(this);
	searchEnglish.addActionListener(this);
	textEnterBox.addActionListener(this);
    }

    public void setLayout() {
	Container panel = getContentPane();

	panel.setLayout(new BorderLayout());
	JPanel topPanel = new JPanel();
	topPanel.setLayout(new BorderLayout());
	JPanel radioPanel = new JPanel();
	radioPanel.setLayout(new GridLayout(2, 1));
	results = new JTextArea(20, 20);
	JScrollPane scrollPane = new JScrollPane(results);
	
	textEnterBox = new JTextField(20);
	searchChinese = new JRadioButton();
	searchEnglish = new JRadioButton("English", true);
	searchType = SEARCH_ENGLISH;

	// radio behaviour
	ButtonGroup grp = new ButtonGroup();
	grp.add(searchEnglish);
	grp.add(searchChinese);


	// geometry
	radioPanel.add(searchEnglish);
	radioPanel.add(searchChinese);
	topPanel.add(textEnterBox, BorderLayout.CENTER);
	topPanel.add(radioPanel, BorderLayout.EAST);
	panel.add(topPanel, BorderLayout.NORTH);
	panel.add(scrollPane, BorderLayout.CENTER);

	searchChinese.setFont(chineseFont);
	searchChinese.setText(ZHONG_WEN);
	results.setFont(largeChineseFont);
    }

    public void actionPerformed(ActionEvent e) {
	Object source = e.getSource();

	if (source == searchChinese) {
	    searchType = SEARCH_CHINESE;
	} else if (source == searchEnglish) {
	    searchType = SEARCH_ENGLISH;
	} else if (source == textEnterBox) {
	    if (searchType == SEARCH_CHINESE) {
		doChineseSearch();
	    } else {
		doEnglishSearch();
	    }
	}
    }

    private void doEnglishSearch() {
	results.setText("");
	Entry entrys[] = entries.doEnglishSearch(textEnterBox.getText());
	for (int n = 0; n < entrys.length; n++) {
	    Entry entry = entrys[n];
	    String line = entry.chinese + " " + entry.accentPinyin() +"\n";
	    System.out.println(line);
	    results.append(line);
	}
    }

    private void doChineseSearch() {
	// nothing yet	
	// but will eventually call entries.doChineseSearch()
    }
}
