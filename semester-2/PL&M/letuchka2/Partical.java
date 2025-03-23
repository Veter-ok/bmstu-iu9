import java.lang.Math; 

public class Partical {
    private int x;
    private int y;

    public Partical(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public double radius() {
        return Math.sqrt(Math.pow(x, 2)+ Math.pow(y, 2));
    }
}
