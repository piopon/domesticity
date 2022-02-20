package com.domesticity.categoriesservice;

import com.domesticity.categoriesservice.api.CategoryController;
import com.domesticity.categoriesservice.api.ErrorPageController;
import com.domesticity.categoriesservice.api.HealthPageController;
import com.domesticity.categoriesservice.api.HomePageController;
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
	private ErrorPageController errorPageController;
	@Autowired
	private HealthPageController healthPageController;
	@Autowired
	private HomePageController homePageController;
	@Autowired
	private CategoryService categoryService;

	@Test
	public void mainStartsApplicationCorrectly() {
		CategoriesServiceApplication.main(new String[] {});
	}

	@Test
	void contextLoadsAllControllers() {
		assertThat(categoryController).isNotNull();
		assertThat(errorPageController).isNotNull();
		assertThat(healthPageController).isNotNull();
		assertThat(homePageController).isNotNull();
	}

	@Test
	void contextLoadsCategoryService() {
		assertThat(categoryService).isNotNull();
	}
}
