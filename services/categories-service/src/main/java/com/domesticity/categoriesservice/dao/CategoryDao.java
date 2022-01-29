package com.domesticity.categoriesservice.dao;

import java.util.List;
import java.util.Optional;
import java.util.UUID;
import com.domesticity.categoriesservice.model.Category;

public interface CategoryDao {

    int addCategory(String id, Category category);

    List<Category> getCategories();

    List<Category> getCategoriesByName(String name);

    List<Category> getCategoriesByColor(String color);

    List<Category> getCategoriesByIcon(String icon);

    Optional<Category> getCategory(String id);

    int deleteCategory(String id);

    int updateCategory(String id, Category newCategory);

    default int addCategory(Category category) {
        final String id = UUID.randomUUID().toString();
        return addCategory(id, category);
    }
}
