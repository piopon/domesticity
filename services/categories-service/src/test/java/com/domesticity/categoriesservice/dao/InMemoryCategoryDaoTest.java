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
    public void setup() {
        testDao = new InMemoryCategoryDao();
        assertEquals(0, testDao.getAllCategories().size());
    }

    @Test
    public void addCategoryWithExplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory("123", Category.empty());

        assertEquals(1, result);
        assertEquals(1, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().get(0).getId().equals("123"));
    }

    @Test
    public void addCategoryWithImplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory(Category.empty());

        assertEquals(1, result);
        assertEquals(1, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().get(0).getName().equals(""));
    }

    @Test
    public void getAllCategoriesShouldRetrieveAllInMemoryList() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        List<Category> actualList = testDao.getAllCategories();

        assertEquals(3, actualList.size());
        assertTrue(actualList.stream().filter(category -> category.getId().equals("123")).count() == 1);
    }

    @Test
    public void getCategoryShouldRetrieveSelectedItemFromMemoryList() {
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
    public void getCategoryShouldReturnEmptyItemIfIdIsNotFound() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        Optional<Category> actualItem = testDao.getCategory("1");

        assertTrue(actualItem.isEmpty());
    }

    @Test
    public void deleteCategoryShouldRemoveExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        int result = testDao.deleteCategory("123");

        assertEquals(1, result);
        assertEquals(2, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().stream().allMatch(category -> category.getName().equals("")));
    }

    @Test
    public void deleteCategoryDoesNothingIfItemDoesNotExist() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        int result = testDao.deleteCategory("1");

        assertEquals(0, result);
        assertEquals(3, testDao.getAllCategories().size());
    }

    @Test
    public void modifyCategoryShouldUpdateExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        int result = testDao.updateCategory("000", new Category("", "name", "colour", "icon"));

        assertEquals(1, result);
        assertEquals(3, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getName().equals("name")).count() == 2);
    }

    @Test
    public void modifyCategoryShouldDoNothingWithNotExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "colour", "icon"));

        int result = testDao.updateCategory("001", new Category("", "name", "colour", "icon"));

        assertEquals(0, result);
        assertEquals(3, testDao.getAllCategories().size());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getName().equals("name")).count() == 1);
    }
}
