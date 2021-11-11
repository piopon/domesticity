package com.domesticity.categoriesservice.model;

import java.util.UUID;

public class Category {
    
    private final UUID id;
    private final String color;
    private final String icon;
    
    public Category(UUID id, String color, String icon) {
        this.id = id;
        this.color = color;
        this.icon = icon;
    }

    public UUID getId() {
        return id;
    }

    public String getColor() {
        return color;
    }

    public String getIcon() {
        return icon;
    }
}
