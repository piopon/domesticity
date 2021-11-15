package com.domesticity.categoriesservice.dao;

import java.util.List;
import java.util.Optional;
import java.util.UUID;
import com.domesticity.categoriesservice.model.Category;

public interface CategoryDao {

    int addCategory(UUID id, Category category);

    List<Category> getCategories();

    Optional<Category> getCategory(UUID id);

    int deleteCategory(UUID id);

    int updateCategory(UUID id, Category newCategory);

    default int addCategory(Category category) {
        final UUID id = UUID.randomUUID();
        return addCategory(id, category);
    }
}
