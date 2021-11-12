package com.domesticity.categoriesservice.model;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Colour {

    private final String red;
    private final String green;
    private final String blue;
    private final String alpha;

    public Colour(@JsonProperty("red") String red,
                  @JsonProperty("green") String green,
                  @JsonProperty("blue") String blue,
                  @JsonProperty("apha") String alpha) {
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
