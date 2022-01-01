package com.domesticity.categoriesservice.dao;

import com.domesticity.categoriesservice.model.Category;

import java.util.List;

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
}
