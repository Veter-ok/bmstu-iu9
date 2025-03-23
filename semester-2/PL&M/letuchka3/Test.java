public class Test {
    public interface countDist {
        double distance(Point a, Point b);
    }
    public static void main(String[] args) {
        countDist distance = new countDist() {
            @Override
            public double distance(Point a, Point b) {
                double xx = (a.x - b.x) * (a.x - b.x);
                double yy = (a.y - b.y) * (a.y - b.y);
                double zz = (a.z - b.z) * (a.z - b.z);
                return Math.sqrt(xx + yy + zz);
            }
        };
        Point a = new Point(10, 1, 20);
        Point b = new Point(14, -21, 21);
        System.out.println(distance.distance(a, b));
    }
}