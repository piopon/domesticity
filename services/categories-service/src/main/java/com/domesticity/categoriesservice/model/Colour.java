package com.domesticity.categoriesservice.model;

import javax.validation.constraints.NotBlank;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Colour {
    private final String name;
    private final String red;
    private final String green;
    private final String blue;
    private final String alpha;

    public Colour(@JsonProperty("name") String name,
                  @JsonProperty("red") String red,
                  @JsonProperty("green") String green,
                  @JsonProperty("blue") String blue,
                  @JsonProperty("alpha") String alpha) {
        this.name = name;
        this.red = red;
        this.green = green;
        this.blue = blue;
        this.alpha = alpha;
    }

    @NotBlank(message = "Colour name cannot be empty")
    public String getName() {
        return name;
    }

    @NotBlank(message = "Red colour value may not be empty")
    public String getRed() {
        return red;
    }

    @NotBlank(message = "Green colour value may not be empty")
    public String getGreen() {
        return green;
    }

    @NotBlank(message = "Blue colour value may not be empty")
    public String getBlue() {
        return blue;
    }

    @NotBlank(message = "Alpha colour value may not be empty")
    public String getAlpha() {
        return alpha;
    }
}
