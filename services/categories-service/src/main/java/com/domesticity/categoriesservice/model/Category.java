package com.domesticity.categoriesservice.model;

import java.util.UUID;

import javax.validation.constraints.NotBlank;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Category {

    private final UUID id;
    private final String icon;
    private final Colour colour;
    @NotBlank(message = "Name may not be empty")
    private final String name;

    public Category(@JsonProperty("id") UUID id,
                    @JsonProperty("name") String name,
                    @JsonProperty("color") Colour colour,
                    @JsonProperty("icon") String icon) {
        this.id = id;
        this.name = name;
        this.colour = colour;
        this.icon = icon;
    }

    public UUID getId() {
        return id;
    }

    public String getName() {
        return name;
    }

    public Colour getColour() {
        return colour;
    }

    public String getIcon() {
        return icon;
    }
}
