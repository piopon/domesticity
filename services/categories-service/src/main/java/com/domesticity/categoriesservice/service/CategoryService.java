package com.domesticity.categoriesservice.service;

import java.util.List;
import java.util.Optional;
import java.util.UUID;

import com.domesticity.categoriesservice.dao.CategoryDao;
import com.domesticity.categoriesservice.model.Category;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;

@Service
public class CategoryService {

    private final CategoryDao categoryDao;

    @Autowired
    public CategoryService(@Qualifier("memory") CategoryDao categoryDao) {
        this.categoryDao = categoryDao;
    }

    public int addCategory(Category category) {
        return categoryDao.addCategory(category);
    }

    public List<Category> getCategories() {
        return categoryDao.getCategories();
    }

    public Optional<Category> getCategory(UUID id) {
        return categoryDao.getCategory(id);
    }
}
