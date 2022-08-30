#!/usr/bin/perl

use Getopt::Std;
use DBI;                             # Script Rule |  [USE]    | TP
use strict;                          # Script Rule |  [USE]    \ STD
use utf8;                            # Script Rule |  [USE]    | STD
use Text::Table ();                  # Script Rule |  [USE]    | TP
use feature 'say';                   # Script Rule |  [USE]    | FEATURE
use open ':std', ':encoding(UTF-8)'; # Script Rule |  [USE]    | ENCODEmy %opts = ();
my %opts = ();

getopt('n:p:h:q:a:u:', \%opts);
my $Database_Name          = $opts{n};  # Script Rule |  [DBN]    | char
my $Database_Port          = $opts{p};  # Script Rule |  [DBP]    | char
my $Database_Host          = $opts{h};  # Script Rule |  [DBH]    | char
my $Database_Query         = $opts{q};  # Script Rule |  [DBQ]    | char
my $Database_Pass          = $opts{a};  # Script Rule |  [DBP]    | char
my $Database_User          = $opts{u};  # Script Rule |  [DPU]    | char
my $Title_Text_Color       = "\033[37m";
my $Start_Color            = "\033[38;5;20m";
my $End_Color              = "\033[38;5;20m";
my $Table_Body_color       = "\033[38;5;55m";
my $conn = "DBI:Pg:dbname=$Database_Name;host=$Database_Host;port=$Database_Port";
my $db = DBI->connect($conn, $Database_User, $Database_Pass, { RaiseError => 1 })or die $DBI::errstr;
my $stmt = qq($Database_Query);
my $result = $db->prepare( $stmt );
my $rv = $result->execute() or die $DBI::errstr;
if($rv < 0) {print $DBI::errstr;}
my @row = $result->fetchrow_array();
my $cols = scalar(@row);
my $counter = 0;
my $fields = join(',', @{ $result->{NAME_lc} });
my @v = split (/,/, $fields);

print "\n\n";

table();
binmode STDOUT, ':encoding(utf8)';



$db->disconnect();


sub table() {
    my @cols = @v;
    my $current_columns = scalar(@cols);
    @cols= grep { $_ ne '' } @cols;
    push @cols,+{align => "center",};
    my $sep = \"│";
    print("\033[38;5;55m");
    my $major_sep = \'│';
    my $tb        = Text::Table->new( $sep, "$Title_Text_Color Row Number $Start_Color", $major_sep,
        (map{ 
            +( 
                ( 
                    ref($_) ? $_ : 
                            "$Title_Text_Color $_\033[38;5;55m $Start_Color" ), 
                                            $sep 
                                                ) 
                                                    } 
                                                        @cols 
                                                                )
                                                                     );
    my $num_cols = @cols;
    my $counter = 0;
    while (my $row = $result->fetchrow_arrayref) {
        $tb->load( [$counter++, @$row]);
    }
    my $make_rule = sub {
        my ($args) = @_;
        my $t  = $args->{left};
        my $r    = $args->{right};
        my $f       = $args->{main_left};
        my $m           = $args->{middle};
        return $tb->rule(
            sub {
                my ( $i, $l ) = @_;
    
                return ( '─' x $l );
            },
            sub {
                my ( $i, $c ) = @_;
                my $l = (
                    ( $i == 0 )               ? $t
                    : ( $i == 1 )               ? $f
                    : ( $i == $num_cols + 1 )       ? $r
                    :                                    $m);
                return $l x $c;},);};
    
    my $start_rule = $make_rule->({
            left      => '┌',
                main_left => '┬',
                    right     => '┐',
                        middle    => '┬',});
    my $mid_rule = $make_rule->({
            left      => '├',
                main_left => '┼',
                    right     => '┤',
                        middle    => '┼',});
    my $end_rule = $make_rule->({
                left      => '└',
                    main_left => '┴',
                        right     => '┘',
                            middle    => '┴',});
    print "$Start_Color", $start_rule, $tb->title, map {
                                                     $mid_rule, $_, 
                                                    }
         "$Table_Body_color", $tb->body();
    print "$End_Color", $end_rule;
    print "\n\n";
}