package com.domesticity.categoriesservice.model;

import java.util.UUID;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Category {

    private final UUID id;
    private final String name;
    private final Colour colour;
    private final String icon;

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
