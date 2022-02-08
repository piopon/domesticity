package com.domesticity.categoriesservice.model;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class CategoryTest {

    @Test
    public void getIdShouldReturnCorrectIdValue() {
        final String id = "test_id1";
        final Category category = new Category(id, "", "", "");

        assertEquals(id, category.getId());
    }

    @Test
    public void getNameShouldReturnCorrectNameValue() {
        final String name = "my_awesomeTest name-123";
        final Category category = new Category("", name, "", "");

        assertEquals(name, category.getName());
    }

    @Test
    public void getColorShouldReturnCorrectColorValue() {
        final String color = "#123456";
        final Category category = new Category("", "", color, "");

        assertEquals(color, category.getColor());
    }

    @Test
    public void getIconShouldReturnCorrectIconValue() {
        final String icon = "beer-outline";
        final Category category = new Category("", "", "", icon);

        assertEquals(icon, category.getIcon());
    }

    @Test
    public void emptyStaticMethodShouldReturnCategoryWithEmptyValues() {
        final Category category = Category.empty();

        assertEquals("", category.getId());
        assertEquals("", category.getName());
        assertEquals("", category.getColor());
        assertEquals("", category.getIcon());
    }
}
