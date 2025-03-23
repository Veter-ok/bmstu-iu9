import java.util.ArrayList;
import java.util.Random;  

public class Universe {
    static private int countParticals;
    private ArrayList<Partical> particals = new ArrayList<>();

    public Universe(int x) {
        Random random = new Random();
        for (int i = 0; i < x; i++){
            countParticals++;
            particals.add(new Partical(random.nextInt(100), random.nextInt(100)));
        }
    }

    public double midRadius() {
        double ful = 0;
        for (Partical elem : particals){
            ful += elem.radius();
        }
        return ful / countParticals;
    }

    public double distantToUniversal(Universe a) {
        return Math.abs(midRadius() - a.midRadius());
    }
}
