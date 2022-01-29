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

    public List<Category> getCategories() {
        return categoryDao.getCategories();
    }

    public List<Category> getCategoriesByName(String name) {
        return categoryDao.getCategoriesByName(name);
    }

    public List<Category> getCategoriesByColor(String color) {
        return categoryDao.getCategoriesByColor(color);
    }

    public List<Category> getCategoriesByIcon(String icon) {
        return categoryDao.getCategoriesByIcon(icon);
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
