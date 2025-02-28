public class Factorial {
    public int count(int n) {
        if (n == 1) return n;
        return n * count(n - 1);
    }
}
