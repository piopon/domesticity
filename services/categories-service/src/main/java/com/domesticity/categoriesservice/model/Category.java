package com.domesticity.categoriesservice.model;

import javax.validation.constraints.NotBlank;

import com.fasterxml.jackson.annotation.JsonProperty;

public class Category {
    private final String id;
    private final String icon;
    private final String colour;
    private final String name;

    public Category(@JsonProperty("id") String id,
                    @JsonProperty("name") String name,
                    @JsonProperty("color") String colour,
                    @JsonProperty("icon") String icon) {
        this.id = id;
        this.name = name;
        this.colour = colour;
        this.icon = icon;
    }

    public String getId() {
        return id;
    }

    @NotBlank(message = "Category name cannot be empty")
    public String getName() {
        return name;
    }

    @NotBlank(message = "Category colour cannot be empty")
    public String getColour() {
        return colour;
    }

    public String getIcon() {
        return icon;
    }

    public static Category empty() {
        return new Category("", "", "", "");
    }
}
