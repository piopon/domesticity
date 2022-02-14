package com.domesticity.categoriesservice;

import com.domesticity.categoriesservice.api.CategoryController;
import com.domesticity.categoriesservice.service.CategoryService;

import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

import static org.assertj.core.api.Assertions.assertThat;

@SpringBootTest
class CategoriesServiceApplicationTest {

	@Autowired
	private CategoryController categoryController;
	@Autowired
	private CategoryService categoryService;

	@Test
	void contextLoadsCategoryController() {
		assertThat(categoryController).isNotNull();
	}

	@Test
	void contextLoadsCategoryService() {
		assertThat(categoryService).isNotNull();
	}
}
