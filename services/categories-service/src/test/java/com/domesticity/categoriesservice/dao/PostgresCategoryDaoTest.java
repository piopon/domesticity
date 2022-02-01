package com.domesticity.categoriesservice.dao;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

import java.util.Optional;

import com.domesticity.categoriesservice.model.Category;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.springframework.jdbc.core.JdbcTemplate;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabase;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabaseBuilder;
import org.springframework.jdbc.datasource.embedded.EmbeddedDatabaseType;

public class PostgresCategoryDaoTest {

    private PostgresCategoryDao testDao;

    @BeforeEach
    void setup() {
        EmbeddedDatabase db = new EmbeddedDatabaseBuilder()
            .generateUniqueName(true)
            .setType(EmbeddedDatabaseType.H2)
            .setScriptEncoding("UTF-8")
            .addScript("schema.sql")
            .addScript("test-data.sql")
            .build();
        JdbcTemplate jdbc = new JdbcTemplate(db);
        testDao = new PostgresCategoryDao(jdbc);
    }

    @Test
    void addCategoryWithExplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory("123", Category.empty());

        assertEquals(1, result);
        assertEquals(4, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().get(3).getId().equals("123"));
    }

    @Test
    void addCategoryWithImplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory(Category.empty());

        assertEquals(1, result);
        assertEquals(4, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().get(3).getName().equals(""));
    }

    @Test
    void getCategoriesShouldRetrieveAllStoredCategories() {
        assertEquals(3, testDao.getAllCategories().size());
        assertEquals("1", testDao.getAllCategories().get(0).getId());
        assertEquals("green", testDao.getAllCategories().get(1).getName());
        assertEquals("#0000FF", testDao.getAllCategories().get(2).getColour());
    }

    @Test
    void getCategoryShouldRetrieveSelectedItemFromDb() {
        Optional<Category> actualItem = testDao.getCategory("1");

        assertTrue(actualItem.isPresent());
        assertEquals("red", actualItem.get().getName());
        assertEquals("#FF0000", actualItem.get().getColour());
        assertEquals("icon1", actualItem.get().getIcon());
    }

    @Test
    void getCategoryShouldReturnEmptyItemIfIdIsNotFound() {
        Optional<Category> actualItem = testDao.getCategory("123");

        assertTrue(actualItem.isEmpty());
    }

    @Test
    void deleteCategoryShouldRemoveExistingItem() {
        int result = testDao.deleteCategory("1");

        assertEquals(1, result);
        assertEquals(2, testDao.getAllCategories().size());
    }

    @Test
    void deleteCategoryDoesNothingIfItemDoesNotExist() {
        int result = testDao.deleteCategory("123");

        assertEquals(0, result);
        assertEquals(3, testDao.getAllCategories().size());
    }

    @Test
    void modifyCategoryShouldUpdateExistingItem() {
        int result = testDao.updateCategory("1", new Category("", "new", "#FFFFFF", ""));

        assertEquals(1, result);
        assertEquals(3, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getName().equals("new")).count() == 1);
    }

    @Test
    void modifyCategoryShouldDoNothingWithNotExistingItem() {
        int result = testDao.updateCategory("123", new Category("", "name", "colour", "icon"));

        assertEquals(0, result);
        assertEquals(3, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getId().equals("123")).count() == 0);
    }
}
