import java.util.Iterator;

public class BooleanSequence implements Iterable<Integer> {
    private int[] data;
    private int length;

    public BooleanSequence(boolean[] data, int n) {
        this.length = n / 8 + (n % 8 == 0 ? 0 : 1);
        this.data = new int[this.length];
        for (int i = 0; i < this.length; i++) {
            for (int j = 0; j < 8; j++) {
                if (i*8 + j < n){
                    this.data[i] += Math.pow(2, 7-j) * (data[i*8 + j] ? 1 : 0);
                }
            }
        }
    }

    public void SetNewData(boolean[] data, int n){
        this.length = n / 8 + (n % 8 == 0 ? 0 : 1);
        this.data = new int[this.length];
        for (int i = 0; i < this.length; i++) {
            for (int j = 0; j < 8; j++) {
                if (i*8 + j < n){
                    this.data[i] += Math.pow(2, 7-j) * (data[i*8 + j] ? 1 : 0);
                }
            }
        }
    }

    public Iterator<Integer> iterator() {
        return new BooleanIterator();
    }

    private class BooleanIterator implements Iterator<Integer> {
        private int currentIndex = 0;

        public boolean hasNext() {
            return currentIndex < length;
        }

        public Integer next() {
            return data[currentIndex++];      
        }
    }

    public static void main(String[] args) {
        boolean[] data = {true, false, true, false, true, false, true, false,
                         true, true, true, false, false, false, false, false};
        BooleanSequence seq = new BooleanSequence(data, data.length);
        
        for (int b : seq) {
            System.out.println(b);
        }
    }
}