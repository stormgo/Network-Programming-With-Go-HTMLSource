#!/usr/bin/perl

#
# Configurable stuff
#

$classcolor = 'purple';
$classnamecolor = 'blue';
$constcolor = 'darkmagenta';
$keywordcolor = 'purple';
$basetypecolor = 'green';
$stringcolor = 'DarkRed';
$comment_1line_color = 'firebrick';
$comment_multiline_color = 'firebrick';
$comment_jdoc_color = 'darkred';

@classwords = ('class', 'interface', 'import', 'package', 
               'public', 'private', 'protected');

@keywords = ('return', 'goto', 'if', 'else', 'case', 'default', 'switch', 
	     'break', 'continue', 'default', 'while', 'do', 'for', 
	     'instanceof', 'abstract', 'static',  "func", "var", "type", "struct",
	     'super', 'synchronized', 'threadsafe', 'synchronised', 
	     'transient', 'final', 'delete', 'new', 'this', 'try', 'catch',
	     );

@basetypes =('void', 'int', 'char', 'byte', 'long', 'boolean');

# 
# End configurable stuff
#

chdir($ARGV[0]) or die "can't chdir to $ARGV[0]";
open(FILE, $ARGV[1]) or die "no such file $ARGV[1]";

@program = <FILE>;

print "<pre><code><font color=\"black\">\n";

$incomment = 0;

FOR:
foreach $line (@program) {

    # fix any &'s
    $line =~ s/\&/\&amp;/g;

    # clean up <> chars
    $line =~ s/\</\&lt;/g;
    $line =~ s/\>/\&gt;/g;

    # check for /* ... */ or JavaDoc comments
    &startComment($line);
    if ($incomment) {
	&endComment($line);
    }

    # split for one line comment // ...
    # ($line, $comment) = split(/\/\//, $line);

    if ($line =~ m/\"/) {
	undef $comment;

	# has at least one string " char in line
	$line = &stringFontify($line);
    } else {
	# split for one line comment // ...
	($line, $comment) = split(/\/\//, $line, 2);

	&setChangeMarker($comment);
	
	$line = &fontify($line);
    }

    # recombine any one line comment
    if (defined($comment)) {
	$comment = "<font color=\"$comment_1line_color\">//".
	           $comment.
                   '</font>';
	print $line.$comment;
    } else {
	print $line;
    }
}

sub setChangeMarker {
    local($comment) = @_;
    # look for CHANGED/END CHANGED and set background
    if ($comment =~ m/CHANGE.*END/i) {
	print "</div>";
    } elsif ($comment =~ m/CHANGE.*START/i) {
	print "<div class=\"changed\">";
    }
}


sub startComment {
    local($line) = @_;

    if ( ! $incomment && $line =~ m/\/\*/) {
	$incomment = 1;
	if ($line =~ m/\/\*\*/) {
	    # /** ... */
	    print "<font color=\"$comment_jdoc_color\">";
	} else {
	    if ($line =~ m/\*\//) {
		# comment starts and finishes on this line
		# but maybe is in the middle of the line :-(
		# so we don't want to comment the whole line,
		# just a bit of it.

		# for now, no comment
		# $incomment = 0;
	    }
	    # /* ... */
	    print "<font color=\"$comment_multiline_color\">";
	}	    
    }
}

sub endComment {
    local($line) = @_;

    print $line;
    if ($line =~ m/\*\//) {
	print "</font>";
	$incomment = 0;
    }
    next FOR;
}

# Strings are messy things
# We want to do no substitutions on chars in a string
# But we can get things like 
#    print("A string with \"quoted\" elements" + x + "\"");
# Some of this can get substituted, some can't
# We can't just work on " since \" is escaped
# So...
#    split into array on \" | " 
#    work though each element, sometimes fontifying, sometimes not
# 
# We also have to look out for // comments, which must be outside a string
sub stringFontify {
    local($line) = @_;

    @parts = split(/(\\\"|\")/, $line);
    $inquote = 0;
    $newline =  '';
    $inlinecomment = 0;
    FOR2:
    foreach $phrase (@parts) {
	# if we are in a // comment, just keep adding the phrase 
	# to the comment
	if ($inlinecomment) {
	    $comment .= $phrase;
	    next FOR2;
	}

	if ($phrase =~ m/^\"$/) {
	    $isquote = 1;
	} else {
	    $isquote = 0;
	}
	if ($inquote) {
	    if ($isquote) {
		# end of string
		$newline .= $qphrase.'"</font>';
		$inquote = 0;
	    } else {
		# add to current string
		$qphrase .= $phrase;
	    }
	} else {
	    if ($isquote) {
		# new string starts
		$inquote = 1;
		$qphrase = "<font color=\"$stringcolor\">\"";
	    } else {
		# not in string

		# check for // comment (max 2 elements
		($phrase, $comment) = split(/\/\//, $phrase, 2);
		if (defined($comment)) {
		    $inlinecomment = 1;
		}

		# fontify
		$phrase = &fontify($phrase);
		$newline .= $phrase;
	    }
	}
    }
    return $newline;
}

sub fontify {
    local($line) = @_;
    
    # class level words
    foreach $class (@classwords) {
        $line =~ s/\b$class\b/<font color="$classcolor">$class<\/font>/g;
    }
    
    # ordinary keywords
    foreach $key (@keywords) {
        $line =~ s/\b$key\b/<font color="$keywordcolor">$key<\/font>/g;
    }
    
    # types
    foreach $type (@basetypes) {
        $line =~ s/\b$type\b/<font color="$basetypecolor">$type<\/font>/g;
    }

    # classname - assume start with uppercase
    $line =~ s/\b([A-Z]\w+)\b/<font color="$classnamecolor">\1<\/font>/g;

    # const
    $line =~ s/\b([A-Z][A-Z_0-9]*)\b/<font color="$constcolor">\1<\/font>/g;

    return $line;
}


print "</font></code></pre>\n";

exit 0;










