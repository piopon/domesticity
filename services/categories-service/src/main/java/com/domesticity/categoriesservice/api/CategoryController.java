package com.domesticity.categoriesservice.api;

import com.domesticity.categoriesservice.model.Category;
import com.domesticity.categoriesservice.service.CategoryService;

public class CategoryController {
    
    private final CategoryService categoryService;

    public CategoryController(CategoryService categoryService) {
        this.categoryService = categoryService;
    }

    public void addCategory(Category category) {
        categoryService.addCategory(category);
    }
}
