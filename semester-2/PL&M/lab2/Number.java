public class Number {
    public long value;

    private Number(long value) {
        this.value = value;
    }

    public Number(double value) {
        this.value = (long) (value * (1L << 32));
    }

    public Number add(Number a) {
        return new Number(this.value + a.value);
    }

    public Number multiply(Number a) {
        long result = (this.value * a.value) >>> 32;
        return new Number(result);
    }

    public String toString() {
        return Double.toString((double) this.value / (1L << 32));
    }
}