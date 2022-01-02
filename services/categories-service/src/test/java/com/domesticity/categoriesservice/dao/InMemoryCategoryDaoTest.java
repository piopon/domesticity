package com.domesticity.categoriesservice.dao;

import com.domesticity.categoriesservice.model.Category;

import java.util.List;
import java.util.Optional;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;

public class InMemoryCategoryDaoTest {

    private InMemoryCategoryDao testDao;

    @BeforeEach
    void setup() {
        testDao = new InMemoryCategoryDao();
        assertEquals(0, testDao.getCategories().size());
    }

    @Test
    void addCategoryWithExplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory("123", Category.empty());

        assertEquals(1, result);
        assertEquals(1, testDao.getCategories().size());
        assertTrue(testDao.getCategories().get(0).getId().equals("123"));
    }

    @Test
    void addCategoryWithImplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory(Category.empty());

        assertEquals(1, result);
        assertEquals(1, testDao.getCategories().size());
        assertTrue(testDao.getCategories().get(0).getName().equals(""));
    }

    @Test
    void getCategoriesShouldRetrieveAllInMemoryList() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        List<Category> actualList = testDao.getCategories();

        assertEquals(3, actualList.size());
        assertTrue(actualList.stream().filter(category -> category.getId().equals("123")).count() == 1);
    }

    @Test
    void getCategoryShouldRetrieveSelectedItemFromMemoryList() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        Optional<Category> actualItem = testDao.getCategory("123");

        assertTrue(actualItem.isPresent());
        assertTrue(actualItem.get().getName().equals("name"));
        assertTrue(actualItem.get().getColour().equals("colour"));
        assertTrue(actualItem.get().getIcon().equals("icon"));
    }

    @Test
    void getCategoryShouldReturnEmptyItemIfIdIsNotFound() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        Optional<Category> actualItem = testDao.getCategory("1");

        assertTrue(actualItem.isEmpty());
    }

    @Test
    void deleteCategoryShouldRemoveExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        int result = testDao.deleteCategory("123");

        assertEquals(1, result);
        assertEquals(2, testDao.getCategories().size());
        assertTrue(testDao.getCategories().stream().allMatch(category -> category.getName().equals("")));
    }

    @Test
    void deleteCategoryDoesNothingIfItemDoesNotExist() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        int result = testDao.deleteCategory("1");

        assertEquals(0, result);
        assertEquals(3, testDao.getCategories().size());
    }
}
