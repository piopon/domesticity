package com.domesticity.categoriesservice.service;

import java.util.List;
import java.util.Optional;

import com.domesticity.categoriesservice.dao.CategoryDao;
import com.domesticity.categoriesservice.model.Category;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

@Service
public class CategoryService {

    private final CategoryDao categoryDao;

    @Autowired
    public CategoryService(@Qualifier("CategoryDaoRepository") CategoryDao categoryDao) {
        this.categoryDao = categoryDao;
    }

    public int addCategory(Category category) {
        return categoryDao.addCategory(category);
    }

    public List<Category> getAllCategories() {
        return categoryDao.getAllCategories();
    }

    public List<Category> getFilteredCategories(String name, String color, String icon) {
        return categoryDao.getFilteredCategories(name, color, icon);
    }

    public Optional<Category> getCategory(String id) {
        return categoryDao.getCategory(id);
    }

    public int deleteCategory(String id) {
        return categoryDao.deleteCategory(id);
    }

    public int updateCategory(String id, Category newCategory) {
        return categoryDao.updateCategory(id, newCategory);
    }
}
