import java.util.ArrayList;

public class SuperNumber implements Comparable<SuperNumber>{
    private int number;

    public SuperNumber(int x){
        this.number = x;
    }

    public int countDigit() {
        ArrayList<Integer> l = new ArrayList<>(); 
        int tmp = 0;
        int num = this.number;
        while(num > 0) {
            tmp = num % 10;
            if (l.indexOf(tmp) == -1){
                l.add(tmp);
            }
            num /= 10;
        }
        return l.size();
    }

    public int compareTo(SuperNumber a) {
        return countDigit() - a.countDigit();
    }

    public String toString() {
        return Integer.toString(this.number);
    }
}
