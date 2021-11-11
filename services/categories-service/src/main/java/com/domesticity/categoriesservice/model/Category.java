package com.domesticity.categoriesservice.model;

import java.util.UUID;

public class Category {

    private final UUID id;
    private final Colour colour;
    private final String icon;

    public Category(UUID id, Colour colour, String icon) {
        this.id = id;
        this.colour = colour;
        this.icon = icon;
    }

    public UUID getId() {
        return id;
    }

    public Colour getColour() {
        return colour;
    }

    public String getIcon() {
        return icon;
    }
}
