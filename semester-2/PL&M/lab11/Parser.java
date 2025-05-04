public class Parser {
    private char sym;
    private String string;
    private int pos;

    public Parser(String string){
        this.string = string;
        this.pos = 0;
        this.sym = string.charAt(0);
    }

    private void Next(){
        if (this.pos < string.length()-1){
            this.sym = string.charAt(++pos);
        }
    }

    // <Parcel> ::= <Call> <Tail>
    public boolean Parcel(){
        System.out.println("<Parcel> ::= <Call> <Tail>");
        boolean resp;
        SkipSpace();
        resp = Call();
        if (!resp) {
            return false;
        }
        return Tail();
    }

    // <Tail>::= : <Parcel> | ε
    public boolean Tail(){
        if (this.sym == ':') {
            System.out.println("<Tail>   ::= :<Parsel>");
            Next();
            return Parcel();
        }
        System.out.println("<Tail>   ::= ε");
        return true;
    }

    // <Call>::= (IDENT <Arg>)
    private boolean Call(){
        boolean resp;
        if (this.sym == '(') {
            System.out.println("<Call>   ::= (IDENT <Arg>)");
            Next();
            resp = Ident();
            SkipSpace();
            resp = Arg();
            if (!resp || this.sym != ')') {
                return GiveError(this.pos);
            }
            Next();
            return true;
        }
        return GiveError(this.pos);
    }

    private boolean Ident(){
        String ident = "";
        while (Character.isDigit(this.sym) || Character.isLetter(this.sym)) {
            ident += this.sym;
            Next();
        }
        System.out.printf("IDENT = %s\n", ident);
        return true;
    }

    // <Arg>::= NUMBER | STRING | <Parsel>
    private boolean Arg(){
        if (Character.isDigit(this.sym)){
            System.out.println("<Arg>    ::= NUMBER");
            return Numberr();
        }else if (Character.isLetter(this.sym)) {
            System.out.println("<Arg>    ::= STRING");
            return Stringg();
        }else if (this.sym == '('){
            System.out.println("<Arg>    ::= <Parsel>");
            return Parcel();
        }
        return GiveError(this.pos);
    }

    private boolean Numberr(){
        String ident = "";
        while (Character.isDigit(this.sym)) {
            ident += this.sym;
            Next();
        }
        System.out.printf("NUMBER = %s\n", ident);
        return true;
    }

    private boolean Stringg(){
        String ident = "";
        while (Character.isLetter(this.sym)) {
            ident += this.sym;
            Next();
        }
        System.out.printf("STRING = %s\n", ident);
        return true;
    }

    private void SkipSpace() {
        while (this.sym == ' '){
            Next();
        }
    }

    private boolean GiveError(int index){
        System.out.printf("Ошибка в позиции %d\n", index);
        return false;
    }

    public static void main(String[] args) {
        String string = "(f 4):(e (g 6))";
        Parser parser = new Parser(string);
        boolean resp = parser.Parcel();
    }
}