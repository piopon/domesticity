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
        assertEquals(0, countAllCategories());
    }

    @Test
    public void addCategoryWithExplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory("123", Category.empty());

        assertEquals(1, result);
        assertEquals(1, countAllCategories());
        assertTrue(testDao.getAllCategories().get(0).getId().equals("123"));
    }

    @Test
    public void addCategoryWithImplicitIdShouldUpdateInMemoryList() {
        int result = testDao.addCategory(Category.empty());

        assertEquals(1, result);
        assertEquals(1, countAllCategories());
        assertTrue(testDao.getAllCategories().get(0).getName().equals(""));
    }

    @Test
    public void getAllCategoriesShouldRetrieveAllInMemoryList() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        List<Category> actualList = testDao.getAllCategories();

        assertEquals(3, actualList.size());
        assertTrue(actualList.stream().filter(category -> category.getId().equals("123")).count() == 1);
    }

    @Test
    public void getEmptyFilteredCategoriesShouldRetrieveAllStoredCategories() {
        testDao.addCategory("007", new Category("", "name1", "color1", "icon1"));
        testDao.addCategory("000", new Category("", "name2", "color2", "icon1"));
        testDao.addCategory("123", new Category("", "name3", "color2", "icon2"));

        List<Category> actualList = testDao.getFilteredCategories(null, null, null);

        assertEquals(3, actualList.size());
        assertEquals("name1", actualList.get(0).getName());
        assertEquals("color2", actualList.get(1).getColor());
        assertEquals("icon2", actualList.get(2).getIcon());
    }

    @Test
    public void getFilteredCategoriesByNameShouldRetrieveCorrectCategories() {
        testDao.addCategory("007", new Category("", "name1", "color1", "icon1"));
        testDao.addCategory("000", new Category("", "name2", "color2", "icon1"));
        testDao.addCategory("123", new Category("", "name3", "color2", "icon2"));

        List<Category> actualList = testDao.getFilteredCategories("name1", null, null);

        assertEquals(1, actualList.size());
        assertEquals("name1", actualList.get(0).getName());
        assertEquals("color1", actualList.get(0).getColor());
        assertEquals("icon1", actualList.get(0).getIcon());
    }

    @Test
    public void getCategoryShouldRetrieveSelectedItemFromMemoryList() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        Optional<Category> actualItem = testDao.getCategory("123");

        assertTrue(actualItem.isPresent());
        assertTrue(actualItem.get().getName().equals("name"));
        assertTrue(actualItem.get().getColor().equals("color"));
        assertTrue(actualItem.get().getIcon().equals("icon"));
    }

    @Test
    public void getCategoryShouldReturnEmptyItemIfIdIsNotFound() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        Optional<Category> actualItem = testDao.getCategory("1");

        assertTrue(actualItem.isEmpty());
    }

    @Test
    public void deleteCategoryShouldRemoveExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        int result = testDao.deleteCategory("123");

        assertEquals(1, result);
        assertEquals(2, countAllCategories());
        assertTrue(testDao.getAllCategories().stream().allMatch(category -> category.getName().equals("")));
    }

    @Test
    public void deleteCategoryDoesNothingIfItemDoesNotExist() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        int result = testDao.deleteCategory("1");

        assertEquals(0, result);
        assertEquals(3, countAllCategories());
    }

    @Test
    public void modifyCategoryShouldUpdateExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        int result = testDao.updateCategory("000", new Category("", "name", "color", "icon"));

        assertEquals(1, result);
        assertEquals(3, countAllCategories());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getName().equals("name")).count() == 2);
    }

    @Test
    public void modifyCategoryShouldDoNothingWithNotExistingItem() {
        testDao.addCategory("007", Category.empty());
        testDao.addCategory("000", Category.empty());
        testDao.addCategory("123", new Category("", "name", "color", "icon"));

        int result = testDao.updateCategory("001", new Category("", "name", "color", "icon"));

        assertEquals(0, result);
        assertEquals(3, countAllCategories());
        assertTrue(testDao.getAllCategories().stream().filter(category -> category.getName().equals("name")).count() == 1);
    }

    private int countAllCategories() {
        return testDao.getAllCategories().size();
    }
}
