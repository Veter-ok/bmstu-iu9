import java.util.*;
import java.util.stream.*;

public class NumberStream {
    private HashMap<Integer, Set<Integer>> numbers;

    public NumberStream() {
        this.numbers = new HashMap<>();
    }

    public void add(int number) {
        Set<Integer> divisors = findDivisors(number);
        this.numbers.put(number, divisors);
    }

    private Set<Integer> findDivisors(int number) {
        Set<Integer> divisors = new HashSet<>();
        if (number == 0) return divisors;
        
        int n = Math.abs(number);
        for (int i = 1; i <= n; i++) {
            if (n % i == 0) {
                divisors.add(i);
            }
        }
        return divisors;
    }

    public Stream<Integer> makeStream(int k) {
        return numbers.entrySet().stream()
            .filter(entry -> entry.getValue().stream()
                .anyMatch(d -> d > 1 && getExponent(entry.getKey(), d) >= k))
            .map(Map.Entry::getKey);
    }

    private int getExponent(int number, int divisor) {
        if (number == 0) return 0;
        int n = Math.abs(number);
        int exponent = 0;
        while (n % divisor == 0) {
            exponent++;
            n /= divisor;
        }
        return exponent;
    }

    public Optional<Integer> findMinSquareFree() {
        return numbers.keySet().stream().filter(x -> isSquareFree(x)).min(Integer::compareTo);
    }

    private boolean isSquareFree(int number) {
        if (number == 0) return false;
        int n = Math.abs(number);
        for (int i = 2; i * i <= n; i++) {
            if (n % (i * i) == 0) {
                return false;
            }
        }
        return true;
    }

    public static void main(String[] args) {
        NumberStream stream = new NumberStream();
        stream.add(100);
        stream.add(13);
        stream.add(40);
        stream.add(30);
        stream.add(49);
        stream.add(1);
        stream.add(0); 
        stream.add(-8);

        stream.makeStream(2).sorted().forEach(System.out::println);
        stream.findMinSquareFree().ifPresent(System.out::println);

        long[] counts = stream.makeStream(2)
            .collect(() -> new long[3],
                    (arr, num) -> {
                        if (num < 0) arr[0]++;
                        else if (num == 0) arr[1]++;
                        else arr[2]++;
                    },
                    (a, b) -> {
                        a[0] += b[0];
                        a[1] += b[1];
                        a[2] += b[2];
                    });
        System.out.printf("%d, %d, %d\n", counts[0], counts[1], counts[2]);
    }
}