package com.domesticity.categoriesservice.model;

public class Colour {

    private final String red;
    private final String green;
    private final String blue;
    private final String alpha;

    public Colour(String red, String green, String blue, String alpha) {
        this.red = red;
        this.green = green;
        this.blue = blue;
        this.alpha = alpha;
    }

    public String getRed() {
        return red;
    }

    public String getGreen() {
        return green;
    }

    public String getBlue() {
        return blue;
    }

    public String getAlpha() {
        return alpha;
    }
}
