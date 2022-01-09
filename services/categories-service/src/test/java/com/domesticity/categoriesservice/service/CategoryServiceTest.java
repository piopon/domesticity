package com.domesticity.categoriesservice.service;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.mockito.Mockito.times;
import static org.mockito.Mockito.verify;
import static org.mockito.Mockito.when;

import java.util.List;

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
        when(dao.getCategories()).thenReturn(List.of(
            new Category("id1", "name1", "color1", "icon1"),
            new Category("id2", "name2", "color2", "icon2")
        ));

        List<Category> result = testService.getCategories();

        assertEquals(2, result.size());
        verify(dao, times(1)).getCategories();
    }
}
