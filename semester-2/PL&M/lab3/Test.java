import java.util.Arrays;
import java.util.Scanner;

public class Test {
    public static void main(String[] args) {
        Scanner reader = new Scanner(System.in);

        System.out.println("Введите количество чисел");
        int n = reader.nextInt();
        SuperNumber l[] = new SuperNumber[n];

        for (int i = 0; i < n; i++){
            System.out.printf("Введите чисел №%d:\n", i+1);
            int num = reader.nextInt();
            l[i] = new SuperNumber(num); 
        }

        Arrays.sort(l);
        for (SuperNumber elem : l ) {
            System.out.printf("%s  -  %d\n", elem, elem.countDigit());
        }

        reader.close();
    }
}  
