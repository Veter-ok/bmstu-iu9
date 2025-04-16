import java.util.ArrayList;
import java.util.Iterator;

public class LineSet implements Iterable<Double[]> {
    private ArrayList<Double[]> lines;
    private ArrayList<Double[]> points;

    public LineSet(ArrayList<Double[]> lines) {
        this.lines = lines;
        this.points = new ArrayList<Double[]>();
        
        for (int i = 0; i < this.lines.size(); i++) {
            for (int j = i + 1; j < this.lines.size(); j++) {
                Double[] intersection = intersection(this.lines.get(i), this.lines.get(j));
                if (intersection != null) {
                    points.add(intersection);
                }
            }
        }
    }

    public void SetLines(ArrayList<Double[]> new_lines) {
        this.lines = new_lines;
        for (int i = 0; i < this.lines.size(); i++) {
            for (int j = i + 1; j < this.lines.size(); j++) {
                Double[] intersection = intersection(this.lines.get(i), this.lines.get(j));
                if (intersection != null) {
                    points.add(intersection);
                }
            }
        }
    }

    private Double[] intersection(Double[] a, Double[] b) {
        double x1 = a[0];
        double y1 = a[1];
        double x2 = a[2];
        double y2 = a[3];
        double x3 = b[0];
        double y3 = b[1];
        double x4 = b[2];
        double y4 = b[3];

        double denom = (x1 - x2) * (y3 - y4) - (y1 - y2) * (x3 - x4);
        
        if (denom == 0) {
            return null;
        }
        double x = ((x1 * y2 - y1 * x2) * (x3 - x4) - (x1 - x2) * (x3 * y4 - y3 * x4)) / denom;
        double y = ((x1 * y2 - y1 * x2) * (y3 - y4) - (y1 - y2) * (x3 * y4 - y3 * x4)) / denom;
        if (isBetween(x, x1, x2) && isBetween(y, y1, y2) && isBetween(x, x3, x4) && isBetween(y, y3, y4)) {
            return new Double[]{x, y};
        }
        
        return null;
    }

    private boolean isBetween(double val, double end1, double end2) {
        return val >= Math.min(end1, end2) && val <= Math.max(end1, end2);
    }

    public Iterator<Double[]> iterator() {
        return new LineSetIter();
    }

    private class LineSetIter implements Iterator<Double[]> {
        private int pos;

        public LineSetIter() { 
            this.pos = 0; 
        }

        public boolean hasNext() { 
            return pos < points.size(); 
        } 

        public Double[] next() {
            Double[] result = points.get(pos);
            pos++;
            return result;
        }
    }

    public static void main(String[] args) {
        ArrayList<Double[]> lines = new ArrayList<Double[]>();
        
        lines.add(new Double[]{2.0, 5.0, 10.0, 12.0});
        lines.add(new Double[]{24.0, 46.0, 31.0, 21.0});
        lines.add(new Double[]{85.0, 42.0, 13.0, 31.0});
        lines.add(new Double[]{40.0, 74.0, 63.0, 27.0});
        
        LineSet linesIter = new LineSet(lines);
        for (Double[] point : linesIter) {
            System.out.println(point[0] + ", " + point[1]);
        }
    }
}