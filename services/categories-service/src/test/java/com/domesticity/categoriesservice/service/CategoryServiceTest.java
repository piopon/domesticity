package com.domesticity.categoriesservice.service;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.dao.CategoryDao;
import com.domesticity.categoriesservice.model.Category;

import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.mockito.InjectMocks;
import org.mockito.Mock;
import org.mockito.MockitoAnnotations;

public class CategoryServiceTest {

    @InjectMocks
    CategoryService testService;

    @Mock
    CategoryDao dao;

    @BeforeEach
    public void setup() {
        MockitoAnnotations.initMocks(this);
    }

    @Test
    public void addCategoryShouldAddNonExistingCategory() {
        Category category = new Category("id1", "name1", "colour1", "icon1");
        testService.addCategory(category);

        verify(dao, times(1)).addCategory(category);
    }

    @Test
    public void getCategoriesShouldReturnAllSavedCategories() {
        when(dao.getAllCategories()).thenReturn(List.of(
            new Category("id1", "name1", "color1", "icon1"),
            new Category("id2", "name2", "color2", "icon2")
        ));

        List<Category> result = testService.getAllCategories();

        assertEquals(2, result.size());
        verify(dao, times(1)).getAllCategories();
    }

    @Test
    public void getCategoryShouldReturnCorrectEntry() {
        when(dao.getCategory("123")).thenReturn(Optional.of(new Category("123","nazwa","#FFFFFF","ball")));

        Optional<Category> result = testService.getCategory("123");

        verify(dao, times(1)).getCategory("123");
        assertTrue(result.isPresent());
        assertEquals("nazwa", result.get().getName());
        assertEquals("#FFFFFF", result.get().getColour());
        assertEquals("ball", result.get().getIcon());
    }

    @Test
    public void deleteCategoryShouldRemoveCategory() {
        testService.deleteCategory("123");

        verify(dao, times(1)).deleteCategory("123");
    }

    @Test
    public void updateCategoryShouldUpdateCategoryEntry() {
        Category category = new Category("", "name1", "colour1", "icon1");
        testService.updateCategory("123", category);

        verify(dao, times(1)).updateCategory("123", category);
    }
}
