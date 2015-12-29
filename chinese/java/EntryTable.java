/**
 * Table of dictionary entries
 */

import java.io.*;
import java.util.Vector;
import java.util.SortedMap;
import java.util.TreeMap;
import java.util.HashMap;

public class EntryTable {

    private Vector entries = new Vector();
    private SortedMap englishSortedEntries;

    public EntryTable(EntryDataBase db) {
	Entry entry = null;
	while ((entry = db.get()) != null) {
	    entries.add(entry);
	}
    }

    public Entry[] doEnglishSearch(String english) {
	Vector matches = new Vector();
	for (int n = 0; n < entries.size(); n++) {
	    if (((Entry) entries.get(n)).englishMatch(english)) {
		System.out.println("Found match");
		matches.add(entries.get(n));
	    }
	}

	// This is the messy bit: to return an array
	// from a Vector we can either return an Object[]
	// array, or build a new array with type
	// coercion to make it an Entry[] array
 
	Entry matchArray[] = new Entry[matches.size()];
	for (int m = 0; m < matchArray.length; m++) {
	    matchArray[m] = (Entry) matches.get(m);
	}
	return matchArray;
    }

    public Entry[] doChineseSearch(String chinese) {
	// not yet implemented
	// return an empty array
	return new Entry[] {};
    }
}
