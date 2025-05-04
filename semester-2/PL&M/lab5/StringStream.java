import java.util.ArrayList;
import java.util.Comparator;
import java.util.HashMap;
import java.util.Map;
import java.util.Optional;
import java.util.stream.Stream;

class NameComparator implements Comparator<String> {
    public int compare(String a, String b) {
        ArrayList<Character> charsA = new ArrayList<>();
        ArrayList<Character> charsB = new ArrayList<>();
        for (int i = 0; i < a.length(); i++) {
            if (charsA.indexOf(a.charAt(i)) != -1) {
                charsA.add(a.charAt(i));
            }
        } 
        for (int i = 0; i < b.length(); i++) {
            if (charsB.indexOf(b.charAt(i)) != -1) {
                charsB.add(b.charAt(i));
            }
        } 
        if (charsA.size() > charsB.size()) {return 1;}
        if (charsA.size() == charsB.size()) {return 0;}
        return -1;
    }
}


public class StringStream {
    private HashMap<String, Integer> strings;
    private int count;

    public StringStream(){
        this.strings = new HashMap<>();
        this.count = 0;
    }

    public void add(String str){
        this.strings.put(str, this.count);
        this.count++;
    }

    public Stream<String> makeStream(int k) {
        ArrayList<String> result = new ArrayList<>();
        strings.entrySet().stream()
            .filter(x -> x.getKey().length() == k && isPalindrom(x.getKey()))
            .forEach(x -> result.add(x.getKey()));
        return result.stream();
    }

    public Optional<Integer> findIndex(int k){
        Optional<Integer> result = Optional.empty();
        Optional<Map.Entry<String, Integer>> tmp = 
            strings.entrySet().stream()
                   .filter(x -> x.getKey().length() == k && isPalindrom(x.getKey()))
                   .findFirst();
        
        if (tmp.isPresent()) {
            result = Optional.ofNullable(tmp.get().getValue());
        }
        return result;
    }

    private boolean isPalindrom(String str){
        for (int i = 0; i < str.length() / 2; i++){
            if (str.charAt(i) != str.charAt(str.length() - 1 - i)){
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        StringStream strings = new StringStream();
        strings.add("avbhgf");
        strings.add("abbba");
        strings.add("aabbaa");
        strings.add("gdhs");
        strings.makeStream(5).sorted(new NameComparator()).forEach(System.out::println);
        System.out.println(strings.findIndex(5).get());
    }
}
