package com.domesticity.categoriesservice.model;

import static org.junit.jupiter.api.Assertions.assertEquals;

import org.junit.jupiter.api.Test;

public class CategoryTest {

    @Test
    void getIdShouldReturnCorrectIdValue() {
        final String id = "test_id1";
        final Category category = new Category(id, "", "", "");

        assertEquals(id, category.getId());
    }

    @Test
    void getNameShouldReturnCorrectNameValue() {
        final String name = "my_awesomeTest name-123";
        final Category category = new Category("", name, "", "");

        assertEquals(name, category.getName());
    }

    @Test
    void getColourShouldReturnCorrectColourValue() {
        final String colour = "#123456";
        final Category category = new Category("", "", colour, "");

        assertEquals(colour, category.getColour());
    }

    @Test
    void getIconShouldReturnCorrectIconValue() {
        final String icon = "beer-outline";
        final Category category = new Category("", "", "", icon);

        assertEquals(icon, category.getIcon());
    }

    @Test
    void emptyStaticMethodShouldReturnCategoryWithEmptyValues() {
        final Category category = Category.empty();

        assertEquals("", category.getId());
        assertEquals("", category.getName());
        assertEquals("", category.getColour());
        assertEquals("", category.getIcon());
    }
}
