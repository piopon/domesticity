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
    public void setup() {
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
    public void addCategoryWithExplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory("123", Category.empty());

        assertEquals(1, result);
        assertEquals(4, countAllCategories());
        assertTrue(testDao.getAllCategories().get(3).getId().equals("123"));
    }

    @Test
    public void addCategoryWithImplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory(Category.empty());

        assertEquals(1, result);
        assertEquals(4, countAllCategories());
        assertTrue(testDao.getAllCategories().get(3).getName().equals(""));
    }

    @Test
    public void getAllCategoriesShouldRetrieveAllStoredCategories() {
        assertEquals(3, countAllCategories());
        assertEquals("1", testDao.getAllCategories().get(0).getId());
        assertEquals("green", testDao.getAllCategories().get(1).getName());
        assertEquals("#0000FF", testDao.getAllCategories().get(2).getColor());
    }

    @Test
    public void getCategoryShouldRetrieveSelectedItemFromDb() {
        Optional<Category> actualItem = testDao.getCategory("1");

        assertTrue(actualItem.isPresent());
        assertEquals("red", actualItem.get().getName());
        assertEquals("#FF0000", actualItem.get().getColor());
        assertEquals("icon1", actualItem.get().getIcon());
    }

    @Test
    public void getCategoryShouldReturnEmptyItemIfIdIsNotFound() {
        Optional<Category> actualItem = testDao.getCategory("123");

        assertTrue(actualItem.isEmpty());
    }

    @Test
    public void deleteCategoryShouldRemoveExistingItem() {
        int result = testDao.deleteCategory("1");

        assertEquals(1, result);
        assertEquals(2, countAllCategories());
    }

    @Test
    public void deleteCategoryDoesNothingIfItemDoesNotExist() {
        int result = testDao.deleteCategory("123");

        assertEquals(0, result);
        assertEquals(3, countAllCategories());
    }

    @Test
    public void modifyCategoryShouldUpdateExistingItem() {
        int result = testDao.updateCategory("1", new Category("", "new", "#FFFFFF", ""));

        assertEquals(1, result);
        assertEquals(3, countAllCategories());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getName().equals("new")).count() == 1);
    }

    @Test
    public void modifyCategoryShouldDoNothingWithNotExistingItem() {
        int result = testDao.updateCategory("123", new Category("", "name", "color", "icon"));

        assertEquals(0, result);
        assertEquals(3, countAllCategories());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getId().equals("123")).count() == 0);
    }

    private int countAllCategories() {
        return testDao.getAllCategories().size();
    }
}
