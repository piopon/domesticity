package com.domesticity.categoriesservice.model;

import javax.validation.constraints.NotBlank;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Colour {
    @NotBlank(message = "Red colour value may not be empty")
    private final String red;
    @NotBlank(message = "Green colour value may not be empty")
    private final String green;
    @NotBlank(message = "Blue colour value may not be empty")
    private final String blue;
    @NotBlank(message = "Alpha colour value may not be empty")
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
